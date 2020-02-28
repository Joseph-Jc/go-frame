package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stock/controller"
	"stock/utils"
)

type ApiRoute struct {
}

var pingController controller.PingController
var studentController controller.StudentController

func (r *ApiRoute) Run(engine *gin.Engine) {
	engine.GET("/ping", pingController.Ping)

	engine.GET("/create-student", studentController.CreateStudent)

	middleware := engine.Group("/api").Use(func() gin.HandlerFunc {
		return func(context *gin.Context) {
			context.JSON(http.StatusOK, utils.Response.Return(utils.CodeSuccess, map[int]interface{}{
				1: "hello",
				2: "123",
			}))
			utils.Log.Trace(map[int]interface{}{
				1: "hello",
				2: "123",
			})
			context.Abort()
		}
	}())

	middleware.GET("/get-student", func(context *gin.Context) {
		context.JSON(200, utils.Response.Return(0, "", "这是student"))
	})
	middleware.GET("/say-hello", func(context *gin.Context) {
		context.JSON(200, utils.Response.Return(0, "", "这是say-hello"))
	})
	engine.GET("/say-hello1", func(context *gin.Context) {
		context.JSON(200, utils.Response.Return(0, "", "这是无中间件的say-hello"))
	})
}
