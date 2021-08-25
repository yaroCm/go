package main

import "fmt"

type Human struct {
	edad     int
	nombre   string
	apellido string
}
type Tutor struct {
	Human
}

func main() {
	tutor := Tutor{Human{nombre: "Jordi"}}

	fmt.Println(tutor.nombre)
}
