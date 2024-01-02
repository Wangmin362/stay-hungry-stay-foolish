package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"testing"
)

// 查看证书
func TestViewCertificate(t *testing.T) {
	//读取证书并解码
	pemTmp, err := os.ReadFile("server.crt")
	if err != nil {
		fmt.Println(err)
		return
	}
	certBlock, restBlock := pem.Decode(pemTmp)
	if certBlock == nil {
		fmt.Println(err)
		return
	}
	//可从剩余判断是否有证书链等，继续解析
	fmt.Println(restBlock)
	//证书解析
	certBody, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	//可以根据证书结构解析
	fmt.Println(certBody.SignatureAlgorithm)
	fmt.Println(certBody.PublicKeyAlgorithm)
}

// 创建证书

// 验证证书
