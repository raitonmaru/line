package main

import (
	"os"
	// "github.com/raitochinyu/LINE-test/handler"
	"./handler"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("起動します。")
	go http.HandleFunc("/", handler.resultHandler)
	go http.HandleFunc("/callback", handler.LINEHandler)

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("ポートが取得できませんでした。")
		port = "5000"
	}
	fmt.Println(port)
	addr := ":" + port

	http.ListenAndServe(addr, nil)
}
