package _1_array

import (
	"container/list"
	"slices"
)

type MyHashSet struct {
	list *list.List
}

func Constructor705() MyHashSet {
	return MyHashSet{list: list.New()}
}

func (this *MyHashSet) Add(key int) {
	if !this.Contains(key) {
		this.list.PushBack(key)
	}
}

func (this *MyHashSet) Remove(key int) {
	head := this.list.Front()
	for head != nil {
		if head.Value.(int) == key {
			this.list.Remove(head)
			return
		}
		head = head.Next()
	}
}

func (this *MyHashSet) Contains(key int) bool {
	head := this.list.Front()
	for head != nil {
		if head.Value.(int) == key {
			return true
		}
		head = head.Next()
	}
	return false
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////// 数组实现 /////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type MyHashSetII struct {
	set []int
}

func Constructor705II() MyHashSet {
	return MyHashSet{}
}

func (this *MyHashSetII) Add(key int) {
	if !this.Contains(key) {
		this.set = append(this.set, key)
	}
}

func (this *MyHashSetII) Remove(key int) {
	for i := 0; i < len(this.set); i++ {
		if this.set[i] == key {
			this.set = append(this.set[:i], this.set[i+1:]...)
			return
		}
	}
}

func (this *MyHashSetII) Contains(key int) bool {
	return slices.Contains(this.set, key)
}
