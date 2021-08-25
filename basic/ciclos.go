package main

import "fmt"

func main() {
	//for normal
	for i := 0; i < 10; i++ {

		fmt.Println(i)
	}

	//for convertido en while
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//detener ciclo
	i = 0
	for i < 10 {
		fmt.Println(i)
		i++
		break
	}

}
