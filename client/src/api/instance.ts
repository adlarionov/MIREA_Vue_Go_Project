import axios from 'axios'
import { API_CONTEXT } from './URLs'

const API = axios.create({
  baseURL: API_CONTEXT,
  headers: {
    'Content-Type': 'application/json',
  },
})

API.interceptors.request.use((config) => {
  return config
})

export default API
