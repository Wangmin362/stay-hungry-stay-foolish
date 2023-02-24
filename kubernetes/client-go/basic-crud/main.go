package main

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	//config, err := clientcmd.BuildConfigFromFlags("", "C:\\Users\\wangmin\\Desktop\\kubeconfig\\222-config")
	config, err := clientcmd.BuildConfigFromFlags("", "C:\\Users\\wangmin\\Desktop\\kubeconfig\\CDYZ-OPS-217-config")
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)

	configmap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-test",
		},
		Data: map[string]string{
			"abc": "def",
		},
	}
	if _, err := clientset.CoreV1().ConfigMaps("gator-cloud").Create(context.Background(), configmap,
		metav1.CreateOptions{}); err != nil {
		panic(err)
	}
}
