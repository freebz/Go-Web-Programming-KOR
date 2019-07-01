// 예제 3-1. 가장 간단한 웹 서버

package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe("", nil)
}
