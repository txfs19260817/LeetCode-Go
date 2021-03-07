package _460_LFU_Cache

import "container/list"

type LFUCache struct {
	cap, minFreq int
	nodes        map[int]*list.Element // key2node
	lists        map[int]*list.List    // freq2list
}

type node struct {
	key, value, freq int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:     capacity,
		minFreq: 0,
		nodes:   make(map[int]*list.Element),
		lists:   make(map[int]*list.List),
	}
}

func (this *LFUCache) Get(key int) int {
	// get value
	v, ok := this.nodes[key]
	// not found
	if !ok {
		return -1
	}
	// remove node from certain freq list
	curNode := v.Value.(*node)
	this.lists[curNode.freq].Remove(v)
	// update freq
	curNode.freq++
	// if certain freq list does not exist, create
	if _, ok := this.lists[curNode.freq]; !ok {
		this.lists[curNode.freq] = list.New()
	}
	// push the node in the front of the list
	listNode := this.lists[curNode.freq].PushFront(curNode)
	// store updated node
	this.nodes[key] = listNode
	// update min freq
	if curNode.freq-1 == this.minFreq && this.lists[curNode.freq-1].Len() == 0 {
		this.minFreq++
	}
	return curNode.value
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}
	if _, ok := this.nodes[key]; ok {
		this.nodes[key].Value.(*node).value = value
		this.Get(key) // update freq
		return
	}
	// delete if full
	if this.cap == len(this.nodes) {
		obsoleteNode := this.lists[this.minFreq].Back()
		delete(this.nodes, obsoleteNode.Value.(*node).key)
		this.lists[this.minFreq].Remove(obsoleteNode)
	}
	this.minFreq = 1 // set minFreq = 1 due to new node insertion
	curNode := &node{
		key:   key,
		value: value,
		freq:  1,
	}
	if _, ok := this.lists[1]; !ok {
		this.lists[1] = list.New()
	}
	listNode := this.lists[1].PushFront(curNode)
	this.nodes[key] = listNode
}
