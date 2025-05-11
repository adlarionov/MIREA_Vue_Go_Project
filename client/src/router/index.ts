import AuthView from '@/views/AuthView.vue'
import NewEventView from '@/views/Events/NewEventView.vue'
import MainView from '@/views/Main/MainView.vue'
import NewRequestView from '@/views/Requests/NewRequestView.vue'
import RequestView from '@/views/Requests/RequestView.vue'
import RoleView from '@/views/Roles/RoleView.vue'
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'main-page',
      component: MainView,
    },
    {
      path: '/requests/:request_id',
      name: 'request',
      component: RequestView,
    },
    {
      path: '/requests/new',
      name: 'request-new',
      component: NewRequestView,
    },
    {
      path: '/events/new',
      name: 'event-new',
      component: NewEventView,
    },
    {
      path: '/roles',
      name: 'role',
      component: RoleView,
    },
    {
      path: '/auth',
      name: 'auth',
      component: AuthView,
    },
  ],
})

export default router
