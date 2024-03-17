package main

import (
	"fmt"
	"os"
	gpath "path"
	"path/filepath"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg/errors"
)

const (
	url          = "https://%s.%s/%s" // https://<bucketName>.<endpoint>/<path>
	syncImageDir = "vx_images"
)

const (
	EndpointKey  = "EndpointKey"
	BucketKey    = "BucketKey"
	OssIDKey     = "OSS_ACCESS_KEY_ID"
	OssSecretKey = "OSS_ACCESS_KEY_SECRET"
	SyncDirKey   = "SyncDirKey"
)

func NewSyncer(syncDir, imageDir string) (*syncer, error) {
	endpoint, err := getEnvVar(EndpointKey)
	if err != nil {
		return nil, err
	}
	bucketName, err := getEnvVar(BucketKey)
	if err != nil {
		return nil, err
	}
	ossId, err := getEnvVar(OssIDKey)
	if err != nil {
		return nil, err
	}
	ossSecret, err := getEnvVar(OssSecretKey)
	if err != nil {
		return nil, err
	}

	client, err := oss.New(fmt.Sprintf("https://%s", endpoint), ossId, ossSecret)
	if err != nil {
		return nil, fmt.Errorf("create aliyun oss client error:%w", err)
	}

	exist, err := client.IsBucketExist(bucketName)
	if err != nil {
		return nil, fmt.Errorf("query %s bucket exist error:%w", bucketName, err)
	}

	if !exist {
		if err := createBucket(bucketName, client); err != nil {
			return nil, nil
		}
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, fmt.Errorf("get %s bucket error:%w", bucketName, err)
	}

	return &syncer{
		client:     client,
		bucket:     bucket,
		endpoint:   endpoint,
		bucketName: bucketName,
		ossId:      ossId,
		ossSecret:  ossSecret,
		syncDir:    syncDir,
		imageDir:   imageDir,
		cacheObjs:  make(map[string]Empty),
	}, nil
}

type Empty struct{}

type syncer struct {
	client     *oss.Client
	bucket     *oss.Bucket
	endpoint   string
	bucketName string
	ossId      string
	ossSecret  string
	syncDir    string
	imageDir   string // 如果设置，那么仅同步名字为指定目录下的文件，否则同步所有文件

	cacheObjs map[string]Empty
}

func getEnvVar(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", errors.Errorf("environment variable %s key not found", key)
	}

	return value, nil
}

func createBucket(bckName string, client *oss.Client) error {
	storageClass := oss.ObjectStorageClass(oss.StorageStandard) // 设置存储类型为标准存储类型，实际上默认就是这个类型
	redundancyType := oss.RedundancyType(oss.RedundancyLRS)     // 设置存储冗余类型为本地冗余类型
	acl := oss.ACL(oss.ACLPublicRead)                           // 设置读写权限为公共可读，但不可写

	if err := client.CreateBucket(bckName, storageClass, redundancyType, acl); err != nil {
		return fmt.Errorf("create bucket error:%w", err)
	}

	return nil
}

// 定时扫描没有上传的文件  一个小时扫描一次
// 判断文件是否已经上传，如果已经上传就不再上传  这里应该使用bucket.ListObject获取桶中的所有对象，减少SDK的调用
// 本地文件删除之后暂时不考虑删除云端的文件，保留备份，以免后面还需要
// TODO 考虑目录的重命名
// TODO 定时同步
// TODO 如何保证图片的安全？ 防止其他人胡乱使用？
func main() {
	syncDir, err := getEnvVar(SyncDirKey)
	if err != nil {
		fmt.Printf("%s\n", syncDir)
		os.Exit(1)
	}

	syncer, err := NewSyncer(syncDir, syncImageDir)
	if err != nil {
		fmt.Printf("%s\n", syncDir)
		os.Exit(1)
	}

	if err := syncer.cacheAllAliyunObjs(); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	if err := syncer.SyncToAliyunOSS(syncDir); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}

func (s *syncer) cacheAllAliyunObjs() error {
	continueToken := ""
	for {
		lsRes, err := s.bucket.ListObjectsV2(oss.ContinuationToken(continueToken))
		if err != nil {
			return err
		}
		// 打印列举结果。默认情况下，一次返回100条记录。
		for _, obj := range lsRes.Objects {
			s.cacheObjs[obj.Key] = Empty{}
		}
		if lsRes.IsTruncated {
			continueToken = lsRes.NextContinuationToken
		} else {
			break
		}
	}

	return nil
}

func (s *syncer) SyncToAliyunOSS(syncDir string) error {
	return filepath.Walk(syncDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if syncDir == path {
			return nil
		}

		if info.IsDir() {
			// 当前目录是想要同步的目录才考虑上传阿里云
			if info.Name() == s.imageDir {
				return s.SyncToAliyunOSS(path)
			}
		}

		if !strings.Contains(path, s.imageDir) {
			return nil
		}

		index := strings.Index(path, syncDir)
		if index < 0 {
			return errors.Errorf("%s目录不正确，基础目录不是%s", path, syncDir)
		}

		realPath := path[len(s.syncDir)+1:]
		split := strings.Split(realPath, "\\")
		realPath = gpath.Join(split...)

		if _, ok := s.cacheObjs[realPath]; ok {
			fmt.Printf("@@@已经同步%s文件到阿里云,访问路径为:%s\n", path, fmt.Sprintf(url, s.bucketName, s.endpoint, realPath))
			return nil
		}

		fmt.Printf("正在同步%s文件到阿里云%s\n", path, realPath)

		if err = s.bucket.PutObjectFromFile(realPath, path); err != nil {
			return err
		}

		fmt.Printf("同步%s文件到阿里云%s成功！！！\n\n", path, realPath)

		return nil
	})
}
