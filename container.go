package di

import (
	"errors"
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
func (c *container) Register(f Factory) error {
	providerType := reflect.TypeOf(f)
	if providerType.Kind() != reflect.Func {
		return errors.New("factory must to be a function")
	}

	if providerType.NumOut() != 2 {
		return errors.New("factory must to return exactly two value")
	}

	depType := providerType.Out(0)
	c.providers[depType] = f
	return nil
}

// Resolve implements Container.
func (c *container) Resolve(target interface{}) error {
	targetType := reflect.TypeOf(target)

	if targetType.Kind() != reflect.Ptr {
		return errors.New("target must be a pointer")
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
		instances: make(map[reflect.Type]interface{}),
	}
}
