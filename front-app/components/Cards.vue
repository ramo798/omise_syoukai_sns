<template>
  <div>
    <v-card
      v-for="item in posts"
      :key="item.post_id"
      class="mx-auto my-12"
      max-width="480"
    >
      <v-list-item class="grow">
        <v-list-item-avatar color="grey darken-3">
          <v-img
            class="elevation-6"
            alt=""
            src="https://avataaars.io/?avatarStyle=Transparent&topType=ShortHairShortCurly&accessoriesType=Prescription02&hairColor=Black&facialHairType=Blank&clotheType=Hoodie&clotheColor=White&eyeType=Default&eyebrowType=DefaultNatural&mouthType=Default&skinColor=Light"
          ></v-img>
        </v-list-item-avatar>

        <v-list-item-content>
          <v-list-item-title>{{ item.users_id }}</v-list-item-title>
        </v-list-item-content>
      </v-list-item>
      <v-progress-linear
        color="teal"
        buffer-value="0"
        v-model="item.assessment"
        stream
      ></v-progress-linear>

      <v-card-title>レストランタイトル</v-card-title>

      <v-card-text>
        <div>{{ item.comment }}</div>
      </v-card-text>

      <v-card-text>
        <div>価格目安：{{ item.price }}円</div>
      </v-card-text>

      <v-divider class="mx-4"></v-divider>
      <v-img
        height="150"
        src="https://cdn.vuetifyjs.com/images/cards/cooking.png"
      ></v-img>
    </v-card>
  </div>
</template>
<script>
export default {
  props: {
    posts: {
      type: Array,
      default: () => {},
    },
  },
}
</script>
