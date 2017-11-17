<template>
    <div>
        <col-head :product="'BTC/USD'" :trades="trades"></col-head>
        <!-- <div class="header is-marginless">
            <div class="has-text-weight-semibold">BTC/USD</div>
            <div class="">{{ trades[0].price }}</div>
        </div> -->
        <div class="columns">
            <order-book class="column is-5" :book="book" :product="product"></order-book>
            <time-sales class="column is-7" :trades="trades" :product="product"></time-sales>
        </div>
    </div>
</template>

<script>
import ColHead from './ColHead.vue'
import OrderBook from "./OrderBook.vue";
import TimeSales from "./TimeSales.vue";

export default {
  computed: {
    product() {
      return this.$store.state.products[0].name;
    },
    book() {
      return this.$store.state.products[0].book;
    },
    trades() {
      return this.$store.state.products[0].trades;
    }
  },
  components: {
    ColHead,
    OrderBook,
    TimeSales
  }
};
</script>

<style>

</style>
