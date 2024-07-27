package main

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func startServer01(handle func(*Conn)) net.Listener {
	ln, err := net.Listen("tcp", ":7856")
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				fmt.Println("[WARNING] ln.Accept", err)
				return
			}
			go handle(NewConn(conn))
		}
	}()
	return ln
}

func TestServer(t *testing.T) {
	const (
		key  = "Bible"
		data = `Then I heard the voice of the Lord saying, “Whom shall I send? And who will go for us?”
And I said, “Here am I. Send me!”
Isaiah 6:8`
	)

	ln := startServer01(func(conn *Conn) {
		// 服务端等待客户端进行传输
		l.Println("接收到客户端的连接")
		_key, reader, err := conn.Receive()
		if err != nil {
			panic(err)
		}
		l.Printf("接收到key=%s的连接, 获取到reader", _key)

		assertEqual(_key, key)
		dataB, err := io.ReadAll(reader)
		if err != nil {
			panic(err)
		}

		l.Printf("接收到key=%s的数据, 数据为：%s", _key, dataB)
		assertEqual(string(dataB), data)

		// 服务端向客户端进行传输
		writer, err := conn.Send(key)
		if err != nil {
			panic(err)
		}
		l.Printf("服务端发送key=%s", key)

		n, err := writer.Write([]byte(data))
		if err != nil {
			panic(err)
		}

		l.Printf("服务端key=%s, 写入数据%s", key, data)

		if n != len(data) {
			panic(n)
		}
		conn.Close()

		l.Printf("服务端key=%s, 写入数据完成", key)
	})

	defer ln.Close()
	select {}
}
