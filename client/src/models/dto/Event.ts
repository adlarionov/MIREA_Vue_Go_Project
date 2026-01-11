import type { ApiModelDto } from '@/models/Api'
import type { OrganizationResponseDto } from '@/models/dto/Organization'

export interface EventResponseDto extends ApiModelDto {
  name: string
  description: string
  Organization: OrganizationResponseDto
}

export type EventRequestDto = Pick<EventResponseDto, 'name' | 'description'>
