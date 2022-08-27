<template>
  <q-scroll-area style="height: 300px">
    <q-card-section>
      <!-- Date -->
      <q-input
        v-model="maintenanceDialogForm.date"
        label="Date"
        mask="date"
        :rules="[(val) => (val && val.length > 0) || 'please enter Date']"
      >
        <template v-slot:append>
          <q-icon name="event" class="cursor-pointer">
            <q-popup-proxy cover transition-show="scale" transition-hide="scale">
              <q-date v-model="maintenanceDialogForm.date" :options="(date) => date < now">
                <div class="row items-center justify-end">
                  <q-btn v-close-popup label="Close" color="primary" flat />
                </div>
              </q-date>
            </q-popup-proxy>
          </q-icon>
        </template>
      </q-input>
      <!-- mileage -->
      <q-input
        v-model.number="maintenanceDialogForm.mileage"
        label="Mileage(KM)"
        type="number"
        suffix="KM"
        step="0.01"
        :rules="[(val) => (val && val > 0 && val < 1000000) || 'please enter vehicle mileage']"
      />
      <!-- shop -->
      <q-input
        v-model.number="maintenanceDialogForm.shop"
        label="Shop"
        type="text"
        :rules="[(val) => (val && val.length > 0) || 'please enter shop']"
      />
      <!-- amount -->
      <q-input
        v-model.number="maintenanceDialogForm.amount"
        label="Amount(NTD$)"
        type="number"
        prefix="NTD$"
        :rules="[(val) => (val && val > 0 && val < 1000000) || 'please enter payment value']"
      />
    </q-card-section>
  </q-scroll-area>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import maintenanceStore from 'stores/maintenance';
import moment from 'moment';

export default defineComponent({
  name: 'MaintenanceBoxInput',

  setup() {
    const { maintenanceDialogForm } = maintenanceStore();

    const now = moment().format('YYYY/MM/DD');

    return {
      maintenanceDialogForm,
      now,
    };
  },
});
</script>
