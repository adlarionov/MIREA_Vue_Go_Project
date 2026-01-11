import axios, { AxiosError } from 'axios'
import { API_CONTEXT } from './URLs'
import router from '@/router'

const API = axios.create({
  baseURL: API_CONTEXT,
  headers: {
    'Content-Type': 'application/json',
  },
})

API.interceptors.response.use(
  (response) => response,
  (responseError: AxiosError) => {
    console.log(responseError, 'error')
    if (responseError.status === 401 || responseError.status === 400) {
      router.push('/auth')
      API.interceptors.request.clear()
      return responseError
    }
  },
)

export const TOKEN_PROVIDE_KEY = 'token_middleware'

export const TokenMiddleware = () => {
  let tokenInterceptorId: number

  return {
    setToken(token: string) {
      tokenInterceptorId = API.interceptors.request.use((request) => {
        request.headers.setAuthorization(`Bearer ${token}`)

        return request
      })
    },
    clearToken() {
      API.interceptors.request.eject(tokenInterceptorId)
    },
  }
}

export default API
