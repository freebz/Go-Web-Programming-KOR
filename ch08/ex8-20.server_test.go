// 예제 8-20. check 패키지를 사용한 server_test.go

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	. "gopkg.in/check.v1"
)

type PostTestSuit struct {}

func init() {
	Suit(&PostTestSuit{})
}

func Test(t *testing.T) { TestingT(t) }

func (s *PostTestSuite) TestHandleGet(c *C) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest(&FakePost{}))
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/post/1", nil)
	mux.ServeHTTP(writer, request)

	c.Check(writer.Code, Equals, 200)
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	c.Check(post.Id, Equals, 1)
}
