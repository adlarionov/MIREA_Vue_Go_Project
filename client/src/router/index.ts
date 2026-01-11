import AuthView from '@/views/Auth/AuthView.vue'
import AllBookingView from '@/views/Bookings/AllBookingView.vue'
import BookingView from '@/views/Bookings/BookingView.vue'
import NewBookingView from '@/views/Bookings/NewBookingView.vue'
import NewEventView from '@/views/Events/NewEventView.vue'
import MainView from '@/views/Main/MainView.vue'
import NotFoundView from '@/views/NotFound/NotFoundView.vue'
import RegisterView from '@/views/Register/RegisterView.vue'
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/:pathMatch(.*)*', name: 'NotFound', component: NotFoundView },
    {
      path: '/',
      name: 'main-page',
      component: MainView,
      redirect: '/bookings',

      children: [
        {
          path: '/bookings',
          name: 'all-booking',
          component: AllBookingView,
        },
        {
          path: '/bookings/:bookingId',
          name: 'booking',
          component: BookingView,
        },
        {
          path: '/bookings/new',
          name: 'booking-new',
          component: NewBookingView,
        },
        {
          path: '/events/new',
          name: 'event-new',
          component: NewEventView,
        },
      ],
    },
    {
      path: '/auth',
      name: 'auth',
      component: AuthView,
    },
    {
      path: '/register',
      name: 'register',
      component: RegisterView,
    },
  ],
})

export default router
