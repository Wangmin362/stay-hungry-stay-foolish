package test

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var client, _ = clientv3.New(clientv3.Config{
	Endpoints:   []string{"172.30.3.230:59101"},
	DialTimeout: time.Duration(5) * time.Second,
})

func TestGetEtcdKey(t *testing.T) {
	response, err := client.Get(context.Background(), "/tenant/info/1000001", clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}
	for _, kv := range response.Kvs {
		fmt.Println(kv.Version, "-->", string(kv.Key), "--->", string(kv.Value))
	}
	//response, err = client.Get(context.Background(), "/pop", clientv3.WithPrefix())
	//if err != nil {
	//	panic(err)
	//}
	//for _, kv := range response.Kvs {
	//	fmt.Println(kv.Version, "-->", string(kv.Key), "--->")
	//}
	//response, err = client.Get(context.Background(), "/pop/product_config/mapping", clientv3.WithPrefix())
	//if err != nil {
	//	panic(err)
	//}
	//for _, kv := range response.Kvs {
	//	fmt.Println(kv.Version, "-->", string(kv.Key), "--->", string(kv.Value))
	//}
}

func TestAuth(t *testing.T) {
	tenantId := "1000015"
	popCode := "544f19eb4ea3bc24a9fdadf5"
	popId := "13380f68-533f-4705-8471-cec9a4bcbdd3"

	//code, _ := base64.StdEncoding.DecodeString(popCode)
	//popCode = string(code)[4:28]
	timestamp, authorization := GetAuth(tenantId, popCode, popId)
	t.Log(popCode)
	t.Log(timestamp)
	t.Log(authorization)
}

// QA环境
func TestQAAuth(t *testing.T) {
	tenantId := "1000006"
	popCode := "29b118e84bccb7e499cfc7b1"
	popId := "090d0c8a-52f8-43c5-b82a-97d72e8da527"

	//code, _ := base64.StdEncoding.DecodeString(popCode)
	//popCode = string(code)[4:28]
	timestamp, authorization := GetAuth(tenantId, popCode, popId)
	t.Log(popCode)
	t.Log(timestamp)
	t.Log(authorization)
}

func TestHuaweiAuth(t *testing.T) {
	tenantId := "1000003"
	popCode := "NWNmMjM1ZDY5ZThlNDI4ODhlYzdkODlmYzNiY2I4YzQ="
	popId := "0424b34b-ba27-41c0-abcf-356b9c1fef96"

	code, _ := base64.StdEncoding.DecodeString(popCode)
	popCode = string(code)[4:28]
	timestamp, authorization := GetAuth(tenantId, popCode, popId)
	t.Log(popCode)
	t.Log(timestamp)
	t.Log(authorization)
}

var httpClient = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
}

func TestTenantAuth(t *testing.T) {
	// chen du
	CdSpsHost := "https://cd-ucss-230.gatorcloud.skyguardmis.com/skgwSps"
	httpGet(CdSpsHost+"/sps/v1/tenant/serviceAuth?version=1",
		"1000079", "11386ffe4384b28d5f7a368d", "db9eff40-f10e-4f19-9fd0-85829d9c0911")

	// beijing
	//BjSpsHost := "https://bj-ucss-230.gatorcloud.skyguardmis.com/skgwSps"
	//httpGet(BjSpsHost+"/sps/v1/tenant/serviceAuth?version=1",
	//	"1000018", "e8b0c396454cbda45725dab0", "eed3ceee-beb0-4dc0-a5b5-ea51300ae2ee")
}

func getUcwiAuth(accessKey, secretKey, timestamp string) string {
	tokenSource := secretKey + timestamp
	hash := hmac.New(sha256.New, []byte(secretKey))
	hash.Write([]byte(tokenSource))
	tokenResult := hex.EncodeToString(hash.Sum(nil))

	auth := fmt.Sprintf("SKG %s:%s", accessKey, tokenResult)
	return auth
}

