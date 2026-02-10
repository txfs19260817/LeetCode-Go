package databricks

import (
	"bytes"
	"encoding/binary"
	"hash/crc32"
	"io"
	"os"
	"sync"
)

// ---------------------------------------------------------------------------
// DataWriter — thread-safe, group-commit append-only file writer
// ---------------------------------------------------------------------------

// pendingWrite is an entry submitted by a Push caller.
type pendingWrite struct {
	data []byte
	done chan struct{} // closed by the batch writer once persisted
}

// DataWriter appends records to a file with group-commit batching.
// Multiple goroutines may call Push concurrently.  Each Push blocks until
// its data is written and fsynced to disk, but concurrent pushes share a
// single write+fsync call, amortising the I/O cost.
//
// Record format on disk (per Push):
//
//	[4B data length (LE uint32)] [data bytes] [4B CRC32 (LE uint32)]
//
// CRC32 (IEEE) covers the 4-byte length header concatenated with the data.
type DataWriter struct {
	file *os.File
	ch   chan *pendingWrite
	wg   sync.WaitGroup
}

// NewDataWriter opens (or creates) filePath for appending and starts the
// background batch-writer goroutine.
func NewDataWriter(filePath string) (*DataWriter, error) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	dw := &DataWriter{
		file: f,
		ch:   make(chan *pendingWrite, 4096),
	}
	dw.wg.Add(1)
	go dw.batchLoop()
	return dw, nil
}

// Push appends data to the file.  It blocks until the data has been written
// and fsynced (possibly as part of a larger batch).
func (dw *DataWriter) Push(data []byte) {
	cp := make([]byte, len(data))
	copy(cp, data) // defensive copy
	pw := &pendingWrite{data: cp, done: make(chan struct{})}
	dw.ch <- pw
	<-pw.done
}

// Close flushes any pending writes and closes the file.
// Callers must ensure no Push calls are in-flight when Close is called.
func (dw *DataWriter) Close() error {
	dw.ch <- nil // sentinel
	dw.wg.Wait()
	return dw.file.Close()
}

// batchLoop is the single background goroutine that serialises all writes.
func (dw *DataWriter) batchLoop() {
	defer dw.wg.Done()
	batch := make([]*pendingWrite, 0, 256)

	for {
		batch = batch[:0]

		// Phase 1: block until at least one entry (or close sentinel).
		pw := <-dw.ch
		if pw == nil {
			return
		}
		batch = append(batch, pw)

		// Phase 2: drain all currently-available entries (non-blocking).
	drain:
		for {
			select {
			case pw := <-dw.ch:
				if pw == nil {
					// Close arrived during drain — flush what we have.
					dw.writeBatch(batch)
					return
				}
				batch = append(batch, pw)
			default:
				break drain
			}
		}

		// Phase 3: write + fsync + unblock callers.
		dw.writeBatch(batch)
	}
}

// writeBatch serialises all entries, writes them in one call, fsyncs,
// then signals every caller.
func (dw *DataWriter) writeBatch(batch []*pendingWrite) {
	var buf bytes.Buffer
	for _, pw := range batch {
		encodeRecord(&buf, pw.data)
	}
	dw.file.Write(buf.Bytes()) //nolint: errcheck
	dw.file.Sync()             //nolint: errcheck
	for _, pw := range batch {
		close(pw.done)
	}
}

// encodeRecord writes one length-prefixed, CRC-protected record into buf.
func encodeRecord(buf *bytes.Buffer, data []byte) {
	var hdr [4]byte
	binary.LittleEndian.PutUint32(hdr[:], uint32(len(data)))

	h := crc32.NewIEEE()
	h.Write(hdr[:])
	h.Write(data)

	buf.Write(hdr[:])
	buf.Write(data)
	binary.Write(buf, binary.LittleEndian, h.Sum32()) //nolint: errcheck
}

// ---------------------------------------------------------------------------
// RecoverRecords — crash recovery
// ---------------------------------------------------------------------------

// RecoverRecords reads a file produced by DataWriter and returns every
// fully-written, CRC-verified record.  If the file ends with an
// incomplete or corrupted record (e.g. due to a crash mid-write), that
// trailing fragment is silently ignored.
func RecoverRecords(filePath string) ([][]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var records [][]byte
	for {
		// Read 4-byte length header.
		var length uint32
		if err := binary.Read(f, binary.LittleEndian, &length); err != nil {
			break
		}
		// Read data payload.
		data := make([]byte, length)
		if _, err := io.ReadFull(f, data); err != nil {
			break
		}
		// Read 4-byte CRC.
		var diskCRC uint32
		if err := binary.Read(f, binary.LittleEndian, &diskCRC); err != nil {
			break
		}
		// Verify CRC (covers length header + data).
		var hdr [4]byte
		binary.LittleEndian.PutUint32(hdr[:], length)
		h := crc32.NewIEEE()
		h.Write(hdr[:])
		h.Write(data)
		if h.Sum32() != diskCRC {
			break // corrupted — stop here
		}
		records = append(records, data)
	}
	return records, nil
}
