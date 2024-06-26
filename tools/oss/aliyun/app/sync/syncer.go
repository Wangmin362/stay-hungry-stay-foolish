package sync

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"

	"github.com/golang/demo/tools"
	"github.com/golang/demo/tools/oss/aliyun/app/sync/cache"
)

func NewSyncer() (*syncer, error) {
	syncDir, err := tools.GetEnvVar(SyncDirKey)
	if err != nil {
		log.Fatalf("%s\n", syncDir)
	}

	endpoint, err := tools.GetEnvVar(EndpointKey)
	if err != nil {
		log.Fatalf("%s\n", syncDir)
	}
	bucketName, err := tools.GetEnvVar(BucketKey)
	if err != nil {
		log.Fatalf("%s\n", syncDir)
	}
	ossId, err := tools.GetEnvVar(OssIDKey)
	if err != nil {
		log.Fatalf("%s\n", syncDir)
	}
	ossSecret, err := tools.GetEnvVar(OssSecretKey)
	if err != nil {
		log.Fatalf("%s\n", syncDir)
	}

	// 创建阿里云OSS客户端
	client, err := oss.New(fmt.Sprintf("https://%s", endpoint), ossId, ossSecret)
	if err != nil {
		return nil, fmt.Errorf("create aliyun oss client error:%w", err)
	}

	// 判断指定的桶是否存在
	exist, err := client.IsBucketExist(bucketName)
	if err != nil {
		return nil, fmt.Errorf("query %s bucket exist error:%w", bucketName, err)
	}

	// 如果桶不存在，就创建这个桶
	if !exist {
		if err := CreateStandardLRSReadPublicBucket(bucketName, client); err != nil {
			return nil, fmt.Errorf("create %s bucket error: %w", bucketName, err)
		}
	}

	// 为当前桶设置防盗链，防止流量盗刷
	if err = SetReferer(client, bucketName,
		[]string{"*wangmin362.github.io", "*localhost:8020",
			"*.jianshu.com", "*jianshu.com",
			"*.bbs.huaweicloud.com", "*bbs.huaweicloud.com",
			"*.juejin.cn", "*juejin.cn",
			"*.segmentfault.com", "*segmentfault.com",
			"*.xie.infoq.cn", "*xie.infoq.cn",
			"*.developer.aliyun.com", "*developer.aliyun.com",
			"*.cloud.tencent.com", "*cloud.tencent.com",
			"*.zhihu.com", "*zhihu.com",
			"*.res.hc-cdn.com", "*res.hc-cdn.com",
		},
		[]string{"*.baidu.com"},
	); err != nil {
		return nil, fmt.Errorf("set bucket referer error: %w", err)
	}

	// 获取当前桶
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, fmt.Errorf("get %s bucket error:%w", bucketName, err)
	}

	// 创建一个文件监听器，当文件、目录发生变化时，我们可以及时知道，而不必每次循环迭代扫描所有文件
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, fmt.Errorf("create file watcher:%w", err)
	}

	// 监听指定的同步目录的文件变化
	if err = watcher.Add(syncDir); err != nil {
		return nil, fmt.Errorf("watch dir %s error:%w", syncDir, err)
	}

	chat, err := NewWeChat()
	if err != nil {
		return nil, fmt.Errorf("init wechat error:%w", err)
	}

	s := &syncer{
		client:     client,
		bucket:     bucket,
		endpoint:   endpoint,
		bucketName: bucketName,
		ossId:      ossId,
		ossSecret:  ossSecret,
		syncDir:    syncDir,
		imageDir:   syncImageDir, // 设置仅针对某些特殊的目录上传文件
		Cache:      cache.NewCache(),
		fsWatcher:  watcher,
		wechat:     chat,
	}

	// 先把当前bucket中所有缓存的文件名存储起来
	if err := s.cacheAllAliOSSObjs(); err != nil {
		return nil, err
	}

	// 全量同步一次
	if err := s.syncDirPic(syncDir); err != nil {
		return nil, err
	}

	return s, nil
}

