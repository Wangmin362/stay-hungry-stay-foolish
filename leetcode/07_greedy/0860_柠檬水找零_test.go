package _0_basic

import (
	"fmt"
	"testing"
)

func lemonadeChange(bills []int) bool {
	charge := [3]int{} // charge 0表示五美元的数量，1表示10美元的数量，2表示20美元的数量
	for _, bill := range bills {
		switch bill {
		case 5: // 不需要找零
			charge[0]++
		case 10: // 需要找一个五元的
			if charge[0] <= 0 { // 如果此时没有5元的，那肯定无法找零
				return false
			}
			charge[0]--
			charge[1]++
		case 20: // TODO 这列需要注意，我可以找三个五元的 或者 一个10元的，一个五元的。我应该优先使用10元的，因为5元是万能的，那种情况下都能找零，而10元只能给20元找零
			// 优先找10元的
			if charge[1] >= 1 && charge[0] >= 1 { // 优先找一个10元，一个五元
				charge[1]--
				charge[0]--
			} else if charge[0] >= 3 { // 如果找不到，那么只能退而求其次找三个五元的
				charge[0] -= 3
			} else { // 如果都无法满足，那么无法实现找零
				return false
			}

			charge[2]++
		}
	}

	return true
}

func TestCharge(t *testing.T) {
	fmt.Println(lemonadeChange([]int{5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5}))
}
