package _1_array

import (
	"container/list"
	"strconv"
	"strings"
	"testing"
)

// 题目：https://leetcode.cn/problems/evaluate-reverse-polish-notation/description/

func evalRPN(tokens []string) int {
	var stack []string
	for _, token := range tokens {
		if strings.Contains("+-*/", token) {
			if len(stack) < 2 {
				return 0
			}
			right := stack[len(stack)-1]
			rightVal, _ := strconv.Atoi(right)
			left := stack[len(stack)-2]
			leftVal, _ := strconv.Atoi(left)
			var res int
			switch token {
			case "+":
				res = leftVal + rightVal
			case "-":
				res = leftVal - rightVal
			case "*":
				res = leftVal * rightVal
			case "/":
				res = leftVal / rightVal
			}
			stack[len(stack)-2] = strconv.Itoa(res)
			stack = stack[:len(stack)-1] // 抵消一个数字，最后一个数字不需要
		} else {
			stack = append(stack, token)
		}
	}
	if len(stack) != 1 {
		return 0
	}

	atoi, _ := strconv.Atoi(stack[0])

	return atoi
}

func evalRPN01(tokens []string) int {
	stack := list.New()
	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			n1 := stack.Remove(stack.Back()).(int)
			n2 := stack.Remove(stack.Back()).(int)
			var res int
			switch token {
			case "+":
				res = n2 + n1
			case "-":
				res = n2 - n1
			case "*":
				res = n2 * n1
			case "/":
				res = n2 / n1
			}
			stack.PushBack(res)
		} else {
			num, _ := strconv.Atoi(token)
			stack.PushBack(num)
		}
	}
	return stack.Remove(stack.Back()).(int)
}

func evalRPN03(tokens []string) int {
	stack := make([]int, 0, len(tokens))

	for _, token := range tokens {
		if token != "+" && token != "-" && token != "*" && token != "/" {
			val, _ := strconv.Atoi(token)
			stack = append(stack, val)
			continue
		}

		one := stack[len(stack)-2]
		two := stack[len(stack)-1]
		switch token {
		case "+":
			stack[len(stack)-2] = one + two
		case "-":
			stack[len(stack)-2] = one - two
		case "*":
			stack[len(stack)-2] = one * two
		case "/":
			stack[len(stack)-2] = one / two
		}
		stack = stack[:len(stack)-1]
	}
	return stack[0]
}

func TestEvalRPN(t *testing.T) {
	var teatdata = []struct {
		s      []string
		expect int
	}{
		{s: []string{"2", "1", "+", "3", "*"}, expect: 9},
		{s: []string{"4", "13", "5", "/", "+"}, expect: 6},
		{s: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, expect: 22},
	}

	for _, test := range teatdata {
		get := evalRPN03(test.s)
		if get != test.expect {
			t.Errorf("s: %v, expect:%v, get:%v", test.s, test.expect, get)
		}
	}
}
