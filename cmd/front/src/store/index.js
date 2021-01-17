import Vue from 'vue'
import Vuex from 'vuex'
import myLocal from "@/utils/mylocal";


Vue.use(Vuex)
//
export default new Vuex.Store({
  state: {
    userInfo: myLocal.getLocal('userInfo') || {}
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
