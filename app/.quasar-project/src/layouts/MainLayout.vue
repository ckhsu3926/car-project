<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="toggleLeftDrawer"
        />

        <q-toolbar-title> Car Project </q-toolbar-title>

        <div>v{{ appVersion }}</div>
      </q-toolbar>
    </q-header>

    <q-drawer v-model="leftDrawerOpen" show-if-above bordered>
      <div class="fit column wrap justify-between">
        <q-list>
          <q-item-label header> Internal Links </q-item-label>

          <EssentialLink
            v-for="link in essentialLinks"
            :key="link.title"
            v-bind="link"
          />
        </q-list>
        <q-btn flat color="secondary" @click="logout">logout</q-btn>
      </div>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>

    <LoginBox />
  </q-layout>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import packageInfo from 'app/package.json';
import EssentialLink from 'components/EssentialLink.vue';
import LoginBox from 'components/fixed/LoginBox.vue';
import loginStore from 'stores/login';

const linksList = [
  {
    title: 'User',
    caption: 'User Info',
    icon: 'person',
    link: '/user',
  },
  {
    title: 'Gas Station',
    caption: 'Gas Station List in common use',
    icon: 'storefront',
    link: '/store',
  },
  {
    title: 'Vehicle List',
    caption: '',
    icon: 'directions_car',
    link: '/vehicle',
  },
  {
    title: 'Refueling Record',
    caption: '',
    icon: 'local_gas_station',
    link: '/refuel',
  },
  {
    title: 'Maintenance Record',
    caption: '',
    icon: 'build',
    link: '/maintain',
  },
];

export default defineComponent({
  name: 'MainLayout',

  components: {
    EssentialLink,
    LoginBox,
  },

  setup() {
    const leftDrawerOpen = ref(false);

    const { logout } = loginStore();

    return {
      appVersion: packageInfo.version,
      essentialLinks: linksList,
      leftDrawerOpen,
      toggleLeftDrawer() {
        leftDrawerOpen.value = !leftDrawerOpen.value;
      },

      logout,
    };
  },
});
</script>
