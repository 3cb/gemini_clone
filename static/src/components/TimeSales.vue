<template>
    <div class="ts-h spin-parent">
        <ul class="spacer has-text-weight-semibold">
            <li class="ts-wrapper">
                <span class="ts-size">Size</span>
                <span class="ts-price">Price</span>
                <span class="ts-time">Time</span>
            </li>
        </ul>
        <ul v-if="trades.length > 0">
            <transition-group
                :css="false"
                @before-enter="beforeEnter"
                @enter="enter"
            >
                <li
                    v-for="(trade, index) in trades"
                    :key="trade.tid"
                    class="ts-wrapper"
                    :class="trade.class"
                >
                    <span class="ts-size">{{ parseFloat(trade.size).toFixed(8) }}</span>
                    <span class="ts-price">{{ trade.price | formatPrice }}</span>
                    <span class="ts-time">{{ trade.time | formatTime }}</span>
                </li>
            </transition-group>
        </ul>
        <spinner
            v-else
            size="large"
            line-fg-color="hsl(217, 71%, 53%)"
            class="spinner"
        ></spinner>
    </div>
</template>

<script>
import Spinner from 'vue-simple-spinner'
import moment from 'moment'

export default {
    props: ['trades'],
    filters: {
        formatTime(time) {
            let t = moment(time).format().split('T')[1].slice(0, -6)
            return t
        },
        formatPrice(price) {
            if (parseFloat(price) < 1) {
                return parseFloat(price).toFixed(5)
            } else {
                return parseFloat(price).toFixed(2)
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
        }
    },
    components: {
        Spinner
    }
}
</script>

<style>
.ts-h {
    height: calc(100vh - 115px);
    overflow: hidden;
}
.ts-wrapper {
    display: inline-flex;
    flex-flow: row nowrap;
    justify-content: space-between;
    width: 100%
}
.ts-size {
    display: inline-flex;
    justify-content: flex-end;
    flex: 0 1 30%;
}
.ts-price {
    display: inline-flex;
    justify-content: flex-end;
    flex: 1 1 25%;
}
.ts-time {
    display: inline-flex;
    justify-content: flex-end;
    flex: 0 1 35%;
}
</style>
