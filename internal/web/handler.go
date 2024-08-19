package web

import (
	"encoding/json"
	"gobooks/internal/service"
	"net/http"
	"strconv"
)

type BookHandler struct {
	service *service.BookService


}
func NewBookHandler(service *service.BookService) *BookHandler {
    return &BookHandler{service: service}
}

func(h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request)  {
	books, err := h.service.GetBooks()
    if err != nil {
        http.Error(w, "Failed to get books", http.StatusInternalServerError)
		return 
    }
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func(h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request)  {
	var book service.Book
    err := json.NewDecoder(r.Body).Decode(&book)
    if err!= nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    err = h.service.CreateBook(&book)
    if err!= nil {
        http.Error(w, "Failed to create book", http.StatusInternalServerError)
        return
    }

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func(h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request)  {
	idQuery := r.PathValue("id")
	id, err := strconv.Atoi(idQuery)
	if err!= nil {
        http.Error(w, "Invalid book ID", http.StatusBadRequest)
        return
    }

    book, err := h.service.GetBookByID(id)
	if err!= nil {
        http.Error(w, "Failed to get book", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func(h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request)  {
	idQuery := r.PathValue("id")
    id, err := strconv.Atoi(idQuery)
    if err!= nil {
        http.Error(w, "Invalid book ID", http.StatusBadRequest)
        return
    }

    var book service.Book
    err = json.NewDecoder(r.Body).Decode(&book)
    if err!= nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    book.ID = id
    err = h.service.UpdateBook(&book)
    if err!= nil {
        http.Error(w, "Failed to update book", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(book)
}