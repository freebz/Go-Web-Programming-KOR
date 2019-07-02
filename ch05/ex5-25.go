// 예제 5-25. 명시적으로 정의된 템플릿 사용하기

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("layout.html")
	t.ExecuteTEmplate(w, "layout", "")
}
