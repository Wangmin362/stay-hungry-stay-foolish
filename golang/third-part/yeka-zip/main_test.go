package main

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"github.com/yeka/zip"
	"io"
	"k8s.io/apimachinery/pkg/util/uuid"
	"os"
	"path"
	"testing"
)

const (
	Password    string = "ec4846fc8a1c4a15"
	CAZipOffset        = 24
	slab               = "D:\\Project\\github\\stay-hungry-stay-foolish\\golang\\third-part\\yeka-zip\\slab-test.zip"
)

func TestZipWithCrypto(t *testing.T) {
	// 1、创建 zip 基础文件，也就是将来生成的zip压缩包
	zipfile := "D:\\Project\\github\\stay-hungry-stay-foolish\\golang\\third-part\\yeka-zip/archive.zip"
	archive, err := os.Create(zipfile)
	if err != nil {
		panic(err)
	}

	// 2、初始化zip写入对象
	zipWriter := zip.NewWriter(archive)

	baseDir := "D:\\Project\\github\\stay-hungry-stay-foolish\\golang\\third-part\\yeka-zip\\test-data"
	dir, err := os.ReadDir(baseDir)
	if err != nil {
		panic(err)
	}

	for _, file := range dir {
		fullname := path.Join(baseDir, file.Name())
		// 这种方式的压缩包不带密码
		//writer, err := zipWriter.Create(file.Name())
		// 这种方式带密码
		writer, err := zipWriter.Encrypt(file.Name(), Password, zip.StandardEncryption)
		if err != nil {
			panic(err)
		}

		openfile, err := os.Open(fullname)
		if err != nil {
			panic(err)
		}
		defer openfile.Close() // nolint

		if _, err := io.Copy(writer, openfile); err != nil {
			panic(err)
		}
	}

	zipWriter.Close()
	archive.Close()

	openfile, err := os.Create(slab)
	if err != nil {
		panic(err)
	}
	defer openfile.Close()

	passwd := fmt.Sprintf("%s%s%s", uuid.NewUUID()[:4], Password, uuid.NewUUID()[:4])
	openfile.Write([]byte(passwd))
	fileraw, err := os.ReadFile(zipfile)
	if err != nil {
		panic(err)
	}
	openfile.Write(fileraw)

	if err := os.Remove(zipfile); err != nil {
		panic(err)
	}
}

func TestUnzip(t *testing.T) {

	caZip, err := os.ReadFile(slab)

	password := caZip[4:20]

	if string(password) != Password {
		t.Fatal("password not equal")
	}

	caRaw := caZip[CAZipOffset:]
	reader, err := zip.NewReader(bytes.NewReader(caRaw), int64(len(caRaw)))
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	for _, file := range reader.File {
		buf.Reset()
		if !file.IsEncrypted() {
			panic(errors.Errorf("expected ca zip to be encrypted"))
		}
		file.SetPassword(string(password))
		rc, err := file.Open()
		if err != nil {
			panic(errors.Wrapf(err, "open %s file error", file.Name))
		}
		defer rc.Close() // nolint

		if _, err := io.Copy(buf, rc); err != nil {
			panic(errors.Wrapf(err, ""))
		}

		fmt.Printf("%s -> %s \n\n", file.Name, buf.String())
	}
}
