package di

type MockService struct {
	value string
}

type IMockService interface {
	Value() string
}

func (s MockService) Value() string {
	return s.value
}

// MockService factory
func NewMockService(c Container) (any, error) {
	return &MockService{
		value: "Hello, DI!",
	}, nil
}
