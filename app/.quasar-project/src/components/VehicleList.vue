<template>
  <div class="row q-gutter-md justify-center">
    <q-card v-for="i in 4" :key="`vehicle-card${i}`" class="vehicle-card">
      <q-card-section class="bg-primary text-white">
        <div class="text-h6">Big Wife</div>
        <div class="text-subtitle2">Honda 2020 CBR650R</div>
      </q-card-section>

      <q-list>
        <q-item clickable>
          <q-item-section avatar>
            <q-icon color="green" name="local_gas_station" />
          </q-item-section>

          <q-item-section>
            <q-item-label>Refueling Record</q-item-label>
            <q-item-label caption>Fill your gas tank.</q-item-label>
          </q-item-section>
        </q-item>

        <q-item clickable>
          <q-item-section avatar>
            <q-icon color="red" name="build" />
          </q-item-section>

          <q-item-section>
            <q-item-label>Maintenance Record</q-item-label>
          </q-item-section>
        </q-item>
      </q-list>
    </q-card>

    <q-btn class="vehicle-card bg-light-blue text-white text-h4" @click="OnAddOpen">ADD</q-btn>
  </div>
</template>

<style lang="sass" scoped>
.vehicle-card
  width: 100%
  height: 190px
  max-width: 250px
</style>

<script lang="ts">
import { defineComponent } from 'vue';
import vehicleStore from 'stores/vehicle';

export default defineComponent({
  name: 'VehicleList',

  setup() {
    const { OnAddOpen } = vehicleStore();

    return {
      OnAddOpen,
    };
  },
});
</script>
