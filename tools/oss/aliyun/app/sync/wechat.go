package sync

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"mime/multipart"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/util"

	"github.com/golang/demo/tools"
)

func NewWeChat() (*WeChat, error) {
	appId, err := tools.GetEnvVar(WeChatAppID)
	if err != nil {
		return nil, fmt.Errorf("获取微信AppID环境变量失败：%w\n", err)
	}
	appSecret, err := tools.GetEnvVar(WeChatAppSecret)
	if err != nil {
		return nil, fmt.Errorf("获取微信AppSecret环境变量失败：%s\n", err)
	}

	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &offConfig.Config{
		AppID:     appId,
		AppSecret: appSecret,
		Cache:     memory,
	}
	oa := wc.GetOfficialAccount(cfg)

	return &WeChat{
		wechat: wc,
		oa:     oa,
	}, nil

}

type WeChat struct {
	wechat *wechat.Wechat
	oa     *officialaccount.OfficialAccount
}

func (w *WeChat) ImageUpload(path string) (string, error) {
	material := w.oa.GetMaterial()
	url, err := material.ImageUpload(path)
	if err != nil {
		APIQuotaErr := "reach max api daily quota limit"
		if strings.Contains(err.Error(), APIQuotaErr) {
			if err := w.ResetQuota(); err != nil {
				return "", errors.Errorf("重置微信API Quota接口限制失败:%w", err)
			}
			// 解除限制之后尝试再次上传
			time.Sleep(time.Minute) // 等一下微信重置
			url, err = material.ImageUpload(path)
			if err != nil {
				return "", fmt.Errorf("upload %s pic to wechat error: %w\n", path, err)
			}
			return url, nil
		}

		return "", fmt.Errorf("upload %s pic to wechat error: %w\n", path, err)
	}
	return url, nil
}

const (
	mediaUploadImageURL = "https://api.weixin.qq.com/cgi-bin/media/uploadimg"
)

func (w *WeChat) ImageUploadByUrl(imageUrl string) (string, error) {
	var accessToken string
	accessToken, err := w.oa.GetMaterial().GetAccessToken()
	if err != nil {
		return "", err
	}

	uri := fmt.Sprintf("%s?access_token=%s", mediaUploadImageURL, accessToken)
	var response []byte

	file, err := http.Get(imageUrl)
	if err != nil {
		return "", fmt.Errorf("get %s image error: %w", url, err)
	}

	imageData, err := io.ReadAll(file.Body)
	if err != nil {
		return "", fmt.Errorf("读取阿里云图片数据失败: %w", err)
	}
	reader := bytes.NewReader(imageData)

	// TODO 上传之前，可能需要把图片的转换一下，因为微信只支持1MB以内的png/jpg图片格式
	response, err = w.PostMultipartForm("media", path.Base(imageUrl), reader, uri)
	if err != nil {
		return "", err
	}

	type resMediaImage struct {
		util.CommonError

		URL string `json:"url"`
	}

	var image resMediaImage
	err = json.Unmarshal(response, &image)
	if err != nil {
		return "", err
	}
	if image.ErrCode != 0 {
		err = fmt.Errorf("UploadImage error : errcode=%v , errmsg=%v", image.ErrCode, image.ErrMsg)

		APIQuotaErr := "reach max api daily quota limit"
		if strings.Contains(string(response), APIQuotaErr) {
			if err := w.ResetQuota(); err != nil {
				return "", errors.Errorf("重置微信API Quota接口限制失败:%w", err)
			}
			// 解除限制之后尝试再次上传
			time.Sleep(time.Minute) // 等一下微信重置
			reader := bytes.NewReader(imageData)
			response, err = w.PostMultipartForm("media", path.Base(imageUrl), reader, uri)
			if err != nil {
				return "", err
			}
		}

		return "", err // 如果是其它错误，直接返回
	}

	return image.URL, nil
}

func (w *WeChat) GetImage(url string) ([]byte, error) {
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

// PostMultipartForm 上传文件或其他多个字段
func (w *WeChat) PostMultipartForm(fieldName, filename string, file io.Reader, uri string) (respBody []byte, err error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, e := bodyWriter.CreateFormFile(fieldName, filename)
	if e != nil {
		err = fmt.Errorf("error writing to buffer , err=%v", e)
		return
	}

	if _, err = io.Copy(fileWriter, file); err != nil {
		return
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, e := http.Post(uri, contentType, bodyBuf)
	if e != nil {
		err = e
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	respBody, err = io.ReadAll(resp.Body)
	return
}

func (w *WeChat) ResetQuota() error {
	openAPI := w.oa.GetOpenAPI()

	return openAPI.ClearQuota()
}
