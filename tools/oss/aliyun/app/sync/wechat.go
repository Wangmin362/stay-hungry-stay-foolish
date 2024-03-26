package sync

import (
	"fmt"
	"io"
	"net/http"

	"github.com/golang/demo/tools"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
)

func NewWeChat() (*WeChat, error) {
	appId, err := tools.GetEnvVar(WeChatAppID)
	if err != nil {
		return nil, fmt.Errorf("获取微信AppID环境变量失败：%w\n", err)
	}
	appSecret, err := tools.GetEnvVar(WeChatAppSecret)
	if err != nil {
		return nil, fmt.Errorf("获取微信AppSecret环境变量失败：%s\n", err)
	}

	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &offConfig.Config{
		AppID:     appId,
		AppSecret: appSecret,
		Cache:     memory,
	}
	oa := wc.GetOfficialAccount(cfg)

	return &WeChat{
		wechat: wc,
		oa:     oa,
	}, nil

}

type WeChat struct {
	wechat *wechat.Wechat
	oa     *officialaccount.OfficialAccount
}

func (w *WeChat) ImageUpload(path string) (string, error) {
	material := w.oa.GetMaterial()
	url, err := material.ImageUpload(path)
	if err != nil {
		return "", fmt.Errorf("upload %s pic to wechat error: %w\n", path, err)
	}

	return url, nil
}

func (w *WeChat) GetImage(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("get %s image error: %w", url, err)
	}

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("read http body error: %w", err)
	}

	return all, nil
}
