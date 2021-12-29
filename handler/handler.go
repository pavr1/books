package handler

import (
	"net/http"
	"strconv"

	"books.com/inventory"
	"books.com/renderer"
	"github.com/gorilla/mux"
)

type Handler interface {
	AddBook(w http.ResponseWriter, r *http.Request)
	ListAllBooks(w http.ResponseWriter, r *http.Request)
	GetBookByTitle(w http.ResponseWriter, r *http.Request)
	EditBook(w http.ResponseWriter, r *http.Request)
}

type HandlerImpl struct {
	Inventory inventory.Inventory
}

func NewHandler(inventory inventory.Inventory) Handler {
	return HandlerImpl{
		Inventory: inventory,
	}
}

func (h HandlerImpl) AddBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	stockStr := vars["stock"]

	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		render := renderer.NewRenderer(http.StatusInternalServerError, err, nil)
		render.Render(w, r)
		return
	}

	err = h.Inventory.AddBook(title, stock)
	if err != nil {
		render := renderer.NewRenderer(http.StatusInternalServerError, err, nil)
		render.Render(w, r)
	} else {
		render := renderer.NewRenderer(http.StatusOK, nil, nil)
		render.Render(w, r)
	}
}

func (h HandlerImpl) ListAllBooks(w http.ResponseWriter, r *http.Request) {
	books := h.Inventory.ListAllBooks()

	render := renderer.NewRenderer(http.StatusOK, nil, books)
	render.Render(w, r)
}

func (h HandlerImpl) GetBookByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	book := h.Inventory.GetBookByTitle(title)

	render := renderer.NewRenderer(http.StatusOK, nil, book)
	render.Render(w, r)
}

func (h HandlerImpl) EditBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	newTitle := vars["newTitle"]
	stockStr := vars["stock"] //0

	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		render := renderer.NewRenderer(http.StatusInternalServerError, err, nil)
		render.Render(w, r)
		return
	}

	err = h.Inventory.EditBook(title, newTitle, stock)
	if err != nil {
		render := renderer.NewRenderer(http.StatusInternalServerError, err, nil)
		render.Render(w, r)
	} else {
		render := renderer.NewRenderer(http.StatusOK, nil, nil)
		render.Render(w, r)
	}
}
