package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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
		log.Printf("%v\n", product)
		if err2 != nil {
			log.Printf("Error reading response body for %v: %v", product, err2)
		}
		w.Write(data)
	})
}
