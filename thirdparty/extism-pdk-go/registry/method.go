package registry

import "errors"

var (
	ErrInvalidArgsNum = errors.New("invalid number of arguments")
	ErrNotAFunction   = errors.New("not a function")
)

type Method interface {
	Call(args []any) (any, error)
}

type FunctionMethod func(args []any) (any, error)

func (m FunctionMethod) Call(args []any) (any, error) {
	return m(args)
}
