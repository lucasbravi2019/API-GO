package api

import (
	"context"
	"encoding/json"
	"net/http"
	"pasteleria/core"

	"go.mongodb.org/mongo-driver/mongo"
)

var routes = core.Routes{
	core.Route{
		Path:    "/",
		Handler: HelloWorld,
		Method:  "GET",
	},
	core.Route{
		Path:    "/movies",
		Handler: GetMovies,
		Method:  "GET",
	},
	core.Route{
		Path:    "/movies",
		Handler: CreateMovie,
		Method:  "POST",
	},
}

type Movie struct {
	Name     string `json:"name"`
	Year     int    `json:"year"`
	Director string `json:"director"`
}

var db *mongo.Database = core.GetDBConnection()

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello world")
}

func GetRoutes() core.Routes {
	return routes
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	var movies []string = []string{"Harry Potter", "Senior de los anillos", "Como si fuera cierto"}
	json.NewEncoder(w).Encode(movies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	newMovie := json.NewDecoder(r.Body)

	defer r.Body.Close()

	var document Movie

	err := newMovie.Decode(&document)

	if err != nil {
		panic(err)
	}

	db.Collection("movies").InsertOne(context.TODO(), document)
}
