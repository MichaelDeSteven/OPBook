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
      path: '/read',
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
    }
  ]
})
