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
			"ErrorMessage": err.Error(),
		})
	}
	c.JSON(http.StatusOK, user)
}

// ユーザの削除
func (uc UserController) DeleteUser(c *gin.Context) {

	err := uc.UserUsecase.DeleteUser(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrorMessage": err.Error(),
		})
	}
	c.JSON(http.StatusNoContent, gin.H{})
}

// ユーザ情報の更新
func (uc UserController) UpdateUser(c *gin.Context) {

	var json entitiy.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ErrorMessage": err.Error()})
		return
	}

	updateUser := entitiy.User{
		ID:    c.Param("user_id"),
		Name:  json.Name,
		Email: json.Email,
	}

	user, err := uc.UserUsecase.UpdateUser(updateUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrorMessage": err.Error(),
		})
	}
	c.JSON(http.StatusOK, user)
}

// ユーザの登録
func (uc UserController) CreateUser(c *gin.Context) {

	var json entitiy.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ErrorMessage": err.Error()})
		return
	}

	createUser := entitiy.User{
		Name:  json.Name,
		Email: json.Email,
	}

	user, err := uc.UserUsecase.CreateUser(createUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrorMessage": err.Error(),
		})
	}
	c.JSON(http.StatusOK, user)
}

// ユーザ情報の一覧取得
func (uc UserController) ListUsers(c *gin.Context) {

	user, err := uc.UserUsecase.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrorMessage": err.Error(),
		})
	}
	c.JSON(http.StatusOK, user)

}
