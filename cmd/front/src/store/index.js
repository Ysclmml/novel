import Vue from 'vue'
import Vuex from 'vuex'
import { getLocal } from '@/utils/mylocal.js'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    userInfo: getLocal('userInfo') || {}
  },
  getters: {
  },
  mutations: {
    setUserInfo: function (state, payload) {
      state.userInfo = payload
    }
  },
  actions: {
  },
  modules: {
  }
})
