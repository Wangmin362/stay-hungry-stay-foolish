package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"net"
	"sync"
)

/* 尊敬的面试官大牛前辈您好，我的核心设计思路如下：
1、仿照HTTP2的流特性原理，不同的key生成不同的stream，也就是说每个key都是单独的stream.
2、由于TCP协议时基于字节流的协议，因此接收方不知道多大的数据是一个完整的数据包，因此我们需要定义上层协议，约定如何判断当前接收到的数据包是否是
一个完整的数据包，这样做也是为了能够解决TCP粘包问题，在应用程可以根据协议内容具体的识别一个数据包
3、协议定义
	|  version   |  key length(变长)  |  key |  data  |
	version: 一个字节，在协议发展的过程中，可能由于需求的改变导致协议发生改变，因此协议的解析也需要发生变化，此时客户端服务端都需要改变，因此增加
		版本字段为了将来可以扩展协议。
	key length: 变长字节，由于题目中没有描述key的应用场景，因此这里假定key可以时无限大的，这样也能更多的适配所有用户。对于key length的每一个
	    字节，每个字节的最高位有特殊函数，不能用于表示数据的长度，0表示当前key的长度只够一个字节，那么只需要读取低7位作为key的长度。1表示当前key
        的长度大于128，因此第二个字节也是key的长度信息，类似的，第二个字节的长度也是这样的函数，一直读取到最后一个字节，这个字节的最高位为0，表示
        这个字节的只有低7位有效。
    key: key字段表示key的数据信息，起大小由前面的key length决定
    data: key之后的所有数据，都认为是真正的数据，直到客户端调用Close()方法
*/

// Conn 是你需要实现的一种连接类型，它支持下面描述的若干接口；
// 为了实现这些接口，你需要设计一个基于 TCP 的简单协议；
type Conn struct {
}

// Send 传入一个 key 表示发送者将要传输的数据对应的标识；
// 返回 writer 可供发送者分多次写入大量该 key 对应的数据；
// 当发送者已将该 key 对应的所有数据写入后，调用 writer.Close 告知接收者：该 key 的数据已经完全写入；
func (conn *Conn) Send(key string) (writer io.WriteCloser, err error) {
	return nil, nil
}

// Receive 返回一个 key 表示接收者将要接收到的数据对应的标识；
// 返回的 reader 可供接收者多次读取该 key 对应的数据；
// 当 reader 返回 io.EOF 错误时，表示接收者已经完整接收该 key 对应的数据；
func (conn *Conn) Receive() (key string, reader io.Reader, err error) {
	return "", nil, nil
}

// Close 关闭你实现的连接对象及其底层的 TCP 连接
func (conn *Conn) Close() {
}

// NewConn 从一个 TCP 连接得到一个你实现的连接对象
func NewConn(conn net.Conn) *Conn {
	return nil
}

// 除了上面规定的接口，你还可以自行定义新的类型，变量和函数以满足实现需求

//////////////////////////////////////////////
///////// 接下来的代码为测试代码，请勿修改 /////////
//////////////////////////////////////////////

// 连接到测试服务器，获得一个你实现的连接对象
func dial(serverAddr string) *Conn {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		panic(err)
	}
	return NewConn(conn)
}

// 启动测试服务器
func startServer(handle func(*Conn)) net.Listener {
	ln, err := net.Listen("tcp", ":0")
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

// 简单断言
func assertEqual[T comparable](actual T, expected T) {
	if actual != expected {
		panic(fmt.Sprintf("actual:%v expected:%v\n", actual, expected))
	}
}

// 简单 case：单连接，双向传输少量数据
func testCase0() {
	const (
		key  = "Bible"
		data = `Then I heard the voice of the Lord saying, “Whom shall I send? And who will go for us?”
And I said, “Here am I. Send me!”
Isaiah 6:8`
	)
	ln := startServer(func(conn *Conn) {
		// 服务端等待客户端进行传输
		_key, reader, err := conn.Receive()
		if err != nil {
			panic(err)
		}
		assertEqual(_key, key)
		dataB, err := io.ReadAll(reader)
		if err != nil {
			panic(err)
		}
		assertEqual(string(dataB), data)

		// 服务端向客户端进行传输
		writer, err := conn.Send(key)
		if err != nil {
			panic(err)
		}
		n, err := writer.Write([]byte(data))
		if err != nil {
			panic(err)
		}
		if n != len(data) {
			panic(n)
		}
		conn.Close()
	})
	//goland:noinspection GoUnhandledErrorResult
	defer ln.Close()

	conn := dial(ln.Addr().String())
	// 客户端向服务端传输
	writer, err := conn.Send(key)
	if err != nil {
		panic(err)
	}
	n, err := writer.Write([]byte(data))
	if n != len(data) {
		panic(n)
	}
	err = writer.Close()
	if err != nil {
		panic(err)
	}
	// 客户端等待服务端传输
	_key, reader, err := conn.Receive()
	if err != nil {
		panic(err)
	}
	assertEqual(_key, key)
	dataB, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	assertEqual(string(dataB), data)
	conn.Close()
}

// 生成一个随机 key
func newRandomKey() string {
	buf := make([]byte, 8)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(buf)
}

// 读取随机数据，并返回随机数据的校验和：用于验证数据是否完整传输
func readRandomData(reader io.Reader, hash hash.Hash) (checksum string) {
	hash.Reset()
	var buf = make([]byte, 23<<20) //调用者读取时的 buf 大小不是固定的，你的实现中不可假定 buf 为固定值
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		_, err = hash.Write(buf[:n])
		if err != nil {
			panic(err)
		}
	}
	checksum = hex.EncodeToString(hash.Sum(nil))
	return checksum
}

