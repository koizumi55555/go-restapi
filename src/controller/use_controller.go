package controller

import (
	"context"
	"errors"
	"koizumi55555/go-restapi/src/entitiy"
	"koizumi55555/go-restapi/src/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	UserUsecase usecase.UserUsecase
}

type ServerController struct {
	ServerUsecase usecase.ServerUsecase
}

func NewUserController(pif usecase.PostgresIf) UserController {
	return UserController{
		UserUsecase: usecase.NewUserUsecase(pif),
	}
}

func NewServerController(clientID, clientSecret string) ServerController {
	return ServerController{
		ServerUsecase: usecase.NewServerUsecase(clientID, clientSecret),
	}
}

// ユーザ情報の取得
func (uc UserController) GetUser(c *gin.Context) {

	user, err := uc.UserUsecase.GetUser(c.Param("user_id"))
	if err == nil {
		c.JSON(http.StatusOK, user)
	} else {
		ErrorHandling(err, c)
	}
}

// ユーザの削除
func (uc UserController) DeleteUser(c *gin.Context) {

	err := uc.UserUsecase.DeleteUser(c.Param("user_id"))
	if err == nil {
		c.JSON(http.StatusNoContent, gin.H{})
	} else {
		ErrorHandling(err, c)
	}
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
	if err == nil {
		c.JSON(http.StatusOK, user)
	} else {
		ErrorHandling(err, c)
	}
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
	if err == nil {
		c.JSON(http.StatusOK, user)
	} else {
		ErrorHandling(err, c)
	}
}

// ユーザ情報の一覧取得
func (uc UserController) ListUsers(c *gin.Context) {

	user, err := uc.UserUsecase.ListUsers()
	if err == nil {
		c.JSON(http.StatusOK, user)
	} else {
		ErrorHandling(err, c)
	}
}

// Googleの認可ログイン画面にリダイレクト
func (s *ServerController) Authorize(c *gin.Context) {
	u, err := s.ServerUsecase.CreateAuthorizationRequestURL()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// 認可リクエストを送信
	c.Redirect(http.StatusMovedPermanently, u.String())

}

// アクセストークンを取得し、レスポンスとしてを返却
func (s *ServerController) Callback(c *gin.Context) {

	code := c.Query("code")
	if code == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	// contextをタイムアウト機能付きで生成
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	token, err := s.ServerUsecase.Exchange(ctx, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// レスポンスとしてを返却
	c.JSON(http.StatusOK, gin.H{"accessToken": token.AccessToken})
}

// エラーハンドリング
func ErrorHandling(err error, c *gin.Context) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrorMessage": err.Error(),
		})
	}
}
