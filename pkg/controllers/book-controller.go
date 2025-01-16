package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/LuisSilva7/go-bookstore/pkg/models"
	"github.com/LuisSilva7/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody(r, newBook)

	book, err := newBook.CreateBook()
	if err != nil {
		res, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
		return
	}

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := models.GetAllBooks()
	if err != nil {
		res, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
		return
	}

	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["bookID"]

	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		res, _ := json.Marshal(map[string]string{"error": "Invalid ID"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}

	book, err := models.GetBookByID(ID)
	if err != nil {
		res, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
		return
	}

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updatedBook := &models.Book{}
	utils.ParseBody(r, updatedBook)

	vars := mux.Vars(r)
	bookID := vars["bookID"]

	ID, err := strconv.ParseInt(bookID, 10, 64)
	if err != nil {
		res, _ := json.Marshal(map[string]string{"error": "Invalid ID"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}

	oldBook, err := models.GetBookByID(ID)
	if err != nil {
		res, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
		return
	}

	if updatedBook.Name != "" {
		oldBook.Name = updatedBook.Name
	}
	if updatedBook.Author != "" {
		oldBook.Author = updatedBook.Author
	}
	if updatedBook.Publication != "" {
		oldBook.Publication = updatedBook.Publication
	}

	finalBook, err := oldBook.UpdateBook()
	if err != nil {
		res, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
		return
	}

	res, _ := json.Marshal(finalBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["bookID"]

	ID, err := strconv.ParseInt(bookID, 10, 64)
	if err != nil {
		res, _ := json.Marshal(map[string]string{"error": "Invalid ID"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}

	_, err = models.GetBookByID(ID)
	if err != nil {
		res, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
		return
	}

	book, err := models.DeleteBook(ID)
	if err != nil {
		res, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
		return
	}

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
