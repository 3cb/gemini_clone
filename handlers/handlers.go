package handlers

import (
	"log"
	"net/http"

	"github.com/3cb/gemini_clone/types"
	"github.com/3cb/ssc"
	"github.com/gorilla/websocket"
)

// WebsocketRequest upgrades incoming http requests to a websocket connection and starts a websocket pool connected to 3 Gemini Exchange websockets
func WebsocketRequest(upgrader *websocket.Upgrader) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Unable to upgrade to websocket connection: %v\n", err)
			return
		}

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
			DataJSON:   types.Message{},
		}

		pp, err := ssc.NewSocketPool(sockets, config)
		if err != nil {
			log.Printf("Error starting new Socket Pool. Cannot start server.")
			return
		}
		go pp.ControlJSON()

		go func() {
			for {
				select {
				default:
					v := <-pp.Pipes.FromPoolJSON
					conn.WriteJSON(v)
					// msg := <-pp.Pipes.FromPoolJSON
					// log.Printf("Message from Pool Controller:\n%v\n", msg)
				}
			}
		}()
	})
}
