<template>
    <div class="ch-wrapper has-text-weight-semibold is-size-5">
        <span class="ch-product">{{ product }}</span>
        <span class="ch-price">
            <span class="icon">
                <i :class="icon" aria-hidden="true"></i>
            </span>
            <span>
                {{ price }}
            </span>
        </span>
    </div>
</template>

<script>
export default {
    props: ['product', 'price', 'icon']
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
    display: inline-flex;
    flex-flow: row nowrap;
    justify-content: flex-end;
    flex: 0 1 0%;
}
</style>
