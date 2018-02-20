package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

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

		pp := ssc.NewPool(sockets, time.Second*30)
		err = pp.Start()
		if err != nil {
			log.Printf("Error starting new Socket Pool. Cannot start server.")
			return
		}

		go func() {
			m := types.Message{}
			for {
				v := <-pp.Outbound
				if v.Type == 1 {
					err := json.Unmarshal(v.Payload, &m)
					if err != nil {
						log.Printf("error receiving message: %v", err)
					}
					slice := strings.Split(v.ID, "/")
					m.Product = slice[len(slice)-1]

					conn.WriteJSON(m)
				} else {
					log.Printf("wrong message type from websocket.")
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
