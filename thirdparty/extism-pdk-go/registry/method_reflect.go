////go:build reflect

// Not supported at this time, due to the use of reflections

package registry

import (
	"fmt"
	"reflect"
)

var errInterface = reflect.TypeOf((*error)(nil)).Elem()

type ReflectiveMethod struct {
	Callable interface{}
}

func (m ReflectiveMethod) Call(args []any) (any, error) {
	callable := reflect.ValueOf(m.Callable)
	callableType := callable.Type()
	if callableType.Kind() != reflect.Func {
		return nil, ErrNotAFunction
	}

	var input []reflect.Value
	for i := 0; i < callableType.NumIn(); i++ {
		expectedArgType := callableType.In(i)
		inputArgType := reflect.TypeOf(args[i])
		if !inputArgType.AssignableTo(expectedArgType) {
			return nil, fmt.Errorf("invalid argument, position: %d, type: %s, expected %s",
				i, inputArgType.Kind(), expectedArgType.Kind())
		}
		input = append(input, reflect.ValueOf(args[i]))
	}

	if len(input) != callableType.NumIn() {
		return nil, ErrInvalidArgsNum
	}

	var ret []any
	var err error
	values := callable.Call(input)
	for _, value := range values {
		if value.Type() != errInterface {
			ret = append(ret, value.Interface())
		} else if v := value.Interface(); v != nil {
			err = v.(error)
		}
	}

	return ret, err
}

type StrictMethod struct {
	M        func(args []any) (any, error)
	ArgsLen  int
	ArgsType []reflect.Type
}

func (m StrictMethod) Call(args []any) (any, error) {
	if len(args) != m.ArgsLen {
		return nil, ErrInvalidArgsNum
	}
	for i, arg := range args {
		inputArgType := reflect.TypeOf(arg)
		expectedArgType := m.ArgsType[i]
		if !inputArgType.AssignableTo(expectedArgType) {
			return nil, fmt.Errorf("invalid argument, position: %d, type: %s, expected %s",
				i, inputArgType.Kind(), expectedArgType.Kind())
		}
	}
	return m.M(args)
}
