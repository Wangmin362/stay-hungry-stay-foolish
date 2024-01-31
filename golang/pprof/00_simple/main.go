package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func Handler(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte(fmt.Sprintf("NumGoroutine=%d", runtime.NumGoroutine())))
}

func main() {
	http.HandleFunc("/demo", Handler)
	go func() {
		if err := http.ListenAndServe("0.0.0.0:9999", nil); err != nil {
			panic(err)
		}
	}()

	go func() {
		for i := 0; i < 100000000; i++ {
			fmt.Println(fmt.Sprintf("NumGoroutine=%d", runtime.NumGoroutine()))
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < 50; i++ {
		go func() {
			time.Sleep(3 * time.Second)
		}()
		time.Sleep(20 * time.Millisecond)
	}
	time.Sleep(7 * time.Hour)
}
