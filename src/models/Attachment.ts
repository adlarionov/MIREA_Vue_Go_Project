export default interface Attachment {
  id: string
  filename: string
}

export interface AttachmentRequestDto {
  download_url: string
  filename: string
}

export type FileType = File & {
  objectURL: string
}
