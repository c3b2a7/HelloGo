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

type ChipType int

func (c ChipType) String() string {
	switch c {
	case UNKNOWN:
		return "UNKNOWN"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}

const (
	UNKNOWN ChipType = 1 << iota
	CPU
	GPU
)
