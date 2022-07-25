<template>
  <q-dialog v-model="isLogin" :persistent="true">
    <q-card>
      <q-card-section>
        <q-input v-model="userInfo.username" label="Username" />
      </q-card-section>

      <q-card-section>
        <q-input v-model="userInfo.password" label="Password" />
      </q-card-section>

      <q-card-actions align="evenly">
        <q-btn flat color="primary" label="Login" @click="login" />
        <q-btn flat color="negative" label="Register" />
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

import loginStore from 'stores/login';

export default defineComponent({
  name: 'LoginBox',

  setup() {
    const { isLogin, userInfo, login } = loginStore();

    return {
      isLogin,
      userInfo,
      login,
    };
  },
});
</script>
