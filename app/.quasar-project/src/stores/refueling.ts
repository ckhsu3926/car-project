import { ref } from 'vue';
import { Dialog } from 'quasar';
import { axiosRequest } from 'boot/axios';

export interface refuelingRecord {
  id: number;
  vehicleID: number;
  date: string;
  station: string;
  octaneNumber: number;
  unitPrice: number;
  count: number;
  value: number;
  mileage: number;
  monitorFuelRecord: number;
}

// dialog
const isRefuelingDialogOpen = ref(false);
const isDialogFormSubmitting = ref(false);
const refuelingDialogMode = ref('add');

// dialog-add
const refuelingDialogForm = ref(<refuelingRecord>{});

const OnAddOpen = (vehicleID: number) => {
  isRefuelingDialogOpen.value = true;
  isDialogFormSubmitting.value = false;
  refuelingDialogMode.value = 'add';
  refuelingDialogForm.value = <refuelingRecord>{};
  refuelingDialogForm.value.vehicleID = vehicleID;
};
const OnAddSubmit = async () => {
  isDialogFormSubmitting.value = true;
  const response = await axiosRequest('POST', '/api/refueling/add', refuelingDialogForm.value);
  if (response.result) {
    isRefuelingDialogOpen.value = false;

    if (response.data instanceof Array) {
      refuelingList.value = <refuelingRecord[]>response.data;
    }
    isDialogFormSubmitting.value = false;
  } else {
    OnAddOpen(refuelingDialogForm.value.vehicleID);
    isDialogFormSubmitting.value = false;
  }
};

// list
const refuelingList = ref(<refuelingRecord[]>[]);
const getRefuelingList = async (vehicleID: number) => {
  const response = await axiosRequest('GET', '/api/refueling/list', { vehicleID });
  if (response.result && response.data instanceof Array) {
    refuelingList.value = <refuelingRecord[]>response.data;
  }
};

// delete
const OnDeleteSubmit = async (vehicleID: number, refuelingRecordID: number) => {
  Dialog.create({
    title: 'Delete Refueling Record',
    message: 'Are you sure to delete refueling record ?',
    cancel: true,
  }).onOk(async () => {
    const response = await axiosRequest('DELETE', '/api/refueling/delete', { vehicleID, refuelingRecordID });
    if (response.result && response.data instanceof Array) {
      refuelingList.value = <refuelingRecord[]>response.data;
    }
  });
};

// edit
const OnEditOpen = (row: refuelingRecord) => {
  isRefuelingDialogOpen.value = true;
  refuelingDialogMode.value = 'edit';
  refuelingDialogForm.value = Object.assign(<refuelingRecord>{}, row);
  isDialogFormSubmitting.value = false;
};
const OnEditSubmit = async () => {
  isDialogFormSubmitting.value = true;
  const response = await axiosRequest('PUT', '/api/refueling/update', refuelingDialogForm.value);
  if (response.result) {
    isRefuelingDialogOpen.value = false;

    if (response.data instanceof Array) {
      refuelingList.value = <refuelingRecord[]>response.data;
    }
    isDialogFormSubmitting.value = false;
  } else {
    OnEditOpen(refuelingDialogForm.value);
    isDialogFormSubmitting.value = false;
  }
};

export default () => {
  return {
    // // dialog
    isRefuelingDialogOpen,
    isDialogFormSubmitting,
    refuelingDialogMode,
    // // add
    refuelingDialogForm,
    OnAddOpen,
    OnAddSubmit,
    // // edit
    OnEditOpen,
    OnEditSubmit,

    // list
    refuelingList,
    getRefuelingList,
    OnDeleteSubmit,
  };
};
