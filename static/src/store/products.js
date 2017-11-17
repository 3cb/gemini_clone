function getProducts(array) {
    let products = []
    for (let i = 0; i < array.length; i++) {
        products.push({
            name: array[i],
            price: '',
            sequence: null,
            trades: [],
            book: []
        })
    }
    return products
}

export { getProducts }