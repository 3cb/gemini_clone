<template>
  <div id="app" class="is-size-6">
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
