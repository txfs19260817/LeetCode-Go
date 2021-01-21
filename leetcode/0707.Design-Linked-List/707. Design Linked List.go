package _707_Design_Linked_List

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{Val: -999, Next: nil}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	p := this
	for i := 0; i < index; i++ {
		p = p.Next
		if p == nil {
			return -1
		}
	}
	return p.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	if this.Val == -999 {
		this.Val = val
		return
	}
	temp := &MyLinkedList{this.Val, this.Next}
	this.Val = val
	this.Next = temp
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.Val == -999 {
		this.Val = val
		return
	}
	p := this
	for p.Next != nil {
		p = p.Next
	}
	p.Next = &MyLinkedList{val, nil}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	pre, q := this, this
	i := 0
	for ; i < index; i++ {
		pre = q
		q = q.Next
		if q == nil {
			if i+1 == index {
				this.AddAtTail(val)
			}
			return
		}
	}
	temp := &MyLinkedList{val, q}
	pre.Next = temp
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 && this.Next != nil {
		this.Val = this.Next.Val
		this.Next = this.Next.Next
		return
	}
	pre, q := this, this
	for i := 0; i < index; i++ {
		pre = q
		q = q.Next
		if q == nil {
			return
		}
	}
	pre.Next = q.Next
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
