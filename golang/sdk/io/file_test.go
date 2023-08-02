package main

import (
	"fmt"
	"os"
	"testing"
)

func TestReadAt(t *testing.T) {
	myfile, err := os.Open("D:\\Project\\github\\stay-hungry-stay-foolish\\golang\\sdk\\io\\my-file")
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 2)
	off := 0
	for {
		n, err := myfile.ReadAt(buf, int64(off))
		if err != nil {
			panic(err)
		}
		if n <= 0 {
			break
		}
		off += n
		fmt.Println(string(buf))
	}

}
