package _1_array

import (
	"strings"
	"testing"
)

// https://leetcode.cn/problems/simplify-path/description/?envType=study-plan-v2&envId=top-interview-150

func simplifyPath(path string) string {
	sp := strings.Split(path, "/")
	spValid := make([]string, 0, len(path))
	for i := 0; i < len(sp); i++ {
		if len(sp[i]) > 0 {
			spValid = append(spValid, sp[i])
		}
	}

	stack := make([]string, 0, len(spValid))
	for _, spv := range spValid {
		switch spv {
		case ".":
			continue
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, spv)
		}
	}

	return "/" + strings.Join(stack, "/")
}

func TestSimplifyPath(t *testing.T) {
	var testdata = []struct {
		path string
		want string
	}{
		{path: "/home/", want: "/home"},
		{path: "/home//foo/", want: "/home/foo"},
		{path: "/home/user/Documents/../Pictures", want: "/home/user/Pictures"},
		{path: "/../", want: "/"},
		{path: "/.../a/../b/c/../d/./", want: "/.../b/d"},
	}
	for _, tt := range testdata {
		get := simplifyPath(tt.path)
		if get != tt.want {
			t.Fatalf("path:%v, want:%v, get:%v", tt.path, tt.want, get)
		}
	}
}
