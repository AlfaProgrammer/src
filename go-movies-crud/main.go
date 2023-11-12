package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "math/random"
	// "strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()
	//addig the first movies manually
	movies = append(movies, Movie{ID: "1", Isbn: "4568225", Title: "Movie1", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "4568542", Title: "Movie2", Director: &Director{Firstname: "Porco", Lastname: "Dio"}})
	//building endpoints
	r.HandleFunc("/movies", getMovies).Methods("GET")
	// r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	// r.HandleFunc("/movies", createMovie).Methods("POST")
	// r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	// r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
