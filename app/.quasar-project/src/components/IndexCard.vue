<template>
  <q-card class="index-card q-ma-sm bg-secondary text-white cursor-pointer" v-ripple @click="$router.push(to)">
    <div class="row justify-center q-mt-md">
      <q-icon :name="icon" />
    </div>

    <q-card-section>
      <div class="text-h6">{{ title }}</div>
      <div class="text-subtitle2">{{ caption }}</div>
    </q-card-section>
  </q-card>
</template>

<style lang="sass" scoped>
.index-card
  width: 100%
  height: 160px
  max-width: 250px

  i
    font-size: 64px
</style>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'IndexCard',
  props: {
    title: {
      type: String,
      required: true,
    },

    caption: {
      type: String,
      default: '',
    },

    to: {
      type: String,
      default: '#',
    },

    icon: {
      type: String,
      default: '',
    },
  },
});
</script>
