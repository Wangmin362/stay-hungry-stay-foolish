package _1_array

import "testing"

// https://leetcode.cn/problems/lru-cache/description/?envType=study-plan-v2&envId=top-interview-150

// map + 双向链表
type LRUCache struct {
	head     *BiNode
	tail     *BiNode
	capacity int
	length   int
	cache    map[int]*BiNode // 缓存
}

type BiNode struct {
	prev *BiNode
	key  int
	val  int
	next *BiNode
}

func Constructor146(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		length:   0,
		cache:    make(map[int]*BiNode),
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	if node == this.head {
		return node.val
	}
	if node == this.tail {
		// 移动尾指针
		this.tail = node.prev
		this.tail.next = nil
		// 移动头指针
		node.next = this.head
		this.head.prev = node
		node.prev = nil
		this.head = node
		return node.val
	}

	node.prev.next = node.next
	node.next.prev = node.prev
	node.next = this.head
	this.head.prev = node
	node.prev = nil
	this.head = node
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if ok { // 存在的话直接更改值
		node.val = value
		// 移动节点到头节点
		if node == this.head {
			return
		}
		if node == this.tail {
			// 移动尾指针
			this.tail = node.prev
			this.tail.next = nil
			// 移动头指针
			node.next = this.head
			this.head.prev = node
			node.prev = nil
			this.head = node
			return
		}

		node.prev.next = node.next
		node.next.prev = node.prev
		node.next = this.head
		this.head.prev = node
		node.prev = nil
		this.head = node

		return
	}

	no := &BiNode{key: key, val: value, next: this.head}
	if this.head == nil {
		this.head = no
		this.tail = no
	} else {
		this.head.prev = no
		this.head = no
	}
	this.cache[key] = no

	if this.length >= this.capacity {
		delete(this.cache, this.tail.key)
		nt := this.tail.prev
		this.tail.prev = nil
		nt.next = nil
		this.tail = nt
	} else {
		this.length++
	}
}

func TestLRU0146(t *testing.T) {
	lru := Constructor146(1)
	lru.Put(1, 1)
	lru.Put(2, 2)
	get := lru.Get(1)
	if get != -1 {
		t.Fatalf("want:%v, get:%v", -1, get)
	}
	get = lru.Get(2)
	if get != 2 {
		t.Fatalf("want:%v, get:%v", 2, get)
	}

	lru.Put(2, 5)
	get = lru.Get(2)
	if get != 5 {
		t.Fatalf("want:%v, get:%v", 5, get)
	}

	lru = Constructor146(2)
	lru.Put(2, 1)
	lru.Put(1, 1)
	lru.Put(2, 3)
	lru.Put(4, 1)
	get = lru.Get(1)
	if get != -1 {
		t.Fatalf("want:%v, get:%v", -1, get)
	}

	get = lru.Get(2)
	if get != 3 {
		t.Fatalf("want:%v, get:%v", 3, get)
	}

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
