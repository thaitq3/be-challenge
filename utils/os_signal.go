package utils

import (
	"os"
	"os/signal"
	"syscall"
)

// WaitShutdownSignal waits for shutdown signal, and if so, then waits for all channels finish or timeout.
func WaitShutdownSignal() {
	waitShutdownSignal(nil)
}

func waitShutdownSignal(inter chan os.Signal) {
	interrupt := inter
	if interrupt == nil {
		interrupt = make(chan os.Signal, 1)
	}

	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Wait for syscall.SIGINT or syscall.SIGTERM
	<-interrupt
}