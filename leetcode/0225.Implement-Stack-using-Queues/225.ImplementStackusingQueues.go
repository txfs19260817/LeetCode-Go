package leetcode

type MyStack struct {
	q []int
}

// Constructor initialize your data structure here.
func Constructor() MyStack {
	return MyStack{make([]int, 0, 100)}
}

// Push element x onto stack.
func (this *MyStack) Push(x int) {
	n := len(this.q)
	this.q = append(this.q, x)
	for i := 0; i < n; i++ {
		this.q = append(this.q, this.q[0])
		this.q = this.q[1:]
	}
}

// Pop removes the element on top of the stack and returns that element.
func (this *MyStack) Pop() int {
	val := this.q[0]
	this.q = this.q[1:]
	return val
}

// Top get the top element.
func (this *MyStack) Top() int {
	return this.q[0]
}

// Empty returns whether the stack is empty.
func (this *MyStack) Empty() bool {
	return len(this.q) == 0
}
