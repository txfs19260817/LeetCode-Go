假设有1m device，每分钟发送heartbeat，以及status （一共有10种status） 你需要设计一个系统可以实时计算上一分钟每一个status 的device 数量：- 实时收集设备健康值（每个设备发送每分钟心跳，1千万的设备体量） - 分钟级别的数据新鲜度 - 如何去重 - 如何降低收集，监控成本
---

## 1) 需求澄清（我会在面试里这样口头确认 + 给出假设）

**已知：**

* ~1M devices（你后面也写了“1千万”，我会按 **1M–10M** 的上限设计）
* 每台设备**每分钟**发一次 heartbeat + status（10 种之一）
* 需要**实时**计算“**上一分钟**每个 status 对应的 device 数量”
* 数据新鲜度：**分钟级**（通常意味着 minute close 后几秒内可读）
* 需要**去重**
* 需要降低**采集/监控成本**

**我做的合理假设（没特别说明就按这个落地）：**

* 每台设备每分钟“理应”发送一次；可能会重试导致重复；可能有少量乱序/延迟（网络抖动）
* 统计口径：**按 event-time** 的 minute bucket 统计（例如 12:01:00–12:01:59）
* “上一分钟”结果允许一个小的 finalize 延迟（比如 **5–10 秒**）以吸收迟到数据
* 10 种 status 中是否包含 OFFLINE？

    * 如果包含：设备没上报则不会被计入任何 status（或由上游显式上报 OFFLINE）
    * 如果不包含而你又想把“没心跳=offline”算进来：需要额外逻辑（我下面也给方案）

---

## 2) 规模估算（SLO/吞吐/存储）

### 吞吐

* **1M devices / min** ⇒ ~16.7K msg/sec 平均
* **10M devices / min** ⇒ ~166.7K msg/sec 平均

假设单条消息（protobuf）~100–200 bytes（device_id + ts + status + seq + 签名）：

* 10M 规模：~16–33 MB/s 级别写入（非常可做）

### 状态（state）规模

去重通常需要“每个 device 记住最近处理到哪儿了”：

* 10M devices ×（last_minute + last_seq + last_status）几十字节/条
  → RocksDB/Flink state 做 **10M 级 key-state**没问题（工程上常见）

---

## 3) 高层架构（推荐：流式处理 + 分钟窗口聚合）

目标是：**分钟级新鲜度 + 去重 + 低成本**，最稳的是：

1. **Ingestion 层**：IoT Gateway / LB → Kafka(Kinesis/MSK)
2. **Stream Processing**：Flink/Kafka Streams（带 state）

    * Stage A：按 device 做**去重/归一**（每 device 每 minute 产出一条“最终状态”）
    * Stage B：按 (minute, status) 做**tumbling window 聚合**，输出 counts
3. **Serving**：Redis / DynamoDB / Cassandra 存“minute → 10 个 status counts”
4. **可选 Raw**：原始事件低成本落 S3（可采样/短期保留），用于回放/审计

```mermaid
flowchart LR
  D[Devices\n(MQTT/HTTPS)] --> G[IoT Gateway / LB]
  G --> Q[(Kafka / Kinesis)]
  Q --> A[Flink Job A\nKeyBy device_id\nDedup + normalize]
  A --> Q2[(Topic: device_minute_status)]
  Q2 --> B[Flink Job B\nKeyBy (minute,status)\n1-min tumbling window]
  B --> S[(Redis/DynamoDB\nminute -> counts[10])]
  Q --> R[(S3 Raw events\noptional)]
  S --> API[Query API / Dashboard]
```

---

## 4) 关键数据模型 & API

### 设备上报消息（建议）

```json
{
  "device_id": "d123",
  "ts_ms": 1700000000123,      // event time
  "minute_id": 28333333,       // floor(ts/60s) 可由服务端算
  "status": 0..9,
  "seq": 102938,               // 单调递增（强烈建议）
  "sig": "..."                 // 可选：防伪/鉴权
}
```

### 聚合结果存储

* Key: `minute_id`
* Value: `counts[10]` + `finalized_at` + `watermark`

查询 API：

* `GET /v1/status_counts?minute_id=...`
* 或 `GET /v1/status_counts/latest`（返回上一分钟 finalized 的那个 minute）

---

## 5) 去重怎么做（核心点，面试必问）

> 你要解决两类重复：**重试重复**、**同一分钟多次上报（或乱序）**。

### 最推荐的去重策略：**device 端提供 seq（或唯一 event_id）**

Flink Job A 做 per-device state：

* `last_seq`（或 `last_event_id`）
* `last_minute_id`（可选）
* `last_status_for_minute`（处理“同一分钟多条”时需要）

处理逻辑（简化版）：

1. `if seq <= last_seq: drop`（硬去重，O(1)）
2. 计算 `minute_id = floor(ts/60s)`
3. 对“同一分钟多条”：

    * 如果你定义“每分钟以**最后一条**为准”：

        * 如果 minute_id == last_minute_id：需要把该 device 在该 minute 的旧 status “撤销”，再加新 status
        * 这意味着 Job A 要输出“delta”（-1 old_status, +1 new_status）或者输出“最终状态”并让下游做覆盖
    * 如果你定义“每分钟以**第一条**为准”：

        * minute_id == last_minute_id 直接丢弃即可（更省事但不一定符合业务）

