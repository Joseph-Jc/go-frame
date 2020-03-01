package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-frame/service"
	"go-frame/utils"
	"strconv"
)

var testService service.TestService

type TestController struct {
}

func (t TestController) Test(context *gin.Context) {
	n, _ := strconv.Atoi(context.DefaultQuery("n", "0"))
	total := testService.GetTotal(n)
	utils.Log.Info(fmt.Sprintf("从0加到%d的结果为：%d", n, total))
	context.String(200, strconv.Itoa(total))
}
