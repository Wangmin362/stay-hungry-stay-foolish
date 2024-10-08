package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var row int
	fmt.Scanln(&row)

	arr := make([]int, row)
	for i := 0; i < row; i++ {
		var num int
		fmt.Scanln(&num)
		arr[i] = num
	}

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		ss := strings.Split(scan.Text(), " ")
		n1, _ := strconv.Atoi(ss[0])
		n2, _ := strconv.Atoi(ss[1])
		fmt.Println(getSum(arr, n1, n2))
	}
}

func getSum(nums []int, begin, end int) int {
	var sum int
	for i := begin; i <= end; i++ {
		sum += nums[i]
	}

	return sum
}
