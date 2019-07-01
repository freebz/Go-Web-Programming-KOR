// 예제 3-9. http.HandleFunc 소스코드

func HandleFunc(pattern string, handler func(ResponseWriter, * Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}
