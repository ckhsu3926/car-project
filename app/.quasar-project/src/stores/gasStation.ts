import { ref } from 'vue';
import { Dialog } from 'quasar';
import { axiosRequest } from 'boot/axios';

export interface gasStation {
  id: number;
  name: string;
}

const list = ref(<gasStation[]>[]);
const getList = async () => {
  const response = await axiosRequest('GET', '/api/user/gas/station/list');
  if (response.result && response.data instanceof Array) {
    list.value = <gasStation[]>response.data;
  }
};

const addStation = async () => {
  Dialog.create({
    title: 'Add Gas Station',
    prompt: {
      model: '',
      type: 'text', // optional
    },
    cancel: true,
  }).onOk(async (data) => {
    const response = await axiosRequest('POST', '/api/user/gas/station/add', { name: data });
    if (response.result && response.data instanceof Array) {
      list.value = <gasStation[]>response.data;
    }
  });
};
const deleteStation = async (id: number, name: string) => {
  Dialog.create({
    title: 'Delete Gas Station',
    message: `Do you want to delete station: ${name} ?`,
    cancel: true,
  }).onOk(async () => {
    const response = await axiosRequest('DELETE', '/api/user/gas/station/delete', { recordID: id });
    if (response.result && response.data instanceof Array) {
      list.value = <gasStation[]>response.data;
    }
  });
};

export default () => {
  return {
    list,
    getList,
    addStation,
    deleteStation,
  };
};
