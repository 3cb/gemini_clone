<template>
<div>
    <div class="level logo is-size-5">
        <div class="level-item">
            <span class="icon">
                <a
                    href="https://github.com/3cb/gemini_clone"
                    target="_blank">
                    <i
                        class="fa fa-github"
                        aria-hidden="true"
                    ></i>
                </a>
            </span>        
            <a href="https://gemini.com/" target="_blank">
                <strong>GEMINI_clone</strong>
            </a>
        </div>
    </div>

    <div class="columns">
        <div class="column is-4">
            <column-one id="col-one" class="col-cont"></column-one>
        </div>
        <div class="column is-4">
            <column-two id="col-two" class="col-cont"></column-two>
        </div>
        <div class="column is-4">
            <column-three id="col-three" class="col-cont"></column-three>
        </div>
    </div>
</div>
</template>

<script>
import ColumnOne from './ColumnOne.vue'
import ColumnTwo from './ColumnTwo.vue'
import ColumnThree from './ColumnThree.vue'
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
    },
    components: {
        ColumnOne,
        ColumnTwo,
        ColumnThree
    }
}
</script>

<style>
.col-cont {
    width: 50%;
}
.col-cont div {
    width: 100%;
}

.logo {
    height: 30px;
    margin-top: 12px;
}
.logo a:hover {
    color: hsl(171, 100%, 41%);
}
.header {
    height: 35px;
}
</style>