// 写入随机数据，并返回随机数据的校验和：用于验证数据是否完整传输
func writeRandomData(writer io.Writer, hash hash.Hash) (checksum string) {
	hash.Reset()
	const (
		dataSize = 500 << 20 //一个 key 对应 500MB 随机二进制数据，dataSize 也可以是其他值，你的实现中不可假定 dataSize 为固定值
		bufSize  = 1 << 20   //调用者写入时的 buf 大小不是固定的，你的实现中不可假定 buf 为固定值
	)
	var (
		buf  = make([]byte, bufSize)
		size = 0
	)
	for i := 0; i < dataSize/bufSize; i++ {
		_, err := rand.Read(buf)
		if err != nil {
			panic(err)
		}
		_, err = hash.Write(buf)
		if err != nil {
			panic(err)
		}
		n, err := writer.Write(buf)
		if err != nil {
			panic(err)
		}
		size += n
	}
	if size != dataSize {
		panic(size)
	}
	checksum = hex.EncodeToString(hash.Sum(nil))
	return checksum
}

// 复杂 case：多连接，双向传输，大量数据，多个不同的 key
func testCase1() {
	var (
		mapKeyToChecksum = map[string]string{}
		lock             sync.Mutex
	)
	ln := startServer(func(conn *Conn) {
		// 服务端等待客户端进行传输
		key, reader, err := conn.Receive()
		if err != nil {
			panic(err)
		}
		var (
			h         = sha256.New()
			_checksum = readRandomData(reader, h)
		)
		lock.Lock()
		checksum, keyExist := mapKeyToChecksum[key]
		lock.Unlock()
		if !keyExist {
			panic(fmt.Sprintln(key, "not exist"))
		}
		assertEqual(_checksum, checksum)

		// 服务端向客户端连续进行 2 次传输
		for _, key := range []string{newRandomKey(), newRandomKey()} {
			writer, err := conn.Send(key)
			if err != nil {
				panic(err)
			}
			checksum := writeRandomData(writer, h)
			lock.Lock()
			mapKeyToChecksum[key] = checksum
			lock.Unlock()
			err = writer.Close() //表明该 key 的所有数据已传输完毕
			if err != nil {
				panic(err)
			}
		}
		conn.Close()
	})
	//goland:noinspection GoUnhandledErrorResult
	defer ln.Close()

	conn := dial(ln.Addr().String())
	// 客户端向服务端传输
	var (
		key = newRandomKey()
		h   = sha256.New()
	)
	writer, err := conn.Send(key)
	if err != nil {
		panic(err)
	}
	checksum := writeRandomData(writer, h)
	lock.Lock()
	mapKeyToChecksum[key] = checksum
	lock.Unlock()
	err = writer.Close()
	if err != nil {
		panic(err)
	}

	// 客户端等待服务端的多次传输
	keyCount := 0
	for {
		key, reader, err := conn.Receive()
		if err == io.EOF {
			// 服务端所有的数据均传输完毕，关闭连接
			break
		}
		if err != nil {
			panic(err)
		}
		_checksum := readRandomData(reader, h)
		lock.Lock()
		checksum, keyExist := mapKeyToChecksum[key]
		lock.Unlock()
		if !keyExist {
			panic(fmt.Sprintln(key, "not exist"))
		}
		assertEqual(_checksum, checksum)
		keyCount++
	}
	assertEqual(keyCount, 2)
	conn.Close()
}

func main() {
	testCase0()
	testCase1()
}
