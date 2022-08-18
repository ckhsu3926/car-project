<template>
  <q-card class="q-ma-xs">
    <q-form @submit="OnAddSubmit">
      <q-card-section class="bg-light-blue text-white">
        <div class="text-h6">Add New Vehicle</div>
      </q-card-section>

      <q-card-section>
        <q-input
          v-model="AddForm.name"
          label="Name"
          :rules="[(val) => (val && val.length > 0) || 'please enter vehicle name']"
        />
        <q-input
          v-model="AddForm.license"
          label="License"
          :rules="[(val) => (val && val.length > 0) || 'please enter vehicle license']"
        />
        <q-input v-model="AddForm.company" label="Company" />
        <q-input v-model="AddForm.model" label="Model" />
      </q-card-section>

      <q-card-actions align="evenly">
        <q-btn flat color="primary" label="Add" type="submit" :disable="IsDialogFormSubmitting" />
      </q-card-actions>
    </q-form>
  </q-card>
</template>

<style lang="sass" scoped>
.q-card
    min-width: 360px
</style>

<script lang="ts">
import { defineComponent } from 'vue';
import vehicleStore from 'stores/vehicle';

export default defineComponent({
  name: 'LoginBoxAdd',

  setup() {
    const { AddForm, OnAddSubmit, IsDialogFormSubmitting } = vehicleStore();

    return {
      AddForm,
      OnAddSubmit,
      IsDialogFormSubmitting,
    };
  },
});
</script>
