package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	dir := "app-logs"
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	for {
		file := fmt.Sprintf(".\\%s\\app测试_%s.txt", dir, time.Now().Format("2006.01.02_15_04_05"))
		_, err := os.Create(file)
		if err != nil {
			os.Exit(1)
		}
		time.Sleep(3 * time.Second)
	}
}
