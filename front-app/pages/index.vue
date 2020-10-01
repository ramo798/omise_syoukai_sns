<template>
  <div>
    <Cards v-bind:posts="posts"></Cards>
  </div>
</template>
<script>
export default {
  async asyncData({ $axios }) {
    const response = await $axios.$get('http://localhost:8080/postlist/all')
    // console.log(response)
    return {
      posts: response,
    }
  },
}
</script>
