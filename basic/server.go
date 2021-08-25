package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, "index.html")
}
