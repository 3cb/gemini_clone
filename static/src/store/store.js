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

        products: getProducts(['btcusd', 'ethusd', 'ethbtc']), // []
        selected_product: '',

        bookDepth: 30
    },
    mutations: {
        startWebsocket(state, socket) {
            if (socket === 'ws1') {
                state.ws1 = new WebSocket("wss://api.gemini.com/v1/marketdata/btcusd")
                state.ws1.onopen = event => {
                    state.ws1Connected = true
                }
            } else if (socket === 'ws2') {
                state.ws2 = new WebSocket("wss://api.gemini.com/v1/marketdata/ethusd")
                state.ws2.onopen = event => {
                    state.ws2Connected = true
                }
            } else if (socket === 'ws3') {
                state.ws3 = new WebSocket("wss://api.gemini.com/v1/marketdata/ethbtc")
                state.ws3.onopen = event => {
                    state.ws3Connected = true
                }
            }
        },
        initBook(state, { product, events, sequence }) {
            let i = _.findIndex(state.products, o => o.name === product)

            state.products[i].sequence = sequence
            let array = _.orderBy(events, o => parseFloat(o.price), ['desc'])
            for (let j = 0; j < array.length; j++) {
                if (array[j].side === 'ask') {
                    state.products[i].book.asks.push([ array[j].price, array[j].remaining ])
                } else if (array[j].side === 'bid') {
                    state.products[i].book.bids.push([ array[j].price, array[j].remaining ])
                }
            }
            state.products[i].book.asks = _.takeRight(state.products[i].book.asks, state.bookDepth * 10)
            state.products[i].book.bids = _.take(state.products[i].book.bids, state.bookDepth * 10)
        },
        updateBook(state, { product, price, remaining, side, sequence }) {
            let i = _.findIndex(state.products, o => o.name === product)
            // if (sequence < state.products[i].sequence) {
            //     return
            // }
            // state.products[i].sequence = sequence
            
            let j = _.findIndex(state.products[i].book[side], a => parseFloat(a[0]) === parseFloat(price))
            if (j === -1) {
                state.products[i].book[side] = _.concat(state.products[i].book[side], [[ price, remaining ]])
                state.products[i].book[side] = _.orderBy(state.products[i].book[side], [a => parseFloat(a[0])], ['desc'])
            } else {
                if (remaining != 0) {
                   state.products[i].book[side][j][1] = remaining
                } else {
                    _.pullAt(state.products[i].book[side], [j])
                }
            }

            // trim length of order book
            if (side === 'asks') {
                _.takeRight(state.products[i].book[side], state.bookDepth * 10)
            } else if (side === 'bids') {
                _.take(state.products[i].book[side], state.bookDepth * 10)
            }
        },
        initTrades(state, { product, data }) {
            console.log(data)
            let i = _.findIndex(state.products, o => o.name === product)
            if (state.products[i].trades.length === 0) {
                for (let j = 0; j < data.length; j++) {
                    state.products[i].trades.push({
                        tid: data[j].tid,
                        price: data[j].price,
                        size: data[j].amount,
                        time: data[j].timestampms,
                        side: data[j].type
                    })
                }
            }
        },
        addTrade(state, { product, tid, price, size, time, side }) {
            let i = _.findIndex(state.products, o => o.name === product)
            state.products[i].trades.unshift({
                tid,
                price,
                size,
                time,
                side
            })
        }
    }
})