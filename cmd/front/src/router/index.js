import Vue from 'vue'
import VueRouter from 'vue-router'
import Index from '../views/index/Index'
import BookDetail from '../views/bookdetail/Index'
import BookComment from '../views/comment/Index'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Index',
    component: Index
  },
  {
    path: '/book/:bookId',
    name: 'BookDetail',
    component: BookDetail,
  },
  {
    path: '/book/comment/:bookId',
    name: 'BookComment',
    component: BookComment,
  },
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
