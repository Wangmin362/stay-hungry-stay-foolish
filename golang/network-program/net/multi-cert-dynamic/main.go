package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"sync"
	"time"
)

type CertManager struct {
	sync.RWMutex
	certificates map[string]tls.Certificate
}

func NewCertManager() *CertManager {
	return &CertManager{
		certificates: make(map[string]tls.Certificate),
	}
}

func (cm *CertManager) AddCert(domain string, cert tls.Certificate) {
	cm.Lock()
	defer cm.Unlock()
	cm.certificates[domain] = cert
}

func (cm *CertManager) RemoveCert(domain string) {
	cm.Lock()
	defer cm.Unlock()
	delete(cm.certificates, domain)
}

func (cm *CertManager) GetCert(domain string) (tls.Certificate, bool) {
	cm.RLock()
	defer cm.RUnlock()
	cert, ok := cm.certificates[domain]
	return cert, ok
}

func main() {
	certManager := NewCertManager()

	// 创建HTTPS服务器
	server := &http.Server{
		Addr: ":8056",
		TLSConfig: &tls.Config{
			GetCertificate: func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
				domain := hello.ServerName
				cert, ok := certManager.GetCert(domain)
				if !ok {
					return nil, nil
				}
				return &cert, nil
			},
		},
	}

	// 处理请求
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, HTTPS World!"))
	})

	log.Println("Server started at https://localhost:8056")

	// 启动服务器
	go func() {
		err := server.ListenAndServeTLS("", "")
		if err != nil {
			log.Fatal("ListenAndServeTLS: ", err)
		}
	}()

	// 通过命令行或其他方式动态添加或删除证书
	// 示例代码中仅供演示，实际应用中可能需要更加安全的方式来管理证书
	// 注意：在生产环境中使用时，请谨慎处理证书的增加和删除操作，确保安全性和可靠性
	// 示例中使用定时器模拟动态添加和删除证书的操作
	go func() {
		ticker := time.NewTicker(time.Minute * 5)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				// 示例中定时每5分钟添加和删除一次证书
				cert, err := tls.LoadX509KeyPair("new_cert.crt", "new_cert.key") // 替换为你的证书文件路径
				if err != nil {
					log.Println("Error loading certificate:", err)
					continue
				}
				certManager.AddCert("example.com", cert)
				log.Println("Added certificate for example.com")
			}
		}
	}()

	// 阻塞主goroutine
	select {}
}
