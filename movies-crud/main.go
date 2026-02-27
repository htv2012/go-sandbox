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
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func updateMovieFields(item Movie, newItem Movie) Movie {
	if newItem.Title != "" {
		item.Title = newItem.Title
	}
	if newItem.Isbn != "" {
		item.Isbn = newItem.Isbn
	}
	if newItem.Director != nil {
		if item.Director == nil {
			item.Director = &Director{}
		}
		if newItem.Director.FirstName != "" {
			item.Director.FirstName = newItem.Director.FirstName
		}
		if newItem.Director.LastName != "" {
			item.Director.LastName = newItem.Director.LastName
		}
	}
	return item
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range movies {
		if item.ID == params["id"] {
			var newItem Movie
			json.NewDecoder(r.Body).Decode(&newItem)
			item = updateMovieFields(item, newItem)
			movies[i] = item
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	movies = append(movies, Movie{
		ID:       "1",
		Isbn:     "43210",
		Title:    "One Live",
		Director: &Director{FirstName: "John", LastName: "Doe"},
	})
	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "32110",
		Title: "Just the Two of Us",
	})

	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Print("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
