package main

import (
	"gin/cron"
	"gin/logs"
	"gin/routers"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go cron.Setup()
	logs.SetupLogger()
	routers.RegisterHandler()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
}
