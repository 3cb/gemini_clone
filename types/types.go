package types

import (
	"fmt"
	"strings"

	"github.com/3cb/ssc"
)

// Message defines the structure of messages received from Gemini websocket
type Message struct {
	Type      string  `json:"type"`
	Product   string  `json:"product"`
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

	TID       int64  `json:"tid"`
	Amount    string `json:"amount"`
	MakerSide string `json:"makerSide"`
}

// JSONRead is a method in the JSONReaderWriter interface that reads from websocket and sends to pool via channel
func (m Message) JSONRead(s *ssc.Socket, Socket2PoolJSON chan<- ssc.JSONReaderWriter) error {
	err := s.Connection.ReadJSON(&m)
	if err != nil {
		return err
	}
	slice := strings.Split(s.URL, "/")
	m.Product = slice[len(slice)-1]
	Socket2PoolJSON <- m
	return nil
}

// JSONWrite is a method in the JSONReaderWriter interface that takes values from the pool via channel and writes them to a websocket
func (m Message) JSONWrite(s *ssc.Socket, Pool2SocketJSON <-chan ssc.JSONReaderWriter) error {
	m, ok := (<-Pool2SocketJSON).(Message)
	if ok == false {
		return fmt.Errorf("wrong data type sent from Pool to websocket goroutine(%v)", s.URL)
	}
	err := s.Connection.WriteJSON(m)
	if err != nil {
		return err
	}
	return nil
}
