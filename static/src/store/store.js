import Vue from 'vue'
import Vuex from 'vuex'
import { getProducts } from './products.js'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        ws1: null,
        ws1Connected: false,
        ws2: null,
        ws2Connected: false,
        ws3: null,
        ws3Connected: false,

        products: getProducts(['btcusd', 'ethusd', 'ethbtc']),
        selected_product: ''
    },
    mutations: {
        startWebsocket(state, socket) {
            if (socket === 'ws1') {
                state.ws1 = new WebSocket("wss://api.gemini.com/v1/marketdata/btcusd")
                state.ws1.onopen = event => {
                    console.log(event)
                    state.ws1Connected = true
                }
            } else if (socket === 'ws2') {
                state.ws2 = new WebSocket("wss://api.gemini.com/v1/marketdata/ethusd")
                state.ws2.onopen = event => {
                    console.log(event)
                    state.ws2Connected = true
                }
            } else if (socket === 'ws3') {
                state.ws3 = new WebSocket("wss://api.gemini.com/v1/marketdata/ethbtc")
                state.ws3.onopen = event => {
                    console.log(event)
                    state.ws3Connected = true
                }
            }
        },
        initBook(state, { product, events, sequence }) {
            let i = _.findIndex(state.products, o => o.name === product)
            
        }
    }
})
