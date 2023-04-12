package http

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
	mux := http.DefaultServeMux
	mux.HandleFunc("/hello1", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello1"))
	})
	mux.HandleFunc("/hello2", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello2"))
	})

	http.ListenAndServe(":9000", mux)

}
