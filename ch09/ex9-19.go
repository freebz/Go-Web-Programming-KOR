// 예제 9-19. DB 구조체

type DB struct {
	mutex *sync.Mutex
	store map[string][3]float64
}
