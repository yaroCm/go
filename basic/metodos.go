package main

import "fmt"

type user struct {
	edad     int
	nombre   string
	apellido string
}

//si no mándas puntero genera una copía
func (this *user) nombreCompleto() string {
	return this.nombre + " " + this.apellido
}
func (this *user) setName(nombre string) {
	this.nombre = nombre
}
func main() {
	var jordi user
	jordi.edad = 23
	fmt.Println(jordi)
	kevin := user{nombre: "Kevin", edad: 18, apellido: "Gonzalez"}
	fmt.Println(kevin)

	lalito := user{15, "Eduardo", "Diaz"}
	lalito.setName("Owen")
	fmt.Println(lalito.nombreCompleto())
}
