package snowflake

import (
	"math/rand"
	"reflect"
	"sort"
	"sync"
	"testing"
	"time"
)

func TestKVStoreOperations(t *testing.T) {
	tests := []struct {
		name     string
		commands []string
		args     []any
		expected []any
	}{
		{
			name:     "Test Case 1",
			commands: []string{"KVStore", "set", "set", "get", "update", "prefixSearch", "deleteKey", "get", "prefixSearch"},
			args: []any{
				[]string{},
				[]any{"aaple", 3},
				[]any{"aap", 2},
				[]string{"aap"},
				[]any{"aap", 5},
				[]string{"aap"},
				[]string{"aap"},
				[]string{"aap"},
				[]string{"aaple"},
			},
			expected: []any{nil, nil, nil, 2, nil, []int{5, 3}, nil, -1, []int{3}},
		},
		{
			name:     "Test Case 2",
			commands: []string{"KVStore", "set", "set", "set", "prefixSearch", "update", "prefixSearch", "deleteKey", "set", "prefixSearch", "get"},
			args: []any{
				[]string{},
				[]any{"foo", 1},
				[]any{"bar", 2},
				[]any{"foobar", 3},
				[]string{"foo"},
				[]any{"foo", 4},
				[]string{"foo"},
				[]string{"bar"},
				[]any{"baz", 5},
				[]string{"ba"},
				[]string{"bar"},
			},
			expected: []any{nil, nil, nil, nil, []int{1, 3}, nil, []int{4, 3}, nil, nil, []int{5}, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var obj *KVStore
			for i, cmd := range tt.commands {
				arg := tt.args[i]
				want := tt.expected[i]

				switch cmd {
				case "KVStore":
					obj = NewKVStore()
				case "set":
					args := arg.([]any)
					obj.Set(args[0].(string), args[1].(int))
				case "get":
					got := obj.Get(arg.([]string)[0])
					if want != nil && got != want.(int) {
						t.Errorf("step %d: Get(%s) = %d, want %d", i, arg.([]string)[0], got, want)
					}
				case "update":
					args := arg.([]any)
					obj.Update(args[0].(string), args[1].(int))
				case "deleteKey":
					obj.DeleteKey(arg.([]string)[0])
				case "prefixSearch":
					got := obj.PrefixSearch(arg.([]string)[0])
					wantSlice := want.([]int)
					sort.Ints(got)
					sort.Ints(wantSlice)
					if !reflect.DeepEqual(got, wantSlice) {
						t.Errorf("step %d: PrefixSearch(%s) = %v, want %v", i, arg.([]string)[0], got, wantSlice)
					}
				}
			}
		})
	}
}

