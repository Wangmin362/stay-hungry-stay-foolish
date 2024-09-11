package main

import (
	"k8s.io/client-go/tools/cache"
	"testing"
)

func TestCacheSync(t *testing.T) {
	done := make(chan struct{})
	if !cache.WaitForCacheSync(done) {
		t.Fatalf("sync error")
	}
	t.Logf("sync successful")
}
