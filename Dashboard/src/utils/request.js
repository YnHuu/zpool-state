import axios from 'axios'
import { useUserStore } from '@/stores'
import Alert from '@/components/Alert'

export const request = axios.create({
  baseURL: import.meta.env.VITE_BASE_URL,
  timeout: 10000
})

request.interceptors.request.use(
  (config) => {
    config.headers['token'] = useUserStore().token
    return config
  },
  (error) => {
    console.log(error)
    return Promise.reject(error)
  }
)

request.interceptors.response.use(
  (response) => {
    const res = response.data
    if (res.code === 200) {
      return res
    } else {
      Error(res)
      return Promise.reject(res)
    }
  },
  (error) => {
    const badMessage = error.message || error
    const code = parseInt(badMessage.toString().replace('Error: Request failed with status code ', ''))
    Error({ code, message: badMessage })
    return Promise.reject(error)
  }
)

function Error(error) {
  if (error.code === 401) {
    useUserStore().delToken()
    location.reload()
  } else {
    console.log(error.msg || error.message)
    Alert({ message: error.msg || error.message || '服务异常', type: "error" })
  }
}

export default request