package closure

// 函数 + 引用外部变量 = 闭包

func Add(value int) func() int {
	return func() int {
		value++
		return value
	}
}

type Player struct {
	hp   int
	name string
}

func genPlayer(defaultHP int) func(name string) *Player {
	return func(name string) *Player {
		return &Player{
			hp:   defaultHP,
			name: name,
		}
	}
}

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
