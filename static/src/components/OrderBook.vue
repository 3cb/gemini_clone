<template>
    <div class="ob spin-parent">
        <ul class="spacer has-text-weight-semibold ob-header">
            <li>
                <span>Price({{ product.slice(3, 6).toUpperCase() }})</span>
                <span class="is-pulled-right">Size({{ product.slice(0, 3).toUpperCase() }})</span>
            </li>
        </ul>
        <div v-if="book.length > 0" class="lower-wrapper">
            <div class="ob-asks-wrapper">
                <ul class="ob-asks">
                    <transition-group
                        :css="false"
                        @before-enter="beforeEnter"
                        @enter="enter"
                        @leave="leave"
                    >
                        <li
                            is="book-row"
                            v-for="level in asks"
                            :level="level"
                            :color="'has-text-danger'"
                            :key="level.price">
                        </li>
                    </transition-group>
                </ul>
            </div>
            <ul class="spacer has-text-weight-semibold">
                <li>
                    <span>{{ spread | formatSpread }}</span>
                    <span class="is-pulled-right">SPREAD</span>
                </li>
            </ul>
            <ul class="ob-bids">
                <transition-group
                    :css="false"
                    @before-enter="beforeEnter"
                    @enter="enter"
                    @leave="leave"
                >
                    <li
                        is="book-row"
                        v-for="level in bids"
                        :level="level"
                        :color="'has-text-success'"
                        :key="level.price">
                    </li>
                </transition-group>
            </ul>
        </div>
        <spinner
            v-else
            size="large"
            line-fg-color="hsl(217, 71%, 53%)"
            class="spinner"
        ></spinner>
    </div>
</template>

<script>
import BookRow from './BookRow.vue'
import Spinner from 'vue-simple-spinner'

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
            return _.chain(this.book)
                        .takeWhile(o => o.side === 'ask')
                        .takeRight(this.$store.state.bookDepth)
                        .value()
        },
        bids() {
            return _.chain(this.book)
                        .takeRightWhile(o => o.side === 'bid')
                        .take(this.$store.state.bookDepth)
                        .value()
        },
        spread() {
            if (this.asks && this.bids) {
                return parseFloat(this.asks[this.asks.length-1].price) - parseFloat(this.bids[0].price)
            } else {
                return 0.00
            }
        }
    },
    methods: {
        beforeEnter: function(el) {
            el.style.opacity = 0
            el.style.backgroundColor = '#7a7a7a'
            },
        enter: function(el, done) {
            Velocity(el,
            {
                opacity: 1,
                backgroundColor: '#363636'
            },
            {
                duration: 700,
                complete: function() {
                done()
                }
            })
        },
        leave: function(el, done) {
            Velocity(el,
            {
                color: '#000000'
            },
            {
                duration: 500,
                complete: function() {
                    done()
                }
            })
        }
    },
    components: {
        BookRow,
        Spinner
    }
}
</script>

<style>
.ob {
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
    width: 100%;
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
