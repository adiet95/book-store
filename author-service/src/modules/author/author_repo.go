package author

import (
	"context"
	"errors"
	"fmt"
	"github.com/adiet95/book-store/author-service/src/libs"
	"github.com/go-redis/redis/v8"
	"time"

	"github.com/adiet95/book-store/author-service/src/database/models"
	"gorm.io/gorm"
)

type author_repo struct {
	db          *gorm.DB
	clientRedis *redis.Client
}

func NewRepo(db *gorm.DB, clientRedis *redis.Client) *author_repo {
	return &author_repo{
		db:          db,
		clientRedis: clientRedis,
	}
}

func (r *author_repo) FindAll(limit, offset int) (*models.Authors, error) {
	var datas *models.Authors
	result := r.db.Model(&datas).Limit(limit).Offset(offset).Find(&datas)
	if result.Error != nil {
		return nil, errors.New("failed obtain datas")
	}
	return datas, nil
}

func (r *author_repo) Save(data *models.Author) (*models.Author, error) {
	var datas models.Authors
	checkName := r.db.Where("LOWER(full_name) LIKE ?", "%"+data.FullName+"%").Find(&datas)
	if checkName.RowsAffected != 0 {
		return nil, errors.New("fullname is exist")
	}

	res := r.db.Create(data)
	if res.Error != nil {
		return nil, errors.New("failed obtain datas")
	}
	return data, nil
}

func (re *author_repo) Update(data *models.Author, id int) (*models.Author, error) {
	res := re.db.Model(&data).Where("author_id = ?", id).Updates(&data)
	if res.Error != nil {
		return nil, errors.New("failed to update data")
	}
	return data, nil
}

func (re *author_repo) Delete(id int) (*models.Author, error) {
	var data *models.Author
	var datas *models.Authors
	res := re.db.Where("author_id = ?", id).Find(&datas)

	if res.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	r := re.db.Model(data).Where("author_id = ?", id).Delete(&data)
	if r.Error != nil {
		return nil, errors.New("failed to delete data")
	}
	return nil, nil
}

func (re *author_repo) FindByName(name string) (*models.Authors, error) {
	var datas *models.Authors
	res := re.db.Order("author_id asc").Where("LOWER(full_name) LIKE ?", "%"+name+"%").Find(&datas)
	if res.Error != nil {
		return nil, errors.New("failed to found data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

func (re *author_repo) FindById(id int) (*models.Author, error) {
	var data *models.Author
	res := re.db.Order("author_id asc").Where("author_id = ?", id).First(&data)
	if res.Error != nil {
		return nil, errors.New("failed to found data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	return data, nil
}

func (re *author_repo) GetRedisKey(ctx context.Context, redisKey string) (*models.Author, error) {
	var cacheKey = fmt.Sprintf("%v:%v", "author-service", redisKey)
	var result models.Author
	payloadBytes, errGetData := re.clientRedis.Get(ctx, cacheKey).Bytes()
	if errGetData != nil {
		return nil, redis.Nil
	}

	if errFromJSON := libs.FromJSON(payloadBytes, &result); errFromJSON != nil {
		return nil, errors.New("failed json redis")
	}
	return &result, nil
}

func (re *author_repo) StoreRedisKey(ctx context.Context, redisKey string, data models.Author) (*models.Author, error) {
	var cacheKey = fmt.Sprintf("%v:%v", "author-service", redisKey)
	var expiredAt = 1 * time.Hour
	var payload, errConvert = libs.Stringify(data)
	if errConvert != nil {
		return nil, errors.New("failed convert json redis")
	}

	err := re.clientRedis.Set(ctx, cacheKey, payload, expiredAt).Err()
	if err != nil {
		return nil, errors.New("err store redis")

	}
	return &data, nil
}

func (re *author_repo) DeleteRedisKey(ctx context.Context, redisKey string) error {
	err := re.clientRedis.Del(ctx, redisKey).Err()
	if err != nil {
		return err
	}
	return nil
}
