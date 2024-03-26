package sync

import (
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"net/http"
	"os"
)

var FileZeroSize = errors.New("file size is empty")

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
		return FileZeroSize
	}

	if err := bucket.PutObjectFromFile(dstBucketKey, filepath); err != nil {
		return fmt.Errorf("保存 %s到阿里云失败:%w", filepath, err)
	}
	return nil
}

func SaveAnyToAliOSS(reader io.Reader, dstBucketKey string, bucket *oss.Bucket) error {
	if err := bucket.PutObject(dstBucketKey, reader); err != nil {
		return err
	}
	return nil
}

func MoveFile(dst, src string, bucket *oss.Bucket) error {
	// 先拷贝、再删除
	if _, err := bucket.CopyObject(src, dst); err != nil {
		return err
	}
	return bucket.DeleteObject(src)
}

func AddObjTag(objKey, tagKey, tagValue string, bucket *oss.Bucket) error {
	tagging, err := bucket.GetObjectTagging(objKey)
	if err != nil {
		return fmt.Errorf("get %s obj tag error: %w", objKey, err)
	}

	tags := tagging.Tags
	tag := oss.Tag{Key: tagKey, Value: tagValue}

	tags = append(tags, tag)
	if err = bucket.PutObjectTagging(objKey, oss.Tagging{Tags: tags}); err != nil {
		return fmt.Errorf("save %s ojb %s=%s tag error: %w", objKey, tagKey, tagValue, err)
	}
	return nil
}

func GetObjTag(objKey, tagKey string, bucket *oss.Bucket) (string, bool) {
	tagging, err := bucket.GetObjectTagging(objKey)
	if err != nil {
		return "", false
	}

	for _, tag := range tagging.Tags {
		if tag.Key == tagKey {
			return tag.Value, true
		}
	}

	return "", false
}

func GetImage(url string) ([]byte, error) {
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
