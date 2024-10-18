package huawei_interview

import (
	"fmt"
	"sort"
	"testing"
)

//给定一个数组data和数组sortBy，对data数组内的数据按照sortBy数组给定的顺序排序， 如果数据不在sortBy数组内，则按升序排列并放在这些数据之后
// 输入：[3,1,1,3,4,5], sortBy: [3,1]  输出：[3,3,1,1,4,5]

func sortByGroup(nums, sortBy []int) []int {
	m := make(map[int]int)
	var idx, n int
	for idx, n = range sortBy {
		m[n] = idx
	}
	idx++

	tmp := make([][2]int, 0, len(nums))
	for _, num := range nums {
		if score, ok := m[num]; ok {
			tmp = append(tmp, [2]int{num, score})
		} else {
			tmp = append(tmp, [2]int{num, idx})
		}
	}
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i][1] == tmp[j][1] {
			return tmp[i][0] < tmp[j][0]
		}
		return tmp[i][1] < tmp[j][1]
	})

	res := make([]int, len(nums))
	for i := 0; i < len(tmp); i++ {
		res[i] = tmp[i][0]
	}
	return res
}

func TestSortByGroup(t *testing.T) {
	fmt.Println(sortByGroup([]int{3, 1, 1, 3, 4, 5}, []int{3, 1}))

}
