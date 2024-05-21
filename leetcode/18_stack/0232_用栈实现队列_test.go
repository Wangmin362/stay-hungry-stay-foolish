package _1_array

import (
	"fmt"
	"testing"
)

// 题目：https://leetcode.cn/problems/implement-queue-using-stacks/description/

type MyQueue struct {
	in  []int
	out []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) Pop() int {
	if len(this.out) == 0 && len(this.in) != 0 {
		for idx := len(this.in) - 1; idx >= 0; idx-- {
			this.out = append(this.out, this.in[idx])
		}
		this.in = []int{}
	}

	idx := len(this.out) - 1
	x := this.out[idx]
	this.out = this.out[:idx]
	return x
}

func (this *MyQueue) Peek() int {
	x := this.Pop()
	this.Push(x)
	return x
}

func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}

func TestMyQueue(t *testing.T) {
	qu := MyQueue{}

	qu.Push(1)
	qu.Push(2)
	fmt.Println(qu.Peek())
	fmt.Println(qu.Pop())
	fmt.Println(qu.Empty())
}
