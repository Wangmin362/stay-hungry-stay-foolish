package main

import "fmt"

func main() {
	var num int
	fmt.Scanf("%d", &num)

	for i := 2; i*i <= num; i++ {
		for num%i == 0 {
			fmt.Printf("%d ", i)
			num = num / i
		}
	}
	if num != 1 {
		fmt.Printf("%d ", num)
	}
}
