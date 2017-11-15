package main

import (
	"log"
	"net/http"

	"github.com/3cb/gemini_clone/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func main() {
	r := mux.NewRouter()

	// serve static files
	r.Handle("/", http.FileServer(http.Dir("./static")))
	r.PathPrefix("/dist/").Handler(http.FileServer(http.Dir("./static")))

	// websocket request handler
	var upgrader = &websocket.Upgrader{}
	r.Handle("/ws", handlers.WebsocketRequest(upgrader))

	// start server
	log.Fatal(http.ListenAndServe(":4000", r))
}
