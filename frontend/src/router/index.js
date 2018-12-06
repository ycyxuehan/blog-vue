import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/views/home/Home'
const Article = ()=> import ('@/views/article/Article')
// import Admin from '@/views/admin/Admin'
const Login = ()=> import ('@/views/login/Login')
const Admin = ()=>import('@/views/admin/Admin')

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/article',
      name: 'Article',
      component: Article
    },
    {
      path: '/admin',
      name: 'Admin',
      component: Admin
    },
    {
      path: '/login',
      name: 'Login',
      component: Login
    },
  ]
})
