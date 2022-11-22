package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type PerfReport struct {
	Id          string    `gorm:"primaryKey,type:uuid"`
	Name        string    `gorm:"unique"` // 报告名字
	PopName     string    // 集群名字
	TenantId    int64     `gorm:"column:tenant_id;default:null" json:"tenant_id"`     // 租户ID
	TenantName  string    `gorm:"column:tenant_name;default:null" json:"tenant_name"` // 租户名字
	ModuleName  string    // 模块名字
	Endpoint    string    // 被测试模块的访问端点
	LimitQps    int64     // 最高测试qps
	HpaReplicas int64     // 入口点的副本数量
	HpaCpu      int64     // 入口点hpa cpu使用率
	HpaMemory   int64     // 入口点hpa memory使用率
	HpaStorage  int64     // 入口点hap 存储使用率
	CpuRequest  int64     // 入口点pod的resource.cpu.request
	CpuLimit    int64     // 入口点pod的resource.cpu.limit
	MemRequest  int64     // 入口点pod的resource.memory.limit
	MemLimit    int64     // 入口点pod的resource.memory.limit
	Status      int64     // 测试类型
	ExtendJson  string    `gorm:"column:extend_json;default:null" json:"extend_json"` // 测试状态
	Remark      string    `gorm:"column:remark;default:null" json:"remark"`           // 备注
	AutoEndTime int64     // 当前任务自动结束时间（单位：分钟），如果时间到了，但是没有手动停止任务，任务会被自动停止
	ReportJson  string    `gorm:"column:report_json;default:null" json:"report_json"` // 性能测试统计报告
	CreateTime  time.Time `gorm:"autoCreateTime"`
	EndTime     time.Time `gorm:"column:end_time;default:null" json:"end_time"`
}

func ReflectMethod(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func TestReflectMethod(t *testing.T) {
	report := &PerfReport{
		Id:         "SDFSDFSFuisdofgjsoipdf647sd65g4s65gf4",
		Name:       "jgoigsdg-sd45646.html",
		PopName:    "cd-pop-222",
		TenantId:   1000001,
		TenantName: "zhangshagn",
		ModuleName: "ucwi",
		CreateTime: time.Now(),
		Status:     4,
		Endpoint:   "http://www.baidu.com",
	}

	method := ReflectMethod(*report)
	fmt.Println(method)
}
