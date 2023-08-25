package main

import (
	"context"
	"github.com/cisco-open/k8s-objectmatcher/patch"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
)

func TestObjectMatcher(t *testing.T) {

	config, err := clientcmd.BuildConfigFromFlags("", "C:\\Users\\wangmin\\Desktop\\kubeconfig\\222-config")
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)

	// 初始创建资源对象
	original := &v1.Service{}
	if err := patch.DefaultAnnotator.SetLastAppliedAnnotation(original); err != nil {
	}
	ctx := context.Background()
	_, err = clientset.CoreV1().Services(original.GetNamespace()).Create(ctx, original, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}

	// 过了一段时间之后，用户重新提交了资源
	modified := &v1.Service{}

	current, err := clientset.CoreV1().Services(modified.GetNamespace()).Get(ctx, modified.GetName(), metav1.GetOptions{})

	patchResult, err := patch.DefaultPatchMaker.Calculate(current, modified)
	if err != nil {
		panic(err)
	}

	if !patchResult.IsEmpty() {
		// patch结果不为空，说明要更新
		if err := patch.DefaultAnnotator.SetLastAppliedAnnotation(modified); err != nil {
		}
		_, err = clientset.CoreV1().Services(modified.GetNamespace()).Update(ctx, modified, metav1.UpdateOptions{})
		if err != nil {
			panic(err)
		}
	}

}
