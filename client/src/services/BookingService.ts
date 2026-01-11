import API from '@/api/instance'
import URLs from '@/api/URLs'
import type { BookingResponseDto, BookingRequestDto } from '@/models/dto/Booking'

export class BookingService {
  static async getBookings() {
    const response = await API.get<BookingResponseDto[]>(URLs.getBookings)

    return response.data
  }

  static async getBookingById(bookingId: number) {
    const response = await API.get<BookingResponseDto>(URLs.getBookingById(bookingId))

    return response.data
  }

  static async createBooking(data: BookingRequestDto) {
    const response = await API.post(URLs.getBookings, data)

    return response.data
  }
}
