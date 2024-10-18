package main

import "sort"

func main() {
}

func pooke(pk []string) [][]string {
	mapp := map[string]int{
		"3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8,
		"9": 9, "10": 10, "J": 11, "Q": 12, "K": 13, "A": 14,
	}
	rmapp := make(map[int]string, len(mapp))
	for k, v := range mapp {
		rmapp[v] = k
	}

	newPk := make([]int, 0, len(pk))
	pkCnt := make(map[int]int)
	for _, p := range pk {
		if p == "2" {
			continue
		}
		pkCnt[mapp[p]]++
	}
	for val := range pkCnt {
		newPk = append(newPk, val)
	}
	sort.Ints(newPk)
	if len(pkCnt) < 5 {
		return nil
	}

	var res [][]string
	for {
		if len(pkCnt) < 5 {
			break
		}

		var sz []string
		for idx := 0; idx < len(newPk); idx++ {
			cnt, ok := pkCnt[newPk[idx]]
			if !ok {
				continue // 说明这种类型的扑克牌已经用完了
			}

			if len(sz) == 0 { // 说明是顺子的第一张牌，不需要比较
				sz = append(sz, rmapp[newPk[idx]])
				if cnt == 1 {
					delete(pkCnt, newPk[idx]) // 用完了就删除
				} else {
					pkCnt[newPk[idx]]--
				}
				continue
			}

			if newPk[idx] == mapp[sz[len(sz)-1]]+1 { // 如果是连续的，就添加进去
				sz = append(sz, rmapp[newPk[idx]])
			} else {
				if len(sz) >= 5 {
					tmp := make([]string, len(sz))
					copy(tmp, sz)
					res = append(res, tmp)
				}

				sz = sz[:0]
				sz = append(sz, rmapp[newPk[idx]])
			}

			if cnt == 1 {
				delete(pkCnt, newPk[idx]) // 用完了就删除
			} else {
				pkCnt[newPk[idx]]--
			}
		}

		if len(sz) >= 5 {
			tmp := make([]string, len(sz))
			copy(tmp, sz)
			res = append(res, tmp)
		}

		sz = sz[:0]
	}

	return res
}
