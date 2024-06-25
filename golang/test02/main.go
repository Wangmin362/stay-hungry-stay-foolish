package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// 读取待压缩的文件内容
	fileContent, err := ioutil.ReadFile("D:\\Skyguard\\2024-06-05-20-20-29")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	// 创建一个 bytes.Buffer 来保存压缩后的内容
	var compressedContent bytes.Buffer

	// 创建一个 gzip.Writer 来压缩文件内容
	gzipWriter := gzip.NewWriter(&compressedContent)
	defer gzipWriter.Close()

	// 将文件内容写入到 gzipWriter 中
	_, err = gzipWriter.Write(fileContent)
	if err != nil {
		log.Fatal("Error compressing file:", err)
	}

	// 完成压缩
	err = gzipWriter.Close()
	if err != nil {
		log.Fatal("Error closing gzip writer:", err)
	}

	// 将压缩后的内容写入到一个新文件中
	err = ioutil.WriteFile("compressed.gz", compressedContent.Bytes(), os.ModePerm)
	if err != nil {
		log.Fatal("Error writing compressed file:", err)
	}

	fmt.Println("File compressed successfully and saved as 'compressed.gz'.")
}
