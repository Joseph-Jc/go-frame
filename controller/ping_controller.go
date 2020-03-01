package controller

import (
	"github.com/gin-gonic/gin"
	"go-frame/utils"
	"net/http"
)

type PingController struct {
}

func (t PingController) Ping(context *gin.Context) {
	context.JSON(http.StatusOK, utils.Response.Return(0, "pong!!!!"))
}
