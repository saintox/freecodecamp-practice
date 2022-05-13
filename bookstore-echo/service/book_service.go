package service

import (
	"context"
	"errors"
	"github.com/labstack/gommon/log"
	"practice-go/entity"
	"practice-go/models"
)

func (b *sqlService) InsertBook(input entity.BookInput) (bool, error) {
	ctx := context.Background()

	books, errValidation := b.sqlRepository.FindByTitleAndAuthor(ctx, input.Title, input.Author)

	if books.ID != 0 {
		return false, errors.New("buku sudah ada di katalog")
	}

	if errValidation != nil {
		return false, errValidation
	}

	book := models.Book{
		Title:       input.Title,
		Author:      input.Author,
		Description: input.Description,
	}

	_, errQuery := b.sqlRepository.InsertBook(ctx, book)

	if errQuery != nil {
		log.Printf(errQuery.Error())
		return false, errQuery
	}

	return true, nil
}
