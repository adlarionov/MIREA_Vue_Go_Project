import API from '@/api/instance'
import type { RegisterRequestDto, LoginRequestDto, LoginResponseDto } from '@/models/dto/Auth'

export default class AuthApi {
  static async login(data: LoginRequestDto) {
    API.post<LoginResponseDto, LoginResponseDto, LoginRequestDto>('/auth/login', data)
  }

  static async register(data: RegisterRequestDto) {
    API.post<string, string, RegisterRequestDto>('/auth/register', data)
  }
}
