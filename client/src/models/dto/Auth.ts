export interface User {
  id: number
  name: string
  login: string
  phone: string
}

export interface LoginRequestDto {
  email: string
  password: string
}

export interface RegisterRequestDto extends LoginRequestDto {
  name: string
  phone: number
}

export interface LoginResponseDto {
  token: string
}
