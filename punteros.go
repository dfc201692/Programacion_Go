package main

import "fmt"

func main() {
	/*
	   1. Es una direccion de memoria
	   2. En lugar del valor, tenemos la direcccion en la que esta el valor
	   3. X, Y => as123d => 5
	   4. X => as123d => 6
	   5. Y ¿? => 6
	   *T => *int, *string, *Struct
	   Valor zero => nil
	*/

	var x, y *int
	entero := 5

	x = &entero
	y = &entero

	*x = 5

	fmt.Println(x)
	fmt.Println(y)

	fmt.Println(*x)
	fmt.Println(*y)

}
