package interfaces

import "fmt"

type Coder interface {
	code()
	debug()
}

type Gopher struct {
	Name string
}

func (g Gopher) code() {
	fmt.Printf("%s code: golang\n", g.Name)
}

func (g Gopher) debug() {
	fmt.Printf("%s debug: golang\n", g.Name)
}

type Javaer struct {
	Name string
}

func (j *Javaer) code() {
	fmt.Printf("%s code: java\n", j.Name)
}

func (j *Javaer) debug() {
	fmt.Printf("%s debug: java\n", j.Name)
}
