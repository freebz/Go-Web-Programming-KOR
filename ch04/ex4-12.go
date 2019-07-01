// 예제 4-12. 쿠키 구조

type Cookie struct {
	Name       string
	Value      string
	Path       string
	Domain     string
	Expires    time.Time
	RawExpires string
	MaxAge     int
	Secure     bool
	Raw        string
	Unparsed   []string
}
