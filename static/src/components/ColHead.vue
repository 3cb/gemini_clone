<template>
    <div class="ch-wrapper">
        <span class="ch-product">{{ product }}</span>
        <span class="ch-price">{{ trades[0].price }}</span>
    </div>
</template>

<script>
export default {
    props: ['product', 'trades']
}
</script>

<style>
.ch-wrapper {
    display: inline-flex;
    flex-flow: row nowrap;
    justify-content: space-around;
    width: 100%;
}
.ch-product {
    justify-content: flex-end;
    flex: 0 1 0%;
}
.ch-price {
    justify-content: flex-end;
    flex: 0 1 0%;
}
</style>
