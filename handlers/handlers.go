package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/3cb/gemini_clone/types"
	"github.com/3cb/ssc"
	"github.com/gorilla/mux"
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
			ServerURLs: sockets,
			IsReadable: true,
			IsWritable: false,
			IsJSON:     true,
			DataJSON:   types.Message{},
		}

		pp, err := ssc.NewSocketPool(config)
		if err != nil {
			log.Printf("Error starting new Socket Pool. Cannot start server.")
			return
		}
		go pp.Control()

		go func() {
			for {
				select {
				default:
					v := <-pp.Pipes.OutboundJSON
					// log.Printf("Message from Pool Controller:\n%v\n", v)
					conn.WriteJSON(v)
				}
			}
		}()
	})
}

// Trades sends request to Gemini API to get Trade history for given product
func Trades() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		product := vars["product"]
		api := fmt.Sprintf("https://api.gemini.com/v1/trades/%v?limit_trades=50", product)
		resp, err := http.Get(api)
		if err != nil {
			log.Printf("Error during http request to Gemini API: %v", err)
		}
		defer resp.Body.Close()
		data, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			log.Printf("Error reading response body for %v: %v", product, err2)
		}
		w.Write(data)
	})
}
