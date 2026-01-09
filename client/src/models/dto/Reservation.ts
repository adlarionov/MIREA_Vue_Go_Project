import type Attachment from './Attachment'
import type { AttachmentRequestDto } from './Attachment'
import type { FormError } from '../Form'
import type { User } from './Auth'

export interface Reservation {
  id: string
  author: User
  event?: string // TODO: check
  created_at: string
}

export interface ReservationDto {
  event_id: string
  attachments: AttachmentRequestDto[]
  user: User
  description: string
  status: string
  comments: ReservationComment[]
  created_at: string
  updated_at: string
}

export interface ReservationComment {
  content: string
}

export interface ExistingReservationComment extends ReservationComment {
  author: User
  created_at: string
}

export interface NewReservationDto {
  event_id: string
  attachments?: Attachment[]
  description?: string
  status?: string
}

export interface NewReservationErrors {
  event_id?: FormError[]
  attachments?: FormError[]
}
