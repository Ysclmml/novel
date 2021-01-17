class LocalStorageOption {
  // 存数据
  setLocal (key, value) {
    // window.localStorage.setItem(key, value)
    window.localStorage.setItem(key, JSON.stringify(value))
  }

  /**
   * 取数据
   * @param key {String}
   * @param {Boolean} parseJson: 是否反序列化
   * @returns {String|Object}
   */
  getLocal (key, parseJson = true) {
    return parseJson ? JSON.parse(window.localStorage.getItem(key)) : window.localStorage.getItem(key)
  }

  // 删数据
  delLocal (key) {
    window.localStorage.removeItem(key)
  }

  // 全部清空
  clearLocal () {
    window.localStorage.clear()
  }
}
const myLocal = new LocalStorageOption()
export default myLocal
