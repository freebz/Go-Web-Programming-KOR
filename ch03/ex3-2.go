// 예제 3-2. 웹 서버에 부가 설정을 함

package main

import (
	"net/http"
)

func main() {
	server := http.Server {
		Addr: "127.0.0.1:8080",
		Handler: nil,
	}
	server.ListenAndServe()
}
