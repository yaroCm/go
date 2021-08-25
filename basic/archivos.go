package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	file_data, err := ioutil.ReadFile("./hola.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file_data))

}
