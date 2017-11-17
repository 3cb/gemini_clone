<template>
    <div>
        <div class="level header is-marginless">
            <div class="level-item">
                <strong>ETH/BTC</strong>
            </div>
        </div>
        <div class="columns">
            <order-book class="column is-5" :book="book" :product="product"></order-book>
            <time-sales class="column is-7" :trades="trades" :product="product"></time-sales>
        </div>
    </div>
</template>

<script>
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
        }
    },
    components: {
        OrderBook,
        TimeSales
    }
}
</script>

<style>

</style>
