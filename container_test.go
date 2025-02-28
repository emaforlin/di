package di

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainerFactory(t *testing.T) {
	container := NewContainer()
	assert.NotNil(t, container)
}

func TestRegisterAndResolve(t *testing.T) {
	container := NewContainer()
	assert.NoError(t, container.Register(NewMockService))

	var service *MockService

	assert.NoError(t, container.Resolve(&service))
	assert.NotNil(t, service)
	assert.Equal(t, "Hello, DI!", service.Value())
}
