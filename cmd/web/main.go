package main

import (
	"log"
	"net/http"

	"github.com/Delaram-Gholampoor-Sagha/WebSocket-Golang/internal/handlers"
)

// port is the port to listen on
const port = ":8080"

// main is the main function
func main() {
	// get our routes
	mux := routes()

	// start websocket functionality
	log.Println("Starting websocket functionality...")
	go handlers.ListenToWsChannel()

	// start the web server
	log.Println("Starting application on port...", port)
	_ = http.ListenAndServe(port, mux)
}
