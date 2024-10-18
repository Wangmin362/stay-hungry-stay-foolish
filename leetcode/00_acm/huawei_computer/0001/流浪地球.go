package main

import (
	"fmt"
	"math"
	"sort"
)

type pair struct {
	time int
	idx  int
}

func main() {
	var total int
	var startNum int
	fmt.Scan(&total)
	fmt.Scan(&startNum)

	var starts []pair
	for i := 0; i < startNum; i++ {
		time, idx := 0, 0
		fmt.Scan(&time)
		fmt.Scan(&idx)
		starts = append(starts, pair{time, idx})
	}

	ee, engines := escapeEarth(total, starts)
	fmt.Println(ee)
	fmt.Println(engines)
}

func escapeEarth(total int, starts []pair) (cnt int, lastEngines []int) {
	// 发动起启动的时刻应该按照从小到大排序
	sort.Slice(starts, func(i, j int) bool {
		return starts[i].time < starts[j].time
	})

	engines := make([]bool, total)
	for t := 0; t <= total; t++ {
		done := true
		for i := 0; i < total; i++ {
			if !engines[i] {
				done = false
				break
			}
		}
		if done {
			sort.Ints(lastEngines)
			return cnt, lastEngines
		}
		cache := make(map[int]struct{})
		for i := 0; i < total; i++ {
			if engines[i] { // 已经启动
				// 旁边两个也设置为true
				if !engines[(i-1+total)%total] {
					cache[(i-1+total)%total] = struct{}{}
				}

				if !engines[(i+1)%total] {
					cache[(i+1)%total] = struct{}{}
				}
			}
		}
		for _, p := range starts {
			if p.time == t {
				cache[p.idx] = struct{}{}
			}
		}

		lastEngines = lastEngines[:0]
		for eng := range cache {
			engines[eng] = true
			lastEngines = append(lastEngines, eng)
		}
		cnt = len(cache)
	}
	return
}

// 初始化的数组记录每个发动机启动的时间，找到最后时刻启动的发动机
func escapeEarth02(total int, starts []pair) (cnt int, lastEngines []int) {
	sort.Slice(starts, func(i, j int) bool {
		return starts[i].idx < starts[j].idx
	})

	engines := make([]int, total) // 记录每个发动机的启动时刻
	for i := 0; i < total; i++ {
		engines[i] = math.MaxInt32 // 初始化设置发动机的启动时刻为一个不可能的最大值
	}

	var t int
	var startedEngine int
	for t = 0; t <= math.MaxInt32; t++ { // 有可能我第一次手动启动发动机的时间很晚很晚
		if startedEngine >= total {
			break
		}
		cache := make(map[int]struct{})
		for i := 0; i < total; i++ {
			if engines[i] != math.MaxInt32 { // 说明这个发动机已经启动，启动周围两个发动机
				if _, ok := cache[(i-1+total)%total]; !ok && engines[(i-1+total)%total] == math.MaxInt32 { // 没有启动才需要启动
					cache[(i-1+total)%total] = struct{}{}
					startedEngine++
				}
				if _, ok := cache[(i+1)%total]; !ok && engines[(i+1)%total] == math.MaxInt32 {
					cache[(i+1)%total] = struct{}{}
					startedEngine++
				}
			}
		}
		for _, p := range starts {
			if p.time == t { // 说明需要手动启动发动机
				if _, ok := cache[p.idx]; !ok && engines[p.idx] == math.MaxInt32 {
					cache[p.idx] = struct{}{}
					startedEngine++
				}
			}
		}

		// 每次最后需要设置需要启动的发动机
		for egn := range cache {
			engines[egn] = t
		}
	}

	t-- // 退一个时刻
	for idx, engTime := range engines {
		if engTime == t {
			lastEngines = append(lastEngines, idx)
		}
	}
	return len(lastEngines), lastEngines
}
