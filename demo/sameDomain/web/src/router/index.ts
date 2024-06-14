import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router'
import Layout from '@/layout/index.vue'
import { h } from 'vue'
import { createDiscreteApi, NIcon } from 'naive-ui'
import { UserCertification } from '@vicons/carbon'
import { getCookie } from '@/utils'

const { loadingBar } = createDiscreteApi(['loadingBar'])

// 菜单配置
export const menuList = [
    {
        path: '/home',
        name: 'home',
        meta: {
            title: '首页',
            icon: () => h(NIcon, null, { default: () => h(UserCertification) })
        },
        component: () => import('@/views/home/index.vue')
    }
]

// 固定路由
const routes: RouteRecordRaw[] = [
    {
        path: '/',
        name: 'layout',
        component: Layout,
        redirect: '/home',
        children: menuList
    },
    {
        path: '/404',
        name: '404',
        component: () => import('@/views/exception/404/index.vue')
    },
    {
        path: '/500',
        name: '500',
        component: () => import('@/views/exception/500/index.vue')
    },
    {
        path: '/:pathMatch(.*)*',
        name: 'notFound',
        redirect: '/404'
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

router.beforeEach((to, _, next) => {
    document.title = `${to.meta.title} - ${import.meta.env.VITE_APP_TITLE}`
    loadingBar.start()
    const token = getCookie('think-sso-token')
    if (token) {
        next()
    } else {
        window.location.href = `${import.meta.env.VITE_SSO_URL}${window.location.href}`
    }
})

router.afterEach(() => {
    loadingBar.finish()
})

export default router