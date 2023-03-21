package net

import (
	"net/http"
	"testing"
)

func Handler(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("hello test"))
}

func TestBasicHttp(t *testing.T) {
	http.HandleFunc("/demo", Handler)
	if err := http.ListenAndServe(":9000", nil); err != nil {
		t.Fatal(err)
	}
}

func TestHttpMux(t *testing.T) {

}
