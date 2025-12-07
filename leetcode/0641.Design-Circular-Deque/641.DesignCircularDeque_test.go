package leetcode

import (
	"testing"
)

func TestDesignCircularDeque(t *testing.T) {
	// Example 1
	k := 3
	obj := Constructor(k)

	// insertLast(1) // return true
	if got := obj.InsertLast(1); got != true {
		t.Errorf("InsertLast(1) = %v, want %v", got, true)
	}
	// insertLast(2) // return true
	if got := obj.InsertLast(2); got != true {
		t.Errorf("InsertLast(2) = %v, want %v", got, true)
	}
	// insertFront(3) // return true
	if got := obj.InsertFront(3); got != true {
		t.Errorf("InsertFront(3) = %v, want %v", got, true)
	}
	// insertFront(4) // return false, the queue is full
	if got := obj.InsertFront(4); got != false {
		t.Errorf("InsertFront(4) = %v, want %v", got, false)
	}
	// getRear() // return 2
	if got := obj.GetRear(); got != 2 {
		t.Errorf("GetRear() = %v, want %v", got, 2)
	}
	// isFull() // return true
	if got := obj.IsFull(); got != true {
		t.Errorf("IsFull() = %v, want %v", got, true)
	}
	// deleteLast() // return true
	if got := obj.DeleteLast(); got != true {
		t.Errorf("DeleteLast() = %v, want %v", got, true)
	}
	// insertFront(4) // return true
	if got := obj.InsertFront(4); got != true {
		t.Errorf("InsertFront(4) = %v, want %v", got, true)
	}
	// getFront() // return 4
	if got := obj.GetFront(); got != 4 {
		t.Errorf("GetFront() = %v, want %v", got, 4)
	}

	// Additional Tests
	// Empty checks
	objEmpty := Constructor(2)
	if !objEmpty.IsEmpty() {
		t.Error("New deque should be empty")
	}
	if objEmpty.DeleteFront() {
		t.Error("DeleteFront on empty deque should return false")
	}
	if objEmpty.DeleteLast() {
		t.Error("DeleteLast on empty deque should return false")
	}
	if objEmpty.GetFront() != -1 {
		t.Error("GetFront on empty deque should return -1")
	}
	if objEmpty.GetRear() != -1 {
		t.Error("GetRear on empty deque should return -1")
	}

	// Operations
	objEmpty.InsertFront(1) // [1]
	objEmpty.InsertLast(2)  // [1, 2]
	if !objEmpty.IsFull() {
		t.Error("Deque should be full")
	}
	if objEmpty.InsertFront(3) { // Fail
		t.Error("InsertFront on full deque should return false")
	}
	if objEmpty.InsertLast(3) { // Fail
		t.Error("InsertLast on full deque should return false")
	}
	objEmpty.DeleteFront() // [2]
	if objEmpty.GetFront() != 2 {
		t.Errorf("GetFront = %v, want %v", objEmpty.GetFront(), 2)
	}
	objEmpty.DeleteLast() // []
	if !objEmpty.IsEmpty() {
		t.Error("Deque should be empty after deleting all elements")
	}
}
