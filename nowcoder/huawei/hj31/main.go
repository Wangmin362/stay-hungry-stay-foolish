package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// https://www.nowcoder.com/practice/81544a4989df4109b33c2d65037c5836

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	word := scan.Text()
	fmt.Println(reverseWords(word))
}

func reverseWords(word string) string {
	isAlpha := func(c byte) bool {
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
			return true
		}
		return false
	}
	slow, fast := len(word)-1, len(word)-1
	var res []string
	for {
		for slow >= 0 && !isAlpha(word[slow]) { // 先把slow移动到第一个字母
			slow--
		}
		if slow < 0 {
			break
		}

		fast = slow // fast从slow的位置开始移动，知道找到第一个非字符
		for fast >= 0 && isAlpha(word[fast]) {
			fast--
		}
		res = append(res, word[fast+1:slow+1])
		if fast < 0 {
			break
		}

		slow = fast
	}

	return strings.Join(res, " ")
}
