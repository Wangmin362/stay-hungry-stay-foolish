package test

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
	"time"
)

const testKey = "/test/wangmin/key"

func TestETCDVersion(t *testing.T) {
	go func() {
		for {
			_, err := client.Put(ctx, testKey, "999955555")
			if err != nil {
				panic(err)
			}
			time.Sleep(time.Second * 3)
		}
	}()

	watchChan := client.Watch(ctx, testKey, clientv3.WithPrefix())
	for resp := range watchChan {
		for _, ev := range resp.Events {
			fmt.Println(string(ev.Kv.Key), string(ev.Kv.Value), ev.Kv.Version)
		}
	}
}
