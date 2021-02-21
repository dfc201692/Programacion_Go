package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	archivo, error := os.Open("./hola.txt")

	if error != nil {
		fmt.Println("Hubo un error")
	}

	scanner := bufio.NewScanner(archivo)
	var i int
	for scanner.Scan() {
		i++
		linea := scanner.Text()
		fmt.Println(i)
		fmt.Println(linea)
	}
}
