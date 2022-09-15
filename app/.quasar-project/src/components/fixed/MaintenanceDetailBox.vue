<template>
  <q-dialog v-model="isMaintenanceDetailDialogOpen">
    <q-card class="q-ma-xs">
      <q-card-section class="bg-light-blue text-white">
        <div class="text-h6">Maintenance Record Detail</div>
      </q-card-section>

      <q-card-section>
        <q-form ref="formRef" @submit="onAddSubmit" @reset="onFormReset">
          <q-table
            :grid="$q.screen.xs"
            :rows="maintenanceDetailList"
            :columns="columns"
            row-key="name"
            :rows-per-page-options="[0]"
            binary-state-sort
            :loading="loading"
          >
            <template #loading>
              <q-inner-loading showing color="primary" />
            </template>

            <template #body="props">
              <q-tr :props="props">
                <q-td key="name" :props="props">
                  {{ props.row.name }}
                </q-td>
                <q-td key="value" :props="props">
                  {{ props.row.value }}
                  <q-popup-edit v-model.number="props.row.value" buttons v-slot="scope" @save="onRowEdit">
                    <q-input type="number" v-model.number="scope.value" maxlength="6" dense autofocus />
                  </q-popup-edit>
                </q-td>
                <q-td key="content" :props="props" class="ellipsis" style="max-width: 240px">
                  {{ props.row.content }}
                  <q-popup-edit v-model="props.row.content" buttons v-slot="scope" @save="onRowEdit">
                    <q-input v-model="scope.value" type="textarea" maxlength="150" dense autofocus counter />
                  </q-popup-edit>
                </q-td>
                <q-td key="actions" :props="props">
                  <q-btn size="md" color="negative" round dense icon="delete" @click="onRowDelete(props.rowIndex)" />
                </q-td>
              </q-tr>
            </template>

            <template #bottom-row>
              <!-- add -->
              <q-tr no-hover>
                <template v-if="!isAdd">
                  <q-td></q-td>
                  <q-td></q-td>
                  <q-td></q-td>
                  <q-td>
                    <q-btn type="submit" size="md" color="secondary" round dense icon="add" @click="isAdd = !isAdd" />
                  </q-td>
                </template>
                <template v-else>
                  <q-td>
                    <q-input
                      v-model="addObject.name"
                      maxlength="45"
                      dense
                      hide-bottom-space
                      :rules="[(val) => (val && val.length > 0) || 'please enter Name']"
                    />
                  </q-td>
                  <q-td>
                    <q-input
                      v-model.number="addObject.value"
                      type="number"
                      maxlength="10"
                      dense
                      hide-bottom-space
                      :rules="[(val) => val > -1 || 'please enter Value']"
                    />
                  </q-td>
                  <q-td>
                    <q-input v-model="addObject.content" type="textarea" maxlength="150" rows="2" dense counter />
                  </q-td>
                  <q-td>
                    <q-btn type="reset" size="md" color="warning" round dense icon="remove" @click="isAdd = !isAdd" />
                    <q-btn type="submit" size="md" color="secondary" round dense icon="add" />
                  </q-td>
                </template>
              </q-tr>
            </template>

            <!-- grid card -->
            <template #item="props">
              <div class="q-pb-md col-xs-12 col-sm-6 col-md-4 col-lg-3 grid-style-transition">
                <q-card>
                  <q-list dense>
                    <!-- name -->
                    <q-item>
                      <q-item-section>
                        <q-item-label>Name</q-item-label>
                      </q-item-section>
                      <q-item-section side>
                        <q-item-label caption>{{ props.row.name }}</q-item-label>
                      </q-item-section>
                    </q-item>
                    <!-- value -->
                    <q-item v-ripple>
                      <q-item-section>
                        <q-item-label>Value</q-item-label>
                      </q-item-section>
                      <q-item-section side>
                        <q-item-label caption>{{ props.row.value }}</q-item-label>
                      </q-item-section>
                      <q-popup-edit v-model.number="props.row.value" buttons v-slot="scope" @save="onRowEdit">
                        <q-input v-model.number="scope.value" type="number" maxlength="6" dense autofocus />
                      </q-popup-edit>
                    </q-item>
                    <!-- content -->
                    <q-item v-ripple>
                      <q-item-section>
                        <q-item-label>Content</q-item-label>
                      </q-item-section>
                      <q-item-section side>
                        <q-item-label caption>{{ props.row.content }}</q-item-label>
                      </q-item-section>
                      <q-popup-edit v-model="props.row.content" buttons v-slot="scope" @save="onRowEdit">
                        <q-input v-model="scope.value" type="textarea" maxlength="150" dense autofocus counter />
                      </q-popup-edit>
                    </q-item>
                    <!-- actions -->
                    <q-item>
                      <q-item-section />
                      <q-item-section side>
                        <q-btn size="sm" color="negative" icon="delete" @click="onRowDelete(props.rowIndex)" />
                      </q-item-section>
                    </q-item>
                  </q-list>
                </q-card>
              </div>

              <template v-if="props.rowIndex === maintenanceDetailList.length - 1">
                <q-btn v-if="!isAdd" class="full-width" color="secondary" label="Add" @click="isAdd = !isAdd" />
                <MaintenanceDetailBoxAddCard
                  v-else
                  v-model:name="addObject.name"
                  v-model:value="addObject.value"
                  v-model:content="addObject.content"
                >
                  <template #actions>
                    <q-item-section>
                      <q-btn type="reset" size="sm" color="negative" icon="delete" @click="isAdd = !isAdd" />
                    </q-item-section>
                    <q-item-section>
                      <q-btn class="q-my-sm" type="submit" size="sm" color="secondary" icon="add" />
                    </q-item-section>
                  </template>
                </MaintenanceDetailBoxAddCard>
              </template>
            </template>

            <template #no-data v-if="$q.screen.xs">
              <q-btn v-if="!isAdd" class="full-width" color="secondary" label="Add" @click="isAdd = !isAdd" />
              <MaintenanceDetailBoxAddCard
                v-else
                v-model:name="addObject.name"
                v-model:value="addObject.value"
                v-model:content="addObject.content"
              >
                <template #actions>
                  <q-item-section>
                    <q-btn class="q-my-sm" type="submit" size="sm" color="secondary" icon="add" />
                  </q-item-section>
                </template>
              </MaintenanceDetailBoxAddCard>
            </template>
          </q-table>
        </q-form>
      </q-card-section>
    </q-card>
  </q-dialog>
