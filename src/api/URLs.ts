export const API_CONTEXT = import.meta.env.VITE_APP_API
export const YA_OUTH_CONTEXT = 'https://oauth.yandex.ru'

const URLs = {
  // requests
  createRequest: '/v3/requests',
  getRequests: (
    request_id?: string,
    event_id?: string,
    author_login?: string,
    created_after?: string,
  ) => {
    let url = '/v3/requests?'
    if (request_id) url += `request_id=${request_id}&`
    if (event_id) url += `event_id=${event_id}&`
    if (author_login) url += `&author_login=${author_login}&`
    if (created_after) url += `&created_after=${created_after}&`

    return url
  },
  getRequestById: (req_id: string) => `/v3/requests/${req_id}`,
  addCommentToRequest: (req_id: string) => `/v3/requests/${req_id}/comments`,
  // events
  createEvent: '/v3/events',
  getEvents: (name_part?: string) => {
    return name_part ? `/v3/events?name_part=${name_part}` : `/v3/events`
  },
  getEventById: (event_id: string) => `/v3/events/${event_id}`,
  deleteEventById: (event_id: string) => `/v3/events/${event_id}`,
  updateEventById: (event_id: string) => `/v3/events/${event_id}`,
  // files
  loadFile: '/v3/files',
  // roles
  updateRole: '/v3/roles',
}

export default URLs
