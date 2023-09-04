package test

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
	"time"
)

/*
测试场景：
1、创建
1.1、模拟DSG连接，其实就是修改/tenant/state
2、删除
3、修改
4、授权未到期
5、授权已经到期
6、增减VPN链接
7、减少VPN链接
8、修改磁盘容量
--> 一下是operator测试
9、修改资源的label
10、修改mount参数
*/

func TestTimeToTimestamp(t *testing.T) {
	fmt.Println(time.Date(2023, 8, 7, 15, 50, 4, 0, time.Local).Unix())
}

func TestTimestampToDate(t *testing.T) {
	tm := time.Unix(1701187200, 0)
	fmt.Println(tm.Format("2006-01-02 15:04:05"))
}

var (
	tenant        = "1000888"
	product       = "swg"
	ctx1          = context.Background()
	etcdClient, _ = clientv3.New(clientv3.Config{
		Endpoints:   []string{"172.22.175.222:59101"},
		DialTimeout: time.Duration(5) * time.Second,
	})
)

func TestProductInfo(t *testing.T) {
	key := fmt.Sprintf("/tenant/info/%s", tenant)
	value := `{"version":120,"status":1,"cae_rs_max":0,"shard_disk_size":80,"tenant_name":"skyguard","secret_key":"NjZmYTY0ZjctMjBmMy00YzRlLTgxMTYtYTE5MjBj","access_key":"ZDRjOWY3ZDgtMWMwMi00","start_time":1693152000,"end_time":1702051200,"delete_time":0}`
	if _, err := etcdClient.Put(ctx1, key, value); err != nil {
		panic(err)
	}
}

func TestDeleteProductInfo(t *testing.T) {
	key := fmt.Sprintf("/tenant/info/%s", tenant)
	if _, err := etcdClient.Delete(ctx1, key); err != nil {
		panic(err)
	}
}

func TestCreateProduct(t *testing.T) {
	value := ""
	if product == "tenantAuth" { //需要在UCSS集群上测试
		tenantId := fmt.Sprintf("/tenant/auth/%s", tenant)
		value = `{"initialized":true,"tenantId":"1000005","service":["dsg","dsa","ucsslite","swg","ucwi"]}`
		if _, err := etcdClient.Put(ctx1, tenantId, value); err != nil {
			panic(err)
		}
		return
	}

	switch product {
	case "dsg":
		value = `{"initialized":false,"auth_status":1,"max_rs":3,"capacity":null,"start_time":1693152000,"end_time":1701100800}`
	case "ucwi":
		value = `{"initialized":true,"auth_status":1,"max_rs":10,"capacity":{"user_count":null,"daily_query":0,"query_speed":1000,"max_bandwidth":null},"start_time":1693152000,"end_time":1702051200}`
	case "ucsslite":
		value = `{"initialized":true,"auth_status":1,"max_rs":1,"capacity":{"user_count":100,"daily_query":null,"query_speed":null,"max_bandwidth":null},"start_time":1693152000,"end_time":1701187200}`
	case "swg":
		value = `{"initialized":true,"auth_status":1,"max_rs":1,"capacity":null,"start_time":1693152000,"end_time":1701100800}`
	case "dsa":
		value = `{"initialized":true,"auth_status":1,"max_rs":1,"capacity":null,"start_time":1693152000,"end_time":1701100800}`
	default:
		panic(errors.New("unknown"))
	}
	tenantId := fmt.Sprintf("/tenant/info/%s/rs/%s", tenant, product)
	if _, err := etcdClient.Put(ctx1, tenantId, value); err != nil {
		panic(err)
	}
}
func TestDeleteProduct(t *testing.T) {
	key := fmt.Sprintf("/tenant/info/%s/rs/%s", tenant, product)
	if _, err := etcdClient.Delete(ctx1, key); err != nil {
		panic(err)
	}
}

// 模拟用户通过IPSec连接云上的VPNServer
func TestSumilateVPNConnection(t *testing.T) {
	key := fmt.Sprintf("/tenant/state/%s/vpn/001", tenant)
	value := `{"vpn_ip":"10.233.97.148","type":"dsg"}`
	if _, err := etcdClient.Put(ctx1, key, value); err != nil {
		panic(err)
	}
}

// 新增VPN连接
func TestCreateVPNConnection(t *testing.T) {
	key := fmt.Sprintf("/tenant/info/%s/vpn/004", tenant)
	value := `{"type":"dsg","if_id":75,"conn_id":"1000005N004","auth_status":0,"pod_id":"","capacity":{"user_count":null,"daily_query":null,"query_speed":null,"max_bandwidth":"2M"},"auth_type":"PSK","pre_shared_key":"b36dbd06-891b-4f96-bd20-c29d227d261b","networks":[{"param_type":"IKE","auth_algorithm":"MD5","encrypt_algorithm":"DES","dh_algorithm":"DH"},{"param_type":"IPSec","auth_algorithm":"MD5","encrypt_algorithm":"DES","dh_algorithm":"DH"}],"router":"linux","remark":"qweqwewq"}`
	//value := `{"type":"swg","if_id":75,"conn_id":"1000005N004","auth_status":0,"pod_id":"","capacity":{"user_count":null,"daily_query":null,"query_speed":null,"max_bandwidth":"2M"},"auth_type":"PSK","pre_shared_key":"b36dbd06-891b-4f96-bd20-c29d227d261b","networks":[{"param_type":"IKE","auth_algorithm":"MD5","encrypt_algorithm":"DES","dh_algorithm":"DH"},{"param_type":"IPSec","auth_algorithm":"MD5","encrypt_algorithm":"DES","dh_algorithm":"DH"}],"router":"linux","remark":"qweqwewq"}`
	if _, err := etcdClient.Put(ctx1, key, value); err != nil {
		panic(err)
	}
}

// 删除VPN连接
func TestDeleteVPNConnection(t *testing.T) {
	key := fmt.Sprintf("/tenant/info/%s/vpn/004", tenant)
	if _, err := etcdClient.Delete(ctx1, key); err != nil {
		panic(err)
	}
}
