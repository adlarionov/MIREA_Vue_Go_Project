export interface ApiModelDto {
  id: number
  updated_at: string
  created_at: string
}

export interface ApiError {
  errorCode?: string
  errorStatus?: number
  message?: string
}
