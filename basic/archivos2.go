package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	archivo, err := os.Open("./hola2.txt")
	//SE EJECUTA AL FINAL SI O SI
	defer archivo.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		linea := scanner.Text()
		fmt.Println(linea)
	}

}
