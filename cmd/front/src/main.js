import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import instance from './utils/request'

// * ----------------------------------------
// * 导入插件
// * ----------------------------------------
import {
  elementUiRegister,
  vueLazyloadReigster
} from '@/plugins/component-register'
import 'amfe-flexible'

// 去除body的size属性
if (document.body) {
  document.body.style.fontSize = ''
}

// 往vue原型添加方法：axios请求方法
Vue.prototype.$axios = instance
Vue.config.productionTip = false

// * ----------------------------------------
// * 注册组件
// * ----------------------------------------
elementUiRegister()
vueLazyloadReigster()

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
