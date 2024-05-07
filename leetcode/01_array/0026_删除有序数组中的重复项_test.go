package _1_array

import (
	"testing"
)

// 解法一：O(N)时间复杂度，O(N)的空间复杂度，额外开辟一个数组，然后把不重复的元素放入到数组，最后在把临时数组拷贝到原始数组
// 解法二：O(N^2)的时间复杂度，直接把数据往前移动

func removeDuplicates01(nums []int) int {
	tmp := make([]int, 0, len(nums))
	for idx, num := range nums {
		if idx == 0 {
			tmp = append(tmp, num)
		} else {
			if num != nums[idx-1] {
				tmp = append(tmp, num)
			}
		}
	}

	for idx, tmpNum := range tmp {
		nums[idx] = tmpNum
	}

	return len(tmp)
}

func removeDuplicates02(nums []int) int {
	length := len(nums)

	if length <= 1 {
		return length
	}

	for idx := 1; idx < length; {
		if nums[idx] == nums[idx-1] {
			// 当前元素的后续所有元素往前移动一格
			// TODO 这里可以优化，找到第一个和当前元素不同的元素，然后该元素的后续所有元素往前移动一格, 优化后为 removeDuplicates03
			for idx := idx; idx+1 < length; idx++ {
				nums[idx] = nums[idx+1]
			}

			length-- // 移动完成之后，数组的总长度减一
		} else {
			idx++
		}
	}
	return length
}

func removeDuplicates03(nums []int) int {
	length := len(nums)

	if length <= 1 {
		return length
	}

	for idx := 1; idx < length; {
		if nums[idx] == nums[idx-1] {
			// 找到第一个和当前元素不同的元素，然后该元素的后续所有元素往前移动一格
			target := -1
			// 找到第一个和当前元素不同的元素
			for i := idx + 1; i < length; i++ {
				if nums[i] == nums[idx] {
					length--
				} else {
					target = i
				}
			}
			// TODO 移动元素
			for i := target; i < len(nums); i++ {
				nums[1] = nums[3]
				nums[2] = nums[4]
				nums[4] = nums[5]
				//	.............
			}

		} else {
			idx++
		}
	}
	return length
}

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
		get := removeDuplicates05(test.array)
		if get != len(test.expect) {
			t.Errorf("expect:%v, get:%v", test.expect, test.array)
		}

		for i := 0; i < get; i++ {
			if test.array[i] != test.expect[i] {
				t.Errorf("expect:%v, get:%v", test.expect, test.array)
			}
		}
	}
}
