package service

import (
	"book-app-go/common/validations"
	"book-app-go/database/model"
	"book-app-go/database/repository"
)

type IBookService interface {
	GetAllBooks() []model.Book
	GetBooksByAuthor(authorName string) []model.Book
	GetBookById(bookId int64) (model.Book, error)
	AddBook(book model.Book) error
	DeleteBookById(bookId int64) error
	UpdateBookPages(bookId int64, newPages int32) error
}

type BookService struct {
	bookRepository repository.IBookRepository
}

func NewBookService(bookRepository repository.IBookRepository) IBookService {
	return &BookService{bookRepository: bookRepository}
}

func (s *BookService) GetAllBooks() []model.Book {
	return s.bookRepository.GetAllBooks()
}

func (s *BookService) GetBooksByAuthor(authorName string) []model.Book {
	return s.bookRepository.GetBooksByAuthor(authorName)
}

func (s *BookService) GetBookById(bookId int64) (model.Book, error) {
	return s.bookRepository.GetBookById(bookId)
}

func (s *BookService) AddBook(book model.Book) error {
	if err := validations.ValidateBook(book); err != nil {
		return err
	}

	return s.bookRepository.AddBook(book)
}

func (s *BookService) DeleteBookById(bookId int64) error {
	return s.bookRepository.DeleteBookById(bookId)
}

func (s *BookService) UpdateBookPages(bookId int64, newPages int32) error {
	if err := validations.ValidatePages(newPages); err != nil {
		return err
	}

	return s.bookRepository.UpdateBookPages(bookId, newPages)
}
