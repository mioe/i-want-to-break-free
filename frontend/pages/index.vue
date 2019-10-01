<template lang="pug">
.view-main
  h1 frontend-nuxt-js
  ul
    li(v-for="l in list" :key="l.id")
      h2 {{ l.title }}
      p {{ l.author }}

</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      list: null
    }
  },

  async mounted() {
    await axios.get('/api/books').then(response => {
      this.list = response.data
    })
  }
}
</script>

<style scoped lang="sass">
@import '@/assets/sass/_variables'

.container
  background: $red
</style>
