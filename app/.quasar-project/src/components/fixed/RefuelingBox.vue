<template>
  <q-dialog v-model="isRefuelingDialogOpen">
    <q-card class="q-ma-xs">
      <q-form @submit="onSubmit">
        <q-card-section class="bg-light-blue text-white">
          <div class="text-h6">{{ refuelingDialogMode === 'add' ? 'Add' : 'Edit' }} New Refueling Record</div>
        </q-card-section>

        <RefuelingBoxInput />

        <q-card-actions align="evenly">
          <q-btn
            flat
            color="primary"
            :label="refuelingDialogMode === 'add' ? 'Add' : 'Edit'"
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
import refuelingStore from 'stores/refueling';
import RefuelingBoxInput from 'components/fixed/RefuelingBoxInput.vue';

export default defineComponent({
  name: 'RefuelingAddBox',
  components: { RefuelingBoxInput },

  setup() {
    const { isRefuelingDialogOpen, refuelingDialogMode, OnAddSubmit, isDialogFormSubmitting, OnEditSubmit } =
      refuelingStore();

    const onSubmit = () => {
      if (refuelingDialogMode.value === 'add') {
        OnAddSubmit();
      } else {
        OnEditSubmit();
      }
    };

    return {
      isRefuelingDialogOpen,
      refuelingDialogMode,
      onSubmit,
      isDialogFormSubmitting,
    };
  },
});
</script>
