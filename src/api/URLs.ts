export const API_CONTEXT = import.meta.env.VITE_APP_API
export const YA_OUTH_CONTEXT = 'https://oauth.yandex.ru'

const URLs = {
  // requests
  createRequest: '/v3/requests',
  getRequests: '/v3/requests',
  getRequestById: (req_id: string) => `/v3/requests/${req_id}`,
  addCommentToRequest: (req_id: string) => `/v3/requests/${req_id}/comments`,
  // events
  createEvent: '/v3/events',
  getEvents: '/v3/events',
  getEventById: (event_id: string) => `/v3/events/${event_id}`,
  deleteEventById: (event_id: string) => `/v3/events/${event_id}`,
  changeEventById: (event_id: string) => `/v3/events/${event_id}`,
  // files
  loadFile: '/v3/files',
  // roles
  changeRole: '/v3/roles',
}

export default URLs
