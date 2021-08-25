package main

import "fmt"

func main() {
	slice := make([]string, 3, 5)
	fmt.Println(cap(slice))

	copia := []int{1, 2, 3, 4}
	destino := make([]int, 4)
	copy(destino, copia)

	fmt.Println(destino)
}
