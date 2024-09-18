package _1_array

import "math/rand"

type RandomizedSet struct {
	m   map[int]int // 用于O(1)的插入和移除
	arr []int       // 用于返回随机数据
}

func Constructor() RandomizedSet {
	return RandomizedSet{m: make(map[int]int), arr: make([]int, 0, 64)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}

	this.m[val] = len(this.arr)
	this.arr = append(this.arr, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.m[val]
	if !ok {
		return false
	}

	last := len(this.arr) - 1
	this.arr[idx] = this.arr[last] // 把最后一个元素移动到要删除的位置，然后删除最后一个元素
	this.m[this.arr[last]] = idx   // 更新索引
	// 删除元素
	delete(this.m, val)
	this.arr = this.arr[:last]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	idx := rand.Intn(len(this.arr))
	return this.arr[idx]
}
