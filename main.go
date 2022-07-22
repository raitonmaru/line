package main

import (
	"github.com/raitochinyu/LINE-test/handler"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("てすと")
	http.HandleFunc("/", handler.HelloHandler)
	http.HandleFunc("/callback", handler.LINEHandler)

	port := "5000"
	addr := ":" + port

	http.ListenAndServe(addr, nil)
}
