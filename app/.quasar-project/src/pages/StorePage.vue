<template>
  <q-page padding>
    <q-table grid title="Gas Stations" :rows="list" :columns="columns" row-key="id" :rows-per-page-options="[0]">
      <template v-slot:top>
        <div class="q-table__title">Gas Station List</div>
        <q-space />
        <q-btn color="secondary" label="Add" @click="addStation" />
      </template>

      <template v-slot:item="props">
        <div class="q-pa-xs col-xs-12 col-sm-6 col-md-4">
          <q-card>
            <q-card-section class="bg-indigo-6 text-white text-center">
              <strong>{{ props.row.name }}</strong>
            </q-card-section>
            <q-separator />
            <q-card-section class="flex flex-center">
              <q-btn color="negative" icon="delete" @click="deleteStation(props.row.id, props.row.name)" />
            </q-card-section>
          </q-card>
        </div>
      </template>
    </q-table>
  </q-page>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import gasStationStore, { gasStation } from 'stores/gasStation';

const columns = [
  {
    name: 'name',
    required: true,
    label: 'Station Name',
    align: 'left',
    field: (row: gasStation) => row.name,
  },
];

export default defineComponent({
  name: 'StorePage',
  setup() {
    const { list, getList, addStation, deleteStation } = gasStationStore();

    getList();

    return {
      columns,
      list,
      addStation,
      deleteStation,
    };
  },
});
</script>
