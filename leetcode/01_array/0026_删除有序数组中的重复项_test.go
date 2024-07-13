package _1_array

import (
	"testing"
)

// 快慢指针  TODO 感觉这种接替思路还是比较难想

// 解题思路：使用双指针解决，刚开始fast, slow分别指向第二个元素，第一个元素不管怎样，肯定都是不用变动的。因此从第二个元素开始。
//	if nums[fast] == nums[slow]{
//      fast++
//  }else{
// 		nums[slow] = nums[fast]
//		slow++
//		fast++
// }

func removeDuplicates04(nums []int) int { // 快慢指针解法一
	if len(nums) <= 0 {
		return len(nums)
	}

	low := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[low] != nums[fast] {
			low++
			nums[low] = nums[fast]
		}
	}
	return low + 1
}

// 0    0    1    1    1    2    2    3    3    4
//     s f 第一步：初始化快慢指针为第二个元素

// 0    0    1    1    1    2    2    3    3    4
//      s    f  第二步：对比nums[fast] 和nums[fast-1]，发现相等，那么fast++,slow不变

// 0    1    1    1    1    2    2    3    3    4
//     		 s    f  第三步：对比nums[fast]和nums[fast-1]，发现不等，那么直接把nums[slow] = nums[fast], slow++, fast++

// 0    1    1    1    1    2    2    3    3    4
//     		 s         f  第四步：对比nums[fast] 和nums[fast-1]，发现相等，那么fast++,slow不变

// 0    1    1    1    1    2    2    3    3    4
//     		 s              f  第五步：对比nums[fast] 和nums[fast-1]，发现相等，那么fast++,slow不变

// 0    1    2    1    1    2    2    3    3    4
//     		      s              f  第六步：对比nums[fast]和nums[fast-1]，发现不等，那么直接把nums[slow] = nums[fast], slow++, fast++

// 0    1    2    1    1    2    2    3    3    4
//     		      s                   f  第七步：对比nums[fast] 和nums[fast-1]，发现相等，那么fast++,slow不变

// 0    1    2    3    1    2    2    3    3    4
//     		           s                   f  第八步：对比nums[fast]和nums[fast-1]，发现不等，那么直接把nums[slow] = nums[fast], slow++, fast++

// 0    1    2    3    1    2    2    3    3    4
//     		           s                        f  第九步：对比nums[fast] 和nums[fast-1]，发现相等，那么fast++,slow不变

// 0    1    2    3    4    2    2    3    3    4
//     		                s                        f  第九步：对比nums[fast]和nums[fast-1]，发现不等，那么直接把nums[slow] = nums[fast], slow++, fast++

// 此时fast指针越界，因此直接推出循环

func removeDuplicates05(nums []int) int { // 快慢指针解法二
	if len(nums) <= 1 {
		return len(nums)
	}

	low := 1 // 相当于下一个不同元素的坑位
	for fast := 1; fast < len(nums); {
		if nums[fast] != nums[fast-1] {
			nums[low] = nums[fast]
			low++
			fast++
		} else {
			fast++
		}
	}
	return low
}

func removeDuplicates06(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	slow := 0
	fast := 1 // 数组的第一个位置肯定不需要动，因此直接把快指针指向第二个元素
	for fast < len(nums) {
		if nums[fast] != nums[fast-1] { // 如果当前元素和前一个元素相同
			slow++
			nums[slow] = nums[fast]
		}

		fast++
	}

	return slow + 1
}

func removeDuplicates07(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	slow, fast := 0, 1
	for fast < len(nums) && slow < len(nums) {
		if nums[fast] != nums[slow] { // 只有当不相等的时候，才交换slow和fast的位置
			slow++
			nums[slow], nums[fast] = nums[fast], nums[slow]
		}
		fast++
	}
	return slow + 1
}

func TestRemoveDuplicates(t *testing.T) {
	var twoSumTest = []struct {
		array  []int
		expect []int
	}{
		{array: []int{3, 3, 3}, expect: []int{3}},
		{array: []int{3, 3, 3, 4}, expect: []int{3, 4}},
		{array: []int{3, 3, 3, 4, 5, 6}, expect: []int{3, 4, 5, 6}},
		{array: []int{}, expect: []int{}},
		{array: []int{3}, expect: []int{3}},
		{array: []int{3, 3, 3, 3, 3}, expect: []int{3}},
		{array: []int{3, 2, 3}, expect: []int{3, 2, 3}},
		{array: []int{3, 2, 2, 3}, expect: []int{3, 2, 3}},
		{array: []int{3, 2, 2, 2, 2, 2, 2, 3}, expect: []int{3, 2, 3}},
		{array: []int{3, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3}, expect: []int{3, 2, 3}},
		{array: []int{0, 1, 2, 2, 3, 0, 4, 2}, expect: []int{0, 1, 2, 3, 0, 4, 2}},
	}

	for _, test := range twoSumTest {
		get := removeDuplicates07(test.array)
		if get != len(test.expect) {
			t.Errorf("expect:%v, get:%v", len(test.expect), get)
		}

		for i := 0; i < get; i++ {
			if test.array[i] != test.expect[i] {
				t.Errorf("expect:%v, get:%v", test.expect, test.array)
			}
		}
	}
}
