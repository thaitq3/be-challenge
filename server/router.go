package server

import (
	"crud-challenge/middleware"
	"github.com/gorilla/mux"
)

var (
	internalShutdown = make(chan struct{})
)

// InternalRouter defines the router for internal endpoints.
func InternalRouter() *mux.Router {
	// accept both /path/ and /path
	r := mux.NewRouter().StrictSlash(true)

	// middleware
	r.Use(middleware.BasicHeader, middleware.RecoverWrap)

	routing(r)

	return r
}
