package time

import (
	"fmt"
	"testing"
	"time"
)

// TODO 为什么简易使用time.Time，为什么time使用字符串简易使用RFC3339标准？
// TODO 如果time需要传递给前端，就是使用RFC3339么？ javascript的时间遵从这个标准么？
func TestTimeFormat(t *testing.T) {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Format(time.RFC3339))
}
