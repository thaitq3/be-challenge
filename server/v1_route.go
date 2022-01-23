package server

import (
	"crud-challenge/handler"
	"crud-challenge/handler/wager"
	"github.com/gorilla/mux"
)

func routing(router *mux.Router) {
	// POST wager
	router.
		HandleFunc("/wagers",handler.GetHandler(&wager.CreateWagerHandler{})).
		Methods("POST")

	// POST buy/:wager_id
	router.
		HandleFunc("/buy/{wager_id}", handler.GetHandler(&wager.PurchaseWagerHandler{})).
		Methods("POST")

	// GET wager
	router.
		HandleFunc("/wagers", handler.GetHandler(&wager.ListWagerHandler{})).
		Methods("GET")
}