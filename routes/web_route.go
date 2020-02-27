package routes

import "github.com/gin-gonic/gin"

type WebRoute struct {
}

func (w *WebRoute) Run(engine *gin.Engine) {
	engine.LoadHTMLGlob("view/*")
	engine.GET("/hello", func(context *gin.Context) {
		context.HTML(200, "index.html", gin.H{
			"title": "say hello!",
		})
	})
}
