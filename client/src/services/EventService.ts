import API from '@/api/instance'
import URLs from '@/api/URLs'
import type { EventRequestDto, EventResponseDto } from '@/models/dto/Event'

export class EventService {
  static async getEvents() {
    const response = await API.get<EventResponseDto[]>(URLs.getEvents)

    return response.data
  }

  static async getEventById(eventId: number) {
    const response = await API.get<EventResponseDto>(URLs.getEventById(eventId))

    return response.data
  }

  static async createEvent(data: EventRequestDto) {
    const response = await API.post(URLs.getVenues, data)

    return response.data
  }
}
