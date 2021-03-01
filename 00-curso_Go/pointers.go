package main

import "fmt"

func main() {
	// var puntero *int

	var x, y *int

	entero := 5

	x = &entero
	y = &entero

	fmt.Println(*x)
	fmt.Println(y)

}
