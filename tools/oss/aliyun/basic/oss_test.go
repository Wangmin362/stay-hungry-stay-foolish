package basic

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
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
func TestCreateBucket(t *testing.T) {
	if err := client.CreateBucket(testBuckName, storageClass, redundancyType, acl); err != nil {
		t.Fatal(err)
	}
}

// 测试创建Bucket并且指定资源组 TODO 似乎阿里云暂时不支持资源组创建
func TestCreateBucketWithResourceGroup(t *testing.T) {
	resourceGroup := oss.PutBucketResourceGroup{
		ResourceGroupId: "blog", // 资源组名
	}
	err := client.PutBucketResourceGroup("my-test-resource-bucket", resourceGroup)
	if err != nil {
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

// 遍历当前拥有所有的桶
func TestListBucket(t *testing.T) {
	buckets, err := client.ListBuckets()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Owner.ID: %s\n", buckets.Owner.ID)
	fmt.Printf("Owner.DisplayName: %s\n", buckets.Owner.DisplayName)
	fmt.Printf("perfix: %s\n", buckets.Prefix)

	for _, bucket := range buckets.Buckets {
		fmt.Printf("Name=%s\n", bucket.Name)
		fmt.Printf("Region=%s\n", bucket.Region)
		fmt.Printf("StorageClass=%s\n", bucket.StorageClass)
		fmt.Printf("Location=%s\n====\n", bucket.Location)
	}
}

// 把文件存储到桶当中，目录分隔符必须是/，不能是\，否则会当成文件名而不是目录
func TestSaveFile(t *testing.T) {
	bucket, err := client.Bucket("gouster-cloud-blog")
	if err != nil {
		t.Fatal(err)
	}

	err = bucket.PutObjectFromFile("dir1/dir2/dir3/20230725174249677_10025.png",
		"20230725174249677_10025.png")
	if err != nil {
		t.Fatal(err)
	}
}

// 遍历桶中的文件
func TestListFile(t *testing.T) {
	bucket, err := client.Bucket("gouster-cloud-blog")
	if err != nil {
		t.Fatal(err)
	}

	objects, err := bucket.ListObjects()
	if err != nil {
		t.Fatal(err)
	}

	for _, obj := range objects.Objects {
		fmt.Printf("StorageClass=%s\n", obj.StorageClass)
		fmt.Printf("Key=%s\n", obj.Key)
		fmt.Printf("ETag=%s\n", obj.ETag)
		fmt.Printf("RestoreInfo=%s\n", obj.RestoreInfo)
		fmt.Printf("Type=%s\n====\n", obj.Type)
	}
}

// 1、测试创建目录，目录必须是以反斜杠结尾；支持创建多级目录
// 2、如何目录重命名？ 官方并不支持使用SDK直接重命名，而是需要先创建新目录，然后拷贝老目录中的东西到新目录中，最后删除旧目录
func TestMkdir(t *testing.T) {
	bucket, err := client.Bucket(testBuckName)
	if err != nil {
		t.Fatal(err)
	}

	// 填写目录名称，目录需以正斜线结尾。 dir1/会创建dir1目录， dir1/dir2则会先创建dir1,然后在dir1目录中创建dir2
	err = bucket.PutObject("dir1/dir2/", bytes.NewReader([]byte("")))
	if err != nil {
		t.Fatal(err)
	}
}
