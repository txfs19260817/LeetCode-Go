from __future__ import annotations

import os
import struct
import threading
import zlib
from io import BytesIO
from queue import Empty, Queue


# ---------------------------------------------------------------------------
# DataWriter — thread-safe, group-commit append-only file writer
# ---------------------------------------------------------------------------

class DataWriter:
    """Append-only file writer with group-commit batching.

    Multiple threads call push() concurrently.  Each push blocks until its
    data has been written and fsynced to disk, but concurrent pushes share
    a single write+fsync, amortising the I/O cost.

    Record format on disk::

        [4B data length (LE uint32)][data bytes][4B CRC32 (LE uint32)]
    """

    def __init__(self, file_path: str) -> None:
        self._file = open(file_path, "ab")  # noqa: SIM115
        self._queue: Queue[tuple[bytes, threading.Event] | None] = Queue(maxsize=4096)
        self._writer = threading.Thread(target=self._batch_loop, daemon=True)
        self._writer.start()

    def push(self, data: bytes) -> None:
        """Append *data* to the file.  Blocks until persisted."""
        event = threading.Event()
        self._queue.put((bytes(data), event))
        event.wait()

    def close(self) -> None:
        """Flush pending writes and close the file."""
        self._queue.put(None)  # sentinel
        self._writer.join()
        self._file.close()

    # -- internal ------------------------------------------------------------

    def _batch_loop(self) -> None:
        batch: list[tuple[bytes, threading.Event]] = []

        while True:
            try:
                # Block for the first item, then opportunistically drain.
                item = self._queue.get() if not batch else self._queue.get_nowait()
            except Empty:
                self._flush(batch)
                batch.clear()
                continue

            if item is None:
                self._flush(batch)
                return

            batch.append(item)

    def _flush(self, batch: list[tuple[bytes, threading.Event]]) -> None:
        if not batch:
            return
        buf = BytesIO()
        for data, _ in batch:
            _encode_record(buf, data)
        self._file.write(buf.getvalue())
        self._file.flush()
        os.fsync(self._file.fileno())
        for _, event in batch:
            event.set()


def _encode_record(buf: BytesIO, data: bytes) -> None:
    hdr = struct.pack("<I", len(data))
    buf.write(hdr)
    buf.write(data)
    crc = zlib.crc32(hdr + data) & 0xFFFFFFFF
    buf.write(struct.pack("<I", crc))


# ---------------------------------------------------------------------------
# Crash recovery
# ---------------------------------------------------------------------------

def recover_records(file_path: str) -> list[bytes]:
    """Read all fully-written, CRC-verified records from *file_path*.

    Any trailing incomplete or corrupted record is silently skipped.
    """
    raw = open(file_path, "rb").read()  # noqa: SIM115
    offset = 0
    records: list[bytes] = []
    while offset + 4 <= len(raw):
        (length,) = struct.unpack_from("<I", raw, offset)
        offset += 4
        if offset + length + 4 > len(raw):
            break  # incomplete data or CRC
        payload = raw[offset : offset + length]
        offset += length
        (disk_crc,) = struct.unpack_from("<I", raw, offset)
        offset += 4
        hdr = struct.pack("<I", length)
        expected = zlib.crc32(hdr + payload) & 0xFFFFFFFF
        if disk_crc != expected:
            break  # corrupted
        records.append(payload)
    return records


# ---------------------------------------------------------------------------
# __main__ tests
# ---------------------------------------------------------------------------

if __name__ == "__main__":
    import tempfile

    # Test 1: single-thread push + recovery
    with tempfile.NamedTemporaryFile(delete=False, suffix=".dat") as tmp:
        path = tmp.name
    dw = DataWriter(path)
    dw.push(b"hello")
    dw.push(b"world")
    dw.close()
    assert recover_records(path) == [b"hello", b"world"]
    os.unlink(path)

    # Test 2: multi-thread ordering
    with tempfile.NamedTemporaryFile(delete=False, suffix=".dat") as tmp:
        path = tmp.name
    dw = DataWriter(path)
    num_threads, pushes_per = 20, 50
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
    records = recover_records(path)
    assert len(records) == num_threads * pushes_per, f"got {len(records)}"
    last_seq: dict[int, int] = {}
    for rec in records:
        tid, seq = struct.unpack("<II", rec)
        if tid in last_seq:
            assert seq > last_seq[tid], f"thread {tid}: {seq} <= {last_seq[tid]}"
        last_seq[tid] = seq
    assert len(last_seq) == num_threads
    os.unlink(path)

    # Test 3: truncated file recovery
    with tempfile.NamedTemporaryFile(delete=False, suffix=".dat") as tmp:
        path = tmp.name
    dw = DataWriter(path)
    dw.push(b"rec1")
    dw.push(b"rec2")
    dw.close()
    with open(path, "ab") as f:
        f.write(b"\x05\x00\x00\x00\xAA")  # partial record
    records = recover_records(path)
    assert len(records) == 2
    assert records == [b"rec1", b"rec2"]
    os.unlink(path)

    # Test 4: corrupted CRC recovery
    with tempfile.NamedTemporaryFile(delete=False, suffix=".dat") as tmp:
        path = tmp.name
    dw = DataWriter(path)
    dw.push(b"good")
    dw.push(b"also-good")
    dw.close()
    raw = bytearray(open(path, "rb").read())
    raw[-1] ^= 0xFF  # flip CRC bits of second record
    open(path, "wb").write(raw)
    records = recover_records(path)
    assert len(records) == 1
    assert records[0] == b"good"
    os.unlink(path)

    print("All tests passed!")
