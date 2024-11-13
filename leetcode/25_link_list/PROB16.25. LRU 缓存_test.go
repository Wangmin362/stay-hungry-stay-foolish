package _1_array

import "container/list"

type LRUCache1625 struct {
	list     *list.List
	capacity int
}

type kv struct {
	key, value int
}

func Constructor1625(capacity int) LRUCache1625 {
	return LRUCache1625{capacity: capacity, list: list.New()}
}

func (this *LRUCache1625) get(key int) *list.Element {
	head := this.list.Front()
	for head != nil {
		if head.Value.(*kv).key == key {
			this.list.MoveToFront(head) // 移动到最前面
			return head
		}
		head = head.Next()
	}
	return nil
}

func (this *LRUCache1625) Get(key int) int {
	ele := this.get(key)
	if ele == nil {
		return -1
	}

	return ele.Value.(*kv).value
}

func (this *LRUCache1625) Put(key int, value int) {
	ele := this.get(key)
	if ele != nil {
		ele.Value.(*kv).value = value
		return
	}

	if this.list.Len() == this.capacity {
		this.list.Remove(this.list.Back())
	}

	this.list.PushFront(&kv{key: key, value: value})
}
