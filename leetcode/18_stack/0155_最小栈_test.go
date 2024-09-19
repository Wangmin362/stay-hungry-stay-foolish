package _1_array

// https://leetcode.cn/problems/min-stack/description/?envType=study-plan-v2&envId=top-interview-150

// 使用一个最小栈保留当前最小的数字
type MinStack struct {
	stack    []int
	minstack []int
}

func Constructor0155() MinStack {
	return MinStack{stack: make([]int, 0, 64), minstack: make([]int, 0, 64)}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.minstack) == 0 {
		this.minstack = append(this.minstack, val)
		return
	}

	last := this.minstack[len(this.minstack)-1]
	if val <= last {
		this.minstack = append(this.minstack, val)
	}
}

func (this *MinStack) Pop() {
	x := this.stack[len(this.stack)-1]
	mx := this.minstack[len(this.minstack)-1]
	if x == mx {
		this.minstack = this.minstack[:len(this.minstack)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	x := this.stack[len(this.stack)-1]
	return x
}

func (this *MinStack) GetMin() int {
	return this.minstack[len(this.minstack)-1]
}
