package registry

var globalRegistrar = NewRegistry(nil)

type Registry struct {
	registry map[string]Method
}

func (r *Registry) GetMethod(name string) Method {
	return r.registry[name]
}

func (r *Registry) RegisterMethod(name string, method Method) {
	if _, ok := r.registry[name]; ok {
		return
	}

	r.registry[name] = method
}

func NewRegistry(m map[string]Method) *Registry {
	registry := Registry{
		registry: make(map[string]Method),
	}

	for name, f := range m {
		registry.RegisterMethod(name, f)
	}

	return &registry
}

func GetMethod(name string) Method {
	return globalRegistrar.registry[name]
}

func RegisterMethod(name string, method Method) {
	globalRegistrar.RegisterMethod(name, method)
}
