package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"log"
	"net"
	"sync"
	"unsafe"
)

/* (思路一) 尊敬的面试官大牛前辈您好，我的设计思路如下：
1、仿照HTTP2的流特性原理，不同的key生成不同的stream，也就是说每个key都是单独的stream.
2、由于TCP协议是基于字节流的协议，因此接收方不知道多大的数据是一个完整的数据包，因此我们需要定义上层协议，约定如何判断当前接收到的数据包是否是
一个完整的数据包，这样做也是为了能够解决TCP粘包问题，在应用程可以根据协议内容具体的识别一个数据包
3、协议定义
	++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	|  version   |  key length(变长)  |  key |  data
	++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	version: 一个字节，在协议发展的过程中，可能由于需求的改变导致协议发生改变，因此协议的解析也需要发生变化，此时客户端服务端都需要改变，因此增加
		版本字段为了将来可以扩展协议。
	key length: 变长字节，由于题目中没有描述key的应用场景，因此这里假定key可以时无限大的，这样也能更多的适配所有用户。对于key length的每一个
	    字节，每个字节的最高位有特殊函数，不能用于表示数据的长度，0表示当前key的长度只够一个字节，那么只需要读取低7位作为key的长度。1表示当前key
        的长度大于128，因此第二个字节也是key的长度信息，类似的，第二个字节的长度也是这样的函数，一直读取到最后一个字节，这个字节的最高位为0，表示
        这个字节的只有低7位有效。
    key: key字段表示key的数据信息，起大小由前面的key length决定
    data: key之后的所有数据，都认为是真正的数据，直到客户端调用Close()方法
*/

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////// 本次的实现采用下面思路二的方式来实现 //////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/* (思路二)  由于您这边对于连接的使用已经明确规定了，并且对于新建自定义连接时，无法区分当前是定义服务端的连接，还是客户端的连接，因此开源的yamux无法
发挥出作用，无法实例化yamux的Server以及Client。但是yamux的思路是行得通的。在数据传输过程中，由于TCP的字节流特性，因此应用层无法区分每个包的边界，
所以这里需要区分当前传输的是key还是数据，以及还需要区分当前传输数据是哪个key的数据。做完以上分析之后，我大概有了以下简单的设计。
	// 以下是对于key的数据包定义
		++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
		|  dataType(0) | version | length(变长)    |  data
		++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
		1、dataType表示当前数据包的类型，占用2位，第一个字节的最高2位表示数据类型 0表示当前数据包是key，1表示当前数据包是data，2标识当前key的
		数据传输完成
		2、version表示当前协议的版本，占用6位，版本范围位[0, 63]，当前版本为0，后续有需求，可以通过修改改变扩展协议
		3、key的长度：变长字节，由于题目中没有描述key的应用场景，因此这里假定key可以是无限大的，这样也能更多的适配所有用户。对于key length的每一个
	    字节，每个字节的最高位有特殊函数，不能用于表示数据的长度，0表示当前key的长度只够一个字节，那么只需要读取低7位作为key的长度。1表示当前key
        的长度大于128，因此第二个字节也是key的长度信息，类似的，第二个字节的长度也是这样的函数，一直读取到最后一个字节，这个字节的最高位为0，表示
        这个字节的只有低7位有效。
		4、data表示当前key的数据

	// 以下是对于data的数据包定义
		++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
		|  dataType(1) | version | SHA3-256 | length(变长)    |  data
		++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
		1、dataType表示当前数据包的类型，占用2位，第一个字节的最高2位表示数据类型 0表示当前数据包是key，1表示当前数据包是data，2标识当前key的
		数据传输完成
		2、version表示当前协议的版本，占用6位，版本范围位[0, 63]，当前版本为0，后续有需求，可以通过修改改变扩展协议
		3、由于每个数据需要和对应的key强绑定，因此发送数据的时候其实需要关联一个key，考虑到key的长度是不固定的，因此这里使用key摘要作为key的标识，
		由于SHA2算法的安全性没有SHA3的安全性好，SHA-3基于Keccak算法设计，经过广泛的安全性评估和审查。同时SHA-3在抗碰撞和抗预像攻击等方面提供
		了更强的安全性保证，因此我更加推荐使用SHA3算法。但是题目中明确要求只能使用go sdk，而sha3实现在golang.org/x/crypto包下面，因此这里
		只能退而求其次，使用SHA2-256算法
		4、data的长度：变长字节，由于题目中没有描述data的应用场景，因此这里假定data可以是无限大的，这样也能更多的适配所有用户。对于data length的每一个
	    字节，每个字节的最高位有特殊函数，不能用于表示数据的长度，0表示当前data的长度只够一个字节，那么只需要读取低7位作为data的长度。1表示当前data
        的长度大于128，因此第二个字节也是data的长度信息，类似的，第二个字节的长度也是这样的函数，一直读取到最后一个字节，这个字节的最高位为0，表示
        这个字节的只有低7位有效。
		5、data表示当前data数据

	// 以下是对于keyDataDone的数据包定义
		++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
		|  dataType(2) | version | SHA3-256 |
		++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
		1、dataType表示当前数据包的类型，占用2位，第一个字节的最高2位表示数据类型 0表示当前数据包是key，1表示当前数据包是data，2标识当前key的
		数据传输完成
		2、version表示当前协议的版本，占用6位，版本范围位[0, 63]，当前版本为0，后续有需求，可以通过修改改变扩展协议
		3、由于每个数据需要和对应的key强绑定，因此发送数据的时候其实需要关联一个key，考虑到key的长度是不固定的，因此这里使用key摘要作为key的标识，
		由于SHA2算法的安全性没有SHA3的安全性好，SHA-3基于Keccak算法设计，经过广泛的安全性评估和审查。同时SHA-3在抗碰撞和抗预像攻击等方面提供
		了更强的安全性保证，因此我更加推荐使用SHA3算法。但是题目中明确要求只能使用go sdk，而sha3实现在golang.org/x/crypto包下面，因此这里
		只能退而求其次，使用SHA2-256算法
