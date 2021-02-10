package main

import "fmt"

func main() {
	/*
		Copia el minimo de elementos en ambos arreglos
	*/
	slice := []int{1, 2, 3, 4, 9, 8, 7, 6}
	copia := make([]int, len(slice))
	copy(copia, slice)
	fmt.Println(slice)
	fmt.Println(copia)
}
