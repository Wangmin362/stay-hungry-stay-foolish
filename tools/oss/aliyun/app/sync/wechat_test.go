package sync

import (
	"os"
	"testing"
)

func TestWeChat_ImageUpload(t *testing.T) {
	path := "test_data/k8s-high-level-arch.png"

	chat, err := NewWeChat()
	if err != nil {
		t.Fatal(err)
	}

	upload, err := chat.ImageUpload(path)
	if err != nil {
		t.Fatal(err)
	}

	image, err := chat.GetImage(upload)
	if err != nil {
		t.Fatal(err)
	}

	//rawFile, err := os.ReadFile(path)
	//if err != nil {
	//	t.Fatal(err)
	//}

	err = os.WriteFile("k8s", image, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	// 微信会对图片做处理，并不完全相等
	//if !bytes.Equal(image, rawFile) {
	//	t.Fatal("image not equal")
	//}
}
