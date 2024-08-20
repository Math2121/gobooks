package service

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Book struct {
	Title  string
	ID     int
	Author string
	Genre  string
}

type BookService struct {
	db *sql.DB
}

func NewBookService(db *sql.DB) *BookService {
	return &BookService{db: db}
}
func (s *BookService) CreateBook(book *Book) error {
	query := "INSERT INTO books (title, author, genre) VALUES (?, ?, ?)"

	result, err := s.db.Exec(query, book.Title, book.Author, book.Genre)

	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	book.ID = int(id)

	return nil
}

func (s *BookService) GetBooks() ([]Book, error) {
	rows, err := s.db.Query("SELECT id, title, author, genre FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (s *BookService) GetBookByID(id int) (*Book, error) {

	row := s.db.QueryRow("SELECT id, title, author, genre FROM books WHERE id =?", id)
	var book Book
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
	if err == sql.ErrNoRows {
		return nil, errors.New("book not found")
	} else if err != nil {
		return nil, err
	}
	return &book, nil
}

func (s *BookService) UpdateBook(book *Book) error {
	query := "UPDATE books SET title =?, author =?, genre =? WHERE id =?"

	_, err := s.db.Exec(query, book.Title, book.Author, book.Genre, book.ID)
	if err == sql.ErrNoRows {
		return errors.New("book not found")
	} else if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	return nil
}

func (s *BookService) SimulateReading(bookId int, duration time.Duration, results chan<- string) {
	book, err := s.GetBookByID(bookId)
	if err != nil {
		results <- fmt.Sprintf("Book not found")
		return
	}

	time.Sleep(duration)
	results <- fmt.Sprintf("Finished reading book %s by %s", book.Title, book.Author)

}

func (s *BookService) SimulateMultipleReadings(bookId []int, duration time.Duration) []string {
	results := make(chan string, len(bookId))

	for _, id := range bookId {
		go func(bookId int) {
			s.SimulateReading(bookId, duration, results)
		}(id)
	}

	var readings []string
	for range bookId {
		readings = append(readings, <-results)
	}
	close(results)
	return readings
}

func (s *BookService) SearchBookByName(name string) ([]Book, error) {
	rows, err := s.db.Query("SELECT id, title, author, genre FROM books WHERE title LIKE?", "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}
