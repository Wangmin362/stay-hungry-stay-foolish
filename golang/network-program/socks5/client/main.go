package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"net/http"
)

// dialSocks5 creates a network connection through SOCKS5 proxy
func dialSocks5(proxyAddr string, targetAddr string) (net.Conn, error) {
	proxyConn, err := net.Dial("tcp", proxyAddr)
	if err != nil {
		return nil, err
	}

	// 协商SOCKS5认证方式，这里要求SOCKS5代理服务器不进行认证
	_, err = proxyConn.Write([]byte{0x05, 0x01, 0x00}) // SOCKS5, 1 method, 0x00: no authentication
	if err != nil {
		return nil, err
	}

	response := make([]byte, 2)
	// 读取服务端响应
	if _, err = io.ReadFull(proxyConn, response); err != nil {
		return nil, err
	}

	// 服务端必须支持SOCKS5协议，并且支持不进行认证的方式
	if response[0] != 0x05 || response[1] != 0x00 {
		return nil, err
	}

	// 客户端向服务端发送代理请求，要求服务端和目标服务器建立好TCP连接
	// 0x05表示使用SOCKS5协议
	// 0x01表示当前使用Connect指令
	// 0x00为SOCKS5协议保留字段，没有实际含义
	// 0x03表示当前客户端想要连接的服务时以域名的方式
	// byte(len(targetAddr))表示当前域名的长度
	// targetAddr...表示当前的域名
	// 0x0080表示当前需要连接的端口
	req := []byte{0x05, 0x01, 0x00, 0x03, byte(len(targetAddr))}
	req = append(req, targetAddr...)
	req = append(req, 0x00, 80) // port 80 in network byte order

	if _, err = proxyConn.Write(req); err != nil {
		return nil, err
	}

	// 读取服务端的响应，获取前三个字节
	if _, err = io.ReadFull(proxyConn, response[:3]); err != nil {
		return nil, err
	}

	// 第二个字节必须是0x00，因为0x00表示SOCKS5代理服务器和目标服务器成功建立连接
	if response[1] != 0x00 {
		return nil, err // Connection failed
	}

	// TODO 下面的目标地址和端口虽然读取出来了，但是都没有啥用处，可以不读取么？
	if response[2] == 0x01 { // IPv4
		buffer := make([]byte, 4+2) // IPv4 address + port
		if _, err = io.ReadFull(proxyConn, buffer); err != nil {
			return nil, err
		}
	} else if response[2] == 0x03 { // Domain name
		length := make([]byte, 1) // 域名
		if _, err = io.ReadFull(proxyConn, length); err != nil {
			return nil, err
		}
		buffer := make([]byte, int(length[0])+2) // 域名 + 端口
		if _, err = io.ReadFull(proxyConn, buffer); err != nil {
			return nil, err
		}
	} else if response[2] == 0x04 { // IPv6
		buffer := make([]byte, 16+2) // IPv6 address + port
		if _, err = io.ReadFull(proxyConn, buffer); err != nil {
			return nil, err
		}
	}

	return proxyConn, nil
}

func main() {
	// Proxy server address
	proxyAddr := "localhost:10089"

	// Target HTTP server address
	targetAddr := "example.com"

	proxyConn, err := dialSocks5(proxyAddr, targetAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer proxyConn.Close()

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "http://"+targetAddr, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Send the HTTP request
	if err := req.Write(proxyConn); err != nil {
		log.Fatal(err)
	}

	// Read the response
	resp, err := http.ReadResponse(bufio.NewReader(proxyConn), req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Print the HTTP response status
	log.Println("Response status:", resp.Status)

	// Optionally, print the response body
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("Response body:", string(body))
}
