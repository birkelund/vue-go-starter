<template>
  <div class="ping">
    <button v-on:click="ping">ping server</button>
    <p>Server has ponged <strong>{{ pongs }}</strong> times.</p>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data: function() {
    return {
      pongs: 0,
    }
  },

  methods: {
    ping() {
      axios.get('/api/ping').then(() => {
        this.pongs++
      })
    }
  }
}
</script>
