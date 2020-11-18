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
    component: () => import(/* webpackChunkName: "userEpgView" */ '../views/UserEPGView.vue'),
  },
  {
    path: '/admin',
    name: 'adminEPGEditView',
    component: () =>
      import(/* webpackChunkName: "adminEPGEditView" */ '../views/AdminEPGEditView.vue'),
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
})

export default router
