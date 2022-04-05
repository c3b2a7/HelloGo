package function

import (
	"fmt"
	"testing"
)

func TestHusky(t *testing.T) {
	var animal Animal
	animal = new(Husky)
	fmt.Printf("%T\n", animal)
	animal.say("Hello, Go")
}

func TestFuncHusky(t *testing.T) {
	var animal Animal = FuncHusky(func(i interface{}) {
		fmt.Println("Husky says: ", i)
	})
	fmt.Printf("%T\n", animal)
	animal.say("Hello, Go")
}

func TestFuncVariable(t *testing.T) {
	// 声明一个函数类型的变量，注意类型为 func()
	var f func()
	// 将函数名赋值给变量 f
	f = sayHello
	// 通过变量 f 直接调用函数
	f()

	//f := sayHello
	//f()
}

func TestVisit(t *testing.T) {
	// 定义一个切片
	list := []string{"li", "fei"}
	// 使用匿名函数打印切片内容
	visit(list, func(value string) {
		fmt.Println(value)
	})
}

func TestAnonymousFunc(t *testing.T) {
	func(name string) {
		fmt.Printf("Hello, %s\n", name)
	}("Go") // 立即调用
}

func TestAnonymousFuncVariable(t *testing.T) {
	hello := func(name string) {
		fmt.Printf("Hello, %s\n", name)
	}
	hello("Go")
}
