package book

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/adiet95/book-store/book-service/src/database/models"
	"github.com/adiet95/book-store/book-service/src/interfaces"
	"github.com/adiet95/book-store/book-service/src/libs"
	"github.com/gin-gonic/gin"
)

type book_ctrl struct {
	svc interfaces.BookService
}

func NewCtrl(reps interfaces.BookService) *book_ctrl {
	return &book_ctrl{svc: reps}
}

func (re *book_ctrl) GetAll(c *gin.Context) {
	v := c.Request.URL.Query().Get("limit")
	limit, _ := strconv.Atoi(v)
	if limit == 0 {
		limit = 10
	}

	val := c.Request.URL.Query().Get("offset")
	offset, _ := strconv.Atoi(val)

	re.svc.GetAll(limit, offset).Send(c)
}

func (re *book_ctrl) Add(c *gin.Context) {
	var data models.Book
	err := json.NewDecoder(c.Request.Body).Decode(&data)
	if err != nil {
		libs.New(err.Error(), 400, true)
		c.Abort()
	}
	re.svc.Add(&data).Send(c)
}

func (re *book_ctrl) Update(c *gin.Context) {
	val := c.Param("id")
	id, _ := strconv.Atoi(val)

	var datas models.Book
	err := json.NewDecoder(c.Request.Body).Decode(&datas)
	if err != nil {
		libs.New(err.Error(), 400, true)
		c.Abort()
	}
	re.svc.Update(&datas, id).Send(c)
}

func (re *book_ctrl) Delete(c *gin.Context) {
	val := c.Param("id")
	v, _ := strconv.Atoi(val)

	re.svc.Delete(v).Send(c)
}

func (re *book_ctrl) Search(c *gin.Context) {
	val := c.Request.URL.Query().Get("name")
	v := strings.ToLower(val)
	re.svc.Search(v).Send(c)
}

func (re *book_ctrl) SearchId(c *gin.Context) {
	val := c.Param("id")
	v, _ := strconv.Atoi(val)
	re.svc.SearchId(v).Send(c)
}
