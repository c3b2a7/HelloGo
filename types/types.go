package types

type StringWrapper string
type Strings []string

type Func func()
type FuncWithArg func(a interface{})

type Animal interface {
	Say(interface{})
}

type Shape interface {
	Area() (int, error)
	Perimeter() (int, error)
}
