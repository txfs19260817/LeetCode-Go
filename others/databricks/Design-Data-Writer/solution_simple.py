from __future__ import annotations

import os
import struct
import tempfile
import threading
from queue import Empty, Queue


class SimpleDataWriter:
    """Group-commit writer with raw-byte append only.

    push(data) blocks until data is written and fsynced.
    No length-prefix, CRC, or crash-recovery logic.
    """

    def __init__(self, file_path: str) -> None:
        # Append-only binary file handle shared by the single writer thread.
        self._file = open(file_path, "ab")  # noqa: SIM115
        # Producer queue: each item is (payload, completion_event).
        # `None` is a sentinel meaning "flush remaining work and stop".
        self._queue: Queue[tuple[bytes, threading.Event] | None] = Queue(maxsize=4096)
        self._writer = threading.Thread(target=self._batch_loop, daemon=True)
        self._writer.start()

    def push(self, data: bytes) -> None:
        # Caller blocks until writer thread persists this payload.
        event = threading.Event()
        self._queue.put((bytes(data), event))
        event.wait()

    def close(self) -> None:
        # Send stop signal, wait for writer to finish, then close file.
        self._queue.put(None)
        self._writer.join()
        self._file.close()

    def _batch_loop(self) -> None:
        batch: list[tuple[bytes, threading.Event]] = []

        while True:
            try:
                # Blocking to wait for the first;
                # otherwise draining queue in a non-blocking way
                item = self._queue.get() if not batch else self._queue.get_nowait()
            except Empty:
                # Queue is empty, flushing
                self._flush(batch)
                batch.clear()
                continue

            if item is None:
                # Close signal: flushing remaining
                self._flush(batch)
                return

            batch.append(item)


    def _flush(self, batch: list[tuple[bytes, threading.Event]]) -> None:
        if not batch:
            return
        # Coalesce payloads into a single contiguous write.
        buf = bytearray()
        for data, _ in batch:
            buf.extend(data)
        self._file.write(buf)
        # Ensure durability for all records in the batch.
        self._file.flush()
        os.fsync(self._file.fileno())
        # Unblock all waiting push callers after durable commit.
        for _, event in batch:
            event.set()


if __name__ == "__main__":
    # Test 1: single-thread append
    with tempfile.NamedTemporaryFile(delete=False, suffix=".dat") as tmp:
        path = tmp.name
    dw = SimpleDataWriter(path)
    dw.push(b"hello")
    dw.push(b"world")
    dw.close()
    assert open(path, "rb").read() == b"helloworld"  # noqa: SIM115
    os.unlink(path)

    # Test 2: multi-thread thread-local ordering
    with tempfile.NamedTemporaryFile(delete=False, suffix=".dat") as tmp:
        path = tmp.name
    dw = SimpleDataWriter(path)
    num_threads, pushes_per = 50, 100
    threads = []

    for tid in range(num_threads):
        def worker(tid: int = tid) -> None:
            for seq in range(pushes_per):
                dw.push(struct.pack("<II", tid, seq))

        t = threading.Thread(target=worker)
        threads.append(t)
        t.start()

    for t in threads:
        t.join()
    dw.close()

    raw = open(path, "rb").read()  # noqa: SIM115
    os.unlink(path)

    record_size = 8
    total_records = num_threads * pushes_per
    assert len(raw) == total_records * record_size

    last_seq: dict[int, int] = {}
    for i in range(total_records):
        rec = raw[i * record_size : (i + 1) * record_size]
        tid, seq = struct.unpack("<II", rec)
        if tid in last_seq:
            assert seq > last_seq[tid], f"thread {tid}: {seq} <= {last_seq[tid]}"
        last_seq[tid] = seq
    assert len(last_seq) == num_threads

    print("All tests passed!")
