package users

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/adiet95/book-store/book-service/src/database/models"
	"github.com/adiet95/book-store/book-service/src/interfaces"
	"github.com/adiet95/book-store/book-service/src/libs"
	"github.com/gin-gonic/gin"
)

type user_ctrl struct {
	svc interfaces.UserService
}

func NewCtrl(reps interfaces.UserService) *user_ctrl {
	return &user_ctrl{svc: reps}
}

func (re *user_ctrl) GetAll(c *gin.Context) {
	claim_user, exist := c.Get("email")
	if !exist {
		libs.New("claim user is not exist", 400, true)
		c.Abort()
	}

	v := c.Request.URL.Query().Get("limit")
	limit, _ := strconv.Atoi(v)

	val := c.Request.URL.Query().Get("offset")
	offset, _ := strconv.Atoi(val)

	result := re.svc.FindEmail(claim_user.(string), limit, offset)
	result.Send(c)
}

func (re *user_ctrl) Add(c *gin.Context) {
	var data models.User
	err := json.NewDecoder(c.Request.Body).Decode(&data)
	if err != nil {
		libs.New(err.Error(), 400, true)
		c.Abort()
	}
	re.svc.Add(&data).Send(c)
}

func (re *user_ctrl) Update(c *gin.Context) {
	claim_user, exist := c.Get("email")
	if !exist {
		libs.New("claim user is not exist", 400, true)
		c.Abort()
	}

	var data models.User
	err := json.NewDecoder(c.Request.Body).Decode(&data)
	if err != nil {
		libs.New(err.Error(), 400, true)
		c.Abort()
	}

	re.svc.Update(&data, claim_user.(string)).Send(c)
}

func (re *user_ctrl) Delete(c *gin.Context) {
	val := c.Request.URL.Query().Get("email")
	re.svc.Delete(val).Send(c)
}

func (re *user_ctrl) Search(c *gin.Context) {
	val := c.Request.URL.Query().Get("email")
	v := strings.ToLower(val)
	re.svc.Search(v).Send(c)
}

func (re *user_ctrl) SearchName(c *gin.Context) {
	val := c.Request.URL.Query().Get("name")
	v := strings.ToLower(val)
	re.svc.SearchName(v).Send(c)
}
