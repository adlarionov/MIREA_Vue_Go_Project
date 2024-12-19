import type Attachment from './Attachment'
import type { AttachmentRequestDto } from './Attachment'
import type { FormError } from './Form'
import type { User } from './User'

export interface Request {
  id: string
  author: User
  event?: string // TODO: check
  created_at: string
}

export interface RequestDto {
  event_id: string
  attachments: AttachmentRequestDto[]
  user: User
  description: string
  status: string
  comments: RequestComment[]
  created_at: string
  updated_at: string
}

export interface NewRequestComment {
  content: string
}

export interface RequestComment {
  content: string
  author: User
  created_id: string
}

export interface NewRequest {
  event_id: string
  attachments: Attachment[]
  description?: string
  status?: string
}

export type NewRequestErrors = {
  event_id?: FormError[]
  attachments?: FormError[]
}
