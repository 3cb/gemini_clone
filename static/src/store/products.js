function getProducts(array) {
    let products = []
    for(let i = 0; i < array.length; i++) {
        products.push({
            name: array[i],
            price: '',
            sequence: null,
            trades: [],
            book: {
                asks: [], // [ price, size ]
                bids: []
            }
        })
    }
    return products
}

export { getProducts }