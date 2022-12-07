package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var client, _ = clientv3.New(clientv3.Config{
	Endpoints:   []string{"172.30.3.222:59101"},
	DialTimeout: time.Duration(5) * time.Second,
})

func TestGetEtcdKey(t *testing.T) {
	response, err := client.Get(context.Background(), "/tenant/info", clientv3.WithPrefix())
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
	response, err = client.Get(context.Background(), "/pop/product_config/mapping", clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}
	for _, kv := range response.Kvs {
		fmt.Println(kv.Version, "-->", string(kv.Key), "--->", string(kv.Value))
	}
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
