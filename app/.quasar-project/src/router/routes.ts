import { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/IndexPage.vue') },
      { path: 'vehicle', component: () => import('pages/VehiclePage.vue') },
      { path: 'user', component: () => import('pages/UserPage.vue') },
      { path: 'store', component: () => import('pages/StorePage.vue') },
      { path: 'vehicle/refueling/:vehicleID', component: () => import('pages/RefuelingPage.vue') },
      { path: 'vehicle/maintenance/:vehicleID', component: () => import('pages/MaintenancePage.vue') },
    ],
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue'),
  },
];

export default routes;
