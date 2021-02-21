package main

import "fmt"

type Human struct {
	name string
}

func (this Human) hablar() string {
	return "Bla bla bla"
}

func (this Tutor) hablar() string {
	return this.Human.hablar() + "Bienvenidos a codigo facilito"
}

type Dummy struct {
	name string
}
type Tutor struct {
	Human
	Dummy
}

func main() {

	tutor := Tutor{Human{"David"}, Dummy{"Carlos"}}

	fmt.Println(tutor.Human.name)
	fmt.Println(tutor.Dummy.name)
	fmt.Println(tutor.hablar())
}