func TestUcwiDlpChannel(t *testing.T) {
	xTenantId := "1000011"
	secretKey := "NDE2MGU0YTktNzI5MS00OTlmLTljZWQtNmM4ZWJi"
	accessKey := "NTc5ZWJjZTktMzkyMC00"
	xTimestamp := strconv.FormatInt(time.Now().Unix(), 10)
	header := map[string]string{
		"x-skg-timestamp": xTimestamp, "Authorization": getUcwiAuth(accessKey, secretKey, xTimestamp),
		"x-tenant-id": xTenantId,
	}
	url := "https://ucwi.cd-pop-222.gatorcloud.skyguardmis.com/skg/v1/dlp/channel"
	for {
		DoHttpRequest(context.Background(), "GET", url, header, nil)
		time.Sleep(2 * time.Millisecond)
	}
}

func TestUcwiDlpChannel1000368(t *testing.T) {
	xTenantId := "1000368"
	secretKey := "YjE5ZGU0OTEtMzZmMS00YjQ0LTkzMDQtMTQzYTJh"
	accessKey := "ZDAyNDdkZTgtYTU0My00"
	xTimestamp := strconv.FormatInt(time.Now().Unix(), 10)
	header := map[string]string{
		"x-skg-timestamp": xTimestamp, "Authorization": getUcwiAuth(accessKey, secretKey, xTimestamp),
		"x-tenant-id": xTenantId,
	}
	url := "https://ucwi.cd-pop-222.gatorcloud.skyguardmis.com/skg/v1/dlp/channel"
	//for {
	DoHttpRequest(context.Background(), "GET", url, header, nil)
	time.Sleep(2 * time.Millisecond)
	//}
}

var tenantId = "1006667"
var ctx = context.Background()

func TestServiceController(t *testing.T) {
	//client.Put(ctx, fmt.Sprintf("/tenant/info/%s", tenantId),
	//	`{"version":208,"status":1,"shard_disk_size":85,"tenant_name":"wangmin-test","secret_key":"aabb","access_key":"ccdd"}`)
	client.Delete(ctx, fmt.Sprintf("/tenant/info/%s", tenantId))

	//// dsg
	//client.Put(ctx, fmt.Sprintf("/tenant/info/%s/rs/dsg", tenantId),
	//	`{"initialized":true,"auth_status":1,"max_rs":23,"start_time":1663862400,"end_time":1670256000}`)
	//client.Delete(ctx, fmt.Sprintf("/tenant/info/%s/rs/dsg", tenantId))
	//client.Put(ctx, fmt.Sprintf("/tenant/info/%s/vpn/001", tenantId),
	//	`{"type":"dsg","conn_id":"`+tenantId+`"N001","if_id":134,"capacity":{"max_bandwidth":"10M"},
	//"auth_type":"PSK","pre_shared_key":"edaf37f0-e12e-40ae-a3ff-e74a2dc777aa","networks":[{"param_type":"IKE","auth_algorithm":"SHA2-256","encrypt_algorithm":"AES-128","dh_algorithm":"DH"},{"param_type":"IPSec","auth_algorithm":"SHA1","encrypt_algorithm":"AES-128","dh_algorithm":"DH"}]}`)
	//client.Put(ctx, fmt.Sprintf("/tenant/state/%s/vpn/001", tenantId), `{"vpn_ip":"10.233.97.162"}`)

	//client.Put(ctx, fmt.Sprintf("/tenant/info/%s/vpn/002", tenantId),
	//	`{"type":"dsg","conn_id":"`+tenantId+`"N001","if_id":134,"capacity":{"max_bandwidth":"10M"},
	//"auth_type":"PSK","pre_shared_key":"edaf37f0-e12e-40ae-a3ff-e74a2dc777aa","networks":[{"param_type":"IKE","auth_algorithm":"SHA2-256","encrypt_algorithm":"AES-128","dh_algorithm":"DH"},{"param_type":"IPSec","auth_algorithm":"SHA1","encrypt_algorithm":"AES-128","dh_algorithm":"DH"}]}`)
	//client.Put(ctx, fmt.Sprintf("/tenant/state/%s/vpn/002", tenantId), `{"vpn_ip":"10.233.97.162"}`)
	//
	//client.Put(ctx, fmt.Sprintf("/tenant/info/%s/vpn/003", tenantId),
	//	`{"type":"dsg","conn_id":"`+tenantId+`"N001","if_id":134,"capacity":{"max_bandwidth":"10M"},
	//"auth_type":"PSK","pre_shared_key":"edaf37f0-e12e-40ae-a3ff-e74a2dc777aa","networks":[{"param_type":"IKE","auth_algorithm":"SHA2-256","encrypt_algorithm":"AES-128","dh_algorithm":"DH"},{"param_type":"IPSec","auth_algorithm":"SHA1","encrypt_algorithm":"AES-128","dh_algorithm":"DH"}]}`)
	//client.Put(ctx, fmt.Sprintf("/tenant/state/%s/vpn/003", tenantId), `{"vpn_ip":"10.233.97.162"}`)

	//client.Put(ctx, fmt.Sprintf("/pop/product_config/mapping/%s", tenantId), "1.1.454")

	// add/delete if tenantId

	//client.Put(ctx, fmt.Sprintf("/pop/product_config/mapping/%s", tenantId), "1.2.3")

	// ucwi
	//client.Put(ctx, fmt.Sprintf("/tenant/info/%s/rs/ucwi", tenantId),
	//	`{"initialized":true,"auth_status":1,"max_rs":12,"capacity":{"daily_query":1200,"query_speed":0},"start_time":1665763200,"end_time":1672416000}`)
	//client.Delete(ctx, fmt.Sprintf("/tenant/info/%s/rs/ucwi", tenantId))
	//client.Put(ctx, fmt.Sprintf("/pop/product_config/mapping/%s", tenantId), "1.2.2")

	// ucsslite
	//client.Put(ctx, fmt.Sprintf("/tenant/info/%s/rs/ucsslite", tenantId),
	//	`{"initialized":true,"auth_status":1,"max_rs":4,"capacity":{"user_count":400},"start_time":1663862400,
	//"end_time":1671120000}`)
	//client.Delete(ctx, fmt.Sprintf("/tenant/info/%s/rs/ucsslite", tenantId))
	//client.Put(ctx, fmt.Sprintf("/pop/product_config/mapping/%s", tenantId), "1.2.3")

	// tenantAuth
	//client.Put(ctx, fmt.Sprintf("/tenant/auth/%s", tenantId), `{"tenantId":"`+tenantId+`","service":["ucwi","dsg"]}`)
	client.Delete(ctx, fmt.Sprintf("/tenant/auth/%s", tenantId))
	//client.Put(ctx, fmt.Sprintf("/pop/product_config/mapping/%s", tenantId), "1.2.3")

}

