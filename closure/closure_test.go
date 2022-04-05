package closure

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	accumulator := Add(1)
	fmt.Println(accumulator()) // 2
	fmt.Println(accumulator()) // 3

	accumulator = Add(10)
	fmt.Println(accumulator()) // 11
	fmt.Println(accumulator()) // 12
	fmt.Println(accumulator()) // 13
}

func TestGenPlayer(t *testing.T) {
	generator := genPlayer(100)
	player := generator("ZhangSan")
	fmt.Printf("%v\n", player)
}
