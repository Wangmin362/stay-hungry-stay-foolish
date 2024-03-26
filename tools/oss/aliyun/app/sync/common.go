package sync

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

	// 微信相关的环境变量
	WeChatAppID      string = "WeChatAppID"
	WeChatAppSecret  string = "WeChatAppSecret"
	WeChatURLTagName string = "webchat/url"
)
