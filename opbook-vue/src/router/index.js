import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import User from '@/view/user/user'
import UserFans from '@/components/user/fans'
import Read from '@/view/document/default_read'
import Email from '@/view/login/email'
import Login from '@/view/login/login'
import Setting from '@/view/user/setting'
import Password from '@/view/user/password'
import BookIndex from '@/view/book/index'
import Intro from '@/view/book/intro'
import Search from '@/view/search/search'
import SearchResult from '@/view/search/result'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '*',
      name: 'HelloWorld',
      component: HelloWorld
    },
    {
      path: '/user/:uid',
      name: 'User',
      component: User
    },
    {
      path: '/fans',
      name: 'UserFans',
      component: UserFans
    },
    {
      path: '/read/:book_identify/:doc_identify',
      name: 'Read',
      component: Read
    },
    {
      path: '/reg',
      name: 'Email',
      component: Email
    },
    {
      path: '/login',
      name: 'Login',
      component: Login
    },
    {
      path: '/setting',
      name: 'Setting',
      component: Setting
    },
    {
        path: '/password',
        name: 'Password',
        component: Password
    },
    {
        path: '/book',
        name: 'BookIndex',
        component: BookIndex
    },
    {
        path: '/introduct/:identify',
        name: 'Intro',
        component: Intro
    },
    {
        path: '/search',
        name: 'Search',
        component: Search
    },
    {
        path: '/search/result',
        name: 'SearchResult',
        component: SearchResult
    }
  ]
})
