import { boot } from 'quasar/wrappers';
import axios, { AxiosResponse, AxiosInstance } from 'axios';

import { Dialog } from 'quasar';
import loginStore from 'stores/login';
const { token, clean } = loginStore();

declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    $axios: AxiosInstance;
  }
}

// Be careful when using SSR for cross-request state pollution
// due to creating a Singleton instance here;
// If any client changes this (global) instance, it might be a
// good idea to move this instance creation inside of the
// "export default () => {}" function below (which runs individually
// for each client)
const api = axios.create({ baseURL: process.env.API });
api.interceptors.request.use((config) => {
  if (token.value) {
    config.headers['token'] = token.value;
  }
  return config;
});

export default boot(({ app }) => {
  // for use inside Vue files (Options API) through this.$axios and this.$api

  app.config.globalProperties.$axios = axios;
  // ^ ^ ^ this will allow you to use this.$axios (for Vue Options API form)
  //       so you won't necessarily have to import axios in each vue file

  app.config.globalProperties.$api = api;
  // ^ ^ ^ this will allow you to use this.$api (for Vue Options API form)
  //       so you can easily perform requests against your app's API
});

async function axiosRequest(
  method: string,
  url: string,
  data: null | unknown = null
): Promise<{ result: number; data: unknown; error: string }> {
  method = method.toLowerCase();

  let req: Promise<AxiosResponse<unknown>>;
  switch (method) {
    case 'post':
      req = api.post(url, data);
      break;
    case 'delete':
      req = api.delete(url, { params: data });
      break;
    case 'put':
      req = api.put(url, data);
      break;
    case 'patch':
      req = api.patch(url, data);
      break;
    case 'get':
    default:
      req = api.get(url, { params: data });
  }

  let res: AxiosResponse | undefined;
  try {
    res = await req;
  } catch (error) {
    if (axios.isAxiosError(error)) {
      res = error.response;

      if (error.response?.status === 403) {
        clean();
      } else {
        Dialog.create({
          title: 'Alert',
          message: res?.data.error,
        });
      }
    }
  }

  return res?.data;
}

export { api, axiosRequest };
