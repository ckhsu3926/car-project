<template>
  <q-dialog v-model="isMaintenanceDialogOpen">
    <q-card class="q-ma-xs">
      <q-form @submit="onSubmit">
        <q-card-section class="bg-light-blue text-white">
          <div class="text-h6">{{ maintenanceDialogMode === 'add' ? 'Add' : 'Edit' }} New Maintenance Record</div>
        </q-card-section>

        <MaintenanceBoxInput />

        <q-card-actions align="evenly">
          <q-btn
            flat
            color="primary"
            :label="maintenanceDialogMode === 'add' ? 'Add' : 'Edit'"
            type="submit"
            :disable="isDialogFormSubmitting"
          />
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import maintenanceStore from 'stores/maintenance';
import MaintenanceBoxInput from 'components/fixed/MaintenanceBoxInput.vue';

export default defineComponent({
  name: 'MaintenanceBox',
  components: { MaintenanceBoxInput },

  setup() {
    const { isMaintenanceDialogOpen, maintenanceDialogMode, OnAddSubmit, isDialogFormSubmitting, OnEditSubmit } =
      maintenanceStore();

    const onSubmit = () => {
      if (maintenanceDialogMode.value === 'add') {
        OnAddSubmit();
      } else {
        OnEditSubmit();
      }
    };

    return {
      isMaintenanceDialogOpen,
      maintenanceDialogMode,
      onSubmit,
      isDialogFormSubmitting,
    };
  },
});
</script>
