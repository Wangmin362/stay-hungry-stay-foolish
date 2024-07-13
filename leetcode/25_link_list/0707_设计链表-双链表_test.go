package _1_array

import "testing"

type dnode struct {
	val  int
	prev *dnode
	next *dnode
}

type dMyLinkedList struct {
	head *dnode
	tail *dnode

	length int
}

func dConstructor() dMyLinkedList {
	return dConstructor()
}

func (this *dMyLinkedList) Get(index int) int {
	if index >= this.length {
		return -1
	}

	idx := 0
	curr := this.head
	for idx < index {
		curr = curr.next
		idx++
	}

	return curr.val
}

func (this *dMyLinkedList) AddAtHead(val int) {
	if this.length == 0 {
		this.head = &dnode{val: val}
		this.tail = this.head
		this.length++
		return
	}

	this.head = &dnode{val: val, next: this.head}
	this.length++
}

func (this *dMyLinkedList) AddAtTail(val int) {
	no := &dnode{val: val, next: nil, prev: this.tail}
	this.tail.next = no
	this.tail = no
	this.length++
}

func (this *dMyLinkedList) AddAtIndex(index int, val int) {
	if index > this.length || index < 0 {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}

	idx := 0
	curr := this.head
	for idx < index-1 { // 遍历到前一个节点
		curr = curr.next
		idx++
	}
	curr.next = &dnode{val: val, next: curr.next, prev: curr}
	this.length++
}

func (this *dMyLinkedList) DeleteAtIndex(index int) {
	if index > this.length || index < 0 {
		return
	}
	if index == 0 {
		this.head = this.head.next
		this.length--
		return
	}

	idx := 0
	curr := this.head
	for idx < index-1 { // 遍历到前一个节点
		curr = curr.next
		idx++
	}
	curr.next = curr.next.next
	curr.next.prev = curr
	this.length--
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
func TestDoubleNode(t *testing.T) {
	obj := Constructor()
	obj.AddAtHead(7)
	obj.AddAtHead(2)
	obj.AddAtHead(1)
	obj.AddAtIndex(3, 0)
	obj.DeleteAtIndex(2)
	obj.AddAtHead(6)
	obj.AddAtTail(4)
	obj.Get(4)
}
