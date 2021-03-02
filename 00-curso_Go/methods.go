package main

import "fmt"

type User struct {
	edad             int
	nombre, apellido string
}

func (usuario User) nombreCompleto() string {
	return usuario.nombre + " " + usuario.apellido
}

func (usuario *User) setName(n string) {
	usuario.nombre = n
}

func main() {

	martin := new(User)

	martin.nombre = "Martin"
	martin.setName("Tincho")

	fmt.Println(martin.nombre)
}
