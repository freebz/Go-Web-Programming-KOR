// 예제 2-14. data.go 파일에서 DB 전역 변수와 init 초기화 함수

var Db * sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "dbname=chitchat sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}
