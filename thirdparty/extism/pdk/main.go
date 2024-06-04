package main

import (
	"errors"
	"github.com/extism/go-pdk"
)

//export greet
func greet() int32 {
	name := string(pdk.Input())
	if name == "Benjamin" {
		pdk.SetError(errors.New("sorry, we don't greet Benjamins"))
		return 1
	}
	greeting := `Hello, ` + name + `!`
	pdk.OutputString(greeting)
	return 0
}

func main() {
}
