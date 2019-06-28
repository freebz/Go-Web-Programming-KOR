// 예제 2-16. thread.go 파일에 선언된 NumReplies 메소드

func(thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("SELECT count(*) FROM posts where thread_id = $1",
		thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan( & count); err != nil {
			return
		}
	}
	rows.Close()
	return
}
