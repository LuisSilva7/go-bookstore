package main

import (
	"log"
	"net/http"

	"github.com/LuisSilva7/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Printf("Server starting on localhost:8888")
	log.Fatal(http.ListenAndServe("localhost:8888", r))
}
