// 예제 8-22. Ginkgo 테스트 수트 파일

package main_test

import (
	. "github.com/onsi/ginkgo"
	. "hithub.com/onsi/gomega"

	"testing"
)

func TEstGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ginkgo Suite")
}
