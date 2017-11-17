<template>
    <div>
        <col-head :product="'ETH/USD'" :trades="trades"></col-head>
        <!-- <nav class="level header is-marginless">
            <div class="level-item">ETH/USD</div>
            <div class="level-item">{{ trades[0].price }}</div>
        </nav> -->
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
            return this.$store.state.products[1].name
        },
        book() {
            return this.$store.state.products[1].book
        },
        trades() {
            return this.$store.state.products[1].trades
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
