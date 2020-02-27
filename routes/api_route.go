package routes

import (
	"github.com/gin-gonic/gin"
	"stock/controller"
)

type ApiRoute struct {
}

var pingController controller.PingController
var studentController controller.StudentController

func (r *ApiRoute) Run(engine *gin.Engine) {
	engine.GET("/ping", pingController.Ping)

	engine.GET("/create-student", studentController.CreateStudent)
}
