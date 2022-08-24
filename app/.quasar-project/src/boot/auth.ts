import { boot } from 'quasar/wrappers';
import loginStore from 'stores/login';

export default boot(async ({}) => {
  const { getAuthInfo } = loginStore();

  await getAuthInfo();
});
