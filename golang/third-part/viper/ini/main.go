package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	App *AppConfig
}

type AppConfig struct {
	PopId         string
	PopCode       string
	DevopsAddr    string
	PopConfigPath string
}

// 大致原理就是viper会把配置文件使用一个map[string]interface{}反序列化，如果想把这个map结果映射到struct中，那么必须使用 mapstructure，而不是json
func main() {
	configFile := "golang/third-part/viper/ini/skg-config-beijing.ini"
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	var config SkgConfig
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}
	fmt.Printf("%v", config)
}

type SkgConfig struct {
	Global GlobalSkgConfig `mapstructure:"global,omitempty"`
	Pop    PopSkgConfig    `mapstructure:"pop,omitempty"`
}

type GlobalSkgConfig struct {
	IsUcssPop    bool   `mapstructure:"is_ucss_pop,omitempty"`
	K8sNamespace string `mapstructure:"k8s_namespace,omitempty"`
	UcssHost     string `mapstructure:"ucss_host,omitempty"`
	Devops       string `mapstructure:"devops_host,omitempty"`
}

type PopSkgConfig struct {
	PopId      string `mapstructure:"pop_id,omitempty"`
	PopCode    string `mapstructure:"pop_code,omitempty"`
	EtcdHost   string `mapstructure:"etcd_host,omitempty"`
	EsHttpHost string `mapstructure:"es_http_host,omitempty"`
	EsHttpAuth string `mapstructure:"es_http_auth,omitempty"`
}
