package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"bookstore/internal/models"
	"bookstore/internal/utils"
)

var (
	NewBook models.Book
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	// Create an instance of Book
	create := &models.Book{}
	utils.ParseBody(r, create)

	book := create.CreateBook()
	if book == nil {
		http.Error(w, "failed to create book", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "failed to marshal response", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := NewBook.GetBooks()
	res, err := json.Marshal(books)
	if err != nil {
		http.Error(w, "failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		http.Error(w, "invalid book ID", http.StatusBadRequest)
		return
	}

	book, result := NewBook.GetBookByID(ID)
	if result.Error != nil {
		http.Error(w, "failed to retrieve book", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	update := &models.Book{}
	utils.ParseBody(r, update)

	vars := mux.Vars(r)
	bookID := vars["id"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		http.Error(w, "invalid book ID", http.StatusBadRequest)
		return
	}

	book, result := update.GetBookByID(ID)
	if result.Error != nil {
		http.Error(w, "failed to retrieve book", http.StatusInternalServerError)
		return
	}
	if book != nil {
		book.Author = update.Author
		book.Name = update.Name
		book.Publication = update.Publication

		book.UpdateBook()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		http.Error(w, "invalid book ID", http.StatusBadRequest)
		return
	}

	book := NewBook.DeleteBook(ID)
	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
