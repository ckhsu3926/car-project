import { ref, computed, reactive } from 'vue';
import { axiosRequest } from 'boot/axios';
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

const login = async () => {
  const response = await axiosRequest('POST', '/api/login', {
    username: userInfo.username,
    password: Md5.hashStr(userInfo.password),
  })
  
  if(response.result && typeof response.data === 'string') {
    token.value = response.data;
    localStorage.setItem('token', response.data);
    localStorage.setItem('username', userInfo.username || '');
  } else {
    userInfo.password = '';
  };
};

const register = async () => {
  const response = await axiosRequest('POST', '/api/register', {
    username: userInfo.username,
    password: Md5.hashStr(userInfo.password),
  })
  if(response.result && typeof response.data === 'string') {
    token.value = response.data;
    localStorage.setItem('token', response.data);
    localStorage.setItem('username', userInfo.username || '');
  } else {
    clean()
  }
};

const logout = async() => {
  await axiosRequest('POST', '/api/logout')
  clean();
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
