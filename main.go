package main

import (
	"encoding/json" // to parse json objects
	"fmt"
	"log"
	"math/rand" // to generate movie id
	"net/http"  // to create and start a server
	"strconv"   // id created using math.rand will be an int and we will have to convert it to string

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
	w.Header().Set("Content-Type", "application/json") // setting the header lets the receiver(browser or postman) know what kind of data it is going to receive.
	json.NewEncoder(w).Encode(movies)                  // create a new encoder which writes to response(w). It encodes movies slice and write it to response and send the data to browser or postman in the form of json

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req := mux.Vars(r)   // retrieve the contents from the body
	movieId := req["id"] // retrieve the value of "id" key from the body

	for index, value := range movies {
		if value.ID == movieId {
			movies = append(movies[:index], movies[index+1:]...) // deleting the movie whose id matches
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	movieId := params["id"]

	for _, value := range movies {
		if value.ID == movieId {
			json.NewEncoder(w).Encode(value)
			break
		}
	}

}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie // to store the incoming request body
	json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(100000000)) // converts the integral id into a string
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// set the headers
	w.Header().Set("Content-Type", "application/json")

	var UpdateMovie Movie
	json.NewDecoder(r.Body).Decode(&UpdateMovie)

	// get the params
	params := mux.Vars(r)
	movieId := params["id"]

	UpdateMovie.ID = movieId

	// loop over movies slice
	for index, value := range movies {
		if value.ID == movieId {
			// delete the movie and add updatedMovie
			movies = append(movies[:index], movies[index+1:]...)
			movies = append(movies, UpdateMovie)
			json.NewEncoder(w).Encode(UpdateMovie)
			return

		}
	}
}

func main() {

	movie1 := Movie{
		ID:       "m001",
		Isbn:     "978-0-123-4567-8",
		Title:    "The Shawshank Redemption",
		Director: &Director{Firstname: "Frank", Lastname: "Darabont"},
	}
	movie2 := Movie{
		ID:       "m002",
		Isbn:     "978-0-987-6543-2",
		Title:    "The Godfather",
		Director: &Director{Firstname: "Francis Ford", Lastname: "Coppola"},
	}
	movie3 := Movie{
		ID:       "m003",
		Isbn:     "978-1-234-5678-9",
		Title:    "The Dark Knight",
		Director: &Director{Firstname: "Christopher", Lastname: "Nolan"},
	}
	movie4 := Movie{
		ID:       "m004",
		Isbn:     "978-2-345-6789-0",
		Title:    "Pulp Fiction",
		Director: &Director{Firstname: "Quentin", Lastname: "Tarantino"},
	}
	movie5 := Movie{
		ID:       "m005",
		Isbn:     "978-3-456-7890-1",
		Title:    "The Lord of the Rings: The Fellowship of the Ring",
		Director: &Director{Firstname: "Peter", Lastname: "Jackson"},
	}

	r := mux.NewRouter()

	movies = append(movies, movie1, movie2, movie3, movie4, movie5)

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server on 8000")

	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}

}
