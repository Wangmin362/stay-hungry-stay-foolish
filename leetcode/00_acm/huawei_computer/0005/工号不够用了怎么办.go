package main

func main() {

}

func pow(x, y int) int {
	res := x
	for i := 1; i < y; i++ {
		res *= res
	}
	return res
}

func gongHao(total, y int) (res int) {
	curr := pow('z'-'a'+1, y)
	if curr >= total {
		return 1 // 至少一个数字
	}

	for curr < total {
		curr *= 10
		res++
	}

	return res
}
