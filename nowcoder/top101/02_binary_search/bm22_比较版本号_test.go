package _2_binary_search

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// https://www.nowcoder.com/practice/2b317e02f14247a49ffdbdba315459e7

// 直接比较
func compare(version1 string, version2 string) int {
	arr1 := strings.Split(version1, ".")
	arr2 := strings.Split(version2, ".")
	length := len(arr1)
	if len(arr1) < len(arr2) {
		length = len(arr2)
	}
	for i := 0; i < length; i++ {
		s1, s2 := 0, 0
		if i <= len(arr1)-1 {
			s1, _ = strconv.Atoi(arr1[i])
		}
		if i <= len(arr2)-1 {
			s2, _ = strconv.Atoi(arr2[i])
		}
		if s1 == s2 {
			continue
		} else if s1 > s2 {
			return 1
		} else {
			return -1
		}
	}

	return 0
}

func TestCompare(t *testing.T) {
	//fmt.Println(compare("1.0.1", "1"))
	//fmt.Println(compare("1.0", "1.0.0"))
	fmt.Println(compare("1.256", "1.87.0"))
}
