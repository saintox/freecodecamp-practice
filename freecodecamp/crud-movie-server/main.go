package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID          string    `json:"id"`
	Isbn        string    `json:"isbn"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Director    *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Gender    int8   `json:"gender"`
}

var movies []Movie

func getMoviesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func getMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var movie Movie

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(100))
	movie.Isbn = strconv.Itoa(rand.Intn(999999))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func updateMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, movie := range movies {
		if movie.ID == params["id"] {
			var movie Movie
			movies = append(movies[:index], movies[index+1:]...)

			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]

			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			break
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "325027", Title: "Movie One", Description: "Movie One Description", Director: &Director{"Mamang", "Racing", 1}})
	movies = append(movies, Movie{ID: "2", Isbn: "482064", Title: "Movie Two", Description: "Movie Two Description", Director: &Director{"Gigachad", "Mamang", 0}})

	r.HandleFunc("/movies", getMoviesHandler).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovieHandler).Methods("GET")
	r.HandleFunc("/movies", createMovieHandler).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovieHandler).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovieHandler).Methods("DELETE")

	fmt.Printf("Starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
