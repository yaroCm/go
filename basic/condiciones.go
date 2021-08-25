package main

import "fmt"

func main() {
	x := 10
	y := 10
	if x == y {
		fmt.Println("El valor: ", x, "es igual que: ", y)
	} else if x > y {
		fmt.Println("El valor: ", x, "es mayor que: ", y)
	} else {
		fmt.Println("El valor: ", y, "es mayor que: ", x)
	}

}
