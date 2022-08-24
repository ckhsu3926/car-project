<template>
  <q-card class="q-ma-xs hide-scrollbar">
    <q-form @submit="OnEditSubmit">
      <q-card-section class="bg-light-blue text-white row justify-between">
        <div class="text-h6">Edit Vehicle</div>
        <q-btn float round icon="done" size="sm" color="secondary" type="submit" :disable="IsDialogFormSubmitting" />
      </q-card-section>

      <q-scroll-area style="height: 60vh">
        <q-card-section>
          <q-input
            v-model="EditForm.name"
            label="Name"
            :rules="[(val) => (val && val.length > 0) || 'please enter vehicle name']"
          />
          <q-input
            v-model="EditForm.license"
            label="License"
            :rules="[(val) => (val && val.length > 0) || 'please enter vehicle license']"
          />
          <q-input v-model="EditForm.company" label="Company" />
          <q-input v-model="EditForm.model" label="Model" />
        </q-card-section>
        <q-card-section>
          <q-input
            v-model.number="EditForm.engineDisplacement"
            label="Engine Displacement(c.c.)"
            type="number"
            step="0.1"
          />
          <q-input v-model="EditForm.engineNumber" label="Engine Number" />
          <q-select v-model="EditForm.defaultOctaneNumber" :options="[0, 92, 95, 98]" label="Default Octane Number" />
        </q-card-section>
        <q-card-section>
          <q-expansion-item expand-separator label="Purchase">
            <q-input v-model.number="EditForm.purchase" label="Purchase" type="number" />
            <q-input v-model="EditForm.purchaseLocation" label="Location" />
            <q-input v-model.number="EditForm.purchaseMileage" label="Mileage" type="number" />
            <q-input v-model="EditForm.purchaseDate" label="Date" mask="date">
              <template v-slot:append>
                <q-icon name="event" class="cursor-pointer">
                  <q-popup-proxy cover transition-show="scale" transition-hide="scale">
                    <q-date v-model="EditForm.purchaseDate" :options="(date) => date < now">
                      <div class="row items-center justify-end">
                        <q-btn v-close-popup label="Close" color="primary" flat />
                      </div>
                    </q-date>
                  </q-popup-proxy>
                </q-icon>
              </template>
            </q-input>
          </q-expansion-item>

          <q-expansion-item expand-separator label="Sold">
            <q-input v-model.number="EditForm.sold" label="Sold" type="number" />
            <q-input v-model.number="EditForm.soldMileage" label="Mileage" type="number" />
            <q-input v-model="EditForm.soldDate" label="Date" mask="date">
              <template v-slot:append>
                <q-icon name="event" class="cursor-pointer">
                  <q-popup-proxy cover transition-show="scale" transition-hide="scale">
                    <q-date v-model="EditForm.soldDate" :options="(date) => date > EditForm.purchaseDate && now > date">
                      <div class="row items-center justify-end">
                        <q-btn v-close-popup label="Close" color="primary" flat />
                      </div>
                    </q-date>
                  </q-popup-proxy>
                </q-icon>
              </template>
            </q-input>
          </q-expansion-item>

          <q-expansion-item expand-separator label="Mileage Reset">
            <q-input v-model.number="EditForm.mileageReset" label="Mileage" type="number" />
          </q-expansion-item>
        </q-card-section>
      </q-scroll-area>
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
import moment from 'moment';

export default defineComponent({
  name: 'LoginBoxEdit',

  setup() {
    const { EditForm, OnEditSubmit, IsDialogFormSubmitting } = vehicleStore();

    const now = moment().format('YYYY/MM/DD');

    return {
      EditForm,
      OnEditSubmit,
      IsDialogFormSubmitting,

      now,
    };
  },
});
</script>
