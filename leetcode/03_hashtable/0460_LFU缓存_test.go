package _0_basic

import (
	"container/list"
	"testing"
)

type entry struct {
	key, value, freq int
}

type LFUCache struct {
	minFreq    int
	capacity   int
	keyToEnt   map[int]*list.Element // 以O(1)的时间找到当前节点
	freqToList map[int]*list.List    // 以O(1)的时间找到当前节点所在列表
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity:   capacity,
		keyToEnt:   make(map[int]*list.Element),
		freqToList: make(map[int]*list.List),
	}
}

func (this *LFUCache) Get(key int) int {
	no, ok := this.keyToEnt[key]
	if !ok {
		return -1
	}

	ent := no.Value.(*entry)
	li := this.freqToList[ent.freq] // 找到以前的旧链表
	li.Remove(no)                   // 从当前列表中移除
	if li.Len() == 0 {
		delete(this.freqToList, ent.freq) // 说明这个频率的链表已经被取完了，此时需要删除链表
		if ent.freq == this.minFreq {     // 说明把最小频率的链表干完了，此时需要维护最小频率
			this.minFreq++
		}
	}

	// 把当前节点放入右边链表的最前面
	ent.freq++
	li, ok = this.freqToList[ent.freq]
	if !ok {
		li = list.New()
		this.freqToList[ent.freq] = li
	}

	// TODO 这里必须要注意，从新放入链表的时候Element发生了变化，此时需要更新keyToEnt这个Map才行
	this.keyToEnt[ent.key] = li.PushFront(ent)
	return ent.value
}

func (this *LFUCache) Put(key int, value int) {
	no, ok := this.keyToEnt[key]
	if !ok { // 说明是新的节点, 放入到频率为1的链表头部即可
		// 如果相等，那么当前节点加入进来之后，肯定会超过容量，因此直接先移除频率最小的节点
		if len(this.keyToEnt) == this.capacity {
			li := this.freqToList[this.minFreq]
			ele := li.Remove(li.Back())
			delete(this.keyToEnt, ele.(*entry).key)
			if li.Len() == 0 { // 说明当前链表节点已经为空，直接移除链表
				delete(this.freqToList, this.minFreq)
			}
		}

		ent := &entry{key: key, value: value, freq: 1}
		this.minFreq = ent.freq // 如果右新插入的节点，那么最小的频率一定是1
		li, ok := this.freqToList[ent.freq]
		if !ok {
			li = list.New()
			this.freqToList[ent.freq] = li
		}
		// 链表和map同时保存
		this.keyToEnt[key] = li.PushFront(ent)
		return
	}

	ent := no.Value.(*entry)
	ent.value = value
	this.Get(ent.key) // 直接复用get逻辑
}

func TestLFU460(t *testing.T) {
	lfu := Constructor(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	get := lfu.Get(1)
	if get != 1 {
		t.Fatalf("want:%v, get:%v", 1, get)
	}
	lfu.Put(3, 3)

	get = lfu.Get(2)
	if get != -1 {
		t.Fatalf("want:%v, get:%v", -1, get)
	}
	get = lfu.Get(3)
	if get != 3 {
		t.Fatalf("want:%v, get:%v", 3, get)
	}
	lfu.Put(4, 4)

	get = lfu.Get(1)
	if get != -1 {
		t.Fatalf("want:%v, get:%v", -1, get)
	}

	get = lfu.Get(3)
	if get != 3 {
		t.Fatalf("want:%v, get:%v", 3, get)
	}

	get = lfu.Get(4)
	if get != 4 {
		t.Fatalf("want:%v, get:%v", 4, get)
	}

}
