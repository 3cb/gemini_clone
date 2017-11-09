package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/3cb/gemini_clone/ssc"
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

	pp := ssc.NewProducerPool(sockets)

	go func() {
		for {
			var message Message

			v := <-pp.DataChan
			err := json.Unmarshal(v.Payload, &message)
			if err != nil {
				log.Printf("Error: %v", err)
			}
			// log.Printf("DATA: %+v", message)
		}
	}()

	time.Sleep(10 * time.Second)
	for _, x := range pp.Producers {
		pp.DisconnectChan <- x
	}

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
