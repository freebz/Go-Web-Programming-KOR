// 예제 9-5.  고루틴 사용 여부에 관계없이 작업을 수행하는 벤치마킹 함수

func BenchmarkPrint2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		print2()
	}
}

func BenchmarkGoPrint2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrint2()
	}
}
