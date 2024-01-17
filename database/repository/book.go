package repository

import (
	"book-app-go/database/model"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type IBookRepository interface {
	GetAllBooks() []model.Book
	GetBooksByAuthor(authorName string) []model.Book
	GetBookById(bookId int64) (model.Book, error)
	AddBook(book model.Book) error
	DeleteBookById(bookId int64) error
	UpdateBookPages(bookId int64, newPages int32) error
}

type BookRepository struct {
	dbPool *pgxpool.Pool
}

func NewBookRepository(dbPool *pgxpool.Pool) IBookRepository {
	return &BookRepository{dbPool: dbPool}
}

func (bookRepository *BookRepository) GetAllBooks() []model.Book {
	ctx := context.Background()
	query := "Select * from books"

	bookRows, err := bookRepository.dbPool.Query(ctx, query)
	if err != nil {
		log.Error("Error while getting book rows")
		return []model.Book{}
	}

	return parseBook(bookRows)
}

func (bookRepository *BookRepository) GetBooksByAuthor(authorName string) []model.Book {
	ctx := context.Background()
	query := `Select * from books where author = $1`

	bookRows, err := bookRepository.dbPool.Query(ctx, query, authorName)
	if err != nil {
		log.Error("Error while getting book by author rows")
		return []model.Book{}
	}

	return parseBook(bookRows)
}

func (bookRepository *BookRepository) GetBookById(bookId int64) (model.Book, error) {
	ctx := context.Background()
	query := `Select * from books where id = $1`

	queryRow := bookRepository.dbPool.QueryRow(ctx, query, bookId)

	var id int64
	var pages int32
	var name, author string

	scanErr := queryRow.Scan(&id, &name, &pages, &author)

	if scanErr != nil {
		errMsg := fmt.Sprintf("Error while getting book with id %v", bookId)
		return model.Book{}, errors.New(errMsg)
	}

	return model.Book{
		Id:     id,
		Name:   name,
		Pages:  pages,
		Author: author,
	}, nil
}

func (bookRepository *BookRepository) AddBook(newBook model.Book) error {
	ctx := context.Background()
	query := `Insert into books (name,pages,author) VALUES ($1,$2,$3)`

	addNewBook, err := bookRepository.dbPool.Exec(ctx, query, newBook.Name, newBook.Pages, newBook.Author)
	if err != nil {
		log.Error("Failed to add new book", err)
		return err
	}

	log.Info(fmt.Sprintf("Book added with %v", addNewBook))

	return nil
}

func (bookRepository *BookRepository) DeleteBookById(bookId int64) error {
	ctx := context.Background()

	query := `Delete from books where id = $1`

	_, err := bookRepository.dbPool.Exec(ctx, query, bookId)
	if err != nil {
		errMsg := fmt.Sprintf("error while deleting book with id %v", bookId)
		return errors.New(errMsg)
	}

	log.Info("Book deleted")

	return nil
}

func (bookRepository *BookRepository) UpdateBookPages(bookId int64, newPages int32) error {
	ctx := context.Background()

	query := `Update books set pages = $1 where id = $2`

	_, err := bookRepository.dbPool.Exec(ctx, query, newPages, bookId)
	if err != nil {
		errMsg := fmt.Sprintf("error while updating book pages with id %v", bookId)
		return errors.New(errMsg)
	}

	message := fmt.Sprintf("Book %v pages updated with %v", bookId, newPages)
	log.Info(message)

	return nil
}

func parseBook(rows pgx.Rows) []model.Book {
	var books = []model.Book{}
	var id int64
	var pages int32
	var name, author string

	for rows.Next() {
		rows.Scan(&id, &name, &pages, &author)
		books = append(books, model.Book{
			Id:     id,
			Name:   name,
			Pages:  pages,
			Author: author,
		})
	}

	return books
}
