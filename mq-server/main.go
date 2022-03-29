package main

import (
	"mq-server/conf"
	"mq-server/services"
)

func main() {
	conf.Init()

	forever := make(chan bool)
	services.CreateTask()
	<-forever // 阻塞主线程
}
