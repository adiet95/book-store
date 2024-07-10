package author

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/adiet95/book-store/author-service/src/database/models"
	"github.com/adiet95/book-store/author-service/src/interfaces"
	"github.com/adiet95/book-store/author-service/src/libs"
	"github.com/gin-gonic/gin"
)

type author_ctrl struct {
	svc interfaces.AuthorService
}

func NewCtrl(reps interfaces.AuthorService) *author_ctrl {
	return &author_ctrl{svc: reps}
}

func (re *author_ctrl) GetAll(c *gin.Context) {
	v := c.Request.URL.Query().Get("limit")
	limit, _ := strconv.Atoi(v)
	if limit == 0 {
		limit = 10
	}

	val := c.Request.URL.Query().Get("offset")
	offset, _ := strconv.Atoi(val)

	re.svc.GetAll(limit, offset).Send(c)
}

func (re *author_ctrl) Add(c *gin.Context) {
	var data models.Author
	err := json.NewDecoder(c.Request.Body).Decode(&data)
	if err != nil {
		libs.New(err.Error(), 400, true)
		c.Abort()
	}
	re.svc.Add(&data).Send(c)
}

func (re *author_ctrl) Update(c *gin.Context) {
	val := c.Param("id")
	id, _ := strconv.Atoi(val)

	var datas models.Author
	err := json.NewDecoder(c.Request.Body).Decode(&datas)
	if err != nil {
		libs.New(err.Error(), 400, true)
		c.Abort()
	}
	re.svc.Update(&datas, id).Send(c)
}

func (re *author_ctrl) Delete(c *gin.Context) {
	val := c.Param("id")
	v, _ := strconv.Atoi(val)

	re.svc.Delete(v).Send(c)
}

func (re *author_ctrl) Search(c *gin.Context) {
	val := c.Request.URL.Query().Get("full_name")
	v := strings.ToLower(val)
	re.svc.Search(v).Send(c)
}

func (re *author_ctrl) SearchId(c *gin.Context) {
	val := c.Param("id")
	v, _ := strconv.Atoi(val)
	re.svc.SearchId(v).Send(c)
}
