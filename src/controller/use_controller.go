package controller

import (
	"koizumi55555/go-restapi/src/entitiy"
	"koizumi55555/go-restapi/src/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase usecase.UserUsecase
}

func NewUserController(pif usecase.PostgresIf) UserController {
	return UserController{
		UserUsecase: usecase.NewUserUsecase(pif),
	}
}

// ユーザ情報の取得
func (uc UserController) GetUser(c *gin.Context) {

	user, err := uc.UserUsecase.GetUser(c.Param("user_id"))
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrorMessaga": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"ID":    user.ID,
		"Name":  user.Name,
		"Email": user.Email,
	})
}

// ユーザの削除
func (uc UserController) DeleteUser(c *gin.Context) {

	err := uc.UserUsecase.DeleteUser(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrorMessaga": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"delete": "succese",
	})
}

// ユーザ情報の更新
func (uc UserController) UpdateUser(c *gin.Context) {

	var json entitiy.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ErrorMessaga": err.Error()})
		return
	}

	updateUser := entitiy.User{}
	updateUser.ID = c.Param("user_id")
	updateUser.Name = json.Name
	updateUser.Email = json.Email
	user, err := uc.UserUsecase.UpdateUser(updateUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrorMessaga": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"ID":    user.ID,
		"Name":  user.Name,
		"Email": user.Email,
	})
}

// ユーザの登録
func (uc UserController) CreateUser(c *gin.Context) {

	var json entitiy.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ErrorMessaga": err.Error()})
		return
	}

	createUser := entitiy.User{}
	createUser.Name = json.Name
	createUser.Email = json.Email

	user, err := uc.UserUsecase.CreateUser(createUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrorMessaga": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"ID":    user.ID,
		"Name":  user.Name,
		"Email": user.Email,
	})
}

// ユーザ情報の一覧取得
func (uc UserController) GetListUsers(c *gin.Context) {

	user, err := uc.UserUsecase.GetListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrorMessaga": err.Error(),
		})
	}

	for i := 0; i < len(user); i++ {
		c.JSON(http.StatusOK, gin.H{
			"ID":    user[i].ID,
			"Name":  user[i].Name,
			"Email": user[i].Email,
		})
	}

}
