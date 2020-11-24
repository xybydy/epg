import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'Index',
    component: () => import(/* webpackChunkName: "index" */ '../views/Index.vue')
  },
  {
    path: '/epg/view',
    name: 'EpgUserView',
    component: () => import(/* webpackChunkName: "epgUserView" */ '../views/EpgUserView.vue')
  },
  {
    path: '/epg/update',
    name: 'EpgUpdateView',
    component: () => import(/* webpackChunkName: "epgUpdateView" */ '../views/EpgUpdateView.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
