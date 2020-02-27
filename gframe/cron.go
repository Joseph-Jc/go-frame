package gframe

import (
	"github.com/robfig/cron"
)

type CronRoute interface {
	Run(*Cron)
}

type Cron struct {
	cron      *cron.Cron
	spec      string
	routeList []CronRoute
}

func NewCron() *Cron {
	return &Cron{
		cron:      cron.New(),
		spec:      "",
		routeList: make([]CronRoute, 0),
	}
}

func (c *Cron) AddRoute(route ...CronRoute) {
	c.routeList = append(c.routeList, route...)
}

func (c *Cron) Start() {
	for _, route := range c.routeList {
		route.Run(c)
	}
	c.cron.Start()
}

func (c *Cron) Stop() {
	c.cron.Stop()
}

func (c *Cron) Schedule(spec string) *Cron {
	c.spec = spec
	return c
}

func (c *Cron) RunFunc(f func()) {
	if err := c.cron.AddFunc(c.spec, f); err != nil {
		panic(err)
	}
}

func (c *Cron) RunJob(j cron.Job) {
	if err := c.cron.AddJob(c.spec, j); err != nil {
		panic(err)
	}
}
