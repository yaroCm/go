package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type mensaje struct {
	msg string
}

func (m mensaje) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, m.msg)
}

func main() {

	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "<h1>Hola mundo</h1>")

		})
	*/
	msg := mensaje{
		msg: "Hola mundo",
	}
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))

	mux.Handle("/", fs)

	mux.HandleFunc("/prueba", HandlerPrueba)

	mux.HandleFunc("/usuario", HandlerUsuario)

	mux.Handle("/hola", msg)

	//http.ListenAndServe(":8000", mux)

	server := http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 2,
	}
	log.Println("Listenning")
	log.Fatal(server.ListenAndServe())
}

func HandlerPrueba(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hola desde prueba</h1>")

}

func HandlerUsuario(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hola desde usuario</h1>")

}
