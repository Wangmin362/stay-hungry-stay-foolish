package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var arr []int
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	for _, s := range strings.Split(scan.Text(), " ") {
		num, _ := strconv.Atoi(s)
		arr = append(arr, num)
	}

	fmt.Println(arr)
}

func emptyStack(nums []int) (res []int) {
	stack := make([]int, 0, len(nums))
	for _, num := range nums {
		if len(stack) == 0 {
			stack = append(stack, num)
			continue
		}

		for len(stack) > 0 {
			top := num
			next := stack[len(stack)-1]
			if top == next {
				num = top * 2
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, num)
				break
			}
		}

		for len(stack) > 2 {
			top := stack[len(stack)-1]
			sum, i := 0, 0
			for i = len(stack) - 2; i >= 0; i-- {
				sum += stack[i]
				if sum == top {
					stack = stack[:i]
					stack = append(stack, top*2)
					break
				}
			}
			if i < 0 {
				break
			}
		}
	}

	for idx := len(stack) - 1; idx >= 0; idx-- {
		res = append(res, stack[idx])
	}
	return res
}
