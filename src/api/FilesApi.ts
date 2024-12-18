import type { NewFile } from '@/models/Attachment'
import API from './instance'
import URLs from './URLs'

export default class FilesApi {
  static async loadFile(file: Blob): Promise<NewFile> {
    const response = await API.put<NewFile>(URLs.loadFile, { body: file })

    return response.data
  }
}
