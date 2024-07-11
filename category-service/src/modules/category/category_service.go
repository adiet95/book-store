package category

import (
	"context"
	"fmt"
	"github.com/adiet95/book-store/category-service/src/database/models"
	"github.com/adiet95/book-store/category-service/src/interfaces"
	"github.com/adiet95/book-store/category-service/src/libs"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type category_service struct {
	category_repo interfaces.CategoryRepo
}

func NewService(reps interfaces.CategoryRepo) *category_service {
	return &category_service{reps}
}

func (r *category_service) GetAll(limit, offset int) *libs.Response {
	data, err := r.category_repo.FindAll(limit, offset)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *category_service) Add(data *models.Category) *libs.Response {
	result, err := re.category_repo.Save(data)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(result, 200, false)
}

func (re *category_service) Update(ctx context.Context, data *models.Category, id int) *libs.Response {
	redisKey := strconv.Itoa(id)
	var cacheKey = fmt.Sprintf("%v:%v", "category-service", redisKey)
	errDel := re.category_repo.DeleteRedisKey(ctx, cacheKey)
	if errDel != nil {
		return libs.New(errDel.Error(), 400, true)
	}

	res, err := re.category_repo.Update(data, id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(res, 200, false)
}

func (re *category_service) Delete(ctx context.Context, id int) *libs.Response {
	redisKey := strconv.Itoa(id)
	var cacheKey = fmt.Sprintf("%v:%v", "category-service", redisKey)
	errDel := re.category_repo.DeleteRedisKey(ctx, cacheKey)
	if errDel != nil {
		return libs.New(errDel.Error(), 400, true)
	}

	data, err := re.category_repo.Delete(id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *category_service) Search(name string) *libs.Response {
	data, err := re.category_repo.FindByName(name)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *category_service) SearchId(id int, ctx context.Context) *libs.Response {
	redisKey := strconv.Itoa(id)
	var cacheKey = fmt.Sprintf("%v:%v", "category-service", redisKey)
	categoryData, err := re.category_repo.GetRedisKey(ctx, cacheKey)
	if err != nil {
		switch {
		case err == redis.Nil:
			dataId, errFind := re.category_repo.FindById(id)
			if errFind != nil {
				return libs.New(errFind.Error(), 400, true)
			}
			data, errStore := re.category_repo.StoreRedisKey(ctx, cacheKey, *dataId)
			if errStore != nil {
				return libs.New(errStore.Error(), 400, true)
			}
			return libs.New(data, 200, false)
		default:
			return libs.New("err: redis not found", 400, true)
		}
	}
	return libs.New(categoryData, 200, false)
}
