<template>
    <div>
        <div class="level header is-marginless">
            <div class="level-item">
                <strong>ETH/USD</strong>
            </div>
        </div>
        <order-book :book="book"></order-book>
    </div>
</template>

<script>
import OrderBook from './OrderBook.vue'

export default {
    computed: {
        book() {
            return this.$store.state.products[1].book
        }
    },
    components: {
        OrderBook
    }
}
</script>

<style>

</style>