*/

type dataType = byte

const (
	keyFrame         dataType = 0      // key数据包
	dataFrame        dataType = 1 << 6 // data数据包
	keyDataDoneFrame dataType = 1 << 7 // 当前key的数据已经传输完成
)

const (
	protoVersion byte = 0
)

var (
	ErrConnClosed                   = fmt.Errorf("connection closed")
	ErrSendDataToClosedStream       = fmt.Errorf("write data in closed stream")
	ErrCloseClosedStream            = fmt.Errorf("try to close closed stream")
	ErrSendDataToKeySendErrorStream = fmt.Errorf("send data to key send error stream")
	ErrCloseKeySendErrorStream      = fmt.Errorf("cose key send error stream")
	ErrReadException                = fmt.Errorf("read exception, should not reach here")
)

func string2bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
func bytes2string(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func sha256Str(k string) ([]byte, [32]byte) {
	h := sha256.New()
	h.Write(string2bytes(k))
	hashSlice := h.Sum(nil)
	var hash [32]byte
	for i := 0; i < 32; i++ {
		hash[0] = hashSlice[0]
	}
	return hashSlice, hash
}
func sha256Byte(k []byte) [32]byte {
	h := sha256.New()
	h.Write(k)
	hashSlice := h.Sum(nil)
	var hash [32]byte
	for i := 0; i < 32; i++ {
		hash[i] = hashSlice[i]
	}
	return hash
}

// Conn 是你需要实现的一种连接类型，它支持下面描述的若干接口；
// 为了实现这些接口，你需要设计一个基于 TCP 的简单协议；
type Conn struct {
	// 底层真正的TCP连接，用发送和接收数据
	conn net.Conn

	// 缓存读取到的key
	recvKey map[[32]byte][]byte
	// 缓存key的io.Reader接口，用于读取真实的数据，直到客户端执行了Close，此时认为这个key的数据接收完成
	recvReaderCh    map[[32]byte]chan []byte // 一个[]byte是一个完整的一个frame的数据
	recvReader      map[[32]byte]io.Reader
	recvReaderClose map[[32]byte]chan struct{}
	recvLock        sync.Mutex
	recvKeyChan     chan [32]byte // 用于生产消费模型，通过channel阻塞调用方

	// 用于发送某个key的数据
	sender map[[32]byte]io.WriteCloser
	// 如果key都发送错了，后面用户还是调用writer发送数据就不应该处理，直接返回错误，因为接收方肯定没有正确接收到key，此时发送这个key的数据时没有意义的
	keySendRes map[[32]byte]error
	// 当前这个key的数据是否已经发送完成，如果用户调用了Close方法，那么说明这个key的数据已经发送完成，所以这里必须为true
	keySendClose map[[32]byte]bool
	sendLock     sync.Mutex
	sendChan     chan *sendReady

	log *log.Logger
}

type KeyReader struct {
	conn       *Conn
	keyHashArr [32]byte
}

func (r *KeyReader) Read(p []byte) (n int, err error) {
	// TODO 如果暂时没有数据，需要阻塞协程
	readCh, ok := r.conn.recvReaderCh[r.keyHashArr]
	if !ok {
		return 0, ErrReadException
	}
	readClose, ok := r.conn.recvReaderClose[r.keyHashArr]
	if !ok {
		return 0, ErrReadException
	}

	select {
	case data := <-readCh: // TODO 如果这里读取的没有接收快，怎么办？
		cnt, err := io.ReadFull(bytes.NewReader(data), p)
		if err != nil {
			return n + cnt, err
		} else {
			n += cnt
		}
	case <-readClose:
		return
	}

	return
}

type sendRes struct {
	n   int
	err error
}

type sendReady struct {
	dataType
	version byte
	length  int64
	hashHex []byte
	data    []byte
	// 用于接收发送数据的返回值，由于发送数据必须一个一个串行发送，因此在一个TCP连接当中同一时刻只可能有一个协程在发送数据，其余想要发送数据的协程
	// 必须阻塞
	sendRes chan sendRes
}

type KeyWriter struct {
	conn       *Conn
	keyHash    []byte
	keyHashArr [32]byte
}

// Write 同一个Key,同一时间只能允许一个协程写入数据，否则数据是乱的，没有办法区分每一帧的数据
func (r *KeyWriter) Write(p []byte) (n int, err error) {
	r.conn.sendLock.Lock()
	defer r.conn.sendLock.Unlock()

	// 1、已经关闭的key stream，不允许再次发送数据。用户必须重新调用Send(key)，重新获取WriterCloser发送数据
	// 2、如果当前这个key找不到，说明这个key被删除，也就是说这个key已经被某个协程Close，此时想要发送数据应该重新Send(key)
	closed, ok := r.conn.keySendClose[r.keyHashArr]
	if !ok || closed {
		return 0, ErrSendDataToClosedStream
	}

	// 1、如果key都发送错了，后面用户还是调用writer发送数据就不应该处理，直接返回错误，因为接收方肯定没有正确接收到key，此时发送这个key的数据是没有意义的
	// 2、如果当前这个key找不到，说明这个key被删除，也就是说这个key已经被某个协程Close，此时想要发送数据应该重新Send(key)
	err, ok = r.conn.keySendRes[r.keyHashArr]
	if !ok || err != nil {
		return 0, ErrSendDataToKeySendErrorStream
	}

	// 封装这个key的数据包，于此同时等待这个包发送的结果
	ready := sendData(r.keyHash, p)
	r.conn.sendChan <- ready
	res := <-ready.sendRes
	return res.n, res.err
}

func (r *KeyWriter) Close() error {
	r.conn.sendLock.Lock()
	defer r.conn.sendLock.Unlock()

	closed, ok := r.conn.keySendClose[r.keyHashArr]
	if !ok || closed {
		return ErrCloseClosedStream
	}

	// 如果Key发送的时候已经发送错误了，此时关闭的动作也不应该支持。
	// 原因是因为客户端拿到这个数据也不能做什么，因为key都没有收到，你告诉我这个key的数据已经发送完成了，也是没有意义的。
	err, ok := r.conn.keySendRes[r.keyHashArr]
	if !ok || err != nil {
		return ErrCloseKeySendErrorStream
	}

	// 当用户调用Close时，表示这个Key的数据发送完成
	// 这里应该考虑如下问题：
	// Q: 1、对于用户已经调用Close关闭某个key的发送数据，用户再次调用如何处理？
	// A: 对于这种情况，根据题目的定义，调用Close就表示用户已经发送完这个key的数据了，用户如果再次发送这个key的数据，协议层应该报错，因为用户已经
	// 违反了接口的约定。既然你还要发送数据，为什么要提前调用这个key的Close方法。
	// Q: 2、如果用户真的提前关闭了Close函数，用户还想再次发送同一个key的数据。应该怎么办？
	// A： 对于这种情况，用户应该再次调用Send(key)方法，重新获取一个WriterCloser，对于协议层来说，这个key和上一次的key没有任何关系，只不过key
	// 的值时相同的，此时协议层应该把当前key当成一个新的stream，进行传输。
	// Q：3、用户调用Close之后，协议曾是否应该调用底层TCP连接的Close方法呢？
	// A：答案很明显，用户调用Close方法，仅仅表明用户对于这个key的数据已经全部发送完成，但是对于其它的key的数据，接口并没有这么约定。所以这里并不
	// 可以关闭底层TCP连接
	// Q: 4、用户调用Close方法之后，协议层因该处理什么动作？
	// A：一：关闭某个key之后，不能再次关闭，否则直接报错。 二、关闭之后不能再次发送数据
	ready := sendKeyDataDone(r.keyHash)
	r.conn.sendChan <- ready
	res := <-ready.sendRes

	// 关闭key stream之后，需要删除key stream的状态
	delete(r.conn.keySendRes, r.keyHashArr)
	delete(r.conn.keySendClose, r.keyHashArr)
	return res.err
}

func sendKey(key string) *sendReady {
	data := string2bytes(key)
	return &sendReady{
		dataType: keyFrame,
		version:  protoVersion,
		length:   int64(len(data)),
		data:     data,
		sendRes:  make(chan sendRes, 1),
	}
}

func sendData(keyHash []byte, data []byte) *sendReady {
	return &sendReady{
		dataType: dataFrame,
		version:  protoVersion,
		hashHex:  keyHash,
		length:   int64(len(data)),
		data:     data,
		sendRes:  make(chan sendRes, 1),
	}
}

func sendKeyDataDone(hashHex []byte) *sendReady {
	return &sendReady{
		dataType: keyDataDoneFrame,
		version:  protoVersion,
		hashHex:  hashHex,
		sendRes:  make(chan sendRes, 1),
	}
}

// Send 传入一个 key 表示发送者将要传输的数据对应的标识；
// 返回 writer 可供发送者分多次写入大量该 key 对应的数据；
// 当发送者已将该 key 对应的所有数据写入后，调用 writer.Close 告知接收者：该 key 的数据已经完全写入；
func (conn *Conn) Send(key string) (writer io.WriteCloser, err error) {
	keyHashSlice, keyHashArr := sha256Str(key)

	conn.sendLock.Lock()
	defer conn.sendLock.Unlock()

	// Q:如果有两个协程同时调用Send(key)想要对于同一个key发送数据如何处理？
	// A: 这种情况必须支持，但是同一个Key应该只发送一次，除非这个Key被Close了，才允许发送第二次。否则认为key已经发送成功了，没有必要再次发送
	// 毕竟从接口定义上来看，相同的key，用户就是想要发送这个key的数据，只不过用户选择了并发的方式来发送数据。这种情况下，每个帧的顺序肯定时不确定
	// 的，这种情况下在协议层没有办法控制，用户自己都不清楚哪个协程先发送了数据，哪个后发送了数据，所以协议曾肯定是无能为力的。但是，协议层能够保证
	// 的是，每个协程发送的数据在TCP连接上肯定是一个完整的数据包，这个是协议层必须要保证的
	var res sendRes
	// 如果这个key的发送有问题，那么允许第二个协程继续发送。如果状态不存在，说明要么这个key从来都没有发送数据，要么这个key已经发送完成了，此时也
	// 可以重新启用这个key,再次发送数据
	if err, ok := conn.keySendRes[keyHashArr]; !ok || err != nil {
		// 发送key的数据
		ready := sendKey(key)
		conn.sendChan <- ready
		res = <-ready.sendRes // 获取Key发送的结果，key发送成功与否，决定后续这个key的数据能否正常发送

		// 记录这个key的发送状态
		conn.keySendRes[keyHashArr] = res.err
		conn.keySendClose[keyHashArr] = false
	}

	// 实现自己的WriterClose  当用户调用Close方法时，需要发送key数据已经发送完的消息
	wr := &KeyWriter{keyHash: keyHashSlice, keyHashArr: keyHashArr, conn: conn}
	conn.sender[keyHashArr] = wr
	return wr, res.err
}

// Receive 返回一个 key 表示接收者将要接收到的数据对应的标识；
// 返回的 reader 可供接收者多次读取该 key 对应的数据；
// 当 reader 返回 io.EOF 错误时，表示接收者已经完整接收该 key 对应的数据；
func (conn *Conn) Receive() (key string, reader io.Reader, err error) {
	// 阻塞，只有收到了任何一个key，才返回reader，让外部可以正常读取数据
	keyHash, ok := <-conn.recvKeyChan
	if !ok {
		return "", nil, ErrConnClosed
	}

	conn.recvLock.Lock()
	defer conn.recvLock.Unlock()

	k := conn.recvKey[keyHash]
	r := conn.recvReader[keyHash]
	return bytes2string(k), r, nil
}

// Close 关闭你实现的连接对象及其底层的 TCP 连接
func (conn *Conn) Close() {
	// TODO 关闭底层TCP连接之前需要做一些清理动作，譬如关闭一些生产消费模型
	if conn.conn != nil {
		conn.conn.Close()
	}
}

// NewConn 从一个 TCP 连接得到一个你实现的连接对象
func NewConn(conn net.Conn) *Conn {
	cn := &Conn{
		conn:            conn,
		recvKey:         make(map[[32]byte][]byte),
		recvReaderCh:    make(map[[32]byte]chan []byte, 256),
		recvReader:      make(map[[32]byte]io.Reader),
		recvReaderClose: make(map[[32]byte]chan struct{}, 256),
		recvKeyChan:     make(chan [32]byte),

		sender:       make(map[[32]byte]io.WriteCloser),
		keySendRes:   make(map[[32]byte]error),
		keySendClose: make(map[[32]byte]bool),
		sendChan:     make(chan *sendReady, 256),

		log: log.Default(),
	}

	cn.log.Printf("init conn")

	// 启动读协程
	cn.log.Println("start receive coroutine")
	go cn.recv()
	// 启动写协程
	cn.log.Println("start send coroutine")
	go cn.send()
	return cn
}

func (conn *Conn) recv() {
	if err := conn.recvLoop(); err != nil {
		// 读取发生错误，直接关闭连接
		conn.exitErr(err)
	}
}

func (conn *Conn) recvLoop() error {
	for {
		// 先读取一个字节的数据，看看是那种数据类型
		buf := make([]byte, 1)
		_, err := io.ReadFull(conn.conn, buf)
		if err != nil {
			// TODO 读取发生错误，直接退出连接
			conn.exitErr(err)
			return err
		}
		conn.log.Printf("[recv] recv first byte 0x%x\n", buf[0])

		dType := buf[0] & 0xC0
		switch dType {
		case keyFrame:
			length, err := conn.readLength()
			if err != nil {
				conn.exitErr(err)
				return err
			}
			conn.log.Printf("[recv keyFrame] recv key length=[%d]\n", length)

			// 接下来读取剩余的数据即可
			keyReader := io.LimitReader(conn.conn, length)
			key, err := io.ReadAll(keyReader)
			if err != nil {
				conn.exitErr(err)
				return err
			}
			keyHash := sha256Byte(key)
			conn.log.Printf("[recv keyFrame] recv key=[%s]， keyHash=%v, init recv metadata\n", key, keyHash)

			conn.recvLock.Lock()
			conn.recvKey[keyHash] = key
			conn.recvReader[keyHash] = &KeyReader{conn: conn, keyHashArr: keyHash}
			conn.recvReaderCh[keyHash] = make(chan []byte, 1)
			conn.recvReaderClose[keyHash] = make(chan struct{})
			conn.recvKeyChan <- keyHash
			conn.recvLock.Unlock()

		case dataFrame:
			// 必须要先接收到key，否则这个数据没有意义
			var keyHash [32]byte
			if _, err = conn.conn.Read(keyHash[:]); err != nil {
				conn.exitErr(err)
				return err
			}
			conn.log.Printf("[recv dataFrame] recv data key hash=[%x]\n", keyHash)

			length, err := conn.readLength()
			if err != nil {
				conn.exitErr(err)
				return err
			}
			conn.log.Printf("[recv dataFrame] recv data length=[%d]\n", length)

			dataReader := io.LimitReader(conn.conn, length)
			data, err := io.ReadAll(dataReader)
			if err != nil {
				conn.exitErr(err)
				return err
			}
			conn.log.Printf("[recv dataFrame] recv data=[%v], sned data to reader channel\n", data)

			rdCh, ok := conn.recvReaderCh[keyHash]
			if !ok || rdCh == nil {
				// TODO 打印错误日志，丢弃当前的数据包
				conn.log.Printf("[recv dataFrame] recv data=[%v], not found reader channel\n", data)
				continue
			}
			rdCh <- data

		case keyDataDoneFrame:
			// TODO 必须要先接收到key，否则这个数据没有意义
			var keyHash [32]byte
			if _, err = conn.conn.Read(keyHash[:]); err != nil {
				// TODO 读取发生错误，直接退出连接
				conn.exitErr(err)
				return err
			}
			conn.log.Printf("[recv keyDataDoneFrame] recv keyHash=[%v] data done signal\n", keyHash)

			// 通知消费数据的协程，这个摘要对应的数据已经发送完成
			closeCh, ok := conn.recvReaderClose[keyHash]
			if !ok || closeCh == nil {
				// TODO 打印错误日志，丢弃当前的数据包
				conn.log.Printf("[recv keyDataDoneFrame] recv keyHash=[%v] , not found reader channel\n", keyHash)
				continue
			}

			// 关闭这个key，通知所有读取这个key的协程退出
			conn.recvLock.Lock()
			close(closeCh)
			delete(conn.recvKey, keyHash)
			delete(conn.recvReaderCh, keyHash)
			delete(conn.recvReader, keyHash)
			delete(conn.recvReaderClose, keyHash)
			conn.recvLock.Unlock()

		default:
			// 不支持这种格式的数据包，直接关闭连接
			conn.log.Printf("[recv] not support data \n")
			conn.exitErr(err)
			return err
		}
	}
}

func (conn *Conn) readLength() (int64, error) {
	// 读取length长度
	buf := make([]byte, 1)
	var lenArr []byte
	hasMoreLen := true
	for hasMoreLen {
		_, err := io.ReadFull(conn.conn, buf)
		if err != nil {
			return 0, err
		}
		lenArr = append(lenArr, buf[0]&0x7F)
		hasMoreLen = buf[0]&0x80 == 0x80
	}
	// 合成length长度，假设这里长度不超过2^63次方，毕竟一般也没有这么大的数据
	// TODO 由于时间有限，这里暂时不处理2^63次方以上的数据长度，后续使用big.NewInt来代替
	var length int64
	for i, l := range lenArr {
		length |= int64(l) << (i * 7)
	}

	return length, nil
}

func (conn *Conn) exitErr(err error) {
	// 清理一些元数据，然后关闭TCP连接，最后退出
}

func (conn *Conn) send() {
	for {
		// TODO 可能需要读取停止通道
		select {
		case frame := <-conn.sendChan:
			// 写入版本和版本
			typeAndVersion := frame.dataType | frame.version
			_, err := conn.conn.Write([]byte{typeAndVersion}) // 这里的n时header发送成功的长度，并非用户期望的数据，这个时候还没有发送数据
			frame.sendRes <- sendRes{n: 0, err: err}
			if err != nil {
				continue
			}
			conn.log.Printf("[send] send type and version = %x\n", typeAndVersion)

			switch frame.dataType {
			case keyFrame: // 构造发送key的数据帧
				// 写入数据长度
				lengthByte := genLengthByte(frame.length)
				_, err = conn.conn.Write(lengthByte)
				frame.sendRes <- sendRes{n: 0, err: err}
				if err != nil {
					continue
				}
				conn.log.Printf("[send keyFrame] send key legnth=%v， lengthByte=%v\n", frame.length, lengthByte)

				// 写入真实数据
				n, err := conn.conn.Write(frame.data)
				frame.sendRes <- sendRes{n: n, err: err}
				if err != nil {
					continue
				}
				conn.log.Printf("[send keyFrame] send key length=%v， data=%v\n", frame.length, frame.data)

			case dataFrame: // 构造发送data的数据帧
				// 写入key的SHA256摘要
				_, err = conn.conn.Write(frame.hashHex)
				frame.sendRes <- sendRes{n: 0, err: err}
				if err != nil {
					continue
				}
				conn.log.Printf("[send dataFrame] send key hash=%v\n", frame.hashHex)

				// 写入数据长度
				lengthByte := genLengthByte(frame.length)
				_, err = conn.conn.Write(lengthByte)
				frame.sendRes <- sendRes{n: 0, err: err}
				if err != nil {
					continue
				}
				conn.log.Printf("[send dataFrame] send data length=%v， data=%v\n", frame.length, frame.data)

				// 写入真实数据
				n, err := conn.conn.Write(frame.data)
				frame.sendRes <- sendRes{n: n, err: err}
				if err != nil {
					continue
				}
				conn.log.Printf("[send dataFrame] send data\n")

			case keyDataDoneFrame: // 构造发送key的数据已经发送完成的数据帧
				// 写入key的SHA256摘要
				_, err = conn.conn.Write(frame.hashHex)
				frame.sendRes <- sendRes{n: 0, err: err}
				if err != nil {
					continue
				}
				conn.log.Printf("[send keyDataDoneFrame] send key=%v data done\n", frame.hashHex)
			}
		}
	}
}

func genLengthByte(length int64) []byte {
	var l []byte
	for length != 0 {
		tmpLen := byte(length & 0x7f)
		length >>= 7
		if length == 0 {
			l = append(l, tmpLen)
		} else {
			l = append(l, tmpLen|0x80) // 把最高位置为1
		}
	}
	return l
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
