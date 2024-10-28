package _1_array

import (
	"strconv"
	"strings"
	"testing"
)

func decodeString(s string) string {
	strStack := make([]string, 0, len(s))
	numStack := make([]int, 0, len(s))
	var res string
	for i := 0; i < len(s); {
		start := i
		if s[i] >= '0' && s[i] <= '9' {
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				i++
			}
			n, _ := strconv.Atoi(s[start:i])
			numStack = append(numStack, n)
		} else if s[i] >= 'a' && s[i] <= 'z' {
			for i < len(s) && s[i] >= 'a' && s[i] <= 'z' {
				i++
			}
			strStack = append(strStack, s[start:i])
		} else if s[i] == ']' {
			var ptn string
			for strStack[len(strStack)-1] != "[" {
				ptn = strStack[len(strStack)-1] + ptn
				strStack = strStack[:len(strStack)-1]
			}
			strStack = strStack[:len(strStack)-1] // 去掉[

			frq := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			var str strings.Builder
			for j := 0; j < frq; j++ {
				str.WriteString(ptn)
			}
			strStack = append(strStack, str.String())
			i++
		} else {
			strStack = append(strStack, "[")
			i++
		}
	}
	for _, str := range strStack {
		res += str
	}

	return res
}
func TestCode(t *testing.T) {
	var testdata = []struct {
		s    string
		want string
	}{
		{s: "3[a]2[bc]", want: "aaabcbc"},
		{s: "3[a2[c]]", want: "accaccacc"},
		{s: "2[abc]3[cd]ef", want: "abcabccdcdcdef"},
		{s: "abc3[cd]xyz", want: "abccdcdcdxyz"},
		{s: "3[z]2[2[y]pq4[2[jk]e1[f]]]ef", want: "zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef"},
	}
	for _, tt := range testdata {
		get := decodeString(tt.s)
		if get != tt.want {
			t.Errorf("s:%v, want:%v, get:%v", tt.s, tt.want, get)
		}
	}
}
