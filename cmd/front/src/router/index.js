import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

// 暂时不拆分路由..., 延迟加载组件, 让组件内的样式能够局部导入生效.
// 下面写路由需要保证静态路由放上面, 带路径参数的放下面, 防止被提前匹配了.
const routes = [
  {
    path: '/',
    name: 'Index',
    component: () => import('@/views/index/Index')
  },
  {
    path: '/book/bookRanking',
    name: 'BookRanking',
    component: () => import('@/views/book_rank/Index')
  },
  {
    path: '/book/bookClass',
    name: 'BookClass',
    component: () => import('@/views/book_class/Index')
  },
  {
    path: '/book/comment/:bookId',
    name: 'BookComment',
    component: () => import('@/views/comment/Index')
  },
  {
    path: '/book/indexList/:bookId',
    name: 'BookIndexList',
    component: () => import('@/views/bookindexes/Index')
  },
  {
    path: '/book/:bookId',
    name: 'BookDetail',
    component: () => import('@/views/bookdetail/Index')
  },
  {
    path: '/book/:bookId/:indexId',
    name: 'BookContent',
    component: () => import('@/views/bookcontent/Index')
  }
]

const router = new VueRouter({
  routes,
  scrollBehavior (to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { x: 0, y: 0 }
    }
  }
})

export default router
