// 예제 2-1. main.go에 선언한 간단한 main 함수

package main

import (
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)

	server := & http.Server {
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
