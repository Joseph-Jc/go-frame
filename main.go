package main

import (
	"os"
	"stock/conf"
	"stock/cron"
	"stock/gframe"
	"stock/model"
	"stock/routes"
)

func main() {
	//配置
	conf.Init()

	//数据库
	model.Init()
	defer func() {
		err := model.DB.Close()
		if err != nil {
			panic(err)
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
