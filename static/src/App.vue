<template>
  <div id="app">
    <dashboard></dashboard>
  </div>
</template>

<script>
import Dashboard from './components/Dashboard.vue'

export default {
  name: 'app',
  components: {
    Dashboard
  }
}
</script>

<style>
body {
  height: 100%;
}

a {
  color: hsl(0, 0%, 21%);
}

dashboard {
  height: 100%;
}
</style>
