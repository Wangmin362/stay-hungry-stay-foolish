package main

import (
	"fmt"
	"github.com/golang/demo/oss/aliyun/app/sync"
	"os"
	"time"
)

func main() {
	home, err := sync.GetEnvVar("USERPROFILE")
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	dir := fmt.Sprintf("%s\\Documents\\test", home)
	if err = os.MkdirAll(dir, os.ModePerm); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	for {
		file := fmt.Sprintf("%s\\%s.txt", dir, time.Now().Format("2006.01.02_15:04:05"))
		_, err = os.Create(file)
		if err != nil {
			os.Exit(1)
		}
		time.Sleep(5 * time.Second)
	}
}
