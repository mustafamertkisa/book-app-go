package service

import (
	"book-app-go/common/validations"
	"book-app-go/database/model"
	"book-app-go/database/repository"
	"book-app-go/service/dto"
	"errors"
)

type IBookService interface {
	GetAllBooks() []model.Book
	GetBooksByAuthor(authorName string) []model.Book
	GetBookById(bookId int64) (model.Book, error)
	AddBook(bookCreate dto.BookCreate) error
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

func (s *BookService) AddBook(bookCreate dto.BookCreate) error {
	if err := validations.ValidateBookCreate(bookCreate); err != nil {
		return err
	}

	return s.bookRepository.AddBook(model.Book{
		Name:   bookCreate.Name,
		Pages:  bookCreate.Pages,
		Author: bookCreate.Author,
	})
}

func (s *BookService) DeleteBookById(bookId int64) error {
	return s.bookRepository.DeleteBookById(bookId)
}

func (s *BookService) UpdateBookPages(bookId int64, newPages int32) error {
	if err := validations.ValidatePages(newPages); err != nil {
		return err
	}

	checkBookExistErr := s.checkBookExist(bookId)
	if checkBookExistErr != nil {
		return errors.New("book not found")
	}

	return s.bookRepository.UpdateBookPages(bookId, newPages)
}

func (s *BookService) checkBookExist(bookId int64) error {
	_, err := s.bookRepository.GetBookById(bookId)
	if err != nil {
		return err
	}

	return nil
}
