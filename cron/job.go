package cron

import "fmt"

func Test1() {
	fmt.Println("Run by func Test1!!!!!")
}

func Test2() {
	fmt.Println("Run by func Test2!!!!!")
}

type Job struct {
}

func (j Job) Run() {
	fmt.Println("这是一个定时任务，每隔一分钟执行------->Run by job!!!!!")
}
