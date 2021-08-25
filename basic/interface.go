package main

import "fmt"

type User interface {
	Permissions() int
	Name() string
}

type Admin struct {
	nombre string
}

func (this Admin) Permissions() int {
	return 5
}

func (this Admin) Name() string {
	return this.nombre
}

type Editor struct {
	nombre string
}

func (this Editor) Permissions() int {
	return 2
}

func (this Editor) Name() string {
	return this.nombre
}
func auth(user User) string {
	if user.Permissions() > 3 {
		return user.Name() + " tiene los permisos necesarios"
	} else {
		return user.Name() + " no  tiene los permisos necesarios"
	}

}

func main() {
	admin := Admin{"Jordi"}
	//fmt.Println(auth(admin))

	editor := Editor{"Kevin"}
	//fmt.Println(auth(editor))

	usuarios := []User{admin, editor}

	for _, usuario := range usuarios {
		fmt.Println(auth(usuario))
	}
}
