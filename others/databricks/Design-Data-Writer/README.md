# Design Data Writer

**Difficulty:** Hard  
**Company:** Databricks  
**Topics:** Concurrency, I/O, Systems Design

Design a thread-safe library that appends data to a file on disk.
Thousands of threads call `push(data)` concurrently; when `push` returns
the data **must** be persisted (write + fsync).

```
class DataWriter {
    public DataWriter(String filePathOnDisk) {}
    public void push(byte[] data) {}
}
```

## Requirements

1. **Thread-level ordering** — If thread A pushes d1 then d2, d1 must
   precede d2 in the file. Interleaving across threads is allowed.
2. **High throughput, low latency** — Minimise the time each thread is
   blocked inside `push`.
3. **Persistence** — Show the on-disk record format (bytes).
4. **Crash recovery** — Discuss / implement how to recover valid records
   after a crash that may leave a partial write at the end of the file.

## Design — Group Commit

The key insight is **group commit** (a.k.a. batched fsync):

1. `Push` puts its data + a "done" signal into a channel and **blocks**.
2. A single background writer goroutine:
   - Waits for at least one entry.
   - Drains all currently-available entries (non-blocking).
   - Serialises the batch into a buffer.
   - Calls `write` + `fsync` **once** for the whole batch.
   - Signals all waiters so their `Push` calls return.

This way N concurrent pushes share **one** fsync — the most expensive
operation — giving O(1) fsyncs per batch instead of O(N).

## Record Format

```
[4 bytes: data length (little-endian uint32)]
[N bytes: data]
[4 bytes: CRC32 of (length bytes ++ data bytes)]
```

Total overhead: 8 bytes per record.

## Crash Recovery

Scan the file record-by-record:
- Read length → read data → read CRC → verify.
- On the first incomplete or CRC-mismatched record, stop.
- Truncate the file at that boundary (all records before it are valid).

## Follow-up

- **Concurrency** — Why a single writer goroutine instead of multiple?
  (Avoids interleaved / torn writes, simplifies fsync batching.)
- **Back-pressure** — Bounded channel (e.g. 4096) naturally applies
  back-pressure when the writer can't keep up.
- **WAL / journaling** — How databases use the same group-commit pattern.
- **Durability vs latency trade-off** — Tuning fsync frequency.
