import API from '@/api/instance'
import type { RegisterRequestDto, LoginRequestDto, LoginResponseDto } from '@/models/dto/Auth'
import type { AxiosRequestConfig } from 'axios'

export default class AuthService {
  static async login(data: LoginRequestDto) {
    const response = await API.post<
      LoginResponseDto,
      AxiosRequestConfig<LoginResponseDto>,
      LoginRequestDto
    >('/auth/login', data)

    return response.data?.token
  }

  static async register(data: RegisterRequestDto) {
    const response = await API.post<string, AxiosRequestConfig<string>, RegisterRequestDto>(
      '/auth/register',
      data,
    )

    return response.data
  }
}
