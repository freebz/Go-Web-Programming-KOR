// 예제 4-14. SetCookie를 이용해 브라우저로 쿠키 전송

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Manning Publications Co",
		HttpOnly: true,
	}
	http.SetCookie(w, &c1);
	http.SetCookie(w, &c2)
}
