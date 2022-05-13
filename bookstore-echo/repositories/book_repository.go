package repositories

import (
	"context"
	"practice-go/models"
)

func (r *repository) InsertBook(ctx context.Context, book models.Book) (models.Book, error) {
	query := `INSERT INTO books(title, author, description)
				VALUES (?,?,?)`

	result, errQuery := r.db.ExecContext(ctx, query, book.Title, book.Author, book.Description)

	if errQuery != nil {
		return book, errQuery
	}

	id, errInsert := result.LastInsertId()

	if errInsert != nil {
		return book, errInsert
	}

	book.ID = int(id)

	return book, nil
}

func (r *repository) FindByTitleAndAuthor(ctx context.Context, title string, author string) (models.Book, error) {
	book := models.Book{}

	query := `SELECT title,author
				FROM books
				WHERE books.title = ? AND books.author = ?
				LIMIT 1`

	result, errQuery := r.db.QueryContext(ctx, query, title, author)

	if errQuery != nil {
		return book, errQuery
	}

	defer result.Close()

	if result.Next() {
		result.Scan(
			&book.Title,
			&book.Author,
		)
		return book, nil
	}

	return book, nil
}
