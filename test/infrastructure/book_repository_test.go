package infrastructure

import (
	"book-app-go/common/postgresql"
	"book-app-go/database/model"
	"book-app-go/database/repository"
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

var bookRepository repository.IBookRepository
var dbPool *pgxpool.Pool
var ctx context.Context
var expectedBooks = []model.Book{
	{
		Id:     1,
		Name:   "To Kill a Mockingbird",
		Pages:  281,
		Author: "Harper Lee",
	},
	{
		Id:     2,
		Name:   "1984",
		Pages:  328,
		Author: "George Orwell",
	},
	{
		Id:     3,
		Name:   "Pride and Prejudice",
		Pages:  279,
		Author: "Jane Austen",
	},
	{
		Id:     4,
		Name:   "The Great Gatsby",
		Pages:  180,
		Author: "F. Scott Fitzgerald",
	},
}

func TestMain(m *testing.M) {
	ctx = context.Background()

	dbPool = postgresql.GetConnectionPool(ctx, postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		DbName:                "bookapp",
		Username:              "postgres",
		Password:              "postgres",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	})

	bookRepository = repository.NewBookRepository(dbPool)

	exitCode := m.Run()
	if exitCode == 0 {
		fmt.Println("✅ All tests passed!")
	} else {
		fmt.Println("❌ Some tests failed!")
	}

	os.Exit(exitCode)
}

func setup(ctx context.Context, dbPool *pgxpool.Pool) {
	TestDataInitialize(ctx, dbPool)
}
func clear(ctx context.Context, dbPool *pgxpool.Pool) {
	TruncateTestData(ctx, dbPool)
}

func TestGetAllBooks(t *testing.T) {
	setup(ctx, dbPool)

	t.Run("GetAllBooks", func(t *testing.T) {
		fetchedBooks := bookRepository.GetAllBooks()
		assert.Equal(t, 4, len(fetchedBooks))
		assert.Equal(t, expectedBooks, fetchedBooks)
	})

	clear(ctx, dbPool)
}

func TestGetBooksByAuthor(t *testing.T) {
	setup(ctx, dbPool)

	t.Run("GetBooksByAuthor", func(t *testing.T) {
		authorName := "Jane Austen"
		fetchedBooks := bookRepository.GetBooksByAuthor(authorName)
		assert.Equal(t, 1, len(fetchedBooks))
		assert.Equal(t, expectedBooks[2], fetchedBooks[0])
	})

	clear(ctx, dbPool)
}

func TestGetBooksById(t *testing.T) {
	setup(ctx, dbPool)

	t.Run("GetBooksById", func(t *testing.T) {
		var bookId int64 = 2
		fetchedBook, _ := bookRepository.GetBookById(bookId)
		assert.Equal(t, expectedBooks[1], fetchedBook)
	})

	clear(ctx, dbPool)
}

func TestAddBook(t *testing.T) {
	var expectedBook = model.Book{
		Id:     1,
		Name:   "White Fang",
		Pages:  256,
		Author: "Jack London",
	}

	var newBook = model.Book{
		Name:   "White Fang",
		Pages:  256,
		Author: "Jack London",
	}
	t.Run("AddBook", func(t *testing.T) {
		bookRepository.AddBook(newBook)
		fetchedBooks := bookRepository.GetAllBooks()

		assert.Equal(t, 1, len(fetchedBooks))
		assert.Equal(t, fetchedBooks[0], expectedBook)
	})

	clear(ctx, dbPool)
}

func TestDeleteById(t *testing.T) {
	setup(ctx, dbPool)
	t.Run("DeleteById", func(t *testing.T) {
		var bookId int64 = 1

		bookRepository.DeleteBookById(bookId)

		_, err := bookRepository.GetBookById(bookId)
		errMsg := fmt.Sprintf("Error while getting book with id %v", bookId)

		assert.Equal(t, errMsg, err.Error())
	})
	clear(ctx, dbPool)
}

func TestUpdateBookPages(t *testing.T) {
	setup(ctx, dbPool)
	t.Run("UpdateBookPages", func(t *testing.T) {
		var bookId int64 = 1
		var newPages int32 = 100

		beforeUpdateGetBook, _ := bookRepository.GetBookById(bookId)

		bookRepository.UpdateBookPages(bookId, newPages)

		afterUpdateGetBook, _ := bookRepository.GetBookById(bookId)

		assert.Equal(t, newPages, afterUpdateGetBook.Pages)
		assert.NotEqual(t, beforeUpdateGetBook.Pages, afterUpdateGetBook.Pages)
	})
	clear(ctx, dbPool)
}
