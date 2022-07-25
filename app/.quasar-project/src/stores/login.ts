import { ref, computed, reactive } from 'vue';

const token = ref(localStorage.getItem('token'));
const isLogin = computed(() => !token.value);

const userInfo = reactive({
  username: '',
  password: '',
});

export const login = () => {
  localStorage.setItem('token', userInfo.username);
  token.value = userInfo.username;
};
const logout = () => {
  localStorage.removeItem('token');
  token.value = null;
  userInfo.username = '';
  userInfo.password = '';
};

export default () => {
  return {
    token,
    isLogin,
    userInfo,
    login,
    logout,
  };
};
