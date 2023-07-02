package src

import (
	"koizumi55555/go-restapi/src/controller"

	"github.com/gin-gonic/gin"
)

func Server(uc controller.UserController, s controller.ServerController) *gin.Engine {

	engine := gin.Default()
	engine.GET("/users/:user_id", uc.GetUser)
	engine.DELETE("/users/:user_id", uc.DeleteUser)
	engine.PUT("/users/:user_id", uc.UpdateUser)
	engine.POST("/users", uc.CreateUser)
	engine.GET("/users", uc.ListUsers)
	engine.GET("/oauth", s.Authorize)
	engine.GET("/auth_redirect", s.Callback)

	return engine
}
