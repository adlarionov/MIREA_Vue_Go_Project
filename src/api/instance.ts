import axios from 'axios'
import { API_CONTEXT } from './URLs'

const API = axios.create(API_CONTEXT)

export default API
