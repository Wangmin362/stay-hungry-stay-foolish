package main

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	decodeString, err := base64.StdEncoding.DecodeString("OTcwODc2ZGY5NTgxNDUyNjk1ODg5ZDBiZWY0ZWE1ZDQ=")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(decodeString)[4:28])
}
