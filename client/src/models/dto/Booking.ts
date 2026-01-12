import type { EventResponseDto } from '@/models/dto/Event'
import type { ApiModelDto } from '@/models/Api'
import type { VenueResponseDto } from '@/models/dto/Venue'

export interface BookingResponseDto extends ApiModelDto {
  price: number
  status: string
  Event: EventResponseDto
  Venue: VenueResponseDto
}

export interface BookingRequestDto extends Omit<BookingResponseDto, 'event' | 'venue'> {
  eventId: number
  venueId: number
}

export interface BookingTableItem extends Pick<BookingResponseDto, 'id' | 'price' | 'status'> {
  eventName: EventResponseDto['name']
  eventDescription: EventResponseDto['description']
  venueName: VenueResponseDto['name']
}
