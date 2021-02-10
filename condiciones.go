package main

import "fmt"

/*
== igual a
!= diferente de
< menor que
> mayor que
>= mayor o igual que
<= menor o igual que
&& AND
|| OR
*/

func main() {
	x := 13
	y := 13
	if x > y {
		fmt.Printf("%d es mayor que  %d \n", x, y)
	} else if x < y {
		fmt.Printf("%d es mayor que %d \n", y, x)
	} else {
		fmt.Println("Son iguales")
	}
}
