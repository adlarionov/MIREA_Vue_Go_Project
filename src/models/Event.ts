import type { FormError } from './Form'

export interface Event {
  id: string
  name: string
  description: string
}

export type NewEvent = Omit<Event, 'event_id'>

export interface UpdatedEvent {
  event_id: string
}

export type NewEventError = {
  name?: FormError[]
  description?: FormError[]
}
