package main

import "fmt"

func main() {
	/*	arreglo := [5]int{5, 6, 7, 8, 9}

		for i := 0; i < len(arreglo); i++ {
			fmt.Println(arreglo[i])
		}*/
	var matriz [3][2]int

	matriz[0][0] = 4
	matriz[0][1] = 1
	matriz[1][0] = 6
	matriz[1][1] = 9
	matriz[2][0] = 7
	matriz[2][1] = 16

	fmt.Println(matriz)

}
