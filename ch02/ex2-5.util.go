// 예제 2-5. util.go 파일에 선언한 세션 유틸리티 함수

func session(w http.ResponseWriter, r * http.Request) (sess data.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = data.Session {Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}
