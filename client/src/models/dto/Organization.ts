import type { RegisterRequestDto } from '@/models/dto/Auth'

export type OrganizationResponseDto = Omit<RegisterRequestDto, 'password'> & Record<'id', number>
