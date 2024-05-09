package _9_binary_search

import (
	"testing"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////下面采用左闭右闭区间解决此问题//////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// 首先采用闭区间解决这个问题，即循环不变量为[left, right]， mid=(right+left)/2, 若nums[mid] = target，直接返回，因为就是这个位置。
// 若nums[mid] > target，那么说明目标插入值应该在mid的左边，所以right = mid-1
// 若nums[mid] < target，那么说明目标插入值应该在Mid的右边, 所以left = mid+1
// 一下使用下面的例子，举例推出循环之后应该怎么返回值

// nums=[1,3,5,6] target=2
// left=0,right=3 => mid=1  nums[1]=3 > target => left=0, right=mid-1=0
// left=0,right=0 => mid=0  nums[0]=1 < target => left=1, right=0  退出循环，返回left

// nums=[1,3,5,6] target=7
// left=0,right=3 =>mid=1  nums[1]=3 < target => left=mid+1=2, right=3
// left=2,right=3 =>mid=2  nums[2]=5 < target => left=mid+1=3, right=3
// left=3,right=3 =>mid=3  nums[3]=6 < target => left=mid+1=4, right=3 退出循环，返回left

// nums=[1,3,5,6] target=-2
// left=0,right=3 =>mid=1  nums[1]=3 > target => left=0, right=mid-1=0
// left=0,right=0 =>mid=0  nums[0]=1 > target => left=0, right=mid-1=-1 退出循环，返回left

func searchInsertAllClose(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] > target { // 中位数大于target，因此直接选取左边区间
			right = mid - 1
		} else if nums[mid] < target { // 中位数小于target，因此选择右边区间
			left = mid + 1
		} else {
			return mid
		}
	}

	return left
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////下面采用左闭右开区间解决此问题//////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// 采用左闭右开解决这个问题，即循环不变量为[left, right)， mid=(right+left)/2, 若nums[mid] = target，直接返回，因为就是这个位置。
// 若nums[mid] > target，那么说明目标插入值应该在mid的左边，所以right = mid
// 若nums[mid] < target，那么说明目标插入值应该在Mid的右边, 所以left = mid+1
// 一下使用下面的例子，举例推出循环之后应该怎么返回值

// nums=[1,3,5,6) target=2
// left=0,right=4 => mid=2  nums[2]=5 > target => left=0, right=mid=2
// left=0,right=2 => mid=1  nums[1]=3 > target => left=0, right=mid=1
// left=0,right=1 => mid=0  nums[0]=1 < target => left=mid+1=1, right=0 退出循环，返回left

// nums=[1,3,5,6] target=7
// left=0,right=4 =>mid=2  nums[2]=5 < target => left=mid+1=3, right=4
// left=3,right=4 =>mid=3  nums[3]=6 < target => left=mid+1=4, right=3 退出循环，返回left

// nums=[1,3,5,6] target=-2
// left=0,right=4 =>mid=2  nums[2]=5 > target => left=0, right=mid=2
// left=0,right=2 =>mid=1  nums[1]=3 > target => left=0, right=mid=1
// left=0,right=1 =>mid=0  nums[0]=1 > target => left=0, right=mid=0 退出循环，返回left

func searchInsertRightClose(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] > target { // 中位数大于target，因此直接选取左边区间
			right = mid
		} else if nums[mid] < target { // 中位数小于target，因此选择右边区间
			left = mid + 1
		} else {
			return mid
		}
	}

	return left
}

func TestSearch(t *testing.T) {
	var twoSumTest = []struct {
		array  []int
		target int
		expect int
	}{
		{array: []int{1, 3, 5, 6}, target: 5, expect: 2},
		{array: []int{1, 3, 5, 6}, target: 2, expect: 1},
		{array: []int{1, 3, 5, 6}, target: 1, expect: 0},
		{array: []int{1, 3, 5, 6}, target: 0, expect: 0},
		{array: []int{1, 3, 5, 6}, target: 7, expect: 4},
		{array: []int{1, 3, 5, 10}, target: 7, expect: 3},
	}

	for _, test := range twoSumTest {
		get := searchInsertRightClose(test.array, test.target)
		if test.expect != get {
			t.Fatalf("arr:%v, target:%v, expect:%v, get:%v", test.array, test.target, test.expect, get)
		}
	}
}
