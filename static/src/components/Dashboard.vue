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

        <div v-if="win.width >= 1450" class="columns">
            <div class="column is-4 col-cont">
                <column-one id="col-one"></column-one>
            </div>
            <div class="column is-4 col-cont">
                <column-two id="col-two"></column-two>
            </div>
            <div class="column is-4 col-cont">
                <column-three id="col-three"></column-three>
            </div>
        </div>

        <div v-else>
          <div class="columns is-centered">
              <div class="column is-6">
                  <column-one></column-one>
              </div>
          </div>
          <div class="columns is-centered">
              <div class="column is-6">
                  <column-two></column-two>
              </div>
          </div>
          <div class="columns is-centered">
              <div class="column is-6">
                  <column-three></column-three>
              </div>
          </div>
        </div>
    </div>
</template>

<script>
import ColumnOne from "./ColumnOne.vue";
import ColumnTwo from "./ColumnTwo.vue";
import ColumnThree from "./ColumnThree.vue";
import xs from "xstream";
import _ from "lodash";
import axios from "axios";
import { getBTCUSD, getETHUSD, getETHBTC } from "../lib/trades.js";

export default {
  data() {
    return {
      producer1: {
        start: listener => {
          this.$store.commit("startWebsocket", "ws1");
          this.$store.state.ws1.onmessage = event => {
            listener.next(event);
          };
        },
        stop: () => {
          this.$store.state.ws1.close();
          this.$store.state.ws1.onclose = event => {
            console.log(event);
          };
        }
      },
      producer2: {
        start: listener => {
          this.$store.commit("startWebsocket", "ws2");
          this.$store.state.ws2.onmessage = event => {
            listener.next(event);
          };
        },
        stop: () => {
          this.$store.state.ws2.close();
          this.$store.state.ws2.onclose = event => {
            console.log(event);
          };
        }
      },
      producer3: {
        start: listener => {
          this.$store.commit("startWebsocket", "ws3");
          this.$store.state.ws3.onmessage = event => {
            listener.next(event);
          };
        },
        stop: () => {
          this.$store.state.ws3.close();
          this.$store.state.ws3.onclose = event => {
            console.log(event);
          };
        }
      },
      // =======  For Debugging -- Remove  ===================================
      mainListener: {
        next: value => {
          // console.log(value)
        },
        error: err => {
          console.error(err);
        },
        complete: () => {
          console.log("Main$ stream complete.");
        }
      },
      // =====================================================================
      initBookListener: {
        next: value => {
          this.$store.commit("initBook", {
            product: value.product,
            events: value.events,
            sequence: value.socket_sequence
          });
        },
        error: err => {
          console.error(err)
        },
        complete: () => {
          console.log("Book initialization stream complete.")
        }
      },
      updateBookListener: {
        next: value => {
          this.$store.commit("updateBook", {
            product: value.product,
            price: value.events[0].price,
            remaining: value.events[0].remaining,
            side: value.events[0].side
          });
        },
        error: err => {
          console.error(err);
        },
        complete: () => {
          console.log("Book update stream complete.")
        }
      },
      tradeListener: {
        next: value => {
          this.$store.commit('addTrade', {
            product: value.product,
            tid: value.events[0].tid,
            price: value.events[0].price,
            size: value.events[0].amount,
            time: value.timestampms,
            side: value.events[0].makerSide
          })
        },
        error: err => {
          console.error(err)
        },
        complete: () => {
          console.log("Trade stream complete.")
        }
      }
    };
  },
  computed: {
    win() {
      return this.$store.state.win
    },
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
      return xs.merge(this.ws1$, this.ws2$, this.ws3$).map(v => {
        return {
          ...JSON.parse(v.data),
          product: _.takeRight(v.target.url.split("/"), 1).join("")
        }
      })
    },
    initBook$() {
      return xs.from(this.main$).filter(v => v.events[0].reason === "initial")
    },
    updateBook$() {
      return xs.from(this.main$)
              .filter(v => v.events.length <= 2)
              .map(v => {
                if (v.events.length === 2) {
                  let x = _.cloneDeep(v)
                  _.reverse(x.events)
                  return x
                }
                return v
              })
    },
    trade$() {
      return xs.from(this.main$)
              .filter(v => v.events.length === 2)
    }
  },
  mounted() {
    this.getTrades()

    this.main$.addListener(this.mainListener)
    this.initBook$.addListener(this.initBookListener)
    this.updateBook$.addListener(this.updateBookListener)
    this.trade$.addListener(this.tradeListener)
  },
  methods: {
    getTrades() {
      axios
        .all([getBTCUSD(), getETHUSD(), getETHBTC()])
        .then(
          axios.spread((btcusd, ethusd, ethbtc) => {
            this.$store.commit("initTrades", {
              data: btcusd.data,
              product: "btcusd"
            });
            this.$store.commit("initTrades", {
              data: ethusd.data,
              product: "ethusd"
            });
            this.$store.commit("initTrades", {
              data: ethbtc.data,
              product: "ethbtc"
            });
          })
        )
        .catch(err => console.error(err));
    }
  },
  components: {
    ColumnOne,
    ColumnTwo,
    ColumnThree
  }
};
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
  color: hsl(217, 71%, 53%);
}
.header {
  height: 35px;
}

order-book {
  width: 50%;
}
time-sales {
  width: 50%;
}

.spin-parent {
  position: relative;
}
.spinner {
    position: absolute;
    top: 50%;
}
</style>