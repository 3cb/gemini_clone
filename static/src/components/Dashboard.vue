<template>
    <div>

    </div>
</template>

<script>
import xs from 'xstream'
import _ from 'lodash'

export default {
    data() {
        return {
            producer1: {
                start: (listener) => {
                    this.$store.commit('startWebsocket', 'ws1')
                    this.$store.state.ws1.onmessage = event => {
                        listener.next(event)
                    }
                },
                stop: () => {
                    this.$store.state.ws1.close()
                    this.$store.state.ws1.onclose = event => {
                        console.log(event)
                    }
                }
            },
            producer2: {
                start: (listener) => {
                    this.$store.commit('startWebsocket', 'ws2')
                    this.$store.state.ws2.onmessage = event => {
                        listener.next(event)
                    }
                },
                stop: () => {
                    this.$store.state.ws2.close()
                    this.$store.state.ws2.onclose = event => {
                        console.log(event)
                    }
                }
            },
            producer3: {
                start: (listener) => {
                    this.$store.commit('startWebsocket', 'ws3')
                    this.$store.state.ws3.onmessage = event => {
                        listener.next(event)
                    }
                },
                stop: () => {
                    this.$store.state.ws3.close()
                    this.$store.state.ws3.onclose = event => {
                        console.log(event)
                    }
                }
            },
            mainListener: {
                next: (value) => {
                    // console.log(value)
                },
                error: (err) => {
                    console.error(err)
                },
                complete: () => {
                    console.log("Main$ stream complete.")
                }
            },
            initBookListener: {
                next: (value) => {
                    console.log(value)
                    this.$store.commit('initBook', {
                        product: value.product,
                        events: value.events,
                        sequence: value.socket_sequence
                    })
                },
                error: (err) => {
                    console.error(err)
                },
                complete: () => {
                    console.log("Book initialization stream complete.")
                }
            },
            updateBookListener: {
                next: (value) => {

                },
                error: (err) => {
                    console.error(err)
                },
                complete: () => {
                    console.log("Book update stream complete.")
                }
            }
        }
    },
    computed: {
        ws1$() {
            return xs.createWithMemory(this.producer1)
        },
        ws2$() {
            return xs.createWithMemory(this.producer2)
        },
        ws3$() {
            return xs.createWithMemory(this.producer3)
        },
        main$() {
            return xs.merge(this.ws1$, this.ws2$, this.ws3$)
                        .map(v => {
                            return {
                                ...JSON.parse(v.data),
                                product: _.takeRight(v.target.url.split('/'), 1).join('')
                            }
                        })
        },
        initBook$() {
            return xs.from(this.main$).filter(v => v.events[0].reason === 'initial')
        },
        updateBook$(){
            return xs.from(this.main$)
        }
    },
    mounted() {
        this.main$.addListener(this.mainListener)
        this.initBook$.addListener(this.initBookListener)
    }
}
</script>

<style>

</style>