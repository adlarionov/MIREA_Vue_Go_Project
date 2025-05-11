import API from './instance'
import URLs from './URLs'
import type { NewRole } from '@/models/dto/Role'

export default class RolesApi {
  static async updateRole(body: NewRole): Promise<unknown> {
    const response = await API.put<unknown>(URLs.updateRole, { ...body })

    return response.data
  }
}
