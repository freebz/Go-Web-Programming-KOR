// 예제 6-7. 데이터베이스 핸들을 생성하는 함수

var Db *sql.DB

func innit() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}
