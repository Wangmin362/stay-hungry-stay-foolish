package _0_basic

import (
	"fmt"
	"sort"
	"testing"
)

// 做题思路：先确定一个维度，然后再确定另外一个维度。  我们可以先确定身高，在确定个数。或者先确定个数，在确定升高。
// 实际上当我们先确定个数，即先按照个数从小到大排序，然后根据排序结果按照身高来排序，这种思想无法实现。所以我们只能先根据身高排序，
// 身高一样就按照个数从小到大排序，按照这种思想排序之后，身高就是有序的，某个元素之前的元素都是大于这个元素的，然后在按照个数进行
// 排序即可
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})

	// TODO 优化为原地排序
	res := make([][]int, 0, len(people))
	for i := 0; i < len(people); i++ {
		k := people[i][1]
		if len(res) == k { // 如果当前元素前面正好有K个元素大于这个元素，直接吧这个元素放在末位即可，非常合适
			res = append(res, people[i])
		} else {
			cnt := len(res) - k // 否则，说明前面比当前大的人数对于K，此时就需要挪动元素
			res = append(res, people[i])
			var j int
			for j = len(res) - 1; j >= len(res)-cnt; j-- {
				res[j] = res[j-1]
			}
			res[j] = people[i]
		}

	}

	return res
}

// 优化为原地排序
func reconstructQueue02(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})

	for i := 0; i < len(people); i++ {
		k := people[i][1]
		length := i
		if length == k { // 如果当前元素前面正好有K个元素大于这个元素，直接吧这个元素放在末位即可，非常合适
			continue
		}

		// 否则，说明前面比当前大的人数对于K，此时就需要挪动元素
		tmp := people[i]
		for j := i; j > k; j-- {
			people[j] = people[j-1]
		}
		people[k] = tmp
	}

	return people
}

func TestReconstructQueue(t *testing.T) {
	fmt.Println(reconstructQueue02([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}))
}
