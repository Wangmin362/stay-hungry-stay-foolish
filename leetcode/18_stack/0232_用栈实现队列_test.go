package _1_array

// 题目：https://leetcode.cn/problems/implement-queue-using-stacks/description/

type MyQueue struct {
	arr []int
}

func Constructor() MyQueue {
	return MyQueue{arr: make([]int, 0)}
}

func (this *MyQueue) Push(x int) {
	this.arr = append(this.arr, x)
}

func (this *MyQueue) Pop() int {
	if len(this.arr) > 0 {
		return 0
	}
	num := this.arr[0]
	this.arr = this.arr[1:]

	return num
}

func (this *MyQueue) Peek() int {
	if len(this.arr) > 0 {
		return 0
	}

	return this.arr[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.arr) == 0
}