func TestIndependentTransactions(t *testing.T) {
	store := NewKVStore()

	// Initial state
	store.Set("apple", 1)

	// Start two concurrent transactions
	tx1 := store.Begin()
	tx2 := store.Begin()

	// Tx1 updates apple, adds apricot
	tx1.Set("apple", 10)
	tx1.Set("apricot", 20)

	// Tx2 deletes apple, adds banana
	tx2.DeleteKey("apple")
	tx2.Set("banana", 30)

	// Verify Isolation:
	// Global store should still see apple=1
	if v := store.Get("apple"); v != 1 {
		t.Errorf("Global store modified prematurely: got %d, want 1", v)
	}

	// Tx1 sees its own changes
	if v := tx1.Get("apple"); v != 10 {
		t.Errorf("Tx1 cannot see own write: got %d, want 10", v)
	}
	if v := tx1.Get("apricot"); v != 20 {
		t.Errorf("Tx1 cannot see own write: got %d, want 20", v)
	}
	// Tx1 should NOT see Tx2 changes
	if v := tx1.Get("banana"); v != -1 {
		t.Errorf("Tx1 saw Tx2 write: got %d, want -1", v)
	}

	// Tx2 sees its own changes
	if v := tx2.Get("apple"); v != -1 {
		t.Errorf("Tx2 cannot see own delete: got %d, want -1", v)
	}
	if v := tx2.Get("banana"); v != 30 {
		t.Errorf("Tx2 cannot see own write: got %d, want 30", v)
	}

	// Commit Tx1
	tx1.Commit()

	// Global store now reflects Tx1
	if v := store.Get("apple"); v != 10 {
		t.Errorf("Global after Tx1 commit: got %d, want 10", v)
	}
	if v := store.Get("apricot"); v != 20 {
		t.Errorf("Global after Tx1 commit: got %d, want 20", v)
	}

	// Tx2 should still see its own view (apple deleted)
	if v := tx2.Get("apple"); v != -1 {
		t.Errorf("Tx2 isolation broken after Tx1 commit: got %d, want -1", v)
	}

	// Commit Tx2
	tx2.Commit()
	// Should overwrite Tx1's update to apple (Last Commit Wins)
	if v := store.Get("apple"); v != -1 {
		t.Errorf("Global after Tx2 commit: got %d, want -1", v)
	}
	if v := store.Get("banana"); v != 30 {
		t.Errorf("Global after Tx2 commit: got %d, want 30", v)
	}
	// Apricot from Tx1 should still be there
	if v := store.Get("apricot"); v != 20 {
		t.Errorf("Global lost Tx1 data: got %d, want 20", v)
	}
}

func TestConcurrentTransactions(t *testing.T) {
	store := NewKVStore()
	var wg sync.WaitGroup

	// Run 50 concurrent transactions
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			tx := store.Begin()
			tx.Set("counter", val)
			time.Sleep(time.Millisecond) // Simulate work
			tx.Commit()
		}(i)
	}

	wg.Wait()

	// Just ensure no panic and some value is set
	if val := store.Get("counter"); val < 0 || val >= 50 {
		t.Errorf("Concurrent writes failed or data corrupted: got %d", val)
	}
}

func TestConcurrentAtomicUpdates(t *testing.T) {
	// This test verifies that Commit is atomic.
	// We update two keys "paira" and "pairb" to the same random value in a transaction.
	// Concurrent readers doing PrefixSearch("pair") must ALWAYS see equal values for A and B.
	// They should never see A=Old, B=New or A=New, B=Old.

	store := NewKVStore()
	store.Set("paira", 0)
	store.Set("pairb", 0)

	var wgWriters sync.WaitGroup
	var wgReaders sync.WaitGroup
	start := make(chan struct{})

	// Writers
	for i := 0; i < 20; i++ {
		wgWriters.Add(1)
		go func(id int) {
			defer wgWriters.Done()
			<-start
			for j := 0; j < 50; j++ {
				val := rand.Intn(1000)
				tx := store.Begin()
				tx.Set("paira", val)
				tx.Set("pairb", val)
				// Sleep randomly to increase chance of interleaving
				if rand.Float32() < 0.1 {
					time.Sleep(time.Microsecond)
				}
				tx.Commit()
			}
		}(i)
	}

	// Readers
	stopReaders := make(chan struct{})
	readerErrors := make(chan string, 100)
	for i := 0; i < 10; i++ {
		wgReaders.Add(1)
		go func() {
			defer wgReaders.Done()
			<-start
			for {
				select {
				case <-stopReaders:
					return
				default:
					// Atomic read of both keys
					vals := store.PrefixSearch("pair")
					if len(vals) != 2 {
						// Only possible if keys deleted or not inserted yet (but we init them)
						// Trie order might vary if implementation changes, but here we expect 2
						continue
					}
					if vals[0] != vals[1] {
						readerErrors <- "Read inconsistent values!"
						return
					}
				}
			}
		}()
	}

	close(start) // release the hounds

	wgWriters.Wait()
	close(stopReaders)
	wgReaders.Wait()

	select {
	case err := <-readerErrors:
		t.Fatal(err)
	default:
		// success
	}
}
