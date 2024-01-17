package validations

import (
	"book-app-go/service/dto"
	"errors"
)

func ValidateBookCreate(book dto.BookCreate) error {
	if book.Name == "" {
		return errors.New("book name cannot be empty")
	}

	if book.Author == "" {
		return errors.New("book author cannot be empty")
	}

	return nil
}

func ValidatePages(pages int32) error {
	if pages <= 0 {
		return errors.New("number of pages must be greater than 0")
	}

	return nil
}