type Empty struct{}

type syncer struct {
	client *oss.Client
	bucket *oss.Bucket

	endpoint   string
	bucketName string
	ossId      string
	ossSecret  string
	syncDir    string // 需要同步的目录
	imageDir   string // 如果设置，那么仅同步名字为指定目录下的文件，否则同步所有文件

	fsWatcher *fsnotify.Watcher

	markdownLock sync.Mutex

	*cache.Cache

	// 用于微信上传照片
	wechat *WeChat
}

func (s *syncer) cacheAllAliOSSObjs() error {
	continueToken := ""
	for {
		lsRes, err := s.bucket.ListObjectsV2(oss.ContinuationToken(continueToken))
		if err != nil {
			return err
		}

		// 打印列举结果。默认情况下，一次返回100条记录。
		for _, obj := range lsRes.Objects {
			//tag, exist := s.ObjExist(obj.Key)
			//if exist && tag.WechatUrl != "" {
			//	if _, err = s.wechat.GetImage(tag.WechatUrl); err == nil {
			//		// 如果当前对象已经缓存了，并且微信的图片链接也是可以使用的，那么直接缓存对象
			//		s.CacheObj(obj.Key, tag)
			//	}
			//}
			//
			//// 否则，说明当前图片没有在微信当中，尝试上传一次到微信中
			//
			//wechatTag, exist := GetObjTag(obj.Key, WeChatURLTagName, s.bucket)
			//if exist && wechatTag != "" { // 先从阿里云中获取这个对象的Tag,看看从Tag中是否能够获取到微信链接
			//	tag.WechatUrl = wechatTag
			//	// 如果获取到了，尝试访问一下这个图片，如果能够正常访问，说明微信链接是对的
			//	if _, err = s.wechat.GetImage(tag.WechatUrl); err == nil {
			//		s.CacheObj(obj.Key, tag)
			//		continue
			//	}
			//}
			//
			//// 否则，如果微信链接获取不到，或者是根本就没存，那就尝试上传一次微信公众号
			//aliOssUrl := fmt.Sprintf(url, s.bucketName, s.endpoint, obj.Key)
			//byUrl, err := s.wechat.ImageUploadByUrl(aliOssUrl)
			//if err != nil {
			//	log.Printf("上传图片到微信错误，%s\n", err)
			//	tag.WechatUrl = ""
			//	s.CacheObj(obj.Key, tag)
			//	continue
			//}
			//
			//// 如果微信上传成功了，就需要更新阿里云的tag
			//if err = AddObjTag(obj.Key, WeChatURLTagName, byUrl, s.bucket); err != nil {
			//	log.Printf("更新阿里云OSS图片Tag错误，%s\n", err)
			//	tag.WechatUrl = ""
			//	s.CacheObj(obj.Key, tag)
			//	continue
			//}
			//
			//tag.WechatUrl = byUrl
			s.CacheObj(obj.Key, cache.Tag{})
		}

		if lsRes.IsTruncated {
			continueToken = lsRes.NextContinuationToken
		} else {
			break
		}
	}

	return nil
}

func (s *syncer) syncDirPic(syncDir string) error {
	return filepath.Walk(syncDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if syncDir == path {
			return nil
		}

		if info.IsDir() {
			if err = s.fsWatcher.Add(path); err != nil { // 子目录也需要监视
				return fmt.Errorf("add watch %s dir error: %w", path, err)
			}
			return nil
		}

		if strings.Contains(path, RecycleBin) {
			return nil
		}

		// 当前文件是普通文件，直接上传到阿里云
		if err = s.saveToAliOss(path); err != nil {
			return err
		}
		return nil
	})
}

