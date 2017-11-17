<template>
    <div class="ob">
        <ul class="spacer has-text-weight-semibold ob-header">
            <li>
                <span>Price({{ product.slice(3, 6).toUpperCase() }})</span>
                <span class="is-pulled-right">Size({{ product.slice(0, 3).toUpperCase() }})</span>
            </li>
        </ul>
        <div class="lower-wrapper">
            <div class="ob-asks-wrapper">
                <ul class="ob-asks">
                    <li
                        is="book-row"
                        v-for="level in book.asks"
                        :level="level"
                        :color="'has-text-danger'"
                        :key="level[0]">
                    </li>
                </ul>
            </div>
            <ul class="spacer has-text-weight-semibold">
                <li>
                    <span>{{ spread | formatSpread }}</span>
                    <span class="is-pulled-right">SPREAD</span>
                </li>
            </ul>
            <ul class="ob-bids">
                <li
                    is="book-row"
                    v-for="level in book.bids"
                    :level="level"
                    :color="'has-text-success'"
                    :key="level[0]">
                </li>
            </ul>
        </div>
    </div>
</template>

<script>
import BookRow from './BookRow.vue'

export default {
    filters: {
        formatSpread(spread) {
            if (spread.toFixed(2) < .01) {
                return spread.toFixed(5)
            } else {
                return spread.toFixed(2)
            }
        }
    },
    props: ['book', 'product'],
    computed: {
        asks() {
            return this.book.asks
        },
        bids() {
            return this.book.bids
        },
        spread() {
            return parseFloat(this.book.asks[this.book.asks.length-1]) - parseFloat(this.book.bids[0])
        }
    },
    components: {
        BookRow
    }
}
</script>

<style>
.ob {
    position: relative;
    width: 100%;
    height: calc(100vh - 115px);
    overflow: hidden;
}
.spacer {
    position: relative;
    border-style: solid;
    border-width: 1px 0px 1px 0px;
    display: inline-flex;
    flex-flow: row nowrap;
    width: 100%;
    color: hsl(217, 71%, 53%);
}
.spacer li {
    width: 100%;
}

.lower-wrapper {
    position: relative;
    height: 100%;
    overflow: hidden;
}
.ob-asks-wrapper {
    position: relative;
    height: calc(50vh - 85px);
}
.ob-asks {
    position: absolute;
    bottom: 0;
    right: 0;
    overflow: hidden;
}
.ob-bids {
    position: relative;
    top: 2px;
}
</style>
