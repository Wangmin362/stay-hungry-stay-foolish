package _1_array

import (
	"container/list"
	"sync"
	"testing"
)

// https://leetcode.cn/problems/lru-cache/description/?envType=study-plan-v2&envId=top-interview-150

type binode struct {
	key  int
	val  int
	prev *binode
	next *binode
}

type LRUCache struct {
	head  *binode
	tail  *binode
	cache map[int]*binode // 保证查询的时间是O(1)的，因此需要维护这个map的状态
	lock  sync.Mutex      // 保证并发安全的

	capacity int
	length   int
}

// 限制容量，LRU的最大容量为capacity，超过容量就需要把队尾元素移除
func Constructor146(capacity int) LRUCache {
	return LRUCache{capacity: capacity, cache: make(map[int]*binode, capacity)}
}

// 每次获取的元素需要把该元素放在对头，因为这是最近访问的
func (this *LRUCache) Get(key int) int {
	this.lock.Lock()
	defer this.lock.Unlock()

	no, ok := this.cache[key]
	if !ok {
		return -1
	}

	// 存在的话需要把这个元素放入到对头

	if no == this.head { // 本来就在对头，不需要操作
		return no.val
	}
	if no == this.tail { // 节点在尾部
		this.tail = no.prev // 移动尾指针
		no.prev.next = nil  // 断开尾部
		no.prev = nil
		no.next = this.head
		this.head.prev = no
		this.head = no
		return no.val
	}

	// 说明节点在中间
	no.prev.next = no.next
	no.next.prev = no.prev
	no.prev = nil
	no.next = this.head
	this.head.prev = no
	this.head = no
	return no.val
}

// 1、如果没有超过容量，那么直接把元素放在对头
// 2、如果放入当前元素之后超过了容量，那么需要移除队尾元素
// 3、如果当前元素就在LRU当中，那么肯定不需要移除元素，但是需要把当前元素放入到对头
func (this *LRUCache) Put(key int, value int) {
	this.lock.Lock()

	if this.length == 0 {
		no := &binode{key: key, val: value}
		this.cache[key] = no
		this.head = no
		this.tail = no
		this.length++
		this.lock.Unlock()
		return
	}

	no, ok := this.cache[key]
	if !ok { // 不存在的话，直接放入对头
		no = &binode{key: key, val: value, next: this.head}
		this.cache[key] = no
		this.head.prev = no
		this.length++
		this.head = no
		if this.length > this.capacity {
			tmp := this.tail
			this.tail = tmp.prev
			this.tail.next = nil
			delete(this.cache, tmp.key) // 移除掉队尾元素
			this.length--
		}
		this.lock.Unlock()
		return
	}
	this.lock.Unlock()

	// 说明当前元素已经存在， 把当前元素放入到对头
	no.val = value
	this.Get(key) // 直接通过这种方式移动到对头
}

type entry struct {
	key, value int
}

type LRUCacheII struct {
	capacity int
	list     *list.List
	mappping map[int]*list.Element
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// 第二种解法，参考灵神，直接使用go内置的双向链表
func Constructor0146II(capacity int) LRUCacheII {
	return LRUCacheII{capacity: capacity, list: list.New(), mappping: make(map[int]*list.Element)}
}

func (this *LRUCacheII) Get(key int) int {
	ele, ok := this.mappping[key]
	if !ok {
		return -1
	}

	this.list.MoveToFront(ele)
	return ele.Value.(*entry).value
}

func (this *LRUCacheII) Put(key int, value int) {
	ele, ok := this.mappping[key]
	if ok {
		ele.Value.(*entry).value = value
		this.list.MoveToFront(ele)
		return
	}

	this.mappping[key] = this.list.PushFront(&entry{key: key, value: value})
	if this.list.Len() > this.capacity {
		ent := this.list.Remove(this.list.Back()).(*entry)
		delete(this.mappping, ent.key)
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

	lru = Constructor146(3)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	lru.Put(4, 4)
	get = lru.Get(4)
	if get != 4 {
		t.Fatalf("want:%v, get:%v", 4, get)
	}
	get = lru.Get(3)
	if get != 3 {
		t.Fatalf("want:%v, get:%v", 3, get)
	}
	get = lru.Get(2)
	if get != 2 {
		t.Fatalf("want:%v, get:%v", 2, get)
	}
	get = lru.Get(1)
	if get != -1 {
		t.Fatalf("want:%v, get:%v", -1, get)
	}
	lru.Put(5, 5)
	get = lru.Get(1)
	if get != -1 {
		t.Fatalf("want:%v, get:%v", -1, get)
	}
	get = lru.Get(2)
	if get != 2 {
		t.Fatalf("want:%v, get:%v", 2, get)
	}
	get = lru.Get(3)
	if get != 3 {
		t.Fatalf("want:%v, get:%v", 3, get)
	}
	get = lru.Get(4)
	if get != -1 {
		t.Fatalf("want:%v, get:%v", -1, get)
	}
	get = lru.Get(5)
	if get != 5 {
		t.Fatalf("want:%v, get:%v", 5, get)
	}

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor146(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
