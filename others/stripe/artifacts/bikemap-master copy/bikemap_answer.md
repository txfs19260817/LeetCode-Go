# BikeMap 模拟面试解答（基于本项目）

## 先对齐题干与项目数据

- 题干写的是 `coords.json`，结构为 `[{ "lat": ..., "lng": ... }]`。
- 本项目实际可用轨迹文件是 `ride-simple.json`，结构是 GeoJSON，坐标在：
  `features[0].geometry.coordinates`，每个点是 `[lon, lat]`。
- 本项目地图渲染接口是：`https://stripe-bikemap.appspot.com/map.png`（见 `2-staticmap.md`）。

下面按项目内容给出一份可落地的 Python 解答。

## Q1 JSON 解析

目标：
1. 从磁盘读取轨迹 JSON。
2. 转成内存对象。
3. 打印前 10 个坐标。
4. 返回坐标供后续复用。

已从 `ride-simple.json` 验证：共有 `495` 个坐标点，前 10 个（`lon, lat`）为：

1. `[-122.2851, 47.55126]`
2. `[-122.28531, 47.55126]`
3. `[-122.28535, 47.55126]`
4. `[-122.28565, 47.55126]`
5. `[-122.28575, 47.55126]`
6. `[-122.28598, 47.55126]`
7. `[-122.2862, 47.55126]`
8. `[-122.28631, 47.55126]`
9. `[-122.28687, 47.55127]`
10. `[-122.28767, 47.55127]`

## Q2 POST 请求并保存 PNG

用 `requests.post()` 把 `staticmap_example.json` 发到接口，二进制响应写入 `map.png`。

## Q3 替换路径为你的坐标

- 保留 Q2 的请求结构。
- 把 `paths` 改成从 Q1 提取的轨迹坐标。
- 发送并保存为 `map_with_path.png`。

## Q4 可选变体

- 增加第二条路径（例如抽样点形成第二条线）。
- 或给每个点加 marker（会比较密，通常建议抽样后再加）。

## Q5 可选变体

- 计算轨迹总距离（Haversine）。
- 或只发送部分坐标（例如每 5 个点取 1 个）以减少 payload。

---

## Python 参考实现（可直接作为面试答案）

```python
from __future__ import annotations

import copy
import json
import math
from pathlib import Path
from typing import Any

import requests

API_URL = "https://stripe-bikemap.appspot.com/map.png"


def load_ride_coords(ride_path: str) -> list[dict[str, float]]:
    """
    读取本项目的 ride-simple.json (GeoJSON) 并转成 staticmap 需要的点格式:
    [{"lat": ..., "lon": ...}, ...]
    """
    data = json.loads(Path(ride_path).read_text(encoding="utf-8"))
    coords = data["features"][0]["geometry"]["coordinates"]  # [lon, lat]

    points: list[dict[str, float]] = []
    for lon, lat in coords:
        points.append({"lat": float(lat), "lon": float(lon)})
    return points


def print_first_n(points: list[dict[str, float]], n: int = 10) -> None:
    for i, p in enumerate(points[:n], start=1):
        print(f"{i:>2}. lat={p['lat']}, lon={p['lon']}")


def post_map(payload: dict[str, Any], output_file: str, url: str = API_URL) -> None:
    resp = requests.post(url, json=payload, timeout=30)
    resp.raise_for_status()

    Path(output_file).write_bytes(resp.content)
    print(f"saved: {output_file} ({len(resp.content)} bytes)")


def load_json(path: str) -> dict[str, Any]:
    return json.loads(Path(path).read_text(encoding="utf-8"))


def build_payload_with_route(
    base_payload: dict[str, Any],
    route_points: list[dict[str, float]],
    color: str = "blue",
) -> dict[str, Any]:
    payload = copy.deepcopy(base_payload)
    payload["paths"] = [
        {
            "color": color,
            "positions": route_points,
        }
    ]
    return payload


def haversine_km(lat1: float, lon1: float, lat2: float, lon2: float) -> float:
    r = 6371.0
    dlat = math.radians(lat2 - lat1)
    dlon = math.radians(lon2 - lon1)
    a = (
        math.sin(dlat / 2) ** 2
        + math.cos(math.radians(lat1))
        * math.cos(math.radians(lat2))
        * math.sin(dlon / 2) ** 2
    )
    return 2 * r * math.asin(math.sqrt(a))


def total_distance_km(points: list[dict[str, float]]) -> float:
    if len(points) < 2:
        return 0.0
    total = 0.0
    for i in range(1, len(points)):
        p1 = points[i - 1]
        p2 = points[i]
        total += haversine_km(p1["lat"], p1["lon"], p2["lat"], p2["lon"])
    return total


def main() -> None:
    # Q1
    points = load_ride_coords("ride-simple.json")
    print(f"loaded points: {len(points)}")
    print_first_n(points, n=10)

    # Q2
    base_payload = load_json("staticmap_example.json")
    post_map(base_payload, "map.png")

    # Q3
    payload_with_route = build_payload_with_route(base_payload, points, color="red")
    post_map(payload_with_route, "map_with_path.png")

    # Q4 (可选): 第二条抽样路径
    sampled = points[::5]
    payload_variant = copy.deepcopy(payload_with_route)
    payload_variant["paths"].append(
        {"color": "purple", "positions": sampled}
    )
    post_map(payload_variant, "map_with_two_paths.png")

    # Q5 (可选): 计算总距离 + 只发部分点
    dist = total_distance_km(points)
    print(f"approx total distance: {dist:.2f} km")

    payload_partial = build_payload_with_route(base_payload, sampled, color="green")
    post_map(payload_partial, "map_with_partial_path.png")


if __name__ == "__main__":
    main()
```

## 运行说明

1. 安装依赖：`pip install requests`
2. 将上面代码保存为 `answer.py`（或拆到你自己的工程结构）
3. 在项目根目录运行：`python answer.py`
4. 输出文件：
   - `map.png`（Q2）
   - `map_with_path.png`（Q3）
   - `map_with_two_paths.png`（Q4 可选）
   - `map_with_partial_path.png`（Q5 可选）

## 面试讲解要点（可口述）

1. 先做数据模型适配：GeoJSON `[lon, lat]` -> API 期望 `{lat, lon}`。
2. 把“读取 JSON / 构建 payload / 发请求 / 保存文件”分函数，便于测试和扩展。
3. 变体题（Q4/Q5）复用主流程，只替换 `paths` 或对点集做变换。

