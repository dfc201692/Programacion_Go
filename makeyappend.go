package main

import "fmt"

func main() {
	slice := make([]int, 3, 5)
	slice = append(slice, 5)
	fmt.Println(slice)
}

//linea 7 al cambiar la palabra reservada len devuelve la longitud del arreglo interno.
//linea 7 al cambiar la palabra reservada cap devuelve la capacidad.
// append es para agregar o añadir un nuevo dato en el arreglo.
