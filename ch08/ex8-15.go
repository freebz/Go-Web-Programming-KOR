// 예제 8-14. 수정된 메인(main) 함수

func main() {
	var err error
	db, err := sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}

	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/post/", handleRequest(&Post{Db: db}))
	server.ListenAndServe()
}
