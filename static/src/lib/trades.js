import axios from 'axios'

function getBTCUSD() {
    return axios.get('/api/trades/btcusd')
}

function getETHUSD() {
    return axios.get('/api/trades/ethusd')
}

function getETHBTC() {
    return axios.get('/api/trades/ethbtc')
}

export { getBTCUSD, getETHUSD, getETHBTC }