func Sha256Bytes(bytes []byte) string {
	sha256Bytes := sha256.Sum256(bytes)
	sha256Str := hex.EncodeToString(sha256Bytes[:])
	return sha256Str
}

func Sha256Str(str string) string {
	return Sha256Bytes([]byte(str))
}

func GetAuth(tenantId, popCode, popId string) (xTimestamp, authorization string) {
	//xTimestamp = strconv.FormatInt(time.Now().Unix(), 10)
	xTimestamp = "2962176541"
	token := Sha256Str(xTimestamp + popCode + tenantId + popId)
	basicAuthStr := strings.Join([]string{xTimestamp, token, tenantId, popId}, ":")
	authorization = "Basic " + base64.StdEncoding.EncodeToString([]byte(basicAuthStr))
	return xTimestamp, authorization
}

func httpGet(url, tenantId, popCode, popId string) {
	xTimestamp, authorization := GetAuth(tenantId, popCode, popId)
	header := map[string]string{
		"x-timestamp": xTimestamp, "x-tenant-id": tenantId, "Authorization": authorization,
		"x-pop-id": popId,
	}
	DoHttpRequest(context.Background(), "GET", url, header, nil)
}

func DoHttpRequest(ctx context.Context, method, url string, headers map[string]string, reqObjPointer interface{}) {
	var reqBodyReader io.Reader
	if reqObjPointer != nil {
		reqBody, err := json.Marshal(reqObjPointer)
		if err != nil {
			return
		}
		reqBodyReader = bytes.NewReader(reqBody)
	}
	req, err := http.NewRequest(method, url, reqBodyReader)
	if err != nil {
		panic(err)
	}
	req = req.WithContext(ctx)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("StatusCode: %d\n", resp.StatusCode)

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, respBody, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(prettyJSON.Bytes()))

	return
}
