package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"
)

func main() {
	file := "D:\\Skyguard\\discovery-2024-06-05-15-03-57"
	outputFile, err := os.Create(file + ".gz")
	if err != nil {
		log.Panicln(err)
	}

	// 创建gzip写入器
	gzipWriter := gzip.NewWriter(outputFile)
	defer gzipWriter.Close()

	// 复制原始文件内容到gzip写入器
	open, err := os.Open(file)
	if err != nil {
		log.Println(err)
	}
	_, err = io.Copy(gzipWriter, open)
	if err != nil {
		log.Println(err)
	}
	defer open.Close()
	defer outputFile.Close()
}
