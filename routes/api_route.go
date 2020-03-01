package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go-frame/controller"
	"go-frame/utils"
)

type ApiRoute struct {
}

var pingController controller.PingController
var studentController controller.StudentController
var testController controller.TestController

func (r *ApiRoute) Run(engine *gin.Engine) {
	engine.GET("/ping", pingController.Ping)

	engine.GET("/create-student", studentController.CreateStudent)

	middleware := engine.Group("/api").Use(func() gin.HandlerFunc {
		return func(context *gin.Context) {
			utils.Log.Info("中间件处理")
			context.Next()
		}
	}())
	{
		middleware.GET("/get-student", func(context *gin.Context) {
			context.JSON(http.StatusOK, utils.Response.Return(utils.CodeSuccess, "", "这是student"))
		})
		middleware.GET("/say-hello", func(context *gin.Context) {
			context.JSON(http.StatusOK, utils.Response.Return(utils.CodeSuccess, "", "这是say-hello"))
		})
	}

	engine.GET("/say-hello1", func(context *gin.Context) {
		context.JSON(http.StatusOK, utils.Response.Return(utils.CodeSuccess, "", "这是无中间件的say-hello"))
	})
	middleware.GET("/test", testController.Test)
}
