package _1_array

import (
	"reflect"
	"slices"
	"sort"
	"testing"
)

// 目前这种写法效率并不是最高的，时间复杂度为O(N^2)
// 分析：主要原因是因为给定数组是无需的，因此只要是有序数组，那么可以增加速度为O(N * M)
func twoSum01(nums []int, target int) []int {
	for idx := range nums {
		myTarget := target - nums[idx]
		for inIdx := range nums {
			if inIdx == idx {
				continue
			}
			if nums[inIdx] == myTarget {
				return []int{idx, inIdx}
			}
		}
	}
	return []int{-1, -1}
}

type myNum struct {
	num int
	idx int
}
type myNums []myNum

func (x myNums) Len() int           { return len(x) }
func (x myNums) Less(i, j int) bool { return x[i].num < x[j].num }
func (x myNums) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// 效率有所提高，但是并非是最高的，时间复杂度依然是O(N*M)
func twoSum02(nums []int, target int) []int {
	transNum := make(myNums, len(nums))
	for idx, num := range nums {
		transNum[idx] = myNum{num: num, idx: idx}
	}
	sort.Sort(transNum)
	for idx := range transNum {
		myTarget := target - transNum[idx].num
		if idx == len(transNum)-1 { // 遍历到最后一个元素时，必须要退出，否则会数组越界
			continue
		}

		newNums := transNum[idx+1:]
		for inIdx := range newNums {
			if newNums[inIdx].num == myTarget {
				return []int{transNum[idx].idx, newNums[inIdx].idx}
			}
		}
	}
	return []int{-1, -1}
}

// 先排序，由于排序之后数组是从小到大的，因此可以使用两个指针算法，此时可以把事件复杂度降为O(nlogn)
func twoSum03(nums []int, target int) []int {
	transNum := make(myNums, len(nums))
	for idx, num := range nums {
		transNum[idx] = myNum{num: num, idx: idx}
	}
	sort.Sort(transNum)
	left := 0
	right := len(nums) - 1

	for {
		if left >= right {
			break
		}
		if transNum[left].num+transNum[right].num == target {
			return []int{transNum[left].idx, transNum[right].idx}
		}
		if transNum[left].num+transNum[right].num > target {
			right--
		}
		if transNum[left].num+transNum[right].num < target {
			left++
		}
	}

	return []int{-1, -1}
}

// 直接使用map，可以降为O(N)
// TODO 为啥这种不用考虑有重复元素的时候？
func twoSum04(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

// 数组排序
func twoSum05(nums []int, target int) []int {
	type wrap struct {
		num int
		idx int
	}
	tmp := make([]wrap, len(nums))
	for idx := range nums {
		tmp[idx] = wrap{num: nums[idx], idx: idx}
	}
	slices.SortFunc(tmp, func(a, b wrap) int {
		if a.num == b.num {
			return 0
		} else if a.num > b.num {
			return 1
		} else {
			return -1
		}
	})

	left, right := 0, len(nums)-1
	for left < right {
		sum := tmp[left].num + tmp[right].num
		if sum == target {
			return []int{tmp[left].idx, tmp[right].idx}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}

	return []int{-1, -1}
}

// map
func twoSum06(nums []int, target int) []int {
	cache := make(map[int]int, len(nums))
	for idx, num := range nums {
		rridx, ok := cache[target-num]
		if ok {
			return []int{rridx, idx}
		} else {
			cache[num] = idx
		}
	}

	return []int{-1, -1}
}

func TestTwoSum(t *testing.T) {
	var twoSumTest = []struct {
		array  []int
		target int
		expect []int
	}{
		{array: []int{3, 2, 4}, target: 6, expect: []int{1, 2}},
		{array: []int{2, 7, 11, 15}, target: 9, expect: []int{0, 1}},
		{array: []int{3, 2, 4}, target: 6, expect: []int{1, 2}},
		{array: []int{3, 3}, target: 6, expect: []int{0, 1}},
		{array: []int{3, 2, 3}, target: 6, expect: []int{0, 2}},
		{array: []int{2, 2}, target: 4, expect: []int{0, 1}},
	}

	for _, test := range twoSumTest {
		sum01 := twoSum06(test.array, test.target)
		if !reflect.DeepEqual(sum01, test.expect) {
			t.Errorf("arr:%v, target:%v, expect:%v, get:%v", test.array, test.target, test.expect, sum01)
		}
	}
}
