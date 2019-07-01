// 예제 3-4. HTTPS를 통해 제공

package main

import (
	"net/http"
)

func main() {
	server := http.Server {
		Addr: "127.0.0.1:8080",
		Handler: nil,
	}
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
