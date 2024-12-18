import axios from 'axios'
import { API_CONTEXT } from './URLs'
import AuthApi from './AuthApi'

const API = axios.create({
  baseURL: API_CONTEXT,
  headers: {
    'Content-Type': 'application/json',
  },
})

API.interceptors.request.use((config) => {
  config.headers['X-Token'] = AuthApi.token
  return config
})

export default API
