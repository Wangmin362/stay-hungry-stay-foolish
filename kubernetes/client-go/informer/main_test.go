package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
	"time"
)

func TestShardInformerFactory1(t *testing.T) {
	config, err := clientcmd.BuildConfigFromFlags("", "config")
	if err != nil {
		t.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		t.Fatal(err)
	}
	factory := informers.NewSharedInformerFactory(clientset, 0)
	stop := make(chan struct{})
	factory.Start(stop) // 此时的factory.informers是空的，不会有任何的资源消耗
	time.Sleep(24 * time.Hour)
}

func TestShardInformerFactory2(t *testing.T) {
	config, err := clientcmd.BuildConfigFromFlags("", "config")
	if err != nil {
		t.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		t.Fatal(err)
	}
	stop := make(chan struct{})
	factory := informers.NewSharedInformerFactory(clientset, 0)
	endpointsInformer := factory.Core().V1().Endpoints().Informer()
	endpointsInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			t.Log("add endpoint")
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			t.Log("update endpoint")
		},
		DeleteFunc: func(obj interface{}) {
			t.Log("delete endpoint")
		},
	})

	factory.Start(stop) // 此时factory中只会包含一个EndpointInformer，所有人应该共享这个factory
	if !cache.WaitForCacheSync(stop, endpointsInformer.HasSynced) {
		runtime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
		return
	}
	factory.Start(stop)

	time.Sleep(24 * time.Hour)
}
