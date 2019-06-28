// 예제 2-2. main.go 파일에 선언한 index 핸들러 함수

func index(w http.ResponseWriter, r * http.Request) {
	files := [] string{"templates/layout.html",
		               "templates/navbar.html",
		               "templates/index.html",}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads(); if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
