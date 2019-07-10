// 예제 9-2. 고루틴 예제 실행을 위한 테스트 파일

package main

import "testing"

func TestPrint1(t *testing.T) {
	print1()
}

func TestGoPrint1(t *testing.T) {
	goPrint1()
}
