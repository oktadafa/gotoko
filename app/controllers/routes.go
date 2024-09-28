package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (server *Server) InitializeRoutes() {
	server.Router = mux.NewRouter()

	server.Router.HandleFunc("/", server.Home).Methods("GET")
	staticFileDirectory := http.Dir("./assets")
	setFileHandler := http.StripPrefix("/public/", http.FileServer(staticFileDirectory))
	server.Router.PathPrefix("/public/").Handler(setFileHandler).Methods("GET")

	server.Router.HandleFunc("/products", server.Products).Methods("GET")
}
