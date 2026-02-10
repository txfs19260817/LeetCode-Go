package databricks

import (
	"encoding/binary"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimpleDataWriter_SingleThread(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.dat")
	dw, err := NewSimpleDataWriter(path)
	require.NoError(t, err)

	dw.Push([]byte("hello"))
	dw.Push([]byte("world"))
	require.NoError(t, dw.Close())

	got, err := os.ReadFile(path)
	require.NoError(t, err)
	assert.Equal(t, []byte("helloworld"), got)
}

func TestSimpleDataWriter_MultiThread_Ordering(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.dat")
	dw, err := NewSimpleDataWriter(path)
	require.NoError(t, err)

	const numThreads = 50
	const pushesPerThread = 100
	const recordSize = 8 // [tid 4B][seq 4B]

	var wg sync.WaitGroup
	for tid := 0; tid < numThreads; tid++ {
		wg.Add(1)
		go func(tid int) {
			defer wg.Done()
			for seq := 0; seq < pushesPerThread; seq++ {
				data := make([]byte, recordSize)
				binary.LittleEndian.PutUint32(data[0:4], uint32(tid))
				binary.LittleEndian.PutUint32(data[4:8], uint32(seq))
				dw.Push(data)
			}
		}(tid)
	}
	wg.Wait()
	require.NoError(t, dw.Close())

	raw, err := os.ReadFile(path)
	require.NoError(t, err)

	totalRecords := numThreads * pushesPerThread
	require.Equal(t, totalRecords*recordSize, len(raw))

	// Parse fixed-size records and verify thread-local ordering.
	lastSeq := make(map[uint32]int)
	for i := 0; i < totalRecords; i++ {
		rec := raw[i*recordSize : (i+1)*recordSize]
		tid := binary.LittleEndian.Uint32(rec[0:4])
		seq := int(binary.LittleEndian.Uint32(rec[4:8]))
		if prev, ok := lastSeq[tid]; ok {
			assert.Greater(t, seq, prev,
				"thread %d: seq %d should be > %d", tid, seq, prev)
		}
		lastSeq[tid] = seq
	}
	assert.Equal(t, numThreads, len(lastSeq))
}
