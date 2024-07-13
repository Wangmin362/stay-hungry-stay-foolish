package _1_array

import "testing"

type node struct {
	val  int
	next *node
}

type MyLinkedList struct {
	head   *node
	length int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.length || index < 0 {
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

func (this *MyLinkedList) AddAtHead(val int) {
	this.head = &node{val: val, next: this.head}
	this.length++
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.length == 0 {
		this.head = &node{val: val}
		this.length++
		return
	}

	curr := this.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &node{val: val}
	this.length++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.length || index < 0 {
		return
	}
	if index == 0 {
		this.head = &node{val: val, next: this.head}
		this.length++
		return
	}

	idx := 0
	curr := this.head
	for idx < index-1 {
		idx++
		curr = curr.next
	}
	curr.next = &node{val: val, next: curr.next}
	this.length++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.length || index < 0 {
		return
	}
	if index == 0 {
		this.head = this.head.next
		this.length--
		return
	}

	idx := 0
	curr := this.head
	for idx < index-1 {
		curr = curr.next
		idx++
	}
	curr.next = curr.next.next
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
func TestSingleNode(t *testing.T) {
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
