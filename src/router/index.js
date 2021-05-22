import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'Index',
    component: () => import('@/views/Index.vue'),
  },
  {
    path: '/epg/view',
    name: 'EpgUserView',
    component: () => import('@/views/EpgUserView.vue'),
  },
  {
    path: '/epg/update',
    name: 'EpgUpdateView',
    component: () => import('@/views/EpgUpdateView.vue'),
  },
  {
    path: '/fatih/tvg',
    name: 'TvgMap',
    component: () => import('@/views/TvgMap.vue'),
  },
  {
    path: '/test',
    component: () => import('@/views/La.vue'),
  },
  {
    path: '/admin',
    component: () => import('@/views/admin/Index.vue'),
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
