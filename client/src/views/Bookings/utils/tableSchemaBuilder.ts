import type { BookingResponseDto, BookingTableItem } from '@/models/dto/Booking'
import type { EventResponseDto, EventTableItem } from '@/models/dto/Event'
import type { VenueResponseDto, VenueTableItem } from '@/models/dto/Venue'

export const flatBookingSchema = (bookings: BookingResponseDto[]): BookingTableItem[] => {
  return bookings.reduce<BookingTableItem[]>((acc, { id, Event, price, status, Venue }) => {
    acc.push({
      id,
      price,
      status,
      eventDescription: Event.description,
      eventName: Event.name,
      venueName: Venue.name,
    })

    return acc
  }, [])
}

export const flatEventSchema = (events: EventResponseDto[]): EventTableItem[] => {
  return events.reduce<EventTableItem[]>((acc, { id, Organization, description, name }) => {
    acc.push({
      id,
      description,
      name,
      organizationName: Organization.name,
      organizationPhone: Organization.phone,
    })

    return acc
  }, [])
}

export const flatVenueSchema = (venues: VenueResponseDto[]): VenueTableItem[] => {
  return venues.reduce<VenueTableItem[]>(
    (acc, { id, description, name, Image, address, capacity }) => {
      acc.push({
        id,
        description,
        name,
        address,
        capacity,
        imageUrl: Image.image_url,
      })

      return acc
    },
    [],
  )
}
