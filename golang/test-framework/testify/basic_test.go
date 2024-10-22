package testify_test

import (
	"github.com/golang/demo/golang/test-framework/testify"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 非常好用的断言工具
func TestAbc(t *testing.T) {
	assert.Equal(t, 1, 1, "must equal")
	assert.NotEqual(t, 1, 2, "must not equal")
}

func TestAdd(t *testing.T) {
	Convey("test add", t, func() {
		x := testify.Add(4, 6)
		So(x, ShouldEqual, 10)
	})
}
