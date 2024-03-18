package sync

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func CreateStandardLRSReadPublicBucket(bckName string, client *oss.Client) error {
	storageClass := oss.ObjectStorageClass(oss.StorageStandard) // 设置存储类型为标准存储类型，实际上默认就是这个类型
	redundancyType := oss.RedundancyType(oss.RedundancyLRS)     // 设置存储冗余类型为本地冗余类型
	acl := oss.ACL(oss.ACLPublicRead)                           // 设置读写权限为公共可读，但不可写

	if err := ; err != nil {
		return err
	}

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
