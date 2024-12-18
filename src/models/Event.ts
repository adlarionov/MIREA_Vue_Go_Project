import type { FormError } from './Form'

export interface Event {
  name: string
  description: string
}

export type NewEvent = Event

export interface UpdatedEvent {
  event_id: string
}

export type NewEventError = {
  name?: FormError[]
  description?: FormError[]
}
