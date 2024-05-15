package _1_array

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/fruit-into-baskets/description/

// 本质上是双重for循环
func totalFruit01(fruits []int) int {
	res := 0
	slow := 0
	for slow < len(fruits) {
		bucket := map[int]struct{}{} // 用来放水果，最多只允许放两种类型的水果

		fast := slow
		cnt := 0
		for fast < len(fruits) {
			bucket[fruits[fast]] = struct{}{}
			if len(bucket) <= 2 {
				fast++
				cnt++
			} else {
				break
			}
		}
		if res < cnt {
			res = cnt
		}

		slow++ // 移动左边界
	}
	return res
}

// 真正的双指针
// TODO 能否写的更优雅一点
func totalFruit02(fruits []int) int {
	res := 0
	bucket := map[int]int{} // 用来放水果，最多只允许放两种类型的水果

	slow := 0
	fast := 0
	cnt := 0
	for fast < len(fruits) {
		if len(bucket) <= 2 {
			fCnt, ok := bucket[fruits[fast]]
			if ok {
				fCnt++
				bucket[fruits[fast]] = fCnt
			} else {
				bucket[fruits[fast]] = 1
			}

			if len(bucket) > 2 && res < cnt {
				res = cnt
			}

			cnt++
			fast++
		} else {
			sCnt := bucket[fruits[slow]]
			if sCnt == 1 {
				delete(bucket, fruits[slow])
			} else {
				sCnt--
				bucket[fruits[slow]] = sCnt
			}
			slow++ // 移动左边界
			cnt--
		}
	}
	if len(bucket) > 2 {
		cnt--
	}
	if res < cnt {
		res = cnt
	}

	return res
}

func TestTotalFruit(t *testing.T) {
	var testdata = []struct {
		array  []int
		expect int
	}{
		//{array: nil, expect: 0},
		//{array: []int{}, expect: 0},
		//{array: []int{1}, expect: 1},
		//{array: []int{1, 2, 1}, expect: 3},
		//{array: []int{0, 1, 2, 2}, expect: 3},
		//{array: []int{1, 2, 3, 2, 2}, expect: 4},
		//{array: []int{1, 2, 3, 2, 2}, expect: 4},
		//{array: []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}, expect: 5},
		{array: []int{0, 1, 2}, expect: 2},
	}

	for _, test := range testdata {
		get := totalFruit02(test.array)
		if get != test.expect {
			t.Errorf("arr:%v,  expect:%v, get:%v", test.array, test.expect, get)
		}
	}
}
