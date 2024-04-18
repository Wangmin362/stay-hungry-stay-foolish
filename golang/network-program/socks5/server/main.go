package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func handleClient(client net.Conn) {
	defer client.Close()

	// Version + Number of authentication methods
	buffer := make([]byte, 2)
	// 上来先读取两个字节
	if _, err := io.ReadFull(client, buffer); err != nil {
		log.Println("Failed to read version and method:", err)
		return
	}

	version, nMethods := buffer[0], buffer[1]
	if version != 5 { // 当前SOCKS代理服务只支持SOCKS5版本的协议
		log.Println("Unsupported SOCKS version:", version)
		return
	}

	// 读取客户端支持的认证方法
	methods := make([]byte, nMethods)
	if _, err := io.ReadFull(client, methods); err != nil {
		log.Println("Failed to read methods:", err)
		return
	}

	// 直接返回客户端这里只支持SOCKS5代理，并且不进行认证
	if _, err := client.Write([]byte{0x05, 0x00}); err != nil {
		log.Println("Failed to write no auth required response:", err)
		return
	}

	// Read the request
	request := make([]byte, 4)
	if _, err := io.ReadFull(client, request); err != nil {
		log.Println("Failed to read request:", err)
		return
	}

	// Only support CONNECT command
	if request[1] != 0x01 { // 这里只支持TCP SOCKS代理，因此只支持CONNECT
		log.Println("Unsupported command:", request[1])
		client.Write([]byte{0x05, 0x07, 0x00, 0x01})
		return
	}

	// Read the target address
	// 读取当前客户端请求的目标代理地址的类型，看时IPv4地址还是IPv6地址，也有可能是域名
	addrType := request[3]
	address := ""
	switch addrType {
	case 0x01: // IP V4 address
		addr := make([]byte, 4)
		io.ReadFull(client, addr)
		address = fmt.Sprintf("%d.%d.%d.%d", addr[0], addr[1], addr[2], addr[3])
	case 0x03: // Domain name
		len := make([]byte, 1)
		io.ReadFull(client, len) // 先读取一个字节，获取域名长度
		addr := make([]byte, len[0])
		io.ReadFull(client, addr)
		address = string(addr)
	case 0x04: // IP V6 address
		addr := make([]byte, 16)
		io.ReadFull(client, addr)
		address = net.IP(addr).String()
	default:
		log.Println("Unsupported address type:", addrType)
		return
	}

	// Read the target port
	port := make([]byte, 2) // 读取端口
	io.ReadFull(client, port)
	destPort := int(port[0])<<8 + int(port[1])

	// 和目标代理地址建立连接
	dest := fmt.Sprintf("%s:%d", address, destPort)
	server, err := net.Dial("tcp", dest)
	if err != nil {
		log.Println("Failed to connect to target:", err)
		// 如果连接失败，就通过0x01告诉代理客户端失败
		client.Write([]byte{0x05, 0x05, 0x00, 0x01}) // Connection refused
		return
	}
	defer server.Close()

	// Send success response
	// 否则，直接相应成功 TODO 这里应该返回这是的IP地址和端口
	client.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})

	// Start forwarding
	go io.Copy(server, client)
	io.Copy(client, server)
}

func main() {
	listener, err := net.Listen("tcp", ":10089")
	if err != nil {
		log.Fatal("Failed to listen on port 1080:", err)
	}
	defer listener.Close()

	log.Println("SOCKS5 proxy server listening on port 1080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection:", err)
			continue
		}
		go handleClient(conn)
	}
}
