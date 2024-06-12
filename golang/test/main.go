package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 检查 Content-Type
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/octet-stream" {
		http.Error(w, "Unsupported media type", http.StatusUnsupportedMediaType)
		return
	}

	// 检查 Content-Encoding
	contentEncoding := r.Header.Get("Content-Encoding")
	if contentEncoding != "gzip" {
		http.Error(w, "Unsupported content encoding", http.StatusUnsupportedMediaType)
		return
	}

	// 读取请求体
	//contentLength := r.ContentLength
	//body := make([]byte, contentLength)
	//_, err := r.Body.Read(body)
	//if err != nil {
	//	http.Error(w, "Failed to read request body", http.StatusInternalServerError)
	//	return
	//}
	all, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to create gzip reader", http.StatusInternalServerError)
		return
	}

	// 解压缩文件
	gzipReader, err := gzip.NewReader(bytes.NewReader(all))
	if err != nil {
		http.Error(w, "Failed to create gzip reader", http.StatusInternalServerError)
		return
	}
	defer gzipReader.Close()

	uncompressedBody, err := ioutil.ReadAll(gzipReader)
	if err != nil {
		http.Error(w, "Failed to decompress gzip body", http.StatusInternalServerError)
		return
	}

	// 打印解压后的内容
	fmt.Println("Received content:")
	fmt.Println(string(uncompressedBody))
	// 处理上传的文件内容
	// 在这里您可以将文件保存到服务器上或执行其他逻辑

	// 返回响应
	fmt.Fprintf(w, "File uploaded successfully!")
}

func main() {
	http.HandleFunc("/sps/v1/ztna/privateApp/discovery/discovery-2024-06-05-15-03-57.gz", uploadFileHandler)
	http.ListenAndServe(":8080", nil)
}
