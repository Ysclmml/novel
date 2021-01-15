// 存数据
function setLocal (key, value) {
  // window.localStorage.setItem(key, value)
  window.localStorage.setItem(key, JSON.stringify(value))
}
// 取数据
function getLocal (key = 'token') {
  // return JSON.parse(window.localStorage.getItem(key))
  return window.localStorage.getItem(key)
}
// 删数据
function delLocal (key = 'token') {
  window.localStorage.removeItem(key)
}
// 全部清空
function clearLocal () {
  window.localStorage.clear()
}

export { setLocal, getLocal, delLocal, clearLocal }