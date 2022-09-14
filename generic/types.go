package generic

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type UInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Float interface{ ~float32 | ~float64 }

type SliceElement interface{ Int | UInt | Float | string }

//type Slice[T interface{ Int | UInt | Float }] []T

type Slice[T SliceElement] []T

func Add[T interface{ Int | UInt | Float }](a, b T) T {
	return a + b
}
