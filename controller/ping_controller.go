package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stock/utils"
)

type PingController struct {
}

func (t *PingController) Ping(context *gin.Context) {
	context.JSON(http.StatusOK, utils.Response.Return(0, "pong!!!!"))
}
