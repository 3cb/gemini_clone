import Vue from 'vue'
import Vuex from 'vuex'
import { getProducts } from './products.js'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        win: {
            width: null,
            height: null
        },
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
        setWin(state, win) {
            state.win = win
        },
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
            events = _.orderBy(events, o => parseFloat(o.price), ['asc'])
            for (let j = 0; j < events.length; j++) {
                state.products[i].book.unshift({
                    price: events[j].price,
                    size: events[j].remaining,
                    side: events[j].side
                })
            }
        },
        updateBook(state, { product, price, remaining, side, sequence }) {
            let i = _.findIndex(state.products, o => o.name === product)
            let j = _.findIndex(state.products[i].book, o => parseFloat(o.price).toFixed(8) === parseFloat(price).toFixed(8))
            if (j < 0) {
                state.products[i].book = _.concat(state.products[i].book, [{
                                                price: price,
                                                size: remaining,
                                                side: side
                                            }])
                state.products[i].book = _.orderBy(state.products[i].book, o => parseFloat(o.price), ['desc'])
            } else {
                if (remaining === '0') {
                    _.pullAt(state.products[i].book, [j])
                } else {
                    state.products[i].book[j].size = remaining
                    state.products[i].book[j].side = side
                }
            }
        },
        initTrades(state, { product, data }) {
            let i = _.findIndex(state.products, o => o.name === product)
            for (let j = 0; j < data.length; j++) {
                state.products[i].trades.push({
                    tid: data[j].tid,
                    price: data[j].price,
                    size: data[j].amount,
                    time: data[j].timestampms,
                    side: data[j].type,
                    class: data[j].type === 'sell' ? 'has-text-danger' : 'has-text-success'
                })
            }
            state.products[i].price = state.products[i].trades[0].price
        },
        addTrade(state, { product, tid, price, size, time, side }) {
            let i = _.findIndex(state.products, o => o.name === product)
            state.products[i].price = price
            state.products[i].trades.unshift({
                tid,
                price,
                size,
                time,
                side,
                class: side === 'bid' ? 'has-text-danger' : 'has-text-success'
            })
        }
    }
})