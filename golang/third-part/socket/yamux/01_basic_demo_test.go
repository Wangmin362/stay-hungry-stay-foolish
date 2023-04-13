package yamux

import (
	"fmt"
	"github.com/hashicorp/yamux"
	"net"
	"testing"
	"time"
)

func Recv(stream net.Conn, id int) {
	for {
		buf := make([]byte, 4)
		n, err := stream.Read(buf)
		if err == nil {
			fmt.Println("ID:", id, ", len:", n, time.Now().Unix(), string(buf))
		} else {
			fmt.Println(time.Now().Unix(), err)
			return
		}
	}
}
func TestServer(t *testing.T) {
	// 建立底层复用连接
	tcpaddr, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:8980")
	tcplisten, _ := net.ListenTCP("tcp", tcpaddr)
	conn, _ := tcplisten.Accept()
	session, _ := yamux.Server(conn, nil)

	id := 1
	for {
		// 建立多个流通路
		stream, err := session.Accept()
		if err == nil {
			fmt.Println("accept")
			id++
			go Recv(stream, id)
		} else {
			fmt.Println("session over.")
			return
		}
	}

}

func TestClient(t *testing.T) {
	// 建立底层复用通道
	conn, _ := net.Dial("tcp", "127.0.0.1:8980")
	session, _ := yamux.Client(conn, nil)

	// 建立应用流通道1
	stream, _ := session.Open()
	stream.Write([]byte("ping"))
	stream.Write([]byte("pong"))
	time.Sleep(1 * time.Second)

	// 建立应用流通道2
	stream1, _ := session.Open()
	stream1.Write([]byte("pong"))
	time.Sleep(1 * time.Second)

	// 清理退出
	time.Sleep(5 * time.Second)
	stream.Close()
	stream1.Close()
	session.Close()
}
