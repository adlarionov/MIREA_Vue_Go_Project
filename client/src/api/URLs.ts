export const API_CONTEXT = import.meta.env.VITE_APP_API

const URLs = {
  // requests
  createRequest: '/requests',
  getRequests: (
    request_id?: string,
    event_id?: string,
    author_login?: string,
    created_after?: string,
  ) => {
    let url = '/events?'
    if (request_id) url += `request_id=${request_id}&`
    if (event_id) url += `event_id=${event_id}&`
    if (author_login) url += `&author_login=${author_login}&`
    if (created_after) url += `&created_after=${created_after}&`

    return url
  },
  getRequestById: (req_id: string) => `/requests/${req_id}`,
  addCommentToRequest: (req_id: string) => `/requests/${req_id}/comments`,
  // events
  createEvent: '/events',
  getEvents: (name_part?: string) => {
    return name_part ? `/events?name_part=${name_part}` : `/events`
  },
  getEventById: (event_id: string) => `/events/${event_id}`,
  deleteEventById: (event_id: string) => `/events/${event_id}`,
  updateEventById: (event_id: string) => `/events/${event_id}`,
  // files
  loadFile: '/files',
  // roles
  updateRole: '/roles',
}

export default URLs
