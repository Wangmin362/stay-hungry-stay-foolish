package sync

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestConvertWindowDirToLinuxDir(t *testing.T) {
	path := "D:\\Notebook\\Vnote\\vx_recycle_bin\\20240319\\file3.md\\file3.md"
	res := ConvertWindowDirToLinuxDir(path)
	if "D:/Notebook/Vnote/vx_recycle_bin/20240319/file3.md/file3.md" != res {
		t.Fatal("error")
	}

	path = "vx_recycle_bin\\20240319\\file3.md\\file3.md"
	res = ConvertWindowDirToLinuxDir(path)
	if "vx_recycle_bin/20240319/file3.md/file3.md" != res {
		t.Fatal("error")
	}
}

func TestWindowPath(t *testing.T) {
	path := "D:\\Notebook\\Vnote\\vx_recycle_bin\\20240319/file3.md/file3.md"
	if _, err := os.Stat(path); err != nil {
		t.Fatal(err)
	}
}

func TestBasePath(t *testing.T) {
	path := "vx_recycle_bin/20240319/abcd/file3.md"
	fmt.Println(filepath.Dir(path))
	fmt.Println(filepath.Base(path))

	path = "https://blog.csdn.net/jiang_xinxing/article/details/71057086/vx_images/7856451324545343.png"
	fmt.Println(filepath.Dir(path))
	fmt.Println(filepath.Base(path))
}
