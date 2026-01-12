import AuthView from '@/views/Auth/AuthView.vue'
import AllBookingView from '@/views/Bookings/AllBookingView.vue'
import BookingView from '@/views/Bookings/BookingView.vue'
import NewBookingView from '@/views/Bookings/NewBookingView.vue'
import AllEventsView from '@/views/Events/AllEventsView.vue'
import EventView from '@/views/Events/EventView.vue'
import NewEventView from '@/views/Events/NewEventView.vue'
import MainView from '@/views/Main/MainView.vue'
import NotFoundView from '@/views/NotFound/NotFoundView.vue'
import RegisterView from '@/views/Register/RegisterView.vue'
import AllVenuesView from '@/views/Venues/AllVenuesView.vue'
import NewVenueView from '@/views/Venues/NewVenueView.vue'
import VenueView from '@/views/Venues/VenueView.vue'
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
          path: '/bookings/:id',
          name: 'booking',
          component: BookingView,
        },
        {
          path: '/bookings/new',
          name: 'booking-new',
          component: NewBookingView,
        },
        {
          path: '/events',
          name: 'all-events',
          component: AllEventsView,
        },
        {
          path: '/events/:id',
          name: 'event',
          component: EventView,
        },
        {
          path: '/events/new',
          name: 'event-new',
          component: NewEventView,
        },
        {
          path: '/venues',
          name: 'all-venues',
          component: AllVenuesView,
        },
        {
          path: '/venues/:id',
          name: 'venue',
          component: VenueView,
        },
        {
          path: '/venues/new',
          name: 'venue-new',
          component: NewVenueView,
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
