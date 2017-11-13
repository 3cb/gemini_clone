package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/3cb/ssc"

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

	// Start new websocket pool to connect to Gemini Websocket API
	config := ssc.PoolConfig{
		IsReadable: true,
		IsWritable: false,
		IsJSON:     true,
	}

	pp, err := ssc.NewSocketPool(sockets, config)
	if err != nil {
		log.Printf("Error starting new Socket Pool. Cannot start server.")
		return
	}

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

// JSONRead is a method in the JSONReaderWriter interface that reads from websocket and sends to pool via channel
func (m Message) JSONRead(s *ssc.Socket, toPoolJSON chan<- ssc.JSONReaderWriter, errorChan chan<- ssc.ErrorMsg) error {
	err := s.Connection.ReadJSON(&m)
	if err != nil {
		return err
	}
	toPoolJSON <- m
	return nil
}

// JSONWrite is a method in the JSONReaderWriter interface that takes values from the pool via channel and writes them to a websocket
func (m Message) JSONWrite(s *ssc.Socket, fromPoolJSON <-chan ssc.JSONReaderWriter, errorChan chan<- ssc.ErrorMsg) error {
	m, ok := (<-fromPoolJSON).(Message)
	if ok == false {
		return fmt.Errorf("wrong data type sent from Pool to websocket goroutine(%v)", s.URL)
	}
	err := s.Connection.WriteJSON(m)
	if err != nil {
		return err
	}
	return nil
}
