package main

import "fmt"

func main() {

	var arreglo [10]int

	fmt.Println(arreglo)

	arreglo2 := [3]int{1, 2}
	fmt.Println(arreglo2)
	fmt.Println(len(arreglo2))
	arreglo2[2] = 30

	for i := 0; i < len(arreglo2); i++ {

		fmt.Println(arreglo2[i])

	}

}
