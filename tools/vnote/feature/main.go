package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	path := "D:/Notebook/Vnote/Blog/10.Kubernetes/06.kube-controller-manager/00.组件"
	ReadDir(path)
}

func ReadDir(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if dir == path {
			return nil
		}

		if !strings.Contains(path, "源码") {
			return nil
		}

		if strings.Contains(path, "总体概览") {
			return nil
		}

		index := strings.Index(info.Name(), "——")
		name := info.Name()[index+6:]
		name = strings.ReplaceAll(name, ".md", "")

		fmt.Println("- [ ] ", name)

		return nil
	})
}

// ReadLinesV3 reads all lines of the file.
func ReadLinesV3(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		// ReadLine is a low-level line-reading primitive.
		// Most callers should use ReadBytes('\n') or ReadString('\n') instead or use a Scanner.
		bytes, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		idx := strings.Index(string(bytes), "=")
		name := string(bytes)[idx+1:]
		name = strings.ReplaceAll(name, "\"", "")
		fmt.Println(fmt.Sprintf("- [ ] %s", name))
	}
	return nil
}
