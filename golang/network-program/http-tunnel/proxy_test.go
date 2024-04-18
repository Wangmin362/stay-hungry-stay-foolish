package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"
)

var proxyUrl = "http://127.0.0.1:7722"

var webUrl = "https://baidu.com"

// go的http代码会自动发送 CONNECT请求
func TestHttpsProxy(t *testing.T) {
	proxyURL, err := url.Parse(proxyUrl)
	if err != nil {
		t.Fatal(err)
	}
	transport := http.Transport{
		Proxy:           http.ProxyURL(proxyURL),
		TLSClientConfig: &tls.Config{},
	}
	client := http.Client{
		Transport: &transport,
		Timeout:   5 * time.Second,
	}

	resp, err := client.Get(webUrl)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode, string(body))
}
