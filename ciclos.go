package main

import "fmt"

func main() {
	var tabla int
	fmt.Println("¿Cual talba deseas ver?")
	fmt.Scanln(&tabla)
	for i := 1; i <= 10; i++ {
		resultado := i * tabla
		fmt.Printf("%d x %d = %d\n", tabla, i, resultado)
	}
}
