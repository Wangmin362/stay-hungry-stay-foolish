package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
	"time"
)

var client, _ = clientv3.New(clientv3.Config{
	Endpoints:   []string{"172.30.3.222:59101"},
	DialTimeout: time.Duration(5) * time.Second,
})

func TestGetEtcdKey(t *testing.T) {
	response := client.Watch(context.Background(), "/tenant/info", clientv3.WithPrefix())

	for resp := range response {
		for _, event := range resp.Events {
			fmt.Println(string(event.Kv.Key))
		}
	}
}
