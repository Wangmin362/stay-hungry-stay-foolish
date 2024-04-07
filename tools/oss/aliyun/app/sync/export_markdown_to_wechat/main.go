package main

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/dlclark/regexp2"
	"github.com/golang/demo/tools"
	"github.com/golang/demo/tools/oss/aliyun/app/sync"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	picPattern  string = `\!\[.*?\]\((.*?)(?: \".*?)?(?: =.*?)?\)`
	tocPattern  string = `[TOC]`
	linkPattern string = `(?<!!)\[(.*?)\]\((.*?)\)`
)

const GenMarkdownDir = "D:\\Markdown"

var bucket *oss.Bucket
var aliOssUrl string

var wechat *sync.WeChat

func init() {
	if err := os.MkdirAll(GenMarkdownDir, os.ModePerm); err != nil {
		log.Fatalf("%s\n", err)
	}

	endpoint, err := tools.GetEnvVar(sync.EndpointKey)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	bucketName, err := tools.GetEnvVar(sync.BucketKey)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	ossId, err := tools.GetEnvVar(sync.OssIDKey)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	ossSecret, err := tools.GetEnvVar(sync.OssSecretKey)
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	aliOssUrl = fmt.Sprintf("https://%s.%s", bucketName, endpoint)

	// 创建阿里云OSS客户端
	client, err := oss.New(fmt.Sprintf("https://%s", endpoint), ossId, ossSecret)
	if err != nil {
		log.Fatalf("create aliyun oss client error:%s", err)
	}

	// 判断指定的桶是否存在
	exist, err := client.IsBucketExist(bucketName)
	if err != nil || !exist {
		log.Fatalf("query %s bucket exist error:%s", bucketName, err)
	}

	// 获取当前桶
	bucket, err = client.Bucket(bucketName)
	if err != nil {
		log.Fatalf("get %s bucket error:%s", bucketName, err)
	}

	wechat, err = sync.NewWeChat()
	if err != nil {
		log.Fatalf("create wechat client error:%s", err)
	}
}

// 用于导出markdown到微信公众号要求的，格式。主要做了以下几个事情：
// 1、把阿里云的图片链接转为微信的，如果这个图片没有上传到微信，那么上传
// 2、把外链改为引用的方式，以明文URL贴在Markdown底部
// 3、把Markdown的[TOC]标记去除掉
// 4、尝试看看能不能把markdown直接通过在线工具，譬如https://markdown.com.cn/转为适合微信公众号的markdown样式
func main() {
	path := "test.md"
	if err := ConvertToWechatFormat(path); err != nil {
		log.Fatal(err)
	}
}

func ConvertToWechatFormat(path string) error {
	if filepath.Ext(path) != ".md" {
		return nil
	}

	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("读取%s文件失败: %s", path, err)
	}

	// 1、去除TOC标识
	file, err = DeleteToc(file)
	if err != nil {
		log.Fatalf("去除%s文件[TOC]标识失败: %s", path, err)
	}

	// 2、替换图片链接为微信链接
	file, err = ReplaceAliToTencent(file)
	if err != nil {
		log.Fatalf("替换%s文件阿里云链接为微信地址失败: %s", path, err)
	}

	// 3、替换外联为明文
	file, err = ReplaceLink(file)
	if err != nil {
		log.Fatalf("替换%s文件外链失败: %s", path, err)
	}

	targetPath := filepath.Join(GenMarkdownDir, filepath.Base(path))

	if err = os.WriteFile(targetPath, file, os.ModePerm); err != nil {
		return err
	}

	fmt.Printf("%s文件处理完成\n", path)

	return nil
}

func ReplaceLink(file []byte) ([]byte, error) {
	fileData := string(bytes.Clone(file))
	re := regexp2.MustCompile(linkPattern, 0)

	index := 1
	refRaw := "\n### 参考链接\n\n"
	ref := refRaw

	match, err := re.FindStringMatch(string(file))
	if err != nil {
		return nil, err
	}

	ma := func(match *regexp2.Match) {
		groups := match.Groups()
		raw := string(groups[0].Captures[0].Runes())  // markdown链接地址
		name := string(groups[1].Captures[0].Runes()) // 链接显示名字
		link := string(groups[2].Captures[0].Runes()) // 链接地址

		target := fmt.Sprintf("***%s***<sup>%d</sup>", name, index)
		fileData = strings.ReplaceAll(fileData, raw, target)
		ref += fmt.Sprintf("%s：%s\n\n", name, link)
		index++
	}

	for match != nil {
		ma(match)
		match, err = re.FindNextMatch(match)
		if err != nil {
			return nil, err
		}
	}

	if ref != refRaw {
		fileData += ref
	}

	return []byte(fileData), nil
}

func ReplaceAliToTencent(file []byte) ([]byte, error) {
	var err error
	fileData := string(bytes.Clone(file))
	re := regexp.MustCompile(picPattern)
	match := re.FindAllSubmatch(file, -1)
	for _, group := range match {
		raw := string(group[0])       // markdown图片格式
		aliOssPic := string(group[1]) // 地址

		if aliOssPic == aliOssUrl {
			log.Printf("存在不正确的%s阿里云地址\n", aliOssPic)
			continue // 直接忽略这个地址
		}

		picKey := aliOssPic[len(aliOssUrl)+1:]
		wechatUrl, exist := sync.GetObjTag(picKey, sync.WeChatURLTagName, bucket)
		// 如果没有找到微信地址，那么就上传到微信
		if !exist || wechatUrl == "" {
			log.Printf("未找到%s图片对应的微信地址，尝试上传到微信\n", aliOssPic)
			wechatUrl, err = wechat.ImageUploadByUrl(aliOssPic)
			if err != nil {
				log.Printf("上传%s图片到微信失败: %s\n", aliOssPic, err)
				return nil, err
			}

			// 上传成功，把微信地址写入到阿里云的Tag中
			if err = sync.AddObjTag(picKey, sync.WeChatURLTagName, wechatUrl, bucket); err != nil {
				log.Printf("写入%s图片对应的微信地址到阿里云失败: %s\n", aliOssPic, err)
				return nil, err
			}

			log.Printf("上传%s图片到微信成功，微信地址为：%s\n", aliOssPic, wechatUrl)
		}

		// 获取一次图片，看看能否正常获取到
		_, err = wechat.GetImage(wechatUrl)
		if err != nil {
			log.Printf("获取%s图片失败：%s, 尝试把阿里云的图片上传到微信\n", wechatUrl, err)
			wechatUrl, err = wechat.ImageUploadByUrl(aliOssPic)
			if err != nil {
				log.Printf("上传%s图片到微信失败: %s\n", aliOssPic, err)
				return nil, err
			}

			// 上传成功，把微信地址写入到阿里云的Tag中
			if err = sync.AddObjTag(picKey, sync.WeChatURLTagName, wechatUrl, bucket); err != nil {
				log.Printf("写入%s图片对应的微信地址到阿里云失败: %s\n", aliOssPic, err)
				return nil, err
			}

			log.Printf("上传%s图片到微信成功，微信地址为：%s\n", aliOssPic, wechatUrl)
		}

		target := fmt.Sprintf("![](%s)", wechatUrl)
		fileData = strings.ReplaceAll(fileData, raw, target)
	}

	return []byte(fileData), nil
}

func DeleteToc(file []byte) ([]byte, error) {
	// 一篇markdown文档，一般只会写一次TOC,所以这里只会选择替换一次
	replace := strings.Replace(string(file), tocPattern, "", 1)
	return []byte(replace), nil
}
