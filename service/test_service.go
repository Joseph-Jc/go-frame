package service

type TestService struct {
}

func (t TestService) GetTotal(n int) (total int) {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum
}
