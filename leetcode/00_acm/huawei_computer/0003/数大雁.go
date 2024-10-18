package main

import "fmt"

func main() {
	//fmt.Println(minDy("quackquack")) // 1
	//fmt.Println(minDy("qaauucqcaa")) // -1
	//fmt.Println(minDy("quacqkuackquack")) // 2
	//fmt.Println(minDy("qququaauqccauqkkcauqqkcauuqkcaaukccakkck")) // 5
	fmt.Println(minDy("quacqkuquacqkacuqkackuack")) // 3
}

/* 其实就是一只大雁叫完之后继续叫，最后的结果就是最少的鸭子
// qququaauqccauqkkcauqqkcauuqkcaaukccakkck
//  q qu auq cauq kcau qkca uqkc auk ca kck 1
//    q   uq  auq  cau  kca  qkc  uk  a  ck 2
//         q   uq   au   ca   kc   k        2
//              q    u    a    c   k        4
//                                          5
*/

func minDy(quack string) int {
	bytes := []byte(quack)
	var res int
	for {
		cnt := 0
		idx := 0
		for {
			q, u, a, c, k := 0, 0, 0, 0, 0
			{

				for ; idx < len(bytes); idx++ {
					if bytes[idx] == 'q' {
						break
					}
				}
				q = idx
				idx++
				for ; idx < len(bytes); idx++ {
					if bytes[idx] == 'u' {
						break
					}
				}
				u = idx
				idx++
				for ; idx < len(bytes); idx++ {
					if bytes[idx] == 'a' {
						break
					}
				}
				a = idx
				idx++
				for ; idx < len(bytes); idx++ {
					if bytes[idx] == 'c' {
						break
					}
				}
				c = idx
				idx++
				for ; idx < len(bytes); idx++ {
					if bytes[idx] == 'k' {
						break
					}
				}
				k = idx
			}
			if idx >= len(bytes) {
				break
			} else {
				idx++
				cnt++ // 找到一叫声
				// 找到了就需要清零
				bytes[q] = '0'
				bytes[u] = '0'
				bytes[a] = '0'
				bytes[c] = '0'
				bytes[k] = '0'
			}
		}

		if cnt == 0 { // 凑不出来大雁的叫生一定是结束了
			break
		} else {
			res++
		}
	}

	if res == 0 {
		return -1
	}
	return res
}
