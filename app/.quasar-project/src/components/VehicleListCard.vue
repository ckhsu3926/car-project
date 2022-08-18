<template>
  <q-card class="vehicle-card">
    <q-card-section class="title bg-primary text-white cursor-pointer" v-ripple @click="onVehicleClick(Vehicle.id)">
      <div class="text-h6 text-capitalize">
        <span class="text-uppercase">{{ Vehicle.license }}</span> {{ Vehicle.name }}
      </div>
      <div class="text-subtitle2 text-uppercase">{{ Vehicle.company }} {{ Vehicle.model }}</div>
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
</template>

<style lang="sass" scoped>
.vehicle-card
  width: 100%
  height: 190px
  max-width: 250px
  .title
    min-height: 86px
</style>

<script lang="ts">
import { defineComponent, PropType } from 'vue';
import vehicleStore, { vehicle } from 'stores/vehicle';

export default defineComponent({
  name: 'VehicleList',

  props: {
    Vehicle: {
      type: Object as PropType<vehicle>,
      required: true,
    },
  },

  setup() {
    const { onVehicleClick } = vehicleStore();

    return {
      onVehicleClick,
    };
  },
});
</script>
