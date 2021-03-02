package main

import "fmt"

type Human struct {
	name string
}

func (this Human) hablar() string {
	return "bla bla blah"
}

type Tutor struct {
	Human
	Dummy
}

type Dummy struct {
	name string
}

func main() {
	tutor := Tutor{Human{"Martin"}, Dummy{"Martin"}}

	// fmt.Println(tutor.name)
	// fmt.Println(tutor.Human.name)
	fmt.Println(tutor.hablar())
}
