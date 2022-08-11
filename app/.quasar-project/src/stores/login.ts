import { ref, computed, reactive } from 'vue';
import { api } from 'boot/axios';
import { Md5 } from 'ts-md5/dist/md5';

const token = ref(localStorage.getItem('token'));
const isLogin = computed(() => !token.value);

const userInfo = reactive({
  username: localStorage.getItem('username'),
  password: '',
});

const clean = () => {
  localStorage.removeItem('token');
  localStorage.removeItem('username');
  token.value = null;
  userInfo.username = '';
  userInfo.password = '';
};

const login = () => {
  api
    .post('/api/login', {
      username: userInfo.username,
      password: Md5.hashStr(userInfo.password),
    })
    .then((response) => {
      token.value = response.data.data;
      localStorage.setItem('token', response.data.data);
      localStorage.setItem('username', userInfo.username || '');
    })
    .catch(() => {
      userInfo.password = '';
    });
};

const register = () => {
  api
    .post('/api/register', {
      username: userInfo.username,
      password: Md5.hashStr(userInfo.password),
    })
    .then((response) => {
      token.value = response.data.data;
      localStorage.setItem('token', response.data.data);
      localStorage.setItem('username', userInfo.username || '');
    })
    .catch(clean);
};

const logout = () => {
  api.post('/api/logout').then(clean).catch(clean);
};

export default () => {
  return {
    token,
    isLogin,
    userInfo,
    login,
    register,
    logout,
  };
};
