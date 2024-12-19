import type { Event, NewEvent, UpdatedEvent } from '@/models/Event'
import API from './instance'
import URLs from './URLs'

export default class EventsApi {
  static async createEvent(body: NewEvent): Promise<unknown> {
    const response = await API.post<unknown>(URLs.createEvent, { ...body })

    return response.data
  }

  static async getEvents(name_part?: string): Promise<Event[]> {
    const response = await API.get<{ events: Event[] }>(URLs.getEvents(name_part))

    return response.data.events
  }

  static async getEventById(event_id: string): Promise<Event> {
    const response = await API.get<Event>(URLs.getEventById(event_id))

    return response.data
  }

  static async deleteEventById(event_id: string): Promise<unknown> {
    const response = await API.delete<unknown>(URLs.deleteEventById(event_id))

    return response.data
  }

  static async updateEventById(event_id: string, body: Partial<Event>): Promise<UpdatedEvent> {
    const response = await API.put<UpdatedEvent>(URLs.updateEventById(event_id), {
      ...body,
    })

    return response.data
  }
}
