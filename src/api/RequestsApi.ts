import type { NewRequest, NewRequestComment, Request, RequestDto } from '@/models/Request'
import API from './instance'
import URLs from './URLs'

export default class RequestsApi {
  static async createRequest(body: NewRequest): Promise<string> {
    const response = await API.post<{ request_id: string }>(URLs.createRequest, {
      body,
    })

    return response.data.request_id
  }

  static async getRequests(
    request_id?: string,
    event_id?: string,
    author_login?: string,
    created_after?: string,
  ): Promise<Request[]> {
    const response = await API.get<{ requests: Request[] }>(
      URLs.getRequests(request_id, event_id, author_login, created_after),
    )

    return response.data.requests
  }

  static async getRequestById(request_id: string): Promise<RequestDto> {
    const response = await API.get<RequestDto>(URLs.getRequestById(request_id))

    return response.data
  }

  static async addCommentToRequest(request_id: string, body: NewRequestComment): Promise<unknown> {
    const response = await API.post<unknown>(URLs.addCommentToRequest(request_id), {
      body,
    })

    return response.data
  }
}
