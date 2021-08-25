package main

import "fmt"

type user struct {
	edad     int
	nombre   string
	apellido string
}

func main() {
	var jordi user
	jordi.edad = 23
	fmt.Println(jordi)
	kevin := user{nombre: "Kevin", edad: 18, apellido: "Gonzalez"}
	fmt.Println(kevin)

	lalito := user{15, "Eduardo", "Diaz"}
	fmt.Println(lalito)
}
