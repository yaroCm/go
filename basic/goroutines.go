package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	go slow("Jordi")
	fmt.Println("Que aburrido este código")
	var wait string
	fmt.Scanln(&wait)

}
func slow(name string) {
	letras := strings.Split(name, "")

	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
