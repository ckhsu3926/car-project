<template>
  <q-dialog v-model="IsVehicleBoxOpen">
    <VehicleBoxAdd v-if="VehicleBoxMode === 'add'" />
    <VehicleBoxEdit v-if="VehicleBoxMode === 'edit'" />
  </q-dialog>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import vehicleStore from 'stores/vehicle';
import VehicleBoxAdd from 'components/fixed/VehicleBoxAdd.vue';
import VehicleBoxEdit from 'components/fixed/VehicleBoxEdit.vue';

export default defineComponent({
  name: 'VehicleBox',

  components: { VehicleBoxAdd, VehicleBoxEdit },

  setup() {
    const { IsVehicleBoxOpen, VehicleBoxMode } = vehicleStore();

    return {
      IsVehicleBoxOpen,
      VehicleBoxMode,
    };
  },
});
</script>
