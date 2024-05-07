package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

// 修改字体颜色的格式
func main() {
	path := "tools/vnote/param.txt"
	if err := ParseParam(path); err != nil {
		log.Fatal(err)
	}
}

const (
	pattern string = `    (--[a-zA-Z-\d]+?) `
)

func ParseParam(path string) error {
	rawData, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	re := regexp.MustCompile(pattern)
	match := re.FindAllSubmatch(rawData, -1)
	for _, group := range match {
		raw := string(group[1])
		if raw == "--config" {
			continue
		}
		fmt.Printf("--[ ] %s\n", raw)
	}

	return nil
}
