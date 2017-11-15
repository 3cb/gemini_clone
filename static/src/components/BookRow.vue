<template>
    <li class="br-row">
        <!-- <span class="br-svg"></span> -->
        <span class="br-size">{{ parseFloat(level[1]).toFixed(8) }}</span>
        <span class="br-price" :class="color">{{ level[0] | formatPrice }}</span>
    </li>
</template>

<script>
export default {
    filters: {
        formatPrice(price) {
            if (parseFloat(price) < 1) {
                return parseFloat(price).toFixed(5)
            } else {
                return parseFloat(price).toFixed(2)
            }
        }
    },
    props: ['level', 'color']
}
</script>

<style>
.br-row {
    display: inline-flex;
    flex-flow: row nowrap;
    justify-content: space-between;
    width: 100%;
    margin-top: -8px;
    /* margin-bottom: -8px; */
}
/* .br-svg {
    display: inline-flex;
    flex-flow: row nowrap;
    justify-content: flex-start;
    flex: 0 1 10%;
} */
.br-size {
    display: inline-flex;
    flex-flow: row nowrap;
    justify-content: flex-end;
    flex: 1 1 45%;
}
.br-price {
    display: inline-flex;
    flex-flow: row nowrap;
    justify-content: flex-end;
    flex: 0 2 40%;
}
</style>