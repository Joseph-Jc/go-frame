package gframe

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

type GinRoute interface {
	Run(*gin.Engine)
}

type Engine struct {
	engine    *gin.Engine
	routeList []GinRoute
}

func NewGin() *Engine {
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "release" {
		logPath := os.Getenv("LOG_PATH")
		errorWriterFile, err := os.OpenFile(
			logPath+"/errors.log",
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			0666)
		if err != nil {
			panic(fmt.Sprintf("Failed to open the log file:%s", err))
		}
		gin.DefaultErrorWriter = io.MultiWriter(errorWriterFile)
		writerFile, err := os.OpenFile(
			logPath+"/gin.log",
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			0666)
		if err != nil {
			panic(fmt.Sprintf("Failed to open the log file:%s", err))
		}
		gin.DefaultWriter = io.MultiWriter(writerFile)
	}
	gin.SetMode(ginMode)

	return &Engine{
		engine:    gin.Default(),
		routeList: make([]GinRoute, 0),
	}
}

func (e *Engine) AddRoute(route ...GinRoute) {
	e.routeList = append(e.routeList, route...)
}

func (e *Engine) Run(port string) {
	for _, route := range e.routeList {
		route.Run(e.engine)
	}
	if err := e.engine.Run(port); err != nil {
		log.Fatal(err)
	}
}
