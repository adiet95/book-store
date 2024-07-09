package category

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/adiet95/book-store/category-service/src/database/models"
	"github.com/adiet95/book-store/category-service/src/interfaces"
	"github.com/adiet95/book-store/category-service/src/libs"
	"github.com/gin-gonic/gin"
)

type category_ctrl struct {
	svc interfaces.CategoryService
}

func NewCtrl(reps interfaces.CategoryService) *category_ctrl {
	return &category_ctrl{svc: reps}
}

func (re *category_ctrl) GetAll(c *gin.Context) {
	v := c.Request.URL.Query().Get("limit")
	limit, _ := strconv.Atoi(v)
	if limit == 0 {
		limit = 10
	}

	val := c.Request.URL.Query().Get("offset")
	offset, _ := strconv.Atoi(val)

	re.svc.GetAll(limit, offset).Send(c)
}

func (re *category_ctrl) Add(c *gin.Context) {
	var data models.Category
	err := json.NewDecoder(c.Request.Body).Decode(&data)
	if err != nil {
		libs.New(err.Error(), 400, true)
		c.Abort()
	}
	re.svc.Add(&data).Send(c)
}

func (re *category_ctrl) Update(c *gin.Context) {
	val := c.Request.URL.Query().Get("id")
	id, _ := strconv.Atoi(val)

	var datas models.Category
	err := json.NewDecoder(c.Request.Body).Decode(&datas)
	if err != nil {
		libs.New(err.Error(), 400, true)
		c.Abort()
	}
	re.svc.Update(&datas, id).Send(c)
}

func (re *category_ctrl) Delete(c *gin.Context) {
	val := c.Request.URL.Query().Get("id")
	v, _ := strconv.Atoi(val)

	re.svc.Delete(v).Send(c)
}

func (re *category_ctrl) Search(c *gin.Context) {
	val := c.Request.URL.Query().Get("name")
	v := strings.ToLower(val)
	re.svc.Search(v).Send(c)
}

func (re *category_ctrl) SearchId(c *gin.Context) {
	val := c.Request.URL.Query().Get("id")
	v, _ := strconv.Atoi(val)
	re.svc.SearchId(v).Send(c)
}
