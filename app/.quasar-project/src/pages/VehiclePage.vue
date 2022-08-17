<template>
  <q-page padding>
    <VehicleListComponent />
    <!-- Refueling Record -->
    <!-- Maintenance Record -->
  </q-page>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import VehicleListComponent from 'components/VehicleList.vue';

export default defineComponent({
  name: 'VehiclePage',
  components: { VehicleListComponent },
});
</script>
