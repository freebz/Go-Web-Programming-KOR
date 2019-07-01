// 예제 4-1. URL 구조

type URL struct {
	Scheme   string
	Opaque   string
	User     *Userinfo
	Host     string
	Path     string
	RawQuery string
	Fragment string
}
