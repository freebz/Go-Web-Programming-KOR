// 예제 8-7. 언마샬 함수를 벤치마킹함

func BenchmarkUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unmarshal("post.json")
	}
}
