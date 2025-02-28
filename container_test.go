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
}
