package _1_array

import (
	"strconv"
	"strings"
	"testing"
)

// 解题思路：使用双栈来实现，一个栈保存数字，一个栈保存操作数，括号也看成是一个操作数
func calculateIII(s string) int {
	// 去除字符串空格
	s = strings.ReplaceAll(s, " ", "")
	if len(s) == 0 {
		return 0
	}
	isNum := func(char byte) bool {
		if char >= '0' && char <= '9' {
			return true
		}
		return false
	}

	// 定义加减乘除优先级，乘除优先级比加减优先级高，同级别优先级因该从左到右计算
	ops := map[byte]int{'+': 0, '-': 0, '*': 1, '/': 1}
	opsStack, numStack := make([]byte, 0, 64), make([]int, 0, 64)
	numStack = append(numStack, 0) // 添加前缀0，防止出现第一个符号是+或者-

	var calc func(priority int)
	calc = func(priority int) {
		if len(opsStack) == 0 || len(numStack) < 2 { // 说明无法操作
			return
		}

		top := opsStack[len(opsStack)-1]
		if top == '(' { // 不能操作
			return
		}

		// 说明操作符栈顶一定是运算符，此时需要看运算符优先级
		if ops[top] < priority { // 前面运算符的优先级比当前第，因此不能优先计算
			return
		}

		opsStack = opsStack[:len(opsStack)-1]
		n1, n2 := numStack[len(numStack)-2], numStack[len(numStack)-1]
		numStack = numStack[:len(numStack)-2]
		var num int
		switch top {
		case '+':
			num = n1 + n2
		case '-':
			num = n1 - n2
		case '*':
			num = n1 * n2
		case '/':
			num = n1 / n2
		}

		numStack = append(numStack, num)

		calc(priority)
	}

	idx := 0
	for idx < len(s) {
		char := s[idx]
		if isNum(char) {
			begin := idx
			for idx < len(s) && isNum(s[idx]) {
				idx++
			}
			n, _ := strconv.Atoi(s[begin:idx])
			numStack = append(numStack, n)
			idx--
		} else if char == '(' {
			opsStack = append(opsStack, '(')
			if s[idx+1] == '-' || s[idx+1] == '+' { // 处理 (-  或者 (+ 情况
				numStack = append(numStack, 0)
			}
		} else if pri, ok := ops[char]; ok { // 说明是允许的操作符
			// 放入操作符之前，需要把之前可以计算的全部计算掉，当然还得看优先级，如果前面的优先级比现在小，就不应该计算
			calc(pri)
			opsStack = append(opsStack, char)
		} else if char == ')' {
			calc(0)
			opsStack = opsStack[:len(opsStack)-1] // 去除掉(括号
		}
		idx++
	}
	calc(0)

	return numStack[len(numStack)-1]
}

func TestCalculateIII(t *testing.T) {
	var testdata = []struct {
		s    string
		want int
	}{
		{s: "(1+(4+5+2)-3)+(6+8)", want: 23},
		{s: "-(+1+(4+5+2)-3)+(6+8)", want: 5},
		{s: "+(1+(4+5+2)-3)+(+6+8)", want: 23},
		{s: "1-(     -2)", want: 3},
		{s: "1-2*5", want: -9},
		{s: "-4+5+9*2", want: 19},
	}
	for _, tt := range testdata {
		get := calculateIII(tt.s)
		if get != tt.want {
			t.Errorf("s:%v, want:%v, get:%v", tt.s, tt.want, get)
		}
	}
}
