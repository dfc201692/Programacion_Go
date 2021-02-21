package main

import "fmt"

type User struct {
	edad             int
	nombre, apellido string
}

func (this User) nombre_completo() string {
	return this.nombre + " " + this.apellido
}

func main() {

	fmt.Println(User{edad: 28, nombre: "Jose", apellido: "Castro"})
	fmt.Println(User{edad: 35, nombre: "Josue", apellido: "Ramirez"})
	fmt.Println(User{edad: 23, nombre: "Camila", apellido: "Giraldo"})
	fmt.Println(User{edad: 39, nombre: "Mario", apellido: "Ruiz"})
	fmt.Println(User{edad: 15, nombre: "Armando", apellido: "Castaño"})
	fmt.Println(User{edad: 92, nombre: "Andres", apellido: "Rodriguez"})

	david := new(User)

	david.nombre = "David"
	david.apellido = "Cruz"

	fmt.Println(david.nombre_completo())
}
