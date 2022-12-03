package main

import (
	"daily/task"
	"github.com/robfig/cron"
)

func main() {
	//创建一个cron实例
	cron2 := cron.New()
	task.Job(cron2)

	//启动/关闭
	cron2.Start()
	defer cron2.Stop()

	//查询语句，保持程序运行，在这里等同于for{}
	select {}
}
