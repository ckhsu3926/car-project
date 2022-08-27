<template>
  <q-scroll-area style="height: 300px">
    <q-card-section>
      <!-- Date -->
      <q-input
        v-model="refuelingDialogForm.date"
        label="Date"
        mask="date"
        :rules="[(val) => (val && val.length > 0) || 'please enter Date']"
      >
        <template v-slot:append>
          <q-icon name="event" class="cursor-pointer">
            <q-popup-proxy cover transition-show="scale" transition-hide="scale">
              <q-date v-model="refuelingDialogForm.date" :options="(date) => date < now">
                <div class="row items-center justify-end">
                  <q-btn v-close-popup label="Close" color="primary" flat />
                </div>
              </q-date>
            </q-popup-proxy>
          </q-icon>
        </template>
      </q-input>
      <!-- station -->
      <q-select
        label="Station"
        :model-value="refuelingDialogForm.station"
        use-input
        hide-selected
        fill-input
        input-debounce="0"
        :options="autoCompleteStationList"
        @filter="stationFilterFn"
        @input-value="stationSetModel"
        :rules="[(val) => (val && val.length > 0) || 'please enter Station']"
      />
      <!-- octaneNumber -->
      <q-select v-model.number="refuelingDialogForm.octaneNumber" label="Octane" :options="[0, 92, 95, 98]" />
      <!-- unitPrice -->
      <q-input
        v-model.number="refuelingDialogForm.unitPrice"
        label="Unit Price(NT/L)"
        type="number"
        suffix="NT/L"
        step="0.1"
        :rules="[(val) => (val && val < 100) || 'please enter unit price']"
      />
      <!-- count -->
      <q-input
        v-model.number="refuelingDialogForm.count"
        label="Count(L)"
        type="number"
        suffix="L"
        step="0.01"
        :rules="[(val) => (val && val < 1000) || 'please enter gas plus count']"
      />
      <!-- value -->
      <q-input
        v-model.number="refuelingDialogForm.value"
        label="Value(NTD)"
        type="number"
        prefix="NTD$"
        :rules="[(val) => (val && val > 0 && val < 1000000) || 'please enter payment value']"
      />
      <!-- mileage -->
      <q-input
        v-model.number="refuelingDialogForm.mileage"
        label="Mileage(KM)"
        type="number"
        suffix="KM"
        step="0.01"
        :rules="[(val) => (val && val > 0 && val < 1000000) || 'please enter vehicle mileage']"
      />
      <!-- monitorFuelRecord -->
      <q-input
        v-model.number="refuelingDialogForm.monitorFuelRecord"
        label="Monitor Fuel Record"
        type="number"
        step="0.1"
        :rules="[(val) => (val && val < 100) || '']"
      />
    </q-card-section>
  </q-scroll-area>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import refuelingStore from 'stores/refueling';
import gasStationStore from 'stores/gasStation';
import moment from 'moment';

export default defineComponent({
  name: 'RefuelingAddBoxInput',

  setup() {
    const { refuelingDialogForm } = refuelingStore();
    const { list: gasStationList } = gasStationStore();

    const now = moment().format('YYYY/MM/DD');

    const stationList = gasStationList.value.map((s) => s.name);
    const autoCompleteStationList = ref(stationList);
    const stationFilterFn = (val: string, update: (_: () => void) => void) => {
      update(() => {
        autoCompleteStationList.value = stationList.filter((v) => v.toLocaleLowerCase().indexOf(val) > -1);
      });
    };
    const stationSetModel = (val: string) => {
      refuelingDialogForm.value.station = val;
    };

    return {
      refuelingDialogForm,
      now,
      autoCompleteStationList,
      stationFilterFn,
      stationSetModel,
    };
  },
});
</script>
