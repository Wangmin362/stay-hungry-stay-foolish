package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	path := "QA.MD"
	ReadLinesV3(path)
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
