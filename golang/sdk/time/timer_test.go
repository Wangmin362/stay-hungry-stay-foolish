package time

import (
	"fmt"
	"testing"
	"time"
)

func TestAfterFunc(t *testing.T) {
	time.AfterFunc(5*time.Second, func() {
		fmt.Println("now")
	})

	time.Sleep(time.Hour)
}
