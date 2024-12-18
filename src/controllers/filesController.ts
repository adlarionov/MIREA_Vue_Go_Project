import FilesApi from '@/api/FilesApi'
import type { NewFile } from '@/models/Attachment'
import { AxiosError } from 'axios'

export default class FilesController {
  static async loadFile(file: Blob): Promise<NewFile> {
    try {
      return FilesApi.loadFile(file)
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
