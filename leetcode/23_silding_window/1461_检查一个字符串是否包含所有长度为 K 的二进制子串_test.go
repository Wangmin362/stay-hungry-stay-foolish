package _1_array

import (
	"math"
	"testing"
)

func hasAllCodes(s string, k int) bool {
	kind := int(math.Pow(2, float64(k)))
	sk := make(map[string]struct{})

	left, right := 0, 0
	for ; right < len(s); right++ {
		if right < k-1 {
			continue
		}
		sk[s[left:right+1]] = struct{}{}
		left++
	}

	return len(sk) == kind
}

func TestHasAllCodes(t *testing.T) {
	var testdata = []struct {
		s    string
		k    int
		want bool
	}{
		{s: "00110110", k: 2, want: true},
		{s: "00110", k: 2, want: true},
		{s: "0011", k: 2, want: false},
		{s: "11", k: 1, want: false},
		{s: "01", k: 1, want: true},
		{s: "10", k: 1, want: true},
	}
	for _, tt := range testdata {
		get := hasAllCodes(tt.s, tt.k)
		if get != tt.want {
			t.Errorf("s:%v, k:%v, want:%v, get:%v", tt.s, tt.k, tt.want, get)
		}
	}
}
