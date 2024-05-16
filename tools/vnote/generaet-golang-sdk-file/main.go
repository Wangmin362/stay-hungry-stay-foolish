package main

import (
	"fmt"
	"github.com/golang/demo/tools/vnote"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	file := "D:/Notebook/Vnote/Blog/06.Golang/08.标准库/02.bufio/00.QA.md"

	err := GenerateFile(file, true)
	if err != nil {
		log.Fatal(err)
	}
	err = GenerateFile(file, false)
	if err != nil {
		log.Fatal(err)
	}
}

const (
	pattern string = `- \[ \] (.+)`
)

func GenerateFile(file string, isFunction bool) error {
	rawData, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	index := strings.Index(string(rawData), "函数")
	data := string(rawData)[index:]
	if !isFunction {
		data = string(rawData)[:index]
	}

	re := regexp.MustCompile(pattern)
	match := re.FindAllSubmatch([]byte(data), -1)
	var names []string
	for _, group := range match {
		raw := string(group[1])
		fmt.Printf("%s\n", raw)
		names = append(names, raw)
	}

	dir := filepath.Dir(file)
	var vxPath string
	if isFunction {
		vxPath = filepath.Join(dir, "函数", "vx.json")
	} else {
		vxPath = filepath.Join(dir, "类型", "vx.json")
	}

	childVX, err := vnote.UnMarshal(vxPath)
	if err != nil {
		return err
	}

	for _, name := range names {
		err := os.WriteFile(filepath.Join(filepath.Dir(vxPath), fmt.Sprintf("%s.md", name)), []byte{}, os.ModePerm)
		if err != nil {
			log.Printf("%+v", err)
			continue
		}
		childVX.Files = append(childVX.Files, &vnote.File{
			CreatedTime:  "2024-04-18T07:24:10Z",
			Id:           "6076",
			ModifiedTime: "2024-04-18T07:24:10Z",
			Name:         fmt.Sprintf("%s.md", name),
			Signature:    "177807976771",
		})

	}

	if err = childVX.PersistentMarshal(vxPath); err != nil {
		return err
	}

	return nil
}
