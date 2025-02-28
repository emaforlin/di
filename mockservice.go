package di

type mockService struct {
	value string
}

func (s mockService) Value() string {
	return s.value
}

type MockService interface {
	Value() string
}

// MockService factory
func NewMockService(c Container) (interface{}, error) {
	return &mockService{
		value: "Hello DI.",
	}, nil
}
