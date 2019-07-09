// 예제 8-13. 새로운 Post 구조체

type Post struct {
	Db      *sql.DB
	Id      int    `json:"id"`
	Content string `json:"Content"`
	Author  string `json:"author"`
}
