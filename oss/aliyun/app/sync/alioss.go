package sync

import (
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func CreateStandardLRSReadPublicBucket(bckName string, client *oss.Client) error {
	storageClass := oss.ObjectStorageClass(oss.StorageStandard) // 设置存储类型为标准存储类型，实际上默认就是这个类型
	redundancyType := oss.RedundancyType(oss.RedundancyLRS)     // 设置存储冗余类型为本地冗余类型
	acl := oss.ACL(oss.ACLPublicRead)                           // 设置读写权限为公共可读，但不可写

	return client.CreateBucket(bckName, storageClass, redundancyType, acl)
}

// SetReferer 设置防盗链，防止流量盗刷
func SetReferer(client *oss.Client, bucketName string, whitelist, blacklist []string) error {
	var setBucketReferer oss.RefererXML
	// 添加Referer白名单，且允许空Referer。Referer参数支持通配符星号（*）和问号（？）。
	setBucketReferer.RefererList = whitelist

	// 添加Referer黑名单。Go SDK 2.2.8及以上版本支持添加Referer黑名单。
	setBucketReferer.RefererBlacklist = &oss.RefererBlacklist{
		Referer: blacklist,
	}

	// 允许空refer的查询
	setBucketReferer.AllowEmptyReferer = true
	boolFalse := true // 允许截断querystring查询
	setBucketReferer.AllowTruncateQueryString = &boolFalse
	return client.SetBucketRefererV2(bucketName, setBucketReferer)
}

func SaveToAliOSS(filepath, dstBucketKey string, bucket *oss.Bucket) error {
	info, err := os.Stat(filepath)
	if err != nil {
		return err
	}
	if info.IsDir() { // 忽略目录
		return nil
	}

	// 如果当前图片的大小为0，暂时先不同步
	if info.Size() <= 0 {
		return nil
	}

	if err := bucket.PutObjectFromFile(dstBucketKey, filepath); err != nil {
		return fmt.Errorf("保存 %s到阿里云失败:%w", filepath, err)
	}
	return nil

}
