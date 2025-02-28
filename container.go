package di

import (
	"errors"
	"fmt"
	"reflect"
)

type Container interface {
	Register(Factory) error
	Resolve(interface{}) error
}

type container struct {
	providers map[reflect.Type]Factory
	instances map[reflect.Type]interface{}
}

// Register implements Container.
func (c *container) Register(factory Factory) error {
	if factory == nil {
		return errors.New("factory function cannot be nil")
	}

	providerType := reflect.TypeOf(factory)
	if providerType.Kind() != reflect.Func {
		return errors.New("factory must to be a function")
	}

	if providerType.NumOut() != 2 {
		return errors.New("factory must to return exactly two value")
	}

	instance, err := factory(c)
	if err != nil {
		return err
	}

	depType := reflect.TypeOf(instance)
	if depType.Kind() != reflect.Ptr {
		return errors.New("factory function must to return a pointer to a struct")
	}

	c.providers[depType] = factory

	fmt.Printf("Providers: %+v", c.providers)
	return nil
}

// Resolve implements Container.
func (c *container) Resolve(target any) error {
	ptr := reflect.ValueOf(target)
	if ptr.Kind() != reflect.Ptr || ptr.IsNil() {
		return errors.New("target must be a non-nil pointer to a pointer")
	}

	targetType := ptr.Elem().Type()
	if targetType.Kind() != reflect.Ptr {
		targetType = reflect.PointerTo(targetType)
	}

	factory, exists := c.providers[targetType]
	if !exists {
		return fmt.Errorf("no provider found for %s", targetType.String())
	}

	instance, err := factory(c)
	if err != nil {
		return err
	}
	ptr.Elem().Set(reflect.ValueOf(instance))
	return nil
}

func NewContainer() *container {
	return &container{
		providers: make(map[reflect.Type]Factory),
		instances: make(map[reflect.Type]any),
	}
}
