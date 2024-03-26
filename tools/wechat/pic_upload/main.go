package main

import (
	"fmt"
	"github.com/golang/demo/tools"
	"github.com/silenceper/wechat/cache"
	"github.com/silenceper/wechat/v2"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"log"
)

const (
	WeChatAppID     string = "WeChatAppID"
	WeChatAppSecret string = "WeChatAppSecret"
)

func main() {
	appId, err := tools.GetEnvVar(WeChatAppID)
	if err != nil {
		log.Fatalf("获取微信AppID环境变量失败：%s\n", err)
	}
	appSecret, err := tools.GetEnvVar(WeChatAppSecret)
	if err != nil {
		log.Fatalf("获取微信AppSecret环境变量失败：%s\n", err)
	}

	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &offConfig.Config{
		AppID:     appId,
		AppSecret: appSecret,
		Cache:     memory,
	}
	oa := wc.GetOfficialAccount(cfg)

	material := oa.GetMaterial()
	url, err := material.ImageUpload("C:\\Users\\David\\Downloads\\k8s-arch.png")
	if err != nil {
		log.Fatalf("上传图片失败，%s\n", err)
	}
	fmt.Printf("图片访问地址为：%s\n", url)
}
