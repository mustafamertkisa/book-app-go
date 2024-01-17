package request

import (
	"book-app-go/service/dto"
)

type AddBookRequest struct {
	Name   string `json:"name"`
	Pages  int32  `json:"pages"`
	Author string `json:"author"`
}

func (addBookRequest AddBookRequest) ParseModel() dto.BookCreate {
	return dto.BookCreate{
		Name:   addBookRequest.Name,
		Pages:  addBookRequest.Pages,
		Author: addBookRequest.Author,
	}
}
