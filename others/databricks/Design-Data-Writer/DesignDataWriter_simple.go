package databricks

import (
	"os"
	"sync"
)

// ---------------------------------------------------------------------------
// SimpleDataWriter — group-commit only, no framing / CRC / recovery
// ---------------------------------------------------------------------------

type simplePendingWrite struct {
	data []byte
	done chan struct{}
}

// SimpleDataWriter is a stripped-down version of DataWriter that focuses
// purely on the group-commit concurrency pattern.  Data is written as raw
// bytes (no length-prefix, no CRC), so there is no crash-recovery support.
type SimpleDataWriter struct {
	file *os.File
	ch   chan *simplePendingWrite
	wg   sync.WaitGroup
}

func NewSimpleDataWriter(filePath string) (*SimpleDataWriter, error) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	dw := &SimpleDataWriter{
		file: f,
		ch:   make(chan *simplePendingWrite, 4096),
	}
	dw.wg.Add(1)
	go dw.batchLoop()
	return dw, nil
}

// Push appends data to the file.  Blocks until written + fsynced.
func (dw *SimpleDataWriter) Push(data []byte) {
	cp := make([]byte, len(data))
	copy(cp, data)
	pw := &simplePendingWrite{data: cp, done: make(chan struct{})}
	dw.ch <- pw
	<-pw.done
}

// Close flushes pending writes and closes the file.
func (dw *SimpleDataWriter) Close() error {
	dw.ch <- nil
	dw.wg.Wait()
	return dw.file.Close()
}

func (dw *SimpleDataWriter) batchLoop() {
	defer dw.wg.Done()
	batch := make([]*simplePendingWrite, 0, 256)

	for {
		batch = batch[:0]

		pw := <-dw.ch
		if pw == nil {
			return
		}
		batch = append(batch, pw)

	drain:
		for {
			select {
			case pw := <-dw.ch:
				if pw == nil {
					dw.writeSimpleBatch(batch)
					return
				}
				batch = append(batch, pw)
			default:
				break drain
			}
		}

		dw.writeSimpleBatch(batch)
	}
}

func (dw *SimpleDataWriter) writeSimpleBatch(batch []*simplePendingWrite) {
	// Concatenate all data into one buffer, single write + single fsync.
	total := 0
	for _, pw := range batch {
		total += len(pw.data)
	}
	buf := make([]byte, 0, total)
	for _, pw := range batch {
		buf = append(buf, pw.data...)
	}
	dw.file.Write(buf) //nolint: errcheck
	dw.file.Sync()     //nolint: errcheck
	for _, pw := range batch {
		close(pw.done)
	}
}
