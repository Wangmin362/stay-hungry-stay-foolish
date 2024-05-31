package _0_algorithm

import (
	"github.com/google/uuid"
	"testing"
)

func TestName(t *testing.T) {
	for i := 0; i <= 99; i++ {
		t.Logf(uuid.New().String())
	}

}
