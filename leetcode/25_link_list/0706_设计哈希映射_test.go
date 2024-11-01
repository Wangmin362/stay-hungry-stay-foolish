package _1_array

import "container/list"

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////链表实现////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type pair struct {
	key, value int
}

type MyHashMap struct {
	list *list.List
}

func Constructor0706() MyHashMap {
	return MyHashMap{list.New()}
}

func (this *MyHashMap) Put(key int, value int) {
	ele := this.get(key)
	if ele != nil {
		ele.Value.(*pair).value = value
		return
	}

	this.list.PushBack(&pair{key: key, value: value})
}

func (this *MyHashMap) get(key int) *list.Element {
	head := this.list.Front()
	for head != nil {
		if head.Value.(*pair).key == key {
			return head
		}
		head = head.Next()
	}
	return nil
}

func (this *MyHashMap) Get(key int) int {
	ele := this.get(key)
	if ele == nil {
		return -1
	}

	return ele.Value.(*pair).value
}

func (this *MyHashMap) Remove(key int) {
	ele := this.get(key)
	if ele == nil {
		return
	}

	this.list.Remove(ele)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////数组实现////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type MyHashMapII struct {
	kv []*pair
}

func Constructor0706II() MyHashMapII {
	return MyHashMapII{}
}

func (this *MyHashMapII) Put(key int, value int) {
	for i := 0; i < len(this.kv); i++ {
		if this.kv[i].key == key {
			this.kv[i].value = value
			return
		}
	}

	this.kv = append(this.kv, &pair{key, value})
}

func (this *MyHashMapII) Get(key int) int {
	for i := 0; i < len(this.kv); i++ {
		if this.kv[i].key == key {
			return this.kv[i].value
		}
	}

	return -1
}

func (this *MyHashMapII) Remove(key int) {
	for i := 0; i < len(this.kv); i++ {
		if this.kv[i].key == key {
			this.kv = append(this.kv[:i], this.kv[i+1:]...)
		}
	}
}
