<template>
  <q-dialog v-model="isLogin" :persistent="true">
    <q-card>
      <q-form ref="formRef" @submit="login">
        <q-card-section>
          <q-input
            v-model="userInfo.username"
            label="Username"
            :rules="[
              (val) => (val && val.length > 0) || 'please enter username',
            ]"
          />
        </q-card-section>

        <q-card-section>
          <q-input
            v-model="userInfo.password"
            type="password"
            label="Password"
            :rules="[
              (val) => (val && val.length > 0) || 'please enter password',
            ]"
          />
        </q-card-section>

        <q-card-actions align="evenly">
          <q-btn flat color="primary" label="Login" type="submit" />
          <q-btn
            flat
            color="negative"
            label="Register"
            @click="handlerRegister"
          />
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import { QForm } from 'quasar';

import loginStore from 'stores/login';

export default defineComponent({
  name: 'LoginBox',

  setup() {
    const { isLogin, userInfo, login, register } = loginStore();

    const formRef = ref<QForm | null>(null);
    const handlerRegister = () => {
      formRef.value?.validate().then((success: boolean) => {
        if (success) {
          register();
        }
      });
    };

    return {
      isLogin,
      userInfo,
      login,
      register,

      formRef,
      handlerRegister,
    };
  },
});
</script>
