import Vue from 'vue'
import VueRouter from 'vue-router'
import ApplicationLayout from '@/views/layout/Application'
import PublicLayout from '@/views/layout/Public'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'About',
    component: () => import(/* webpackChunkName: "about" */ '@/views/About')
  },
  {
    path: '/',
    name: 'Public',
    component: PublicLayout,
    children: [
      {
        path: 'login',
        name: 'LogIn',
        component: () => import(/* webpackChunkName: "signin" */ '@/views/Login')
      },
      {
        path: 'register',
        name: 'Register',
        component: () => import(/* webpackChunkName: "registration" */ '@/views/Registration')
      }
    ]
  },

  {
    path: '/',
    name: 'ApplicationLayout',
    component: ApplicationLayout,
    children: [
      {
        path: 'profile',
        name: 'Profile',
        component: () => import(/* webpackChunkName: "profile" */ '@/views/Profile')
      }
    ]
  },
  {
    path: '*',
    component: () => import(/* webpackChunkName: "notfound" */ '@/views/NotFound')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: '',
  routes
})

export default router