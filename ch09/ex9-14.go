// 예제 9-14. 거리 함수

func distance(p1 [3]float64, p2 [3]float64) float64 {
	return math.Sqrt(sq(p2[0]-p1[0]) + sq(p2[1]-p1[2]) + sq(p2[2]-p1[2]))
}

func sq(n float64) float64 {
	return n * n
}
