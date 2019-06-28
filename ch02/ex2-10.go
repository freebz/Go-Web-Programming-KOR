// 예제 2-10. generateHTML 함수

func generateHTML(w http.ResponseWriter, data interface {}, fn ...string) {
	var files[] string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}
