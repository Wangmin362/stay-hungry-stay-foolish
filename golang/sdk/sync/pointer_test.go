package sync

import (
	"fmt"
	"net"
	"sync/atomic"
	"testing"
)

// atomic.Pointer是
func TestBase01(t *testing.T) {
	type ServerConn struct {
		Connection net.Conn
		ID         string
		Open       bool
	}

	aPointer := atomic.Pointer[ServerConn]{}
	s := ServerConn{ID: "first_conn"}
	aPointer.Store(&s)

	fmt.Println(aPointer.Load())

	aValue := atomic.Value{}
	aValue.Store(&s)
	conn, ok := aValue.Load().(*ServerConn)
	if !ok {
		panic("assert is not ok")
	}
	fmt.Println(conn)
}
