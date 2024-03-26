package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/golang/demo/tools"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/domain/openapi"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"log"
	"testing"
)

var oa *officialaccount.OfficialAccount

func init() {
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
	oa = wc.GetOfficialAccount(cfg)
}

// 查询上传永久素材，限制为：{"errcode":0,"errmsg":"ok","quota":{"daily_limit":1000,"used":1001,"remain":0}}
func TestQueryAPIQuota(t *testing.T) {
	openAPI := oa.GetOpenAPI()

	api := "/cgi-bin/media/uploadimg"
	quota, err := openAPI.GetAPIQuota(openapi.GetAPIQuotaParams{CgiPath: api})
	if err != nil {
		log.Fatalf("查询APIQuota失败：%s\n", err)
	}
	marshal, _ := json.Marshal(quota)
	fmt.Println(string(marshal))
}

func TestQueryClearQuota(t *testing.T) {
	openAPI := oa.GetOpenAPI()

	api := "/cgi-bin/clear_quota"
	quota, err := openAPI.GetAPIQuota(openapi.GetAPIQuotaParams{CgiPath: api})
	if err != nil {
		log.Fatalf("查询APIQuota失败：%s\n", err)
	}
	marshal, _ := json.Marshal(quota)
	fmt.Println(string(marshal))
}
