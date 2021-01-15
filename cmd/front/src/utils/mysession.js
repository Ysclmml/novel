// 存数据
function setSession (key, value) {
  window.sessionStorage.setItem(key, JSON.stringify(value))
}
// 取数据
function getSession (key) {
  return JSON.parse(window.sessionStorage.getItem(key))
}
// 删数据
function delSession (key) {
  window.sessionStorage.removeItem(key)
}
export { setSession, getSession, delSession }
