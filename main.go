package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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
	// setting the content type of response to be in json
	w.Header().Set("content-Type", "application/json")

	// Encode the movies list into the response writer
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	// setting the content type of response to be in json
	w.Header().Set("content-Type", "application/json")

	// Get the request params
	params := mux.Vars(r)

	// Delete the movie in movies using the "id" from the request params
	for index, item := range movies {

		if item.ID == params["id"] {
			// calling the append version of appending multiple slices of a list,thus ...
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	// setting the content type of response to be in json
	w.Header().Set("content-Type", "application/json")

	// Get the request params
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	// setting the content type of response to be in json
	w.Header().Set("content-Type", "application/json")

	var movie Movie

	//Decode the json request into the Movie object movie,the details of the movie will be in body of that request
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// setting the content type of response to be in json
	w.Header().Set("content-Type", "application/json")

	var updatedMovie Movie
	//Decode the json request into the Movie object movie,the details of the movie will be in body of that request
	_ = json.NewDecoder(r.Body).Decode(&updatedMovie)
	updatedMovie.ID = strconv.Itoa(rand.Intn(10000000))

	params := mux.Vars(r)

	// Delete the movie in movies using the "id" from the request params
	for index, item := range movies {

		if item.ID == params["id"] {
			// calling the append version of appending multiple slices of a list,thus ...
			tmp := append(movies[:index], updatedMovie)
			movies = append(tmp, movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()

	//Enter a movie by hardcoding it
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "Arunkumar", Lastname: "M"}})
	movies = append(movies, Movie{ID: "2", Isbn: "56789", Title: "Movie two", Director: &Director{Firstname: "Lalitha", Lastname: "Murugesan"}})

	// Define handler funcs for each route: endpoints,functions,methodType
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	// Start server
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
