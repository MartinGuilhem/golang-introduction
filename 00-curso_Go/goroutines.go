package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go miNombreLentooo("hola mundo")
	fmt.Println("Que aburridooo")
	go func() {
		var wait string
		fmt.Scanln(&wait)
	}()
}

func miNombreLentooo(name string) {
	letras := strings.Split(name, "")

	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
