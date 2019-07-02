// 예제 5-11. 템플릿을 포함하는 핸들러

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t1.html", "t2.html")
	t.Execute(w, "Hello World!")
}
