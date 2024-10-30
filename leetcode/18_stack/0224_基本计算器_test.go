package _1_array

import (
	"strconv"
	"strings"
	"testing"
)

// 解题思路：使用双栈来实现，一个栈保存数字，一个栈保存操作数，括号也看成是一个操作数
func calculate(s string) int {
	isNum := func(char byte) bool {
		if char >= '0' && char <= '9' {
			return true
		}
		return false
	}

	// 去除所有空格
	s = strings.ReplaceAll(s, " ", "")

	opStack, numStack := make([]byte, 0, 64), make([]int, 0, 64)
	numStack = append(numStack, 0) // 操作数添加一个零，防止开头第一个符号为负数

	var calc func()
	calc = func() {
		if len(opStack) == 0 || len(numStack) < 2 { // 无法进行操作
			return
		}
		top := opStack[len(opStack)-1]
		if top == '+' || top == '-' {
			opStack = opStack[:len(opStack)-1]
			n1, n2 := numStack[len(numStack)-2], numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-2]
			var num int
			if top == '+' {
				num = n1 + n2
			} else {
				num = n1 - n2
			}
			numStack = append(numStack, num)
		} else {
			return // 说明是左括号
		}
		calc() // 如果还可以计算，继续计算
	}

	idx := 0
	for idx < len(s) {
		char := s[idx]
		//b := fmt.Sprintf("%c", char)
		//fmt.Println(b)
		if isNum(char) {
			begin := idx                        // 记录第一个位置
			for idx < len(s) && isNum(s[idx]) { // 找到最后一个位置
				idx++
			}
			n, _ := strconv.Atoi(s[begin:idx]) // 找到了当前的数字
			numStack = append(numStack, n)
			idx--
		} else if char == '(' {
			opStack = append(opStack, '(')
			if s[idx+1] == '-' || s[idx+1] == '+' { // 处理(-  或者(+的情况
				numStack = append(numStack, 0) // 添加前置零，方便计算
			}
		} else if char == '+' || char == '-' {
			// 放入之前必须把栈中可以计算的全部计算了
			calc()
			opStack = append(opStack, char)
		} else if char == ')' {
			calc()
			opStack = opStack[:len(opStack)-1] // 去掉左括号
		} else if char == ' ' {

		}
		idx++
	}
	calc()

	return numStack[len(numStack)-1]
}

func TestCalculate(t *testing.T) {
	var testdata = []struct {
		s    string
		want int
	}{
		{s: "(1+(4+5+2)-3)+(6+8)", want: 23},
		{s: "-(+1+(4+5+2)-3)+(6+8)", want: 5},
		{s: "+(1+(4+5+2)-3)+(+6+8)", want: 23},
		{s: "1-(     -2)", want: 3},
	}
	for _, tt := range testdata {
		get := calculate(tt.s)
		if get != tt.want {
			t.Errorf("s:%v, want:%v, get:%v", tt.s, tt.want, get)
		}
	}
}
