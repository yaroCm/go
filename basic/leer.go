package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var edad int
	fmt.Println("Coloca tu edad: ")
	fmt.Scanf("%d\n", &edad)

	fmt.Println("Mi edad es ", edad)

	//version 2
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre ")
	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Tu nombre es: " + text)
	}

}
