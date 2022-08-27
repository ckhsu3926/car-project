<template>
  <q-page padding>
    <MaintenanceRecordTable :vehicleID="vehicleID" :list="maintenanceList" />
  </q-page>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import maintenanceStore from 'stores/maintenance';
import MaintenanceRecordTable from 'components/MaintenanceRecordTable.vue';

export default defineComponent({
  name: 'UserPage',
  components: { MaintenanceRecordTable },
  setup() {
    const route = useRoute();
    const router = useRouter();
    const vehicleID = Number(route.params.vehicleID);

    if (!vehicleID) {
      router.push('/vehicle');
    }

    const { maintenanceList, getMaintenanceList } = maintenanceStore();
    getMaintenanceList(vehicleID);

    return { vehicleID, maintenanceList };
  },
});
</script>
