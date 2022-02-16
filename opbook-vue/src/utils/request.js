import axios from 'axios' // 引入axios

const service = axios.create({
  baseURL: '/api',
  timeout: 99999
})
// http request 拦截器
service.interceptors.request.use(
    config => {
        config.headers = {
            'Content-Type': 'application/json',
            'x-token': localStorage.getItem('token'),
            ...config.headers
        }
        return config
    },
    error => {
        console.log(error)
    }
)

service.interceptors.response.use(
    // response => {
    //     console.log(response)
    // },
    // error => {
    //     console.log(error)
    // }
)
export default service
