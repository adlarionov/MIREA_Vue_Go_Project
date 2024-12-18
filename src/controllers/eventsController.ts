import EventsApi from '@/api/EventsApi'
import type { Event, NewEvent, UpdatedEvent } from '@/models/Event'
import { AxiosError } from 'axios'

export default class EventsController {
  static async createEvent(body: NewEvent): Promise<unknown> {
    try {
      return EventsApi.createEvent(body)
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

  static async getEvents(name_part?: string): Promise<Event[]> {
    try {
      return EventsApi.getEvents(name_part)
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

  static async getEventById(event_id: string): Promise<Event> {
    try {
      return EventsApi.getEventById(event_id)
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

  static async deleteEventById(event_id: string): Promise<unknown> {
    try {
      return EventsApi.deleteEventById(event_id)
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

  static async updateEventById(event_id: string, body: Partial<Event>): Promise<UpdatedEvent> {
    try {
      return EventsApi.updateEventById(event_id, body)
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
