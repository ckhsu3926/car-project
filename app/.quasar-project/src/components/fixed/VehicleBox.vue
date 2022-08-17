<template>
  <q-dialog v-model="IsVehicleBoxOpen">
    <VehicleBoxAdd v-if="VehicleBoxMode === 'add'" />
  </q-dialog>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import vehicleStore from 'stores/vehicle';
import VehicleBoxAdd from 'components/fixed/VehicleBoxAdd.vue';

export default defineComponent({
  name: 'VehicleBox',

  components: { VehicleBoxAdd },

  setup() {
    const { IsVehicleBoxOpen, VehicleBoxMode } = vehicleStore();

    return {
      IsVehicleBoxOpen,
      VehicleBoxMode,
    };
  },
});
</script>
