import { ref, reactive } from 'vue';
import { axiosRequest } from 'boot/axios';

const IsVehicleBoxOpen = ref(false);
const VehicleBoxMode = ref('add');

const AddForm = reactive({
  name: '',
  license: '',
  company: '',
  model: '',
});
const IsAddFormSubmitting = ref(false);

const OnAddOpen = () => {
  IsVehicleBoxOpen.value = true;
  VehicleBoxMode.value = 'add';
  AddForm.name = '';
  AddForm.license = '';
  AddForm.company = '';
  AddForm.model = '';

  return;
};
const OnAddSubmit = async () => {
  IsAddFormSubmitting.value = true;
  const response = await axiosRequest('POST', '/api/vehicle/add', {
    name: AddForm.name,
    license: AddForm.license,
    company: AddForm.company,
    model: AddForm.model,
  })
    if(response.result) {
      // TODO
      console.log(response);
      IsVehicleBoxOpen.value = false;
      IsAddFormSubmitting.value = false;
    } else {
      OnAddOpen();
      IsAddFormSubmitting.value = false;
    }

  return;
};

export default () => {
  return {
    IsVehicleBoxOpen,
    VehicleBoxMode,

    AddForm,
    OnAddOpen,
    OnAddSubmit,
    IsAddFormSubmitting,
  };
};
