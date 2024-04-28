package main

import (
	"bufio"
	"fmt"
	"github.com/golang/demo/tools/vnote"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	qaPath := "D:/Notebook/Vnote/Blog/10.Kubernetes/特性开关/QA.md"
	vnotePath := "D:/Notebook/Vnote/Blog/10.Kubernetes/特性开关/vx.json"
	err := GenerateDir(qaPath, vnotePath)
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateDir(qaPath, vnotePath string) error {
	vx, err := vnote.UnMarshal(vnotePath)
	if err != nil {
		return err
	}
	fmt.Println(vx)

	f, err := os.Open(qaPath)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	idx := 0
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

		line := string(bytes)
		if !strings.HasPrefix(line, "- [ ] ") {
			continue
		}

		line = line[7:]
		dir := filepath.Dir(qaPath)
		line = fmt.Sprintf("%03d.%s", idx, line)
		mkDir := filepath.Join(dir, line)
		fmt.Printf("%s -> %s\n", line, mkDir)
		if vx.Folders == nil {
			vx.Folders = make([]*vnote.Folder, 0)
		}
		if err = os.MkdirAll(mkDir, os.ModePerm); err != nil {
			return err
		}

		vx.Folders = append(vx.Folders, &vnote.Folder{Name: line})

		idx++
	}

	if err = vx.PersistentMarshal(vnotePath); err != nil {
		return err
	}

	return nil
}
