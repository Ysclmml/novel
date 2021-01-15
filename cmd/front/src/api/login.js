import instance from '../utils/request.js'
// 登录
function userInfo (data) {
  return instance({
    url: '/v1/book/listBookSetting',
    method: 'POST',
    data
  })
}
// 获取短信验证码
function getCode (params) {
  return instance({
    url: 'v1_0/sms/codes',
    method: 'GET',
    params
  })
}
export { userInfo, getCode }
