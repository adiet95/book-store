package interfaces

import (
	"context"
	"github.com/adiet95/book-store/author-service/src/database/models"
	"github.com/adiet95/book-store/author-service/src/libs"
)

type AuthorRepo interface {
	FindAll(limit, offset int) (*models.Authors, error)
	Save(data *models.Author) (*models.Author, error)
	Update(data *models.Author, id int) (*models.Author, error)
	Delete(id int) (*models.Author, error)
	FindByName(name string) (*models.Authors, error)
	FindById(id int) (*models.Author, error)
	GetRedisKey(ctx context.Context, redisKey string) (*models.Author, error)
	StoreRedisKey(ctx context.Context, redisKey string, data models.Author) (*models.Author, error)
	DeleteRedisKey(ctx context.Context, redisKey string) error
}

type AuthorService interface {
	GetAll(limit, offset int) *libs.Response
	Add(data *models.Author) *libs.Response
	Update(ctx context.Context, data *models.Author, id int) *libs.Response
	Delete(ctx context.Context, id int) *libs.Response
	Search(name string) *libs.Response
	SearchId(ctx context.Context, id int) *libs.Response
}
