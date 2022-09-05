import { ref } from 'vue';
import { Dialog } from 'quasar';
import { axiosRequest } from 'boot/axios';

export interface maintenanceRecord {
  id: number;
  vehicleID: number;
  date: string;
  shop: string;
  mileage: number;
  amount: number;
}
export interface maintenanceRecordDetail {
  name: string;
  value: number;
  content: string;
}

// list
const maintenanceList = ref(<maintenanceRecord[]>[]);
const getMaintenanceList = async (vehicleID: number) => {
  const response = await axiosRequest('GET', '/api/maintenance/record/list', { vehicleID });
  if (response.result && response.data instanceof Array) {
    maintenanceList.value = <maintenanceRecord[]>response.data;
  }
};

// dialog
const isMaintenanceDialogOpen = ref(false);
const isDialogFormSubmitting = ref(false);
const maintenanceDialogMode = ref('add');

// dialog-add
const maintenanceDialogForm = ref(<maintenanceRecord>{});
const clearMaintenanceDialogForm = () => {
  maintenanceDialogForm.value = <maintenanceRecord>{};
};

const OnAddOpen = (vehicleID: number) => {
  isMaintenanceDialogOpen.value = true;
  isDialogFormSubmitting.value = false;
  maintenanceDialogMode.value = 'add';
  clearMaintenanceDialogForm();
  maintenanceDialogForm.value.vehicleID = vehicleID;
};
const OnAddSubmit = async () => {
  isDialogFormSubmitting.value = true;
  const response = await axiosRequest('POST', '/api/maintenance/record/create', maintenanceDialogForm.value);
  if (response.result) {
    isMaintenanceDialogOpen.value = false;

    if (response.data instanceof Array) {
      maintenanceList.value = <maintenanceRecord[]>response.data;
    }
    isDialogFormSubmitting.value = false;
  } else {
    OnAddOpen(maintenanceDialogForm.value.vehicleID);
    isDialogFormSubmitting.value = false;
  }
};

// edit
const OnEditOpen = (row: maintenanceRecord) => {
  isMaintenanceDialogOpen.value = true;
  maintenanceDialogMode.value = 'edit';
  maintenanceDialogForm.value = Object.assign(<maintenanceRecord>{}, row);
  isDialogFormSubmitting.value = false;
};
const OnEditSubmit = async () => {
  isDialogFormSubmitting.value = true;
  const response = await axiosRequest('PUT', '/api/maintenance/record/update', maintenanceDialogForm.value);
  if (response.result) {
    isMaintenanceDialogOpen.value = false;

    if (response.data instanceof Array) {
      maintenanceList.value = <maintenanceRecord[]>response.data;
    }
    isDialogFormSubmitting.value = false;
  } else {
    OnEditOpen(maintenanceDialogForm.value);
    isDialogFormSubmitting.value = false;
  }
};

// delete
const OnDeleteSubmit = async (vehicleID: number, maintenanceRecordID: number) => {
  Dialog.create({
    title: 'Delete Maintenance Record',
    message: 'Are you sure to delete maintenance record ?',
    cancel: true,
  }).onOk(async () => {
    const response = await axiosRequest('DELETE', '/api/maintenance/record/delete', { vehicleID, maintenanceRecordID });
    if (response.result && response.data instanceof Array) {
      maintenanceList.value = <maintenanceRecord[]>response.data;
    }
  });
};

// detail
const isMaintenanceDetailDialogOpen = ref(false);
const dialogMaintenanceID = ref(0);
const maintenanceDetailList = ref(<maintenanceRecordDetail[]>[]);
const onRowClick = async (row: maintenanceRecord) => {
  dialogMaintenanceID.value = row.id;
  const response = await axiosRequest('GET', '/api/maintenance/record/detail/', {
    maintenanceRecordID: dialogMaintenanceID.value,
  });
  if (response.result && response.data instanceof Array) {
    maintenanceDetailList.value = <maintenanceRecordDetail[]>response.data;
    isMaintenanceDetailDialogOpen.value = true;
  }
};
const updateMaintenanceDetailList = async () => {
  const response = await axiosRequest(
    'POST',
    `/api/maintenance/record/detail/?maintenanceRecordID=${dialogMaintenanceID.value}`,
    maintenanceDetailList.value
  );
  if (response.result && response.data instanceof Array) {
    maintenanceDetailList.value = <maintenanceRecordDetail[]>response.data;
  }
};

export default () => {
  return {
    // // dialog
    isMaintenanceDialogOpen,
    isDialogFormSubmitting,
    maintenanceDialogMode,
    // // add
    maintenanceDialogForm,
    OnAddOpen,
    OnAddSubmit,
    // // edit
    OnEditOpen,
    OnEditSubmit,

    // list
    maintenanceList,
    getMaintenanceList,
    OnDeleteSubmit,

    // detail
    onRowClick,
    isMaintenanceDetailDialogOpen,
    maintenanceDetailList,
    updateMaintenanceDetailList,
  };
};
