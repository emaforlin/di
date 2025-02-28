package di

type Provider interface {
	Provide(container *Container) (interface{}, error)
}
