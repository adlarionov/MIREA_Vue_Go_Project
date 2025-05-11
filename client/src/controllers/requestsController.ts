import RequestsApi from '@/api/RequestsApi'
import type { NewRequest, NewRequestComment, Request, RequestDto } from '@/models/dto/Reservation'
import { AxiosError } from 'axios'

export default class RequestsController {
  static async createRequest(body: NewRequest): Promise<string> {
    try {
      return RequestsApi.createRequest(body)
    } catch (error: unknown) {
      if (error instanceof AxiosError) {
        throw {
          errorCode: error.code,
          errorStatus: error.status,
          message: error.message,
        }
      }
      throw {
        message: String(error),
      }
    }
  }

  static async getRequests(
    request_id?: string,
    event_id?: string,
    author_login?: string,
    created_after?: string,
  ): Promise<Request[]> {
    try {
      return RequestsApi.getRequests(request_id, event_id, author_login, created_after)
    } catch (error: unknown) {
      if (error instanceof AxiosError) {
        throw {
          errorCode: error.code,
          errorStatus: error.status,
          message: error.message,
        }
      }
      throw {
        message: String(error),
      }
    }
  }

  static async getRequestById(request_id: string): Promise<RequestDto> {
    try {
      return RequestsApi.getRequestById(request_id)
    } catch (error: unknown) {
      if (error instanceof AxiosError) {
        throw {
          errorCode: error.code,
          errorStatus: error.status,
          message: error.message,
        }
      }
      throw {
        message: String(error),
      }
    }
  }

  static async addCommentToRequest(request_id: string, body: NewRequestComment): Promise<unknown> {
    try {
      return RequestsApi.addCommentToRequest(request_id, body)
    } catch (error: unknown) {
      if (error instanceof AxiosError) {
        throw {
          errorCode: error.code,
          errorStatus: error.status,
          message: error.message,
        }
      }
      throw {
        message: String(error),
      }
    }
  }
}
