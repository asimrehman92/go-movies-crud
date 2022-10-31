package main

import (
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"log"
	"math/rand"
	"net/http" // used to access the request and response object of the api
	"strconv"  // package used to covert string into int type

	"github.com/gorilla/mux" // used to get the params from the route
)

//mux is the pakg for creating these routes
// response format
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

//passing a point of request that you will send
//from your postman to this function
//and w is the response writer so when we send a response from this function
//it will be w
func getMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	//basically we want to set the content type as json, the
	//thing is that our golang struct needs to be able to convert
	//the json coming into it
	// send the response
	json.NewEncoder(w).Encode(movies)
}

//when we are deleting a movie we are actually passing
//an id of a movie
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//now get some params that id that pass from postman which
	//will go as a paramsto this function delete movie
	//and that param which will be the id present inside mux.vars(r)
	//inside the request which will be part of the request
	//the pointer to the request sending here
	// get the userid from the request params, key is "id"

	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	//return all the existing movies means remaining movies
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-TYpe", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	//while creating a movie we will send something in a body the entire
	//movie we will goint send so now we want to decode the body
	//decode that json body

	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	//pseudo code of our steps
	//set json content type
	w.Header().Set("Content-Type", "application/json")
	//params acess
	params := mux.Vars(r)
	//loop over or range over thge movies
	//delete the movie with th id that you have sent
	//add a new moview the movie that we send in the body of position

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}

func main() {
	//We are using gorilla/mux to create the router
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438337", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})
	//handlefunc create get movies route and the method using is get
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	// log for error if server doesnt start
	//to strat the server listen and serve
	log.Fatal(http.ListenAndServe(":8000", r))
}

// when we goto our postman means when we will hit the server or when we hit the api /movies and we want to get the movies
//in the beginning there wont be any movie so we want that we want couople of movies
//so we know our appi working properly
