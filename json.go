package main

import (
	"encoding/json"
	"net/http"
)

type Curso struct {
	Title        string `json:"Titulo"`
	NumberVideos int    `json:"Numero_videos"`
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		cursos := Cursos{
			Curso{"Curso de Go", 30},
			Curso{"Curso de php", 20},
			Curso{"Curso de nodejs", 15},
			Curso{"Curso de swift", 40},
			Curso{"Curso de JS", 40},
		}

		json.NewEncoder(w).Encode(cursos)
	})
	http.ListenAndServe(":8000", nil)
}
