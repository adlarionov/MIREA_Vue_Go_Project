import { API_CONTEXT } from '@/api/URLs'

export const formImageUrl = (image_url: string): string => `${API_CONTEXT}${image_url.slice(1)}`
