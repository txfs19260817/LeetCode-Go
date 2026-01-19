package leetcode

import "container/list"

type FirstUnique struct {
	v2node map[int]*list.Element // 记录元素在队列中的位置
	dll    *list.List            // 队列存储元素
	unique map[int]bool          // 记录数字是否唯一
}

func Constructor(nums []int) FirstUnique {
	fu := FirstUnique{
		v2node: make(map[int]*list.Element),
		dll:    list.New(),
		unique: make(map[int]bool),
	}
	for _, num := range nums {
		fu.Add(num)
	}
	return fu
}

func (this *FirstUnique) ShowFirstUnique() int {
	if this.dll.Len() > 0 {
		return this.dll.Front().Value.(int) // 返回队列的第一个值
	}
	return -1
}

func (this *FirstUnique) Add(value int) {
	if _, ok := this.unique[value]; !ok { // 情况 1:该值未在数据结构中，应该添加
		this.unique[value] = true
		element := this.dll.PushBack(value)
		this.v2node[value] = element
	} else if this.unique[value] { // 情况 2:该值已经出现过一次，现在变成非唯一，应该从队列中移除
		this.unique[value] = false
		this.dll.Remove(this.v2node[value])
		delete(this.v2node, value)
	}
}
