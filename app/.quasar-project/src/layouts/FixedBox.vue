<template>
  <div>
    <LoginBox />
    <VehicleBox />
    <RefuelingBox />
    <MaintenanceBox />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import LoginBox from 'components/fixed/LoginBox.vue';
import VehicleBox from 'components/fixed/VehicleBox.vue';
import RefuelingBox from 'components/fixed/RefuelingBox.vue';
import MaintenanceBox from 'components/fixed/MaintenanceBox.vue';

export default defineComponent({
  name: 'FixedBox',

  components: {
    LoginBox,
    VehicleBox,
    RefuelingBox,
    MaintenanceBox,
  },
});
</script>
