<template>
  <q-page padding>
    <RefuelingRecordTable :vehicleID="vehicleID" :list="refuelingList" />
  </q-page>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import refuelingStore from 'stores/refueling';
import gasStationStore from 'stores/gasStation';
import RefuelingRecordTable from 'components/RefuelingRecordTable.vue';

export default defineComponent({
  name: 'UserPage',
  components: { RefuelingRecordTable },
  setup() {
    const route = useRoute();
    const router = useRouter();
    const vehicleID = Number(route.params.vehicleID);

    if (!vehicleID) {
      router.push('/vehicle');
    }

    const { refuelingList, getRefuelingList } = refuelingStore();
    getRefuelingList(vehicleID);

    const { getList } = gasStationStore();
    getList();

    return { vehicleID, refuelingList };
  },
});
</script>
