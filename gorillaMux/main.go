package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/users", handlerGet).Methods("GET")
	r.HandleFunc("/api/users", handlerPost).Methods("POST")
	r.HandleFunc("/api/users", handlerPut).Methods("PUT")
	r.HandleFunc("/api/users", handlerDelete).Methods("DELETE")

	server := &http.Server{
		Addr:              ":8000",
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 2,
	}
	log.Println("Listenning")
	server.ListenAndServe()
}

func handlerGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola mundo con gorilla desde get")
}
func handlerPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola mundo con gorilla desde post")
}
func handlerPut(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola mundo con gorilla desde put")
}
func handlerDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola mundo con gorilla desde delete")
}
