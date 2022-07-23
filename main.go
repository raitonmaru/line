package main

import (
	"os"
	"github.com/raitochinyu/LINE-test/handler"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("てすと")
	http.HandleFunc("/", handler.HelloHandler)
	http.HandleFunc("/callback", handler.LINEHandler)

	port := os.Getenv("PORT")
	addr := ":" + port

	http.ListenAndServe(addr, nil)
}
