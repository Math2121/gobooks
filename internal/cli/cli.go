package cli

import (
	"fmt"
	"gobooks/internal/service"
	"os"
	"strconv"
	"time"
)

type BookCli struct {
	bookService *service.BookService
}

func NewBookCli(bookService *service.BookService) *BookCli {
	return &BookCli{bookService: bookService}
}

func (cli *BookCli) Run() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command: create, read, update, delete, or list")
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Usage books search")
			return
		}
		bookName := os.Args[2]
		cli.searchBooks(bookName)
	case "simulate":
		if len(os.Args) < 3 {
			fmt.Println("Usage books simulate <book_id> <duration>")
			return
		}
		booksIds := os.Args[2:]
		cli.simulatReading(booksIds)

	}
}

func (cli *BookCli) searchBooks(bookName string) {
	books, err := cli.bookService.SearchBookByName(bookName)
	if err != nil {
		fmt.Printf("Error searching books: %v\n", err)
		return
	}

	if len(books) == 0 {
		fmt.Println("No books found with that name.")
		return
	}

	for _, book := range books {
		fmt.Printf("%d: %s by %s (%s)\n", book.ID, book.Title, book.Author, book.Genre)
	}
}

func (cli *BookCli) simulatReading(bookIdStr []string) {
	var booksIdStr []int

	for _, idStr := range bookIdStr {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Printf("Invalid book ID: %s\n", idStr)
			continue
		}
		booksIdStr = append(booksIdStr, id)
	}
	responses := cli.bookService.SimulateMultipleReadings(booksIdStr, 2*time.Second)
	for _, response := range responses {
		fmt.Println(response)
	}
}
