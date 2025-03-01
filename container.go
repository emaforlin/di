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

	fmt.Printf("\nProviders: %+v\n", c.providers)
	return nil
}

// Resolve implements Container.
func (c *container) Resolve(target any) error {
	targetType := reflect.TypeOf(target)
	if targetType.Kind() != reflect.Ptr || targetType == nil {
		return errors.New("target must be a non-nil pointer to a pointer")
	}

	elemType := targetType.Elem()

	if instance, found := c.instances[elemType]; found {
		reflect.ValueOf(target).Elem().Set(reflect.ValueOf(instance))
		return nil
	}

	provider, found := c.providers[elemType]
	if !found {
		return errors.New("no provider found for " + elemType.String())
	}

	instance, err := provider(c)
	if err != nil {
		return err
	}
	c.instances[elemType] = instance
	reflect.ValueOf(target).Elem().Set(reflect.ValueOf(instance))
	return nil
}

func NewContainer() *container {
	return &container{
		providers: make(map[reflect.Type]Factory),
		instances: make(map[reflect.Type]any),
	}
}
