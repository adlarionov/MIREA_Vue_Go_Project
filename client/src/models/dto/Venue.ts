import type { ApiModelDto } from '@/models/Api'
import type { ImageResponseDto } from '@/models/dto/Image'

export interface VenueResponseDto extends ApiModelDto {
  name: string
  description: string
  capacity: number
  address: string
  Image: ImageResponseDto
}

export type VenueRequestDto = Record<
  'venue',
  Pick<VenueResponseDto, 'address' | 'capacity' | 'description' | 'name'>
>

export interface VenueTableItem
  extends Pick<VenueResponseDto, 'address' | 'id' | 'description' | 'name' | 'capacity'> {
  imageUrl: ImageResponseDto['image_url']
}