</template>

<script lang="ts">
import { defineComponent, ref, onUpdated } from 'vue';
import { Dialog, QForm } from 'quasar';
import maintenanceStore, { maintenanceRecordDetail } from 'stores/maintenance';
import MaintenanceDetailBoxAddCard from 'components/fixed/MaintenanceDetailBoxAddCard.vue';

const columns = [
  {
    name: 'name',
    required: true,
    label: 'Name',
    align: 'left',
    field: (row: maintenanceRecordDetail) => row.name,
    sortable: true,
  },
  {
    name: 'value',
    required: true,
    label: 'Value($NTD)',
    align: 'left',
    field: (row: maintenanceRecordDetail) => row.value,
    sortable: true,
  },
  {
    name: 'content',
    label: 'Content',
    align: 'left',
    field: (row: maintenanceRecordDetail) => row.content,
  },
  {
    name: 'actions',
    label: 'Actions',
    align: 'left',
  },
];

export default defineComponent({
  name: 'MaintenanceDetailBox',
  components: {
    MaintenanceDetailBoxAddCard,
  },
  setup() {
    const { isMaintenanceDetailDialogOpen, maintenanceDetailList, updateMaintenanceDetailList } = maintenanceStore();

    const formRef = ref(<QForm | null>null);
    const onFormReset = () => {
      addObject.value = <maintenanceRecordDetail>{};
    };

    const loading = ref(false);
    const onRowEdit = async () => {
      loading.value = true;
      await updateMaintenanceDetailList();
      loading.value = false;
    };
    const onRowDelete = async (rowIndex: number) => {
      Dialog.create({
        title: 'Delete Maintenance Detail Record',
        message: 'Are you sure to delete maintenance detail record ?',
        cancel: true,
      }).onOk(async () => {
        loading.value = true;
        maintenanceDetailList.value.splice(rowIndex, 1);
        await updateMaintenanceDetailList();
        loading.value = false;
      });
    };

    const isAdd = ref(false);
    const addObject = ref(<maintenanceRecordDetail>{});
    const onAddSubmit = async () => {
      loading.value = true;
      maintenanceDetailList.value.push(addObject.value);
      await updateMaintenanceDetailList();
      if (formRef.value) {
        addObject.value = <maintenanceRecordDetail>{};
        formRef.value.reset();
      }
      loading.value = false;
    };

    onUpdated(() => {
      isAdd.value = false;
    });

    return {
      isMaintenanceDetailDialogOpen,
      maintenanceDetailList,

      columns,
      loading,
      isAdd,
      onRowEdit,
      onRowDelete,
      formRef,
      onFormReset,
      addObject,
      onAddSubmit,
    };
  },
});
</script>
