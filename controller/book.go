package controller

import (
	"book-app-go/controller/request"
	"book-app-go/controller/response"
	"book-app-go/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	bookService service.IBookService
}

func NewBookController(bookService service.IBookService) *BookController {
	return &BookController{
		bookService: bookService,
	}
}

func (bookController *BookController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/books/:id", bookController.GetBookById)
	e.GET("/api/v1/books", bookController.GetAllBooks)
	e.POST("/api/v1/books", bookController.AddBook)
	e.PUT("/api/v1/books/:id", bookController.UpdateBookPages)
	e.DELETE("/api/v1/books/:id", bookController.DeleteBookById)
}

func (bookController *BookController) GetBookById(c echo.Context) error {
	param := c.Param("id")
	productId, _ := strconv.Atoi(param)

	book, err := bookController.bookService.GetBookById(int64(productId))
	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			ErrorDescription: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.ParseResponse(book))
}

func (bookController *BookController) GetAllBooks(c echo.Context) error {
	author := c.QueryParam("author")
	if len(author) == 0 {
		allBooks := bookController.bookService.GetAllBooks()
		return c.JSON(http.StatusOK, response.ParseResponseList(allBooks))
	}

	authorBooks := bookController.bookService.GetBooksByAuthor(author)

	return c.JSON(http.StatusOK, response.ParseResponseList(authorBooks))
}

func (bookController *BookController) AddBook(c echo.Context) error {
	var addBookRequest request.AddBookRequest

	bindErr := c.Bind(&addBookRequest)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: bindErr.Error(),
		})
	}

	err := bookController.bookService.AddBook(addBookRequest.ParseModel())
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
			ErrorDescription: err.Error(),
		})
	}

	return c.NoContent(http.StatusCreated)
}

func (bookController *BookController) UpdateBookPages(c echo.Context) error {
	param := c.Param("id")
	productId, _ := strconv.Atoi(param)

	newPages := c.QueryParam("newPages")
	if len(newPages) == 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: "Parameter newPages is required",
		})
	}

	convertedPages, err := strconv.ParseInt(newPages, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: "NewPrice Format Disrupted!",
		})
	}

	bookController.bookService.UpdateBookPages(int64(productId), int32(convertedPages))

	return c.NoContent(http.StatusOK)
}

func (bookController *BookController) DeleteBookById(c echo.Context) error {
	param := c.Param("id")
	productId, _ := strconv.Atoi(param)

	err := bookController.bookService.DeleteBookById(int64(productId))
	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			ErrorDescription: err.Error(),
		})
	}

	return c.NoContent(http.StatusOK)
}
