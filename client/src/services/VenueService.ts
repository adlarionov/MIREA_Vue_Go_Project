import API from '@/api/instance'
import URLs from '@/api/URLs'
import type { VenueRequestDto, VenueResponseDto } from '@/models/dto/Venue'
import type { AxiosRequestConfig } from 'axios'

export class VenueService {
  static async getVenues() {
    const response = await API.get<VenueResponseDto[]>(URLs.getVenues)

    return response.data
  }

  static async getVenueById(venueId: number) {
    const response = await API.get<VenueResponseDto>(URLs.getVenueById(venueId))

    return response.data
  }

  static async createVenue(data: VenueRequestDto) {
    const response = await API.post<number, AxiosRequestConfig<number>, VenueRequestDto>(
      URLs.getVenues,
      data,
    )

    return response.data
  }

  static async addImageToVenue(venueId: number, image: FormData) {
    const response = await API.post(URLs.addImageToVenueById(venueId), image, {
      headers: { 'Content-Type': 'multipart/form-data' },
    })

    return response.data
  }
}
