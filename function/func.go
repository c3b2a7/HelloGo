package function

import "fmt"

type Animal interface {
	say(interface{})
}

type Husky struct {
}

func (h Husky) say(i interface{}) {
	fmt.Println("Husky says: ", i)
}

type FuncHusky func(interface{})

func (f FuncHusky) say(i interface{}) {
	f(i)
}

func sayHello() {
	fmt.Println("Hello, Go")
}

func visit(list []string, consumer func(string)) {
	// 遍历切片中的元素，并将值作为参数传给回调函数
	for _, value := range list {
		// 调用回调函数
		consumer(value)
	}
}
