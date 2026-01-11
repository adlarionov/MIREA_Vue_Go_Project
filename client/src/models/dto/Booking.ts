import type { EventResponseDto } from '@/models/dto/Event'
import type { ApiModelDto } from '@/models/Api'
import type { VenueResponseDto } from '@/models/dto/Venue'

export interface BookingResponseDto extends ApiModelDto {
  price: number
  status: string
  event: EventResponseDto
  venue: VenueResponseDto
}

export interface BookingRequestDto extends Omit<BookingResponseDto, 'event' | 'venue'> {
  event_id: number
  venue_id: number
}
