<template>
  <v-app dark>
    <v-bottom-navigation :value="value" color="primary">
      <v-btn @click="Change(1)">
        <span>Recents</span>
        <v-icon>mdi-history</v-icon>
      </v-btn>
      <v-btn @click="Change(2)">
        <span>Favorites</span>
        <v-icon>mdi-heart</v-icon>
      </v-btn>
      <v-btn @click="Change(3)">
        <span>Nearby</span>
        <v-icon>mdi-map-marker</v-icon>
      </v-btn>
    </v-bottom-navigation>
    <v-container v-if="menutoggle == 2">
      <Cards></Cards>
    </v-container>
  </v-app>
</template>

<script>
export default {
  data: () => ({ menutoggle: 2, value: 1 }),
  methods: {
    Change(num) {
      this.menutoggle = num
    },
  },
}
</script>
