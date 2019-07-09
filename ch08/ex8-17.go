// 예제 8-17. 새로운 fetch 메소드

func (post *Post) fetch(id int) (err error) {
	err = post.Db.QueryRow("select id, content, author from posts where id = $1",
		id).Scan(&post.Id, &post.Content, &post.Author)
	return
}
