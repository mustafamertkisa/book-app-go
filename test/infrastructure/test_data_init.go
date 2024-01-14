package infrastructure

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

var query = `INSERT INTO books (name, pages, author)
VALUES 
  ('To Kill a Mockingbird', 281, 'Harper Lee'),
  ('1984', 328, 'George Orwell'),
  ('Pride and Prejudice', 279, 'Jane Austen'),
  ('The Great Gatsby', 180, 'F. Scott Fitzgerald');
`

func TestDataInitialize(ctx context.Context, dbPool *pgxpool.Pool) {
	insertBooksResult, err := dbPool.Exec(ctx, query)
	if err != nil {
		log.Error(err)
	} else {
		log.Info(fmt.Sprintf("Books data created with %v rows", insertBooksResult.RowsAffected()))
	}
}
