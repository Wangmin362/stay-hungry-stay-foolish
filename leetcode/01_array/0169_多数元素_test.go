package _1_array

import "sort"

// https://leetcode.cn/problems/majority-element/description/?envType=study-plan-v2&envId=top-interview-150

// 解法一：map统计  时间复杂度O(n), 空间复杂度O(n)
// 解法二：先排序，在统计， 时间复杂度为O(longn), 空间复杂度O(1)
// 解法三：摩尔投票，其实就是从数组中找出众数

func majorityElement01(nums []int) int {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	res, maxCnt := 0, 0
	for k, cnt := range m {
		if cnt > maxCnt {
			res = k
			maxCnt = cnt
		}
	}

	return res
}

func majorityElement02(nums []int) int {
	sort.Ints(nums)
	half := len(nums) >> 1

	cnt := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			cnt++
		} else {
			cnt = 1
		}
		if cnt > half {
			return nums[i]
		}
	}
	return nums[0]
}

func majorityElement03(nums []int) int {
	votex, x := 0, 0
	for _, num := range nums {
		if votex == 0 { // 说明当前没有投票的人
			x = num
			votex++
			continue
		}

		if num == x {
			votex++
		} else {
			votex--
		}
	}
	return x
}
