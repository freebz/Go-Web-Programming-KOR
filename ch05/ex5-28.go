// 예제 5-28. 다른 템플릿 파일에서 동일 템플릿 파일을 사용하는 핸들러

func process(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("layout.html", "red_hello.html")
	} else {
		t, _ = template.ParseFiles("layout.html", "blue_hello.html")
	}
	t.ExecuteTemplate(w, "layout", "")
}
