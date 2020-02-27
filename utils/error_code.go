package utils

const (
	CodeSuccess = 0 //成功
	CodeFail    = 1 //失败
)

var MessageMap = map[int]string{
	0: "请求成功",
	1: "请求失败",
}
