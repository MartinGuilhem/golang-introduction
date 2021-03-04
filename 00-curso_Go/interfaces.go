package main

import "fmt"

type User interface {
	Permisos() int // 1 - 5
	Nombre() string
}

type Admin struct {
	nombre string
}

type Editor struct {
	nombre string
}

// Es un usuario xq esta implementando los metodos de la interfaz.
func (this Admin) Permisos() int {
	return 5
}

func (this Admin) Nombre() string {
	return this.nombre
}

//es un usuario xq esta implementando los metodos de la interfaz
func (this Editor) Permisos() int {
	return 3
}

func (this Editor) Nombre() string {
	return this.nombre
}

func auth(user User) string {
	permisos := user.Permisos()

	if permisos > 4 {
		return user.Nombre() + " tiene permisos de Administrador"
	} else if permisos == 3 {
		return user.Nombre() + " tiene permisos de Editor"
	}
	// switch user.Permisos(){
	// 	case
	// }
	return ""
}

func main() {
	// admin := Admin{"Martin"}
	// editor := Editor{"Fulano"}

	// fmt.Println(auth(admin))
	// fmt.Println(auth(editor))

	usuarios := []User{Admin{"Martin"}, Editor{"Fulano"}}

	for _, usuario := range usuarios {
		fmt.Println(auth(usuario))
	}

}
