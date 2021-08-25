package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Curso struct {
	Titulo       string `json:"titulo"`
	NumeroVideos int    `json:"numeroVideos"`
}
type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		curso := Cursos{
			Curso{"Golang", 15},
			Curso{"JS", 30},
			Curso{"Java", 20},
		}
		fmt.Println(curso)
		json.NewEncoder(w).Encode(curso)

	})
	http.ListenAndServe(":8000", nil)
}
