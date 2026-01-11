export const API_CONTEXT = import.meta.env.VITE_APP_API

const URLs = {
  // auth
  login: '/auth/login',
  register: '/auth/register',
  // requests
  getBookings: '/bookings',
  getBookingById: (bookingId: number) => `/bookings/${bookingId}`,
  // events
  getEvents: '/events',
  getEventById: (eventId: number) => {
    return `/events/${eventId}`
  },
  // Venues
  getVenues: '/venues',
  getVenueById: (venueId: number) => `/venues/${venueId}`,
  addImageToVenueById: (venueId: number) => `/venues/${venueId}/images`,
}

export default URLs
