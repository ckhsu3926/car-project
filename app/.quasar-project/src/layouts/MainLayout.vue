<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn flat dense round icon="menu" aria-label="Menu" @click="toggleLeftDrawer" />

        <q-toolbar-title>
          <q-btn flat to="/">Car Project</q-btn>
        </q-toolbar-title>

        <div>v{{ appVersion }}</div>
      </q-toolbar>
    </q-header>

    <q-drawer v-model="leftDrawerOpen" show-if-above bordered>
      <div class="fit column wrap justify-between">
        <q-list>
          <q-item-label header> Internal Links </q-item-label>

          <EssentialLink v-for="link in essentialLinks" :key="link.title" v-bind="link" />
        </q-list>
        <q-btn flat color="secondary" @click="logout">logout</q-btn>
      </div>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>

    <FixedBox />
  </q-layout>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import packageInfo from 'app/package.json';
import EssentialLink from 'components/EssentialLink.vue';
import FixedBox from 'layouts/FixedBox.vue';
import loginStore from 'stores/login';

const linksList = [
  {
    title: 'User',
    caption: 'User Info',
    icon: 'person',
    to: '/user',
    disabled: true,
  },
  {
    title: 'Gas Station',
    caption: 'Gas Station List in common use',
    icon: 'storefront',
    to: '/store',
  },
  {
    title: 'Vehicle List',
    caption: '',
    icon: 'directions_car',
    to: '/vehicle',
  },
];

export default defineComponent({
  name: 'MainLayout',

  components: {
    EssentialLink,
    FixedBox,
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