func (s *syncer) saveToAliOss(path string) error {
	info, err := os.Stat(path)
	if err != nil { // 如果这里出错，一般都是文件不存在造成的，直接忽略这个错误
		log.Printf("[warning] statistic %s error: %s", path, err)
		return nil
	} else if info.Size() <= 0 { // // 如果当前图片的大小为0，暂时先不同步
		return nil
	}

	// 当前文件路径必须包含指定的路径才是需要同步的文件,否则直接跳过
	if !strings.Contains(path, s.imageDir) {
		return nil
	}

	// 如果目录不正确，直接跳过
	index := strings.Index(path, s.syncDir)
	if index < 0 {
		return errors.Errorf("%s目录不正确，基础目录不是%s", path, s.syncDir)
	}

	// 获取子路径
	realPath := path[len(s.syncDir)+1:]
	dstBucketKey := tools.ConvertWindowDirToLinuxDir(realPath)

	tag, exist := s.ObjExist(dstBucketKey)
	if exist { // 说明当前文件已经同步
		// TODO 这里的写法为什么会导致程序卡死？ 和go的遍历目录实现有关？
		//if err = os.Remove(path); err != nil { // 如果已经同步成功，直接删除文件
		//	log.Printf("删除%s图片失败：%s", path, err)
		//} else {
		//	log.Printf("删除%s图片成功", path)
		//}
		return nil
	}

	// TODO 是否需要排查图片以外的数据

	if err := SaveToAliOSS(path, dstBucketKey, s.bucket); err != nil {
		if errors.Is(err, FileZeroSize) {
			return nil
		}
		return fmt.Errorf("同步%s文件到阿里云错误: %w", path, err)
	}

	ext := filepath.Ext(path)
	// TODO 微信素材接口仅支持jpg/png格式，大小必须在1MB以下  支持图片格式转换，大小转换
	if (ext != ".jpg" && ext != ".png") || info.Size() > 2<<20 {
		s.CacheObj(dstBucketKey, tag)

		log.Printf("同步%s文件到阿里云%s成功!，当前图片不是jpg或者png图片，或者大小朝超过1MB\n", path, dstBucketKey)
		return nil
	}

	imageUrl, err := s.wechat.ImageUpload(path)
	if err != nil {
		log.Printf("上传%s图片到微信公众号失败: %s\n", path, err)
		imageUrl = "" // 如果微信上传失败了，我们选择忽略，因为重点还是需要把图片上传到阿里云，只要阿里云同步成功了就好说，微信这里迟早会同步成功
	}

	// 如果微信上传成功了，就需要更新阿里云的tag
	if err = AddObjTag(dstBucketKey, WeChatURLTagName, imageUrl, s.bucket); err != nil {
		log.Printf("更新阿里云OSS图片Tag错误，%s\n", err)
		tag.WechatUrl = ""
	}

	_ = os.Remove(path) // 如果已经同步成功，直接删除文件

	tag.WechatUrl = imageUrl
	s.CacheObj(dstBucketKey, tag)

	log.Printf("同步%s文件到阿里云%s成功!\n", path, dstBucketKey)
	return nil
}

func (s *syncer) moveFile(dst, src string) error {
	if err := MoveFile(dst, src, s.bucket); err != nil {
		return err
	}

	// 对象移动之后，标签也会跟着移动，查询目标的Tag即可
	tag, exist := GetObjTag(dst, WeChatURLTagName, s.bucket)
	if !exist {
		tag = ""
		return nil
	}

	s.Replace(dst, src, cache.Tag{WechatUrl: tag})
	return nil
}

func (s *syncer) Run() {
	defer s.fsWatcher.Close()

	// 增量同步
	go s.watchDirTask()

	for {
		if err := s.syncDirPic(s.syncDir); err != nil {
			log.Printf("%s\n", err)
		}

		if err := s.replaceDirPic(s.syncDir); err != nil {
			log.Printf("%s\n", err)
		}
		time.Sleep(60 * time.Minute)
	}
}
