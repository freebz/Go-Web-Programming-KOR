// 예제 4-8. 클라이언트에 보낼 응답 메시지를 작성함

package main

import (
	"net/http"
)

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web Programming</title></head>
<body><h1>Hello World</h1></body>
<html>`
	w.Write([]byte(str))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/write", writeExample)
	server.ListenAndServe()
}
