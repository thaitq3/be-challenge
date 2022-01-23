package server

import (
	"crud-challenge/config"
	"time"
	"crud-challenge/utils"
)

func Serve() {
	config.LoadConfig()
	InitDependencies()

	Start()

	utils.WaitShutdownSignal()
	// actions on shutdown
	utils.WaitOrTimeout(time.Minute* 3, Stop())
}
