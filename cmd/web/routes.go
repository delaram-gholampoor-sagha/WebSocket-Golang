package main

import (
	"net/http"

	"github.com/Delaram-Gholampoor-Sagha/WebSocket-Golang/internal/handlers"
	"github.com/bmizerany/pat"
)

// routes handles application routes
func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/send", http.HandlerFunc(handlers.SendAlertToConnectedUsers))

	mux.Get("/ws", http.HandlerFunc(handlers.WsEndPoint))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return mux

}
