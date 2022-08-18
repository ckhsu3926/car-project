<template>
  <div class="row q-gutter-md justify-center">
    <VehicleListCard v-for="vehicle in vehicleList" :key="`vehicleListCard${vehicle.id}`" :Vehicle="vehicle" />

    <q-btn class="add-card bg-light-blue text-white text-h4" @click="OnAddOpen">ADD</q-btn>
  </div>
</template>

<style lang="sass" scoped>
.add-card
  width: 100%
  height: 190px
  max-width: 250px
</style>

<script lang="ts">
import { defineComponent } from 'vue';
import vehicleStore from 'stores/vehicle';
import VehicleListCard from 'components/VehicleListCard.vue';

export default defineComponent({
  name: 'VehicleList',
  components: { VehicleListCard },

  setup() {
    const { OnAddOpen, vehicleList, getVehicliList } = vehicleStore();

    getVehicliList();

    return {
      OnAddOpen,
      vehicleList,
    };
  },
});
</script>
