import type { BookingResponseDto } from '@/models/dto/Booking'
import type { EventResponseDto } from '@/models/dto/Event'
import type { VenueResponseDto } from '@/models/dto/Venue'

export interface ColumnType<T extends string> {
  field: T
  header: string
  type: 'rich_text' | 'text' | 'link'
}

export type BookingColumnType = ColumnType<keyof BookingResponseDto>
export type EventColumnType = ColumnType<keyof EventResponseDto>
export type VenueColumnType = ColumnType<keyof VenueResponseDto>

export type ColumnTypeMap = BookingColumnType | EventColumnType | VenueColumnType
export type TableDataMap = BookingResponseDto | EventResponseDto | VenueResponseDto

export const bookingTableSchema: BookingColumnType[] = [
  { field: 'id', header: 'ID', type: 'text' },
  { field: 'price', header: 'Цена', type: 'text' },
  { field: 'status', header: 'Статус', type: 'text' },
]

export const eventTableSchema: EventColumnType[] = [
  { field: 'id', header: 'ID', type: 'text' },
  { field: 'name', header: 'Название', type: 'text' },
  { field: 'description', header: 'Описание', type: 'rich_text' },
]

export const venueTableSchema: VenueColumnType[] = [
  { field: 'id', header: 'ID', type: 'text' },
  { field: 'name', header: 'Название', type: 'text' },
  { field: 'description', header: 'Описание', type: 'rich_text' },
  { field: 'address', header: 'Адрес', type: 'link' },
  { field: 'capacity', header: 'Вместимость', type: 'text' },
]
