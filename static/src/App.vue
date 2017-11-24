<template>
  <div id="app" class="is-size-6">
    <dashboard></dashboard>
  </div>
</template>

<script>
import Dashboard from './components/Dashboard.vue'
import xs from 'xstream'

export default {
  name: 'app',
  data() {
    return {
      producer: {
        start: (listener) => {
          window.onresize = event => {
            listener.next(event)
          }
        },
        stop: () => {
          console.log("No longer listening to window resize.")
        }
      },
      winListener: {
        next: (value) => {
          this.$store.commit('setWin', value)
        },
        error: (err) => {
          console.error(err)
        },
        complete: () => {
          console.log("Window resize stream complete.")
        }
      }
    }
  },
  computed: {
    window$() {
      return xs.createWithMemory(this.producer).map(v => {
        return {
          width: v.target.innerWidth,
          height: v.target.innerHeight
        }
      })
    }
  },
  mounted() {
    this.$store.commit('setWin', { width: window.innerWidth, height: window.innerHeight })
    this.window$.addListener(this.winListener)
  },
  components: {
    Dashboard
  }
}
</script>

<style>
html {
  background-color: hsl(0, 0%, 21%);
}
body {
  height: 100%;
  color: whitesmoke;
}

a {
  color: whitesmoke;
}

dashboard {
  height: 100%;
}
</style>
