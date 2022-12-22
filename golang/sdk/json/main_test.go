package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type ProductInfo struct {
	Initialized bool   `json:"initialized"`           // 是否已初始化，初始为false，为ture时SKServiceController方可创建资源
	AuthStatus  *int32 `json:"auth_status,omitempty"` // 0: 取消授权，1：已授权
	MaxRs       int32  `json:"max_rs"`                // 最大replicas
	AccessKey   string `json:"access_key,omitempty"`  // 用于skproxy代理ucwi-api接口时，通过该值获取租户id,skproxy内部维护一个dict，存放所有租户access_key对应的tenant id
	StartTime   int32  `json:"start_time,omitempty"`  // 授权开始时间
	EndTime     int32  `json:"end_time,omitempty"`    // 授权结束时间
}

func TestJsonTest(t *testing.T) {
	auth := int32(1)
	pro := &ProductInfo{
		Initialized: false,
		AuthStatus:  &auth,
		MaxRs:       0,
	}

	marshal, _ := json.Marshal(pro)
	fmt.Println(string(marshal))
}
