import Vue from 'vue'
import VueRouter from 'vue-router'
import ApplicationLayout from '@/views/layout/Application'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    component: ApplicationLayout,
    children: [
      {
        path: '',
        name: 'About',
        component: () => import(/* webpackChunkName: "about" */ '@/views/About')
      },
      {
        path: 'profile',
        name: 'Profile',
        component: () => import(/* webpackChunkName: "profile" */ '@/views/Profile')
      }
    ]
  },
  {
    path: '/login',
    name: 'LogIn',
    component: () => import(/* webpackChunkName: "signin" */ '@/views/Login')
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import(/* webpackChunkName: "registration" */ '@/views/Registration')
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
