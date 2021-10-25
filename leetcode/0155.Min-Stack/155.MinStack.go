package _155_Min_Stack

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{minStack: []int{1 << 32}}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if top := this.GetMin(); val < top {
		this.minStack = append(this.minStack, val)
	} else {
		this.minStack = append(this.minStack, top)
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
