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

func TestDataWriter_SingleThread(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.dat")
	dw, err := NewDataWriter(path)
	require.NoError(t, err)

	dw.Push([]byte("hello"))
	dw.Push([]byte("world"))
	require.NoError(t, dw.Close())

	records, err := RecoverRecords(path)
	require.NoError(t, err)
	assert.Equal(t, [][]byte{[]byte("hello"), []byte("world")}, records)
}

func TestDataWriter_MultiThread_Ordering(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.dat")
	dw, err := NewDataWriter(path)
	require.NoError(t, err)

	const numThreads = 50
	const pushesPerThread = 100

	var wg sync.WaitGroup
	for tid := 0; tid < numThreads; tid++ {
		wg.Add(1)
		go func(tid int) {
			defer wg.Done()
			for seq := 0; seq < pushesPerThread; seq++ {
				// Encode [tid (4B)][seq (4B)]
				data := make([]byte, 8)
				binary.LittleEndian.PutUint32(data[0:4], uint32(tid))
				binary.LittleEndian.PutUint32(data[4:8], uint32(seq))
				dw.Push(data)
			}
		}(tid)
	}
	wg.Wait()
	require.NoError(t, dw.Close())

	records, err := RecoverRecords(path)
	require.NoError(t, err)
	assert.Equal(t, numThreads*pushesPerThread, len(records))

	// For each thread, seq numbers must be strictly increasing.
	lastSeq := make(map[uint32]int)
	for _, rec := range records {
		tid := binary.LittleEndian.Uint32(rec[0:4])
		seq := int(binary.LittleEndian.Uint32(rec[4:8]))
		if prev, ok := lastSeq[tid]; ok {
			assert.Greater(t, seq, prev,
				"thread %d: seq %d should be > %d", tid, seq, prev)
		}
		lastSeq[tid] = seq
	}
	// Every thread must have contributed.
	assert.Equal(t, numThreads, len(lastSeq))
}

func TestDataWriter_EmptyPush(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.dat")
	dw, err := NewDataWriter(path)
	require.NoError(t, err)

	dw.Push([]byte{})
	dw.Push([]byte("after-empty"))
	require.NoError(t, dw.Close())

	records, err := RecoverRecords(path)
	require.NoError(t, err)
	assert.Equal(t, [][]byte{{}, []byte("after-empty")}, records)
}

func TestRecoverRecords_TruncatedFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.dat")
	dw, err := NewDataWriter(path)
	require.NoError(t, err)

	dw.Push([]byte("record1"))
	dw.Push([]byte("record2"))
	require.NoError(t, dw.Close())

	// Append a partial record to simulate a crash mid-write:
	// length = 5 but only 1 byte of data follows.
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	require.NoError(t, err)
	f.Write([]byte{0x05, 0x00, 0x00, 0x00, 0xAA}) //nolint: errcheck
	f.Close()

	records, err := RecoverRecords(path)
	require.NoError(t, err)
	// Only the 2 fully valid records should be recovered.
	require.Equal(t, 2, len(records))
	assert.Equal(t, []byte("record1"), records[0])
	assert.Equal(t, []byte("record2"), records[1])
}

func TestRecoverRecords_CorruptedCRC(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.dat")
	dw, err := NewDataWriter(path)
	require.NoError(t, err)

	dw.Push([]byte("good"))
	dw.Push([]byte("also-good"))
	require.NoError(t, dw.Close())

	// Corrupt the last 4 bytes (CRC of second record).
	data, err := os.ReadFile(path)
	require.NoError(t, err)
	data[len(data)-1] ^= 0xFF // flip bits in CRC
	require.NoError(t, os.WriteFile(path, data, 0644))

	records, err := RecoverRecords(path)
	require.NoError(t, err)
	// Second record's CRC is bad → only first record survives.
	require.Equal(t, 1, len(records))
	assert.Equal(t, []byte("good"), records[0])
}
