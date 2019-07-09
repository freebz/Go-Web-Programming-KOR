// 예제 8-12. handlePost에 전달할 때 사용하는 인터페이스

type Text interface {
	fetch(id int) (err error)
	create() (err error)
	update() (err error)
	delete() (err error)
}
