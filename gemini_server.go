package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// serve static files
	r.Handle("/", http.FileServer(http.Dir("./static")))
	r.PathPrefix("/dist/").Handler(http.FileServer(http.Dir("./static")))

	sockets := []string{
		"wss://api.gemini.com/v1/marketdata/btcusd",
		"wss://api.gemini.com/v1/marketdata/ethusd",
		"wss://api.gemini.com/v1/marketdata/ethbtc",
	}

	logWS := log.New(os.Stdout, "", 0)
	reconnect := make(chan string)
	disconnect := make(chan string)
	for _, v := range sockets {
		go connnectWS(v, reconnect, disconnect, logWS)
	}

	go func() {
		for {
			ws := <-reconnect
			go connnectWS(ws, reconnect, disconnect, logWS)
		}
	}()

	// time.Sleep(10 * time.Second)
	// for _, x := range sockets {
	// 	disconnect <- x
	// }

	// start server
	log.Fatal(http.ListenAndServe(":4000", r))
}

// Message defines the structure of messages received from Gemini websocket
type Message struct {
	Type      string  `json:"type"`
	EventID   int     `json:"eventId"`
	Sequence  int     `json:"socket_sequence"`
	Events    []Event `json:"events"`
	Timestamp int     `json:"timestampms"`
}

// Event defines the structure of Events field contained in Message struct
type Event struct {
	Type      string `json:"type"`
	Reason    string `json:"reason"`
	Price     string `json:"price"`
	Delta     string `json:"delta"`
	Remaining string `json:"remaining"`
	Side      string `json:"side"`
}
