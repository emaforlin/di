package di

type Factory func(Container) (interface{}, error)
