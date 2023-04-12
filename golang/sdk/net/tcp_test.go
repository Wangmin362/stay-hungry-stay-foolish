package net

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"net/url"
	"testing"
)

func TestDial(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:6200")
	if err != nil {
		t.Fatal(err)
	}

	url, err := url.Parse("http://127.0.0.1:6200/")
	if err != nil {
		t.Fatal(err)
	}

	req := &http.Request{
		Method: "CONNECT",
		// URL:    &url.URL{Opaque: address},
		URL:    url,
		Header: map[string][]string{"x-tenant-id": []string{"1000001"}},
	}

	if err = req.Write(conn); err != nil {
		t.Fatal(err)
	}

	br := bufio.NewReader(conn)
	response, err := http.ReadResponse(br, req)
	if err != nil {
		t.Fatal(err)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(body))

}
