package controller

import (
	"github.com/gin-gonic/gin"
	"go-frame/model"
	"go-frame/utils"
)

type StudentController struct {
}

func (s StudentController) CreateStudent(context *gin.Context) {
	student := model.CreateStudent()
	context.JSON(200, utils.Response.Success(student))
}
