package go_monkey

import (
	"reflect"
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/agiledragon/gomonkey/v2/test/fake"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

// MyClient 是一个示例结构体，包含一个公有方法
type MyClient struct {
	Name string
}

// GetData 是一个公有方法，返回数据
func (c *MyClient) GetData() string {
	// 原始实现可能会从外部获取数据，这里为了简化直接返回
	return "real data from external source"
}

// TODO 这种方法会导致gomonkey打桩失败，也就是说这种方法不能打桩
func (c MyClient) GetData22() string {
	// 原始实现可能会从外部获取数据，这里为了简化直接返回
	return "real data from external source data22"
}

func FetchData(c *MyClient) string {
	// 使用 MyClient 的 GetData 方法获取数据
	return c.GetData()
}

// TODO 使用gomonkey时，需要禁止golang内联优化，否则会造成gomonkey打桩失败
// 禁止内联优化的方法是在编译时添加 -gcflags "-l -N" 标志
func TestFetchData(t *testing.T) {
	client := &MyClient{Name: "TestClient"}

	// 使用 gomonkey 打桩 MyClient 的 GetData 方法
	patches := ApplyMethod(reflect.TypeOf(client), "GetData", func(_ *MyClient) string {
		// 返回我们模拟的测试数据
		return "mocked data"
	})
	defer patches.Reset()

	// 调用 FetchData，这时 GetData 方法会返回打桩的模拟数据
	result := FetchData(client)
	//result := client.GetData()

	// 验证结果是否是打桩的返回值
	assert.Equal(t, "mocked data", result)
}

func TestFetchData22(t *testing.T) {
	client := &MyClient{Name: "TestClient"}

	// 使用 gomonkey 打桩 MyClient 的 GetData22 方法
	patches := ApplyMethod(reflect.TypeOf(client), "GetData22", func(_ *MyClient) string {
		// 返回我们模拟的测试数据
		return "mocked data22"
	})
	defer patches.Reset()

	// 调用 FetchData，这时 GetData 方法会返回打桩的模拟数据
	result := client.GetData22()

	// 验证结果是否是打桩的返回值
	assert.Equal(t, "mocked data22", result)
}

func TestGoMonkeyTest(t *testing.T) {
	slice := fake.NewSlice()

	Convey("TestApplyMethod", t, func() {

		Convey("for succ", func() {
			err := slice.Add(1)
			So(err, ShouldEqual, nil)
			patches := ApplyMethod(reflect.TypeOf(&slice), "Add", func(_ *fake.Slice, _ int) error {
				return nil
			})
			defer patches.Reset()
			err = slice.Add(1)
			So(err, ShouldEqual, nil)
			err = slice.Remove(1)
			So(err, ShouldEqual, nil)
			So(len(slice), ShouldEqual, 0)
		})

		Convey("for already exist", func() {
			err := slice.Add(2)
			So(err, ShouldEqual, nil)
			patches := ApplyMethod(reflect.TypeOf(&slice), "Add", func(_ *fake.Slice, _ int) error {
				return fake.ErrElemExsit
			})
			defer patches.Reset()
			err = slice.Add(1)
			So(err, ShouldEqual, fake.ErrElemExsit)
			err = slice.Remove(2)
			So(err, ShouldEqual, nil)
			So(len(slice), ShouldEqual, 0)
		})
	})
}
