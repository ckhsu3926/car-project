<template>
  <q-table :grid="$q.screen.xs" :rows="list" :columns="columns" row-key="id" :rows-per-page-options="[20, 50]">
    <template v-slot:top>
      <div class="q-table__title">Maintenance Record</div>
      <q-space />
      <q-btn color="secondary" label="Add" @click="OnAddOpen(vehicleID)" />
    </template>

    <!-- table -->
    <template v-slot:body-cell-actions="props">
      <q-td :props="props">
        <q-btn size="md" color="warning" dense icon="settings" class="q-mr-sm" @click="OnEditOpen(props.row)" />
        <q-btn
          size="md"
          color="negative"
          dense
          icon="delete"
          @click="OnDeleteSubmit(props.row.vehicleID, props.row.id)"
        />
      </q-td>
    </template>
    <!-- grid card -->
    <template v-slot:item="props">
      <div class="q-pa-xs col-xs-12 col-sm-6 col-md-4 col-lg-3 grid-style-transition">
        <q-card>
          <q-list dense>
            <q-item v-for="col in props.cols" :key="col.name">
              <q-item-section>
                <q-item-label>{{ col.label }}</q-item-label>
              </q-item-section>
              <q-item-section side>
                <q-item-label v-if="col.name != 'actions'" caption>{{ col.value }}</q-item-label>
                <div v-else>
                  <q-btn
                    size="sm"
                    color="warning"
                    dense
                    icon="settings"
                    class="q-mr-md"
                    @click="OnEditOpen(props.row)"
                  />
                  <q-btn
                    size="sm"
                    color="negative"
                    dense
                    icon="delete"
                    @click="OnDeleteSubmit(props.row.vehicleID, props.row.id)"
                  />
                </div>
              </q-item-section>
            </q-item>
          </q-list>
        </q-card>
      </div>
    </template>
  </q-table>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue';
import maintenanceStore, { maintenanceRecord } from 'stores/maintenance';

const columns = [
  {
    name: 'date',
    required: true,
    label: 'Date',
    align: 'left',
    field: (row: maintenanceRecord) => row.date,
  },
  {
    name: 'mileage',
    required: true,
    label: 'Mileage(KM)',
    align: 'left',
    field: (row: maintenanceRecord) => row.mileage,
  },
  {
    name: 'shop',
    required: true,
    label: 'Shop',
    align: 'left',
    field: (row: maintenanceRecord) => row.shop,
  },
  {
    name: 'amount',
    required: true,
    label: 'Amount(NTD$)',
    align: 'left',
    field: (row: maintenanceRecord) => row.amount,
  },
  {
    name: 'actions',
    label: 'Actions',
    align: 'left',
  },
];

export default defineComponent({
  name: 'MaintenanceRecordTable',
  props: {
    vehicleID: {
      type: Number,
      required: true,
    },
    list: {
      type: Array as PropType<Array<maintenanceRecord>>,
      required: true,
    },
  },

  setup() {
    const { OnAddOpen, OnEditOpen, OnDeleteSubmit } = maintenanceStore();

    return {
      columns,
      OnAddOpen,
      OnEditOpen,
      OnDeleteSubmit,
    };
  },
});
</script>