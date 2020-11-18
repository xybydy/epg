import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'index',
    component: () => import(/* webpackChunkName: "index" */ '../views/Index.vue'),
  },
  {
    path: '/epg',
    name: 'userEpgView',
    component: () => import(/* webpackChunkName: "userEpgView" */ '../views/UserEpgView.vue'),
  },
  {
    path: '/admin',
    name: 'adminEpgEditView',
    component: () =>
      import(/* webpackChunkName: "adminEPGEditView" */ '../views/AdminEpgEditView.vue'),
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
})

export default router
