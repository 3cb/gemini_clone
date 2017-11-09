package ssc

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

// ProducerPool is a collection of websocket connections combined with 3 channels used to send and received message to and from the goroutines that control them
type ProducerPool struct {
	Producers      []*Producer
	FailureChan    chan FailureMSG
	DisconnectChan chan *Producer
	DataChan       chan ProducerData
}

// Producer type defines a websocket connection
type Producer struct {
	URL         string
	Connection  *websocket.Conn
	isConnected bool
	OpenedAt    time.Time
	ClosedAt    time.Time
	ReconTry    int // Number of reconnects attempted
}

// ProducerData wraps []byte and Producer instance together so receiver can identify the source
type ProducerData struct {
	Producer *Producer
	Payload  []byte
}

// FailureMSG wraps an error message with Producer instance so receiver can try reconnect and/or log error
type FailureMSG struct {
	Producer *Producer
	Error    error
}

// NewProducerPool creates a new instance of ProducerPool and returns a pointer to it
func NewProducerPool(urls []string) *ProducerPool {
	failure := make(chan FailureMSG)
	disconnect := make(chan *Producer)
	data := make(chan ProducerData)

	producers := []*Producer{}
	for _, v := range urls {
		p := &Producer{URL: v}
		go Connect(p, failure, disconnect, data)
		producers = append(producers, p)
	}

	pool := &ProducerPool{
		producers,
		failure,
		disconnect,
		data,
	}
	return pool
}

// Connect connects to websocket given a url string and three channels from ProducerPool type.
// Creates a goroutine to receive and send data as well as to listen for failures and calls to disconnect
func Connect(p *Producer, failure chan<- FailureMSG, disconnect <-chan *Producer, data chan<- ProducerData) {
	c, _, err := websocket.DefaultDialer.Dial(p.URL, nil)
	if err != nil {
		log.Printf("Error connecting to websocket: \n%v\n%v", p.URL, err)
		failure <- FailureMSG{p, err}
		return
	}
	p.Connection = c
	p.isConnected = true
	p.OpenedAt = time.Now()

	// Start goroutine to listen to websocket.
	// Closes connection and sends failure message on error.
	// If disconnect message is received websocket failure message is sent with nil error value and socket is Closed.
	go func() {
		defer p.Connection.Close()
		for {
			select {
			case v := <-disconnect:
				if v == p {
					log.Printf("Close Message Received from Controller. Closing websocket at: %v", p.URL)
					failure <- FailureMSG{p, nil}
					return
				}
			default:
				_, msg, err := p.Connection.ReadMessage()
				if err != nil {
					log.Printf("Error reading from websocket(%v): ", err)
					failure <- FailureMSG{p, err}
					return
				}
				data <- ProducerData{p, msg}
			}
		}
	}()
}
