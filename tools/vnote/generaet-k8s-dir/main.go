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
	qaPath := "D:/Notebook/Vnote/Blog/10.Kubernetes/05.KubeAPIServer/07.准入控制/00.内建准入控制器/00.QA.md"
	vnotePath := "D:/Notebook/Vnote/Blog/10.Kubernetes/05.KubeAPIServer/07.准入控制/00.内建准入控制器/vx.json"
	err := GenerateDir(qaPath, vnotePath)
	if err != nil {
		log.Fatal(err)
	}
}

var pattern []string = []string{
	"Kubernetes %s插件是什么？",
	"为什么需要Kubernetes %s插件？",
	"Kubernetes %s插件是为了解决什么问题？",
	"什么场景下需要Kubernetes %s插件？",
	"如何正确使用Kubernetes %s插件？",
	"Kubernetes %s插件原理",
	"Kubernetes %s插件使用注意事项",
	"Kubernetes %s插件发展历程",
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
		//if idx <= 2 {
		//	idx++ // 跳过已经创建好的目录
		//	continue
		//}

		line = line[6:]
		dir := filepath.Dir(qaPath)
		rawLine := line
		line = fmt.Sprintf("%02d.%s", idx, line)
		mkDir := filepath.Join(dir, line)
		fmt.Printf("%s -> %s\n", line, mkDir)
		if vx.Folders == nil {
			vx.Folders = make([]*vnote.Folder, 0)
		}
		if err = os.MkdirAll(mkDir, os.ModePerm); err != nil {
			return err
		}

		// 目录下生成vx.json文件 通过父目录的vx.json文件生成
		childVX, err := vnote.UnMarshal(vnotePath)
		if err != nil {
			return err
		}
		// 清空数据
		childVX.Folders = []*vnote.Folder{}
		childVX.Files = []*vnote.File{}

		vx.Folders = append(vx.Folders, &vnote.Folder{Name: line})

		// 子目录下下放入QA，以及各个问题，并修改上面的vx.json文件
		ptnQA, err := os.ReadFile("D:/Notebook/Vnote/Blog/10.Kubernetes/04.特性开关/00.CrossNamespaceVolumeDataSource/00.QA.md")
		if err != nil {
			return nil
		}
		ptnQA = []byte(strings.ReplaceAll(string(ptnQA), "CrossNamespaceVolumeDataSource", rawLine))
		ptnQA = []byte(strings.ReplaceAll(string(ptnQA), "特性", "插件"))

		childVX.Files = append(childVX.Files, &vnote.File{
			CreatedTime:  "2024-04-18T07:24:10Z",
			Id:           "6076",
			ModifiedTime: "2024-04-18T07:24:10Z",
			Name:         "00.QA.md",
			Signature:    "177807976771",
		})

		for idx, q := range pattern {
			fName := fmt.Sprintf("%02d.%s.md", idx+1, fmt.Sprintf(q, rawLine))
			qPath := filepath.Join(mkDir, fName)
			if err = os.WriteFile(qPath, []byte{}, os.ModePerm); err != nil {
				return err
			}
			childVX.Files = append(childVX.Files, &vnote.File{
				CreatedTime:  "2024-04-18T07:24:10Z",
				Id:           "6076",
				ModifiedTime: "2024-04-18T07:24:10Z",
				Name:         fName,
				Signature:    "177807976771",
			})

			appQ := fmt.Sprintf("\n你是一个Kubernetes高级专家，请写一篇文章，详细阐述%s 请在文章中尽可能多的帮我添加emoji以增强文章趣味\n", fmt.Sprintf(q, rawLine))
			ptnQA = append(ptnQA, []byte(appQ)...)
		}

		if err = os.WriteFile(filepath.Join(mkDir, "00.QA.md"), ptnQA, os.ModePerm); err != nil {
			return err
		}

		childVXPath := filepath.Join(mkDir, "vx.json")
		if err = childVX.PersistentMarshal(childVXPath); err != nil {
			return err
		}

		idx++
	}

	if err = vx.PersistentMarshal(vnotePath); err != nil {
		return err
	}

	return nil
}
