// 예제 8-5. 벤치마크 테스팅

package main

import (
	"testing"
)

// benchmarking the decode function
func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decode("post.json")
	}
}
