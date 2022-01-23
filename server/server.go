package server

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"gopkg.in/tylerb/graceful.v1"

	log "github.com/sirupsen/logrus"

	"crud-challenge/utils"
)

// Start starts all servers.
func Start() {

	go startInternalServer()
}

// Stop stops all servers.
func Stop() <-chan struct{} {
	return utils.WaitChannels(
		stopInternalServer(),
	)
}

// startInternalServer to serve internal api calls.
func startInternalServer() {
	defer doAPIPanicRecovery("internal-server")

	serveAddr := fmt.Sprintf("0.0.0.0:%s", "8080")
	log.Info("Starting internal server on: ", serveAddr)

	// Use graceful to do proper clean up before server exit
	srv := &graceful.Server{
		Server: &http.Server{
			Addr:           serveAddr,
			Handler:        InternalRouter(),
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
		Timeout: time.Second * 10,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Error(err)
		panic("startInternalServer failed err: " + err.Error())
	}

	log.Info("internal server exits")

	close(internalShutdown)
}

func stopInternalServer() <-chan struct{} {
	return internalShutdown
}

func doAPIPanicRecovery(serviceTag string) {
	if r := recover(); r != nil {
		logMessage := fmt.Sprintf("%s%s service got exception, failed with error %s %s", "alert", serviceTag, r, string(debug.Stack()))
		log.Error(logMessage)
	}
}
