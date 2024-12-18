import RolesApi from '@/api/RolesApi'
import type { NewRole } from '@/models/Role'
import { AxiosError } from 'axios'

export default class RolesController {
  static async updateRole(body: NewRole): Promise<unknown> {
    try {
      return RolesApi.updateRole(body)
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
