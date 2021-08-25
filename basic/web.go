package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/jordi", func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "Jordi Cruz") })
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hola mundo")
	fmt.Println("El servidor corriendo")
}
