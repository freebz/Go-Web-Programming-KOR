// 예제 2-12. 스레드 구조

package data

import (
	"time"
)

type Thread struct {
	Id      int
	Uuid    string
	Topic   string
	UserId  int
	CreatedAt time.Time
}
