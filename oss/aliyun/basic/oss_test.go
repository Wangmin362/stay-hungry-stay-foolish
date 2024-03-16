package basic

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"testing"
)

const (
	chengduOssEndpoint = "https://oss-cn-chengdu.aliyuncs.com"
	// bucket名字
	testBuckName = "gopher-bucket-test"
)

var (
	provider oss.EnvironmentVariableCredentialsProvider
	client   *oss.Client

	storageClass   = oss.ObjectStorageClass(oss.StorageStandard) // 设置存储类型为标准存储类型，实际上默认就是这个类型
	redundancyType = oss.RedundancyType(oss.RedundancyLRS)       // 设置存储冗余类型为本地冗余类型
	acl            = oss.ACL(oss.ACLPublicRead)                  // 设置读写权限为公共可读，但不可写
	id             []byte
	key            []byte
)

func init() {
	var err error
	// 设置OSS_ACCESS_KEY_ID
	id, err = os.ReadFile("secret-id")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	if err = os.Setenv("OSS_ACCESS_KEY_ID", string(id)); err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	// 设置OSS_ACCESS_KEY_SECRET
	key, err = os.ReadFile("secret-key")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	if err = os.Setenv("OSS_ACCESS_KEY_SECRET", string(key)); err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	// 通过环境变量获取认证
	provider, err = oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	client, err = oss.New(chengduOssEndpoint, "", "", oss.SetCredentialsProvider(&provider))
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
}

// 测试创建Bucket
func TestBucketBucket(t *testing.T) {
	if err := client.CreateBucket(testBuckName, storageClass, redundancyType, acl); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteBucket(t *testing.T) {
	if err := client.DeleteBucket(testBuckName, storageClass, redundancyType, acl); err != nil {
		t.Fatal(err)
	}
}

// 获取当前可用的所有区域，以及区域的详细信息
func TestRegin(t *testing.T) {
	list, err := client.DescribeRegions()
	if err != nil {
		t.Fatal(err)
	}
	for _, region := range list.Regions {
		// 打印所有支持地域的信息。
		fmt.Printf("Region:%s\n", region.Region)
		// 打印所有支持地域对应的外网访问（IPv4）Endpoint。
		fmt.Printf("Region Internet Endpoint:%s\n", region.InternetEndpoint)
		// 打印所有支持地域对应的内网访问（经典网络或VPC网络）Endpoint。
		fmt.Printf("Region Internal Endpoint:%s\n", region.InternalEndpoint)
		// 打印所有支持地域对应的传输加速域名（全地域上传下载加速）Endpoint。
		fmt.Printf("Region Accelerate Endpoint:%s\n", region.AccelerateEndpoint)
	}
	fmt.Println("List Describe Regions Success")
}

func TestListBucket(t *testing.T) {
	buckets, err := client.ListBuckets()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Owner: %s\n", buckets.Owner)
	fmt.Printf("perfix: %s\n", buckets.Prefix)

	for _, bucket := range buckets.Buckets {
		fmt.Printf("Name=%s\n", bucket.Name)
		fmt.Printf("Region=%s\n", bucket.Region)
		fmt.Printf("StorageClass=%s\n", bucket.StorageClass)
		fmt.Printf("Location=%s\n====\n", bucket.Location)
	}
}

func TestSaveFile(t *testing.T) {
	bucket, err := client.Bucket(testBuckName)
	if err != nil {
		t.Fatal(err)
	}

	err = bucket.PutObjectFromFile("Kubernetes\\CRI\\vx_images\\20220614103441287_22373.png", "D:\\Notebook\\Vnote\\Kubernetes\\CRI\\vx_images\\20220614103441287_22373.png")
	if err != nil {
		t.Fatal(err)
	}
}
