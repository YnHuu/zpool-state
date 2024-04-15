import { createRouter, createWebHashHistory } from 'vue-router'
import { useUserStore } from '@/stores'
import Layout from '@/layout'

import NProgress from 'nprogress'
import 'nprogress/nprogress.css'
NProgress.configure({ showSpinner: false })


export const routes = [
  {
    path: '/',
    redirect: '/zfs',
    hidden: true
  },
  {
    path: '/login',
    component: () => import('@/views/login'),
    hidden: true
  },
  {
    path: '/zfs',
    component: Layout,
    redirect: '/zfs/zpool',
    meta: { title: 'ZFS信息', icon: 'zfs', single: true },
    children: [
      {
        path: 'zpool',
        component: () => import('@/views/zpool'),
      }
    ]
  },
  {
    path: '/disk',
    component: Layout,
    redirect: '/disk/info',
    meta: { title: '硬盘信息', icon: 'disk', single: true },
    children: [
      {
        path: 'info',
        meta: { keepalive: true },
        component: () => import('@/views/disk'),
      }
    ]
  },
  {
    path: '/cron',
    component: Layout,
    redirect: '/cron/crontab',
    meta: { title: '定时任务', icon: 'cron', single: true },
    children: [
      {
        path: 'crontab',
        component: () => import('@/views/crontab'),
      }
    ]
  },
  {
    path: '/wechat',
    component: Layout,
    redirect: '/wechat/msg',
    meta: { title: '微信推送', icon: 'wechat', single: true },
    children: [
      {
        path: 'msg',
        component: () => import('@/views/wechat'),
      }
    ]
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

router.beforeEach((to, _from, next) => {
  NProgress.start()
  if (to.path == '/login' || useUserStore().token != '') {
    next()
  } else {
    next("/login");
  }
})

router.afterEach(() => {
  NProgress.done()
})

export default router
