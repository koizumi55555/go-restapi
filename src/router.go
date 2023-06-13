package src

import (
	"koizumi55555/go-restapi/src/controller"

	"github.com/gin-gonic/gin"
)

func Server(uc controller.UserController) *gin.Engine {

	engine := gin.Default()
	engine.GET("/users/:user_id", uc.GetUser)
	engine.DELETE("/users/:user_id", uc.DeleteUser)
	engine.PUT("/users/:user_id", uc.UpdateUser)
	engine.POST("/users", uc.CreateUser)
	engine.GET("/users", uc.GetListUsers)

	return engine
}
