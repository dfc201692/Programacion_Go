package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go mi_nombre_lentooo("David")
	fmt.Println("Que aburrido")
	var wait string
	fmt.Scan(&wait)
}

func mi_nombre_lentooo(name string) {
	letras := strings.Split(name, "")

	for _, letra := range letras {
		time.Sleep(1000 * time.Microsecond)
		fmt.Println(letra)
	}

}
