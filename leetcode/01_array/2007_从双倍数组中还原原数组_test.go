package _1_array

import "sort"

func findOriginalArray(changed []int) []int {
	if len(changed)%2 == 1 {
		return nil
	}

	m := make(map[int]int)
	for _, num := range changed {
		m[num]++
	}
	sort.Ints(changed)
	res := make([]int, 0, len(changed)/2)
	for i := 0; i < len(changed); i++ {
		if i == 0 {
			res = append(res, changed[i]) // 第一个元素肯定在数组当中
			m[changed[i]]--
			if cnt, ok := m[changed[i]*2]; ok && cnt >= 1 {
				m[changed[i]*2]--
			} else {
				return nil
			}
		}
		if cnt, ok := m[changed[i]]; ok && cnt >= 1 {
			res = append(res, changed[i]) // 第一个元素肯定在数组当中
			m[changed[i]]--
			if cnt, ok = m[changed[i]*2]; ok && cnt >= 1 {
				m[changed[i]*2]--
			} else {
				return nil
			}
		}
	}
	return res
}
