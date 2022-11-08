package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var client, _ = clientv3.New(clientv3.Config{
	Endpoints:   []string{"172.30.3.230:59101"},
	DialTimeout: time.Duration(5) * time.Second,
})

func TestGetEtcdKey(t *testing.T) {
	response, err := client.Get(context.Background(), "/tenant/auth", clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}
	for _, kv := range response.Kvs {
		fmt.Println(kv.Version, "-->", string(kv.Key), "--->", string(kv.Value))
	}
	response, err = client.Get(context.Background(), "/pop", clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}
	for _, kv := range response.Kvs {
		fmt.Println(kv.Version, "-->", string(kv.Key), "--->")
	}
	response, err = client.Get(context.Background(), "/pop/product_config/mapping", clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}
	for _, kv := range response.Kvs {
		fmt.Println(kv.Version, "-->", string(kv.Key), "--->", string(kv.Value))
	}
}
