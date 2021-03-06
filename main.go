package main

import (
	"fmt"
	"go-frame/conf"
	"go-frame/cron"
	"go-frame/gframe"
	"go-frame/model"
	"go-frame/routes"
	"go-frame/utils"
	"os"
)

func main() {
	//配置
	conf.Init()
	//日志
	file := utils.InitLog()
	defer func() {
		if err := file.Close(); err != nil {
			panic(fmt.Sprintf("Failed to close the log file:%s", err))
		}
	}()

	//数据库
	model.Init()
	defer func() {
		err := model.DB.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed to close the DB:%s", err))
		}
	}()

	//定时任务
	c := gframe.NewCron()
	c.AddRoute(&cron.Route{})
	c.Start()
	defer c.Stop()

	//gin
	engine := gframe.NewGin()
	engine.AddRoute(&routes.ApiRoute{}, &routes.WebRoute{})
	engine.Run(":" + os.Getenv("PORT"))

}
