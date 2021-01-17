import axios from 'axios'
import myLocal from './mylocal.js'
import Message from "element-ui/packages/message/src/main";

// 创建axios副本
let instance = axios.create({
  baseURL: process.env.VUE_APP_URL,
  // 跨域携带token
  // withCredentials: true
})

// 请求拦截器
instance.interceptors.request.use(function (config) {
  // 在发送请求之前做些什么
  // 请求拦截器中，如果有token，请求中添加token
  if (myLocal.getLocal()) {
    config.headers.token = myLocal.getLocal('token')
  }
  return config
}, function (error) {
  // 对请求错误做些什么
  return Promise.reject(error)
})

// 响应拦截器
instance.interceptors.response.use(function (response) {
  // 对响应数据做点什么
  const data = response.data
  if (data.code === 200) {
    return response.data
  }
  // todo: 处理错误码
  Message.info(response.data.msg)
}, function (error) {
  // 对响应错误做点什么
  return Promise.reject(error)
})

export default instance
