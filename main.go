package main

import (
	"github.com/LucasFreitasRocha/probe-challenge-go/src/main/config"
	"github.com/LucasFreitasRocha/probe-challenge-go/src/main/config/logger"
)

func main() {
	logger.Info("About to start user application")
	config.InitApp()
}
