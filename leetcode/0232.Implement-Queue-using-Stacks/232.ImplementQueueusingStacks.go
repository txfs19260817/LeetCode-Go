package _232_Implement_Queue_using_Stacks

type MyQueue struct {
	a, b []int
}

// Constructor initialize your data structure here.
func Constructor() MyQueue {
	return MyQueue{make([]int, 0, 100), make([]int, 0, 100)}
}

// Push element x to the back of queue.
func (this *MyQueue) Push(x int) {
	this.a = append(this.a, x)
}

// Pop removes the element from in front of queue and returns that element.
func (this *MyQueue) Pop() int {
	val := this.Peek()
	this.b = this.b[:len(this.b)-1]
	return val
}

// Peek gets the front element.
func (this *MyQueue) Peek() int {
	if len(this.b) == 0 {
		for len(this.a) > 0 {
			this.b = append(this.b, this.a[len(this.a)-1])
			this.a = this.a[:len(this.a)-1]
		}
	}
	return this.b[len(this.b)-1]
}

// Empty returns whether the queue is empty.
func (this *MyQueue) Empty() bool {
	return len(this.a) == 0 && len(this.b) == 0
}
