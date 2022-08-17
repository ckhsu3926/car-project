<template>
  <div>
    <LoginBox />
    <VehicleBox />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import LoginBox from 'components/fixed/LoginBox.vue';
import VehicleBox from 'components/fixed/VehicleBox.vue';

export default defineComponent({
  name: 'FixedBox',

  components: {
    LoginBox,
    VehicleBox,
  },
});
</script>
