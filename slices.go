package main

import "fmt"

func main() {
	matriz := [4]int{1, 2, 3, 4}
	slice := matriz[0:2]
	fmt.Println(slice)
}

//puntero al arreglo
//longitud del arreglo al que apunta
//capacidad
