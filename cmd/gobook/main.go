package main

import (
	"database/sql"
	"gobooks/internal/service"
	"gobooks/internal/web"
	"log"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
)


func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	if err!= nil {
        log.Fatal(err)
    }

	defer db.Close()

	bookService  := service.NewBookService(db)
	bookHandler := web.NewBookHandler(bookService)

	router := http.NewServeMux()
	router.HandleFunc("GET /api/books", bookHandler.GetBooks)
	router.HandleFunc("GET /api/books/{id}", bookHandler.GetBookByID)
	router.HandleFunc("POST /api/books", bookHandler.CreateBook)
	router.HandleFunc("PUT /api/books/{id}", bookHandler.UpdateBook)

    log.Fatal(http.ListenAndServe(":8080", router))  // set listening port to :8080
	
}