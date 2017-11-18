<template>
    <div>
        <col-head :product="'ETH/BTC'" :price="price" :icon="'fa fa-btc'"></col-head>
        <div class="columns">
            <order-book class="column is-5" :book="book" :product="product"></order-book>
            <time-sales class="column is-7" :trades="trades" :product="product"></time-sales>
        </div>
    </div>
</template>

<script>
import ColHead from './ColHead.vue'
import OrderBook from './OrderBook.vue'
import TimeSales from './TimeSales.vue'

export default {
    computed: {
        product() {
            return this.$store.state.products[2].name
        },
        book() {
            return this.$store.state.products[2].book
        },
        trades() {
            return this.$store.state.products[2].trades
        },
        price() {
            return this.$store.state.products[2].price
        }
    },
    components: {
        ColHead,
        OrderBook,
        TimeSales
    }
}
</script>

<style>

</style>
