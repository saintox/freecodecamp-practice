package repositories

import (
	"context"
	"database/sql"
	"practice-go/configs"
	"practice-go/models"
)

type (
	repository struct {
		db *sql.DB
	}

	SqlRepository interface {
		InsertBook(ctx context.Context, book models.Book) (models.Book, error)
		FindByTitleAndAuthor(ctx context.Context, title string, author string) (models.Book, error)
		CreateUser(ctx context.Context, user models.User) (models.User, error)
		FindByEmail(ctx context.Context, email string) (models.User, error)
		//ReadUser(ctx context.Context, id int) (models.User, error)
		//UpdateUser(ctx context.Context, id int, user models.User) (bool, error)
		//DeleteUser(ctx context.Context, id int) (bool, error)
	}
)

func NewRepository() *repository {
	return &repository{configs.GetConnection()}
}