**工程上更常用**：以“最后一条为准”+ 少量 delta（因为重试/重复不多）。

### 如果设备端没有 seq/event_id

就只能用 “(device_id, minute_id)” 做幂等键：

* 保存 `last_minute_id` 和 `status_for_that_minute`
* 乱序/迟到会更难处理，且 state 更新更频繁
  **我会在面试里强调：强烈建议协议加 seq**，这会大幅降低复杂度和成本。

---

## 6) 分钟级新鲜度怎么保证（窗口 + watermark）

Job B 做 (minute_id, status) 的**1-min tumbling window**：

* 基于 event-time
* 允许迟到 `allowed_lateness = 5s`（或 10s）
* minute 结束 + lateness 后输出“最终 counts”
* Serving 端的 `latest` 永远返回 **最近 finalized 的 minute**

这样你可以明确说：

* 新鲜度：**minute close 后 ~5–10 秒可用**
* 迟到数据：在 lateness 内会修正（或者输出修正版本）

---

## 7) 摄像头很多怎么办？瓶颈在哪？

### 主要瓶颈 1：**状态存储（per-device dedup state）**

* 10M key-state 会落到 RocksDB，本质瓶颈是 **磁盘 I/O / compaction**
* 缓解：

    * keyBy device_id 均匀分区（足够多 partitions）
    * state TTL（只保留最近 X 天/小时的 last_seq/last_minute，过期清理）
    * checkpoint 频率合理（例如 1–5 分钟）+ incremental checkpoint 到 S3
    * SSD + 合理的 RocksDB tuning（block cache, compaction, write buffer）

### 主要瓶颈 2：**Shuffle/聚合成本**

* 如果你直接把原始事件按 (minute,status) 聚合，去重会让 counts 不准
* 所以先按 device 去重（Job A），再聚合（Job B）是为了减少下游噪声
* 分区建议：

    * Kafka partitions：按 device_id hash（保证同设备有序）
    * Flink 并行度：与 partitions 对齐

### 主要瓶颈 3：**网络/入口成本**

* 大量设备每分钟上报，入口 LB + TLS 终止成本不小
* 缓解：

    * MQTT（长连接）比每分钟 HTTPS 建连便宜
    * protobuf + gzip（或者仅 protobuf）
    * 多区域接入（就近 POP），后端 Kafka 跨 AZ 聚合

---

## 8) 如何降低采集、监控成本（你这题的“加分点”）

### 采集成本

* **协议层**：MQTT 长连接 + protobuf（最直接降钱）
* **语义层**（如果业务允许）：

    * heartbeat 仍每分钟，但 status 只在变化时发送（再由服务端“carry forward last status”补齐分钟状态）
    * 这能把 status 事件从 “每设备每分钟” 降到 “按变化频率”，成本大幅下降
    * 但你题目写“每分钟发 status”，我会把这个当 optional 优化提出

### 存储成本

* 不存每条原始事件到数据库（太贵）
* 只存：

    1. 分钟聚合结果（minute→10 counts）
    2. 必要的 per-device 去重 state（在 stream processor 的 state backend）
* 原始事件：**S3**（可采样/短期保留），需要时回放

### 监控成本

* 监控目标聚焦在 “pipeline 健康”：

    * Kafka lag / Flink backpressure / checkpoint duration
    * 每分钟 ingest 数 vs 预期数（异常报警）
    * 去重丢弃率（duplicate rate）异常报警
* 不要做“每设备实时监控”（会爆炸）；设备级问题用抽样/按需查询

---

## 9) 可靠性与一致性（面试官常追问）

* **至少一次 vs 恰好一次**

    * Kafka + Flink 支持 **exactly-once**（事务 sink / checkpoint）
    * 即便是 at-least-once，也可以用“设备 seq + 幂等更新”保证逻辑效果接近 exactly-once
* **数据丢失**

    * Gateway→Kafka：确认写入（acks=all），多 AZ replication
    * Flink checkpoint 到 S3，故障恢复可继续
* **迟到/乱序**

    * watermark + allowed lateness（你可以给出具体数字）
* **热点 key**

    * device_id hash 一般均匀；(minute,status) 聚合只有 10 个 status，可能会热点

        * 解决：先做分片聚合（key=(minute,status,shard)），最后再 reduce 合并 10 个 status（两层聚合）

---

## 10)（可选加分）如果要把“离线设备”也算进某个 status

如果要求“上一分钟每个 status 的设备数”且希望总和=总设备数：

* 需要知道全量设备集合 `TotalDevices`
* 维护 `last_seen_minute` per device（Job A 已经有）
* 在 minute finalize 时：

    * `offline = TotalDevices - seen_in_last_minute`
    * 或者：若 status 本身包含 OFFLINE，就由服务端补一条 OFFLINE 事件（更复杂但可做）

---

如果你愿意，我可以继续按“面试官追问模式”对你发问并逼你做取舍（比如：**exactly-once 你要不要？lateness 设多少？离线怎么算？Kafka 还是 Kinesis？sink 用 Redis 还是 DynamoDB？**），这样你练一遍就能在 Verkada 面试里把这题讲得很像真的做过。
