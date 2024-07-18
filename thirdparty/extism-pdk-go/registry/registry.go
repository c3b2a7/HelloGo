package registry

import (
	"encoding/json"
)

var globalRegistrar = NewRegistry(nil)

type Registry struct {
	registry map[string]ExportMethod
}

type ExportMethod func(args []json.RawMessage) (any, error)

func (m ExportMethod) Call(args []json.RawMessage) (any, error) {
	return m(args)
}

func (r *Registry) GetMethod(name string) ExportMethod {
	return r.registry[name]
}

func (r *Registry) RegisterMethod(name string, method ExportMethod) {
	if _, ok := r.registry[name]; ok {
		return
	}

	r.registry[name] = method
}

func NewRegistry(m map[string]ExportMethod) *Registry {
	registry := Registry{
		registry: make(map[string]ExportMethod),
	}

	for name, f := range m {
		registry.RegisterMethod(name, f)
	}

	return &registry
}

func GetMethod(name string) ExportMethod {
	return globalRegistrar.registry[name]
}

func RegisterMethod(name string, method ExportMethod) {
	globalRegistrar.RegisterMethod(name, method)
}
