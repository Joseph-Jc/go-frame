package cron

import "go-frame/gframe"

type Route struct {
}

func (r *Route) Run(cron *gframe.Cron) {
	////每3秒执行1次
	//cron.Schedule("*/3 * * * * *").RunFunc(Test1) //每隔3秒运行
	////不设置Schedule会复用上一条的Schedule设置
	//cron.RunFunc(Test2)
	//使用job方式运行，只需实现Run方法
	cron.Schedule("0 */1 * * * *").RunJob(Job{}) //每隔一分钟运行
}
