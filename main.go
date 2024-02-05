package main

import (
	"context"
	"gin/cron"
	"gin/logs"
	"gin/routers"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// 开启定时任务
	go cron.Setup()

	// 设置日志
	logs.SetupLogger()

	// 注册路由
	srv := routers.RegisterHandler()

	// 服务
	go srv.ListenAndServe()

	// 监听外部信号退出
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutdown Server ...")
	srv.Shutdown(context.Background())

}
