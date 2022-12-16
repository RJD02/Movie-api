package router

import (
	"github.com/RJD02/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	r.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movie/{id}", controller.MarkWatched).Methods("PUT")
	r.HandleFunc("/api/movie/{id}", controller.DeleteMovie).Methods("DELETE")
	r.HandleFunc("/api/movies", controller.DeleteMovies).Methods("DELETE")

	return r
}
