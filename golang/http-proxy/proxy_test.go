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

var webUrl = "http://172.30.3.206"

// HTTPS请求无法代理
//var webUrl = "https://baidu.com"

// 直接请求，不使用代理
func TestJustDoReq(t *testing.T) {
	resp, err := http.Get(webUrl)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// 通过代理请求，这种方式只能代理HTTP请求，无法代理HTTPS请求，因为无法建立TLS握手连接
/*
// 直接连接
GET / HTTP/1.1
Host: staight.github.io
Connection: keep-alive

// http 代理连接
GET http://staight.github.io/ HTTP/1.1
Host: staight.github.io
Proxy-Connection: keep-alive
*/
// TODO 思考，为什么HTTP代理的包和普通HTTP请求包会不一样？
func TestSecond(t *testing.T) {
	proxy, _ := url.Parse(proxyUrl)
	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 5,
	}

	resp, err := client.Get(webUrl)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode, string(body))
}
