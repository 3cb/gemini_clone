package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func connnectWS(wsURL string, reconnect chan<- string, disconnect <-chan string, l *log.Logger) {
	c, resp, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		l.Printf("%v -- Error connecting to websocket: %v", time.Now(), err)
		reconnect <- wsURL
	}
	l.Printf("%v -- Request to %v: %v", time.Now(), wsURL, resp.Status)

	// start goroutine to listen to websocket. Closes connection and sends reconnect message to main function on error
	go func() {
		defer c.Close()
		for {
			select {
			case v := <-disconnect:
				if v == wsURL {
					l.Printf("%v -- Closing websocket: %v", time.Now(), wsURL)
					return
				}
			default:
				var message Message

				err2 := c.ReadJSON(&message)
				if err2 != nil {
					l.Printf("%v -- Error from websocket(%v): ", time.Now(), err2)
					reconnect <- wsURL
					return
				}
				// l.Printf("%v -- Message from %v:\n%+v", time.Now(), wsURL, message)
			}
		}
	}()

	return
}
