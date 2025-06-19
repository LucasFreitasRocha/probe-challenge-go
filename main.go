package main

import (
	"github.com/LucasFreitasRocha/probe-challenge-go/config"
	"github.com/LucasFreitasRocha/probe-challenge-go/config/logger"
)

func main() {
	logger.Info("About to start user application")
	config.InitApp()
}
