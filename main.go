package main

import (
	"log"
	"net/http"

	"books.com/handler"
	"books.com/inventory"
	"github.com/gorilla/mux"
)

func main() {
	inventory := inventory.NewInventory()
	handler := handler.NewHandler(inventory)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/AddBook/{title}/{stock}", handler.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/ListAllBooks", handler.ListAllBooks)
	router.HandleFunc("/EditBook/{title}/{newTitle}/{stock}", handler.EditBook).Methods(http.MethodPost)
	router.HandleFunc("/GetBookByTitle/{title}", handler.GetBookByTitle)

	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatalln("ListenAndServer Error", err)
	}
}
