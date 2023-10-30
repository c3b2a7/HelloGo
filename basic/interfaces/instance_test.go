package interfaces

import "testing"

func Test(t *testing.T) {
	gopher1 := Gopher{Name: "gopher1"}
	gopher1.code()
	gopher1.debug()
	gopher2 := &Gopher{Name: "gopher2"}
	gopher2.code()
	gopher2.debug()

	javaer1 := Javaer{Name: "javaer1"}
	javaer1.code()
	javaer1.debug()
	javaer2 := &Javaer{Name: "javaer2"}
	javaer2.code()
	javaer2.debug()

	var gopherCoder1 Coder = Gopher{Name: "gopherCoder1"}
	gopherCoder1.code()
	gopherCoder1.debug()
	var gopherCoder2 Coder = &Gopher{Name: "gopherCoder2"}
	gopherCoder2.code()
	gopherCoder2.debug()

	//var javaerCoder1 Coder = Javaer{Name: "javaerCoder1"}
	//javaerCoder1.code()
	//javaerCoder1.debug()
	var javaerCoder2 Coder = &Javaer{Name: "javaerCoder2"}
	javaerCoder2.code()
	javaerCoder2.debug()
}
