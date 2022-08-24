<template>
  <q-page class="row content-center justify-center">
    <IndexCard v-for="(item, index) in linksList" :key="`indexCardItem${index}`" v-bind="item" />
  </q-page>
</template>

<script lang="ts">
import IndexCard from 'components/IndexCard.vue';
import { defineComponent } from 'vue';

const linksList = [
  {
    title: 'User',
    caption: 'User Info',
    icon: 'person',
    to: '/user',
  },
  {
    title: 'Gas Station',
    caption: 'Gas Station List in common use',
    icon: 'storefront',
    to: '/store',
  },
  {
    title: 'Vehicle',
    caption: 'Vehicle List',
    icon: 'directions_car',
    to: '/vehicle',
  },
];

export default defineComponent({
  name: 'IndexPage',
  components: { IndexCard },
  setup() {
    return { linksList };
  },
});
</script>
