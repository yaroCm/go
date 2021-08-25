package main

import "fmt"

func main() {
	matriz := []int{1, 2, 3, 4}
	fmt.Println(matriz)

	arreglo := [3]int{1, 2, 3}
	slice := arreglo[:3]
	fmt.Println(slice)
}
