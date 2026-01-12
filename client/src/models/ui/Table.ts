import type { BookingTableItem } from '@/models/dto/Booking'
import type { EventTableItem } from '@/models/dto/Event'
import type { VenueTableItem } from '@/models/dto/Venue'

export interface ColumnType<T extends string> {
  field: T
  header: string
  type: 'rich_text' | 'text' | 'link'
}

export type BookingColumnType = ColumnType<keyof BookingTableItem>
export type EventColumnType = ColumnType<keyof EventTableItem>
export type VenueColumnType = ColumnType<keyof VenueTableItem>

export type ColumnTypeMap = BookingColumnType | EventColumnType | VenueColumnType
export type TableDataMap = BookingTableItem | EventTableItem | VenueTableItem

export const bookingTableSchema: BookingColumnType[] = [
  { field: 'id', header: 'ID', type: 'text' },
  { field: 'price', header: 'Цена', type: 'text' },
  { field: 'status', header: 'Статус', type: 'text' },
  { field: 'eventName', header: 'Событие', type: 'text' },
  { field: 'venueName', header: 'Локация', type: 'text' },
]

export const eventTableSchema: EventColumnType[] = [
  { field: 'id', header: 'ID', type: 'text' },
  { field: 'name', header: 'Название', type: 'text' },
  { field: 'description', header: 'Описание', type: 'rich_text' },
  { field: 'organizationName', header: 'Организация', type: 'text' },
  { field: 'organizationPhone', header: 'Телефон', type: 'text' },
]

export const venueTableSchema: VenueColumnType[] = [
  { field: 'id', header: 'ID', type: 'text' },
  { field: 'name', header: 'Название', type: 'text' },
  { field: 'description', header: 'Описание', type: 'rich_text' },
  { field: 'address', header: 'Адрес', type: 'link' },
  { field: 'capacity', header: 'Вместимость', type: 'text' },
]
