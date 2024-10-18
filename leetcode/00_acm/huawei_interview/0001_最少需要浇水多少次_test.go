package huawei_interview

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

// 数组water表示一排瓶子的水位高度。小明往这些瓶子内浇水，1次操作可以使1个瓶子的水位增加1。给定一个整数cnt，
// 输入：water: [7,1,9,10] ,cnt=3  输出：3

func minWaterCount(water []int, cnt int) int {
	sort.Ints(water)

	tmp := make([]int, 0, len(water)-cnt+1)
	for i := 0; i <= len(water)-cnt; i++ {
		// 考虑当前这三瓶水，求需要多少次可以使得三瓶水一样高。
		// 显然，只需要保证每个瓶子浇水浇到最高的瓶子即可
		t0, t1, t2 := water[i], water[i+1], water[i+2]
		total := t2*2 - t1 - t0
		tmp = append(tmp, total)
	}

	res := math.MaxInt
	for i := 0; i < len(tmp); i++ {
		res = min(res, tmp[i])
	}

	return res
}

func TestMinWaterCount(t *testing.T) {
	fmt.Println(minWaterCount([]int{7, 1, 9, 10}, 3))
}
