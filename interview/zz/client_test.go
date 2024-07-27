package main

import (
	"io"
	"log"
	"testing"
)

var l = log.Default()

func TestClient(t *testing.T) {
	const (
		key  = "Bible"
		data = `Then I heard the voice of the Lord saying, “Whom shall I send? And who will go for us?”
And I said, “Here am I. Send me!”
Isaiah 6:8`
	)

	conn := dial(":7856")
	// 客户端向服务端传输
	s, _ := sha256Str(key)
	l.Printf("发送[%d]长度的key，key：[%s], keyhash=%v", len(key), key, s)
	writer, err := conn.Send(key)
	if err != nil {
		panic(err)
	}

	l.Printf("写入[%d]长度的数据，数据为：[%s]", len(data), data)
	n, err := writer.Write([]byte(data))
	if n != len(data) {
		panic(n)
	}

	l.Printf("key=%s的数据发送完成", key)
	err = writer.Close()
	if err != nil {
		panic(err)
	}
	// 客户端等待服务端传输

	l.Printf("开始接收数据")
	_key, reader, err := conn.Receive()
	if err != nil {
		panic(err)
	}
	l.Printf("接收到key=%s的reader, 后续开始读取数据", _key)

	assertEqual(_key, key)
	dataB, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	l.Printf("读取到key=%s的数据, 数据为：%s", _key, dataB)

	assertEqual(string(dataB), data)
	conn.Close()
}
