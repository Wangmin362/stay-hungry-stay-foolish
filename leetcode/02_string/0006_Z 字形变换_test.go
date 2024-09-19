package _1_array

import (
	"strings"
	"testing"
)

// https://leetcode.cn/problems/zigzag-conversion/description/?envType=study-plan-v2&envId=top-interview-150

// 直接模拟字符串Z字形就可以
func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	res := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]byte, 0)
	}
	flag := 1 // 从上往下读取
	idx := 0  // 从第一行开始放字符串
	for _, c := range s {
		res[idx] = append(res[idx], byte(c))
		if idx == numRows-1 {
			flag = -1
		} else if idx == 0 {
			flag = 1
		}

		idx += flag
	}

	re := strings.Builder{}
	for i := 0; i < numRows; i++ {
		re.Write(res[i])
	}

	return re.String()
}

func TestConvert0006(t *testing.T) {
	var testdata = []struct {
		s       string
		numRows int
		want    string
	}{
		{s: "PAYPALISHIRING", numRows: 3, want: "PAHNAPLSIIGYIR"},
	}

	for _, tt := range testdata {
		get := convert(tt.s, tt.numRows)
		if get != tt.want {
			t.Fatalf("strs:%v, numRows:%v, want:%v, get:%v", tt.s, tt.numRows, tt.want, get)
		}
	}
}
