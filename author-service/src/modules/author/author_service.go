package author

import (
	"context"
	"fmt"
	"github.com/adiet95/book-store/author-service/src/database/models"
	"github.com/adiet95/book-store/author-service/src/interfaces"
	"github.com/adiet95/book-store/author-service/src/libs"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type author_service struct {
	author_repo interfaces.AuthorRepo
}

func NewService(reps interfaces.AuthorRepo) *author_service {
	return &author_service{reps}
}

func (r *author_service) GetAll(limit, offset int) *libs.Response {
	data, err := r.author_repo.FindAll(limit, offset)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *author_service) Add(data *models.Author) *libs.Response {
	result, err := re.author_repo.Save(data)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(result, 200, false)
}

func (re *author_service) Update(ctx context.Context, data *models.Author, id int) *libs.Response {
	key := strconv.Itoa(id)
	redisKey := fmt.Sprintf("%v:%v", "author-service", key)
	errDel := re.author_repo.DeleteRedisKey(ctx, redisKey)
	if errDel != nil {
		return libs.New(errDel.Error(), 400, true)
	}

	res, err := re.author_repo.Update(data, id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(res, 200, false)
}

func (re *author_service) Delete(ctx context.Context, id int) *libs.Response {
	key := strconv.Itoa(id)
	redisKey := fmt.Sprintf("%v:%v", "author-service", key)
	errDel := re.author_repo.DeleteRedisKey(ctx, redisKey)
	if errDel != nil {
		return libs.New(errDel.Error(), 400, true)
	}

	data, err := re.author_repo.Delete(id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *author_service) Search(name string) *libs.Response {
	data, err := re.author_repo.FindByName(name)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *author_service) SearchId(ctx context.Context, id int) *libs.Response {
	redisKey := strconv.Itoa(id)
	authorData, err := re.author_repo.GetRedisKey(ctx, redisKey)
	if err != nil {
		switch {
		case err == redis.Nil:
			dataStore, errFind := re.author_repo.FindById(id)
			if errFind != nil {
				return libs.New(errFind.Error(), 400, true)
			}
			data, errStore := re.author_repo.StoreRedisKey(ctx, redisKey, *dataStore)
			if errStore != nil {
				return libs.New(errStore.Error(), 400, true)
			}
			return libs.New(data, 200, false)
		default:
			return libs.New("err: redis not found", 400, true)
		}
	}
	return libs.New(authorData, 200, false)
}
