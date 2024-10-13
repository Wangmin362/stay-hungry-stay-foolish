package _0_basic

import (
	"sort"
	"testing"
)

// 贪心思想：优先和小行星碰撞，获取小行星的质量
func asteroidsDestroyed(mass int, asteroids []int) bool {
	sort.Ints(asteroids)
	for i := 0; i < len(asteroids); i++ {
		if mass < asteroids[i] {
			return false
		}
		mass += asteroids[i]
	}

	return true
}

func TestAsteroidsDestroyed(t *testing.T) {
	var testsdata = []struct {
		mass      int
		asteroids []int
		want      bool
	}{
		{mass: 10, asteroids: []int{3, 9, 19, 5, 21}, want: true},
	}
	for _, tt := range testsdata {
		get := asteroidsDestroyed(tt.mass, tt.asteroids)
		if get != tt.want {
			t.Fatalf("mass:%v, asteroids:%v, want:%v, get:%v", tt.mass, tt.asteroids, tt.want, get)
		}
	}
}
