package response

import "book-app-go/database/model"

type ErrorResponse struct {
	ErrorDescription string `json:"errorDescription"`
}

type BookResponse struct {
	Name   string `json:"name"`
	Pages  int32  `json:"pages"`
	Author string `json:"author"`
}

func ParseResponse(book model.Book) BookResponse {
	return BookResponse{
		Name:   book.Name,
		Pages:  book.Pages,
		Author: book.Author,
	}
}

func ParseResponseList(books []model.Book) []BookResponse {
	var bookResponseList = []BookResponse{}
	for _, book := range books {
		bookResponseList = append(bookResponseList, ParseResponse(book))
	}

	return bookResponseList
}
