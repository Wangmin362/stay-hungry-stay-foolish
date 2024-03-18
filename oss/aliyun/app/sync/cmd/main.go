package main

import (
	"fmt"
	"github.com/golang/demo/oss/aliyun/app/sync"
	"github.com/pkg/errors"
	"os"
)

const (
	EndpointKey  = "EndpointKey"
	BucketKey    = "BucketKey"
	OssIDKey     = "OSS_ACCESS_KEY_ID"
	OssSecretKey = "OSS_ACCESS_KEY_SECRET"
	SyncDirKey   = "SyncDirKey"
)

func getEnvVar(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", errors.Errorf("environment variable %s key not found", key)
	}

	return value, nil
}

// 1、本地文件删除之后暂时不考虑删除云端的文件，保留备份，以免后面还需要
// TODO 2、考虑目录的重命名暂时不处理，后续写一个定时任务，直接清楚阿里云OSS中没有使用的图片
// TODO 3、修正文件移动后，路径不对的问题
// TODO 如何保证图片的安全？ 防止其他人胡乱使用？
func main() {
	var err error
	syncDir, err := getEnvVar(SyncDirKey)
	if err != nil {
		fmt.Printf("%s\n", syncDir)
		os.Exit(1)
	}

	endpoint, err := getEnvVar(EndpointKey)
	if err != nil {
		fmt.Printf("%s\n", syncDir)
		os.Exit(1)
	}
	bucketName, err := getEnvVar(BucketKey)
	if err != nil {
		fmt.Printf("%s\n", syncDir)
		os.Exit(1)
	}
	ossId, err := getEnvVar(OssIDKey)
	if err != nil {
		fmt.Printf("%s\n", syncDir)
		os.Exit(1)
	}
	ossSecret, err := getEnvVar(OssSecretKey)
	if err != nil {
		fmt.Printf("%s\n", syncDir)
		os.Exit(1)
	}

	syncer, err := sync.NewSyncer(syncDir, endpoint, bucketName, ossId, ossSecret)
	if err != nil {
		fmt.Printf("%s\n", syncDir)
		os.Exit(1)
	}

	syncer.Run()
}
