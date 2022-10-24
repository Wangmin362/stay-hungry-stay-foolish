package redis_cluster

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	redis "github.com/go-redis/redis/v9"
)

var clusterClient *redis.ClusterClient
var ctx = context.Background()

// TODO 如何在K8S集群外部测试K8S内部的Redis集群, 目前似乎无法测试，由于redis启用了集群模式，因此即便连接了主节点，redis集群也会让
// TODO 客户端连接具体的某个节点，儿一般该节点的IP在K8S内部，所以在K8S外部无法连接该节点
func init() {
	log.SetFlags(log.Llongfile | log.Lshortfile)
	// 连接redis集群
	clusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{ // 填写master主机
			"192.168.21.22:30001",
			"192.168.21.22:30002",
			"192.168.21.22:30003",
		},
		Password:     "123456",              // 设置密码
		DialTimeout:  50 * time.Microsecond, // 设置连接超时
		ReadTimeout:  50 * time.Microsecond, // 设置读取超时
		WriteTimeout: 50 * time.Microsecond, // 设置写入超时
	})
	// 发送一个ping命令,测试是否通
	s := clusterClient.Do(ctx, "ping").String()
	fmt.Println(s)
}

func TestConnByRedisCluster(t *testing.T) {
	// 测试一个set功能
	s := clusterClient.Set(ctx, "name", "barry", time.Second*60).String()
	fmt.Println(s)
}
func TestPipe(t *testing.T) {
	// 测试管道发送多条命令.
	pipe := clusterClient.Pipeline()
	for i := 0; i < 10; i++ {
		pipe.Set(ctx, "go"+strconv.Itoa(i), strconv.Itoa(i), time.Second*300)
	}
	// 真正执行发送操作.
	result, err := pipe.Exec(ctx)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

// 验证上面是否拿到数据
func TestGetKey(t *testing.T) {
	for i := 0; i < 10; i++ {
		ret := clusterClient.Get(ctx, "go"+strconv.Itoa(i)).String()
		fmt.Println(ret)
	}
}
