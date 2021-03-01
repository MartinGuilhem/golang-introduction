package main

import "fmt"

type User struct {
	edad             int
	nombre, apellido string
}

func main() {

	// var martin User

	// fmt.Println(User{nombre: "Martin", apellido: "Guilhem", edad: 32})

	// martin := User{nombre: "Martin", apellido: "Guilhem"}

	martin := User{32, "Martin", "Guilhem"}

	fmt.Println(martin)

}
