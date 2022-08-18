import { ref, reactive } from 'vue';
import { axiosRequest } from 'boot/axios';
import { Notify } from 'quasar';

export interface vehicle {
  id: number;
  name: string;
  license: string;
  company: string;
  model: string;
}
export interface vehicleDetail {
  id: number;
  name: string;
  license: string;
  company: string;
  model: string;
  engineDisplacement: number;
  engineNumber: string;
  defaultOctaneNumber: number;
  purchase: number;
  purchaseDate: string;
  purchaseLocation: string;
  purchaseMileage: number;
  sold: number;
  soldDate: string;
  soldMileage: number;
  mileageReset: number;
}

// dialog
const IsVehicleBoxOpen = ref(false);
const VehicleBoxMode = ref('add');

// dialog-add
const AddForm = reactive(<vehicle>{});
const clearAddForm = () => {
  AddForm.name = '';
  AddForm.license = '';
  AddForm.company = '';
  AddForm.model = '';
};
const IsDialogFormSubmitting = ref(false);

const OnAddOpen = () => {
  IsVehicleBoxOpen.value = true;
  VehicleBoxMode.value = 'add';
  clearAddForm();

  return;
};
const OnAddSubmit = async () => {
  IsDialogFormSubmitting.value = true;
  const response = await axiosRequest('POST', '/api/vehicle/add', {
    name: AddForm.name,
    license: AddForm.license,
    company: AddForm.company,
    model: AddForm.model,
  });
  if (response.result) {
    IsVehicleBoxOpen.value = false;
    IsDialogFormSubmitting.value = false;

    if (response.data instanceof Array) {
      vehicleList.value = response.data;
    }
  } else {
    OnAddOpen();
    IsDialogFormSubmitting.value = false;
  }

  return;
};

// dialog-edit
const EditForm = ref(<vehicleDetail>{});
const OnEditSubmit = async () => {
  IsDialogFormSubmitting.value = true;
  const response = await axiosRequest('PUT', '/api/vehicle/edit', EditForm.value);
  if (response.result) {
    IsDialogFormSubmitting.value = false;

    if (response.data instanceof Object) {
      EditForm.value = <vehicleDetail>response.data;
    }
    Notify.create({ type: 'positive', message: 'Success!' });
    getVehicliList();
  } else {
    IsDialogFormSubmitting.value = false;
    Notify.create({ type: 'negative', message: 'Failed!' });
  }

  return;
};

// list
const vehicleList = ref(<vehicle[]>{});
const getVehicliList = async () => {
  const response = await axiosRequest('GET', '/api/vehicle/list');
  if (response.result && response.data instanceof Array) {
    vehicleList.value = response.data;
  }

  return;
};
const onVehicleClick = async (id: 0) => {
  const response = await axiosRequest('GET', '/api/vehicle/', { vehicleID: id });
  if (response.result && response.data instanceof Object) {
    IsVehicleBoxOpen.value = true;
    VehicleBoxMode.value = 'edit';
    EditForm.value = <vehicleDetail>response.data;
  }

  return;
};

export default () => {
  return {
    // dialog
    IsVehicleBoxOpen,
    VehicleBoxMode,
    // add
    AddForm,
    OnAddOpen,
    OnAddSubmit,
    IsDialogFormSubmitting,
    // edit
    EditForm,
    OnEditSubmit,

    // list
    vehicleList,
    getVehicliList,
    onVehicleClick,
  };
};
