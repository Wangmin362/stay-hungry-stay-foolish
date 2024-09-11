package _0_basic

import (
	"reflect"
	"testing"
)

// https://leetcode.cn/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/description/?envType=problem-list-v2&envId=two-pointers&difficulty=EASY

// 同向指针，写的不优雅，需要判断很多情况
func trainingPlan(actions []int) []int {
	slow, fast := 0, 0

	// 找到第一个偶数
	for slow < len(actions) && actions[slow]%2 == 1 {
		slow++
	}

	// 找到slow之后第一个奇数
	fast = slow + 1
	for fast < len(actions) && actions[fast]%2 == 0 {
		fast++
	}

	for fast < len(actions) {
		actions[slow], actions[fast] = actions[fast], actions[slow]
		slow++ // 移动到下一个需要放奇数的位置
		fast++
		for fast < len(actions) && actions[fast]%2 == 0 { // 移动到下一个奇数位置
			fast++
		}
	}

	return actions
}

// 碰撞指针
func trainingPlan01(actions []int) []int {
	left, right := 0, len(actions)-1
	for left < right {
		for left < right && actions[left]%2 == 1 { // 找到第一个偶数
			left++
		}
		for left < right && actions[right]%2 == 0 { // 找到第一个奇数
			right--
		}
		if left >= right {
			return actions
		}
		actions[left], actions[right] = actions[right], actions[left]
		left++
		right--
	}
	return actions
}

func TestTrainingPlan(t *testing.T) {
	var testData = []struct {
		actions []int
		want    []int
	}{
		{actions: []int{1, 2, 3, 4, 5}, want: []int{1, 3, 5, 4, 2}},
		{actions: []int{2, 4, 6}, want: []int{2, 4, 6}},
	}

	for _, tt := range testData {
		get := trainingPlan01(tt.actions)
		if !reflect.DeepEqual(get, tt.want) {
			t.Fatalf("actions:%v, want:%v, get:%v", tt.actions, tt.want, get)
		}
	}
}
