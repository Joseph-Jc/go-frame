package utils

import "github.com/gin-gonic/gin"

var Response JsonResult

type JsonResult struct {
}

//返回json
//参数1：code，参数2：data，参数3：message
func (j *JsonResult) Return(params ...interface{}) interface{} {
	var code int = CodeSuccess
	var data interface{}
	var message interface{} = ""
	switch len(params) {
	case 1:
		code = params[0].(int)
	case 2:
		code = params[0].(int)
		data = params[1]
	case 3:
		code = params[0].(int)
		data = params[1]
		message = params[2]
	}
	if _, ok := MessageMap[code]; ok && message == "" {
		message = MessageMap[code]
	}
	return gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	}
}

//返回成功json
//参数1：data，参数2：code，参数3：message
func (j *JsonResult) Success(params ...interface{}) interface{} {
	var data interface{}
	var code int = CodeSuccess
	var message interface{} = ""
	switch len(params) {
	case 1:
		data = params[0]
	case 2:
		data = params[0]
		code = params[1].(int)
	case 3:
		data = params[0]
		code = params[1].(int)
		message = params[2]
	}
	if _, ok := MessageMap[code]; ok && message == "" {
		message = MessageMap[code]
	}
	return gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	}
}

//返回失败json
//参数1：code，参数2：message，参数3：data
func (j *JsonResult) Fail(params ...interface{}) interface{} {
	var data interface{}
	var code int = CodeFail
	var message interface{} = ""
	switch len(params) {
	case 1:
		code = params[0].(int)
	case 2:
		code = params[0].(int)
		message = params[1]
	case 3:
		code = params[0].(int)
		message = params[1]
		data = params[2]
	}
	if _, ok := MessageMap[code]; ok && message == "" {
		message = MessageMap[code]
	}
	return gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	}
}
