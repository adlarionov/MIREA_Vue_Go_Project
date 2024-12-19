import type { NewFile } from '@/models/Attachment'
import API from './instance'
import URLs from './URLs'

export default class FilesApi {
  static async loadFile(file: Blob): Promise<NewFile> {
    const response = await API.put<NewFile>(URLs.loadFile, file, {
      headers: {
        'Content-Type': 'application/octet-stream',
      },
    })

    return response.data
  }
}
