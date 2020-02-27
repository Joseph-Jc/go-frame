package gframe

import (
	"github.com/gin-gonic/gin"
	"log"
)

type GinRoute interface {
	Run(*gin.Engine)
}

type Engine struct {
	engine     *gin.Engine
	routeList []GinRoute
}

func NewGin() *Engine {
	return &Engine{
		engine:     gin.Default(),
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
