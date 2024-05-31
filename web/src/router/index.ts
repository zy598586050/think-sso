import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Layout from '@/layout/index.vue'
import { h } from 'vue'
import { createDiscreteApi, NIcon } from 'naive-ui'
import { ChatbubbleEllipsesOutline } from '@vicons/ionicons5'
import { useUserStore } from '@/store'

const { loadingBar } = createDiscreteApi(['loadingBar'])

// 菜单配置
export const menuList = [
    {
        path: '/generate',
        name: 'generate',
        meta: {
            title: '生成',
            icon: () => h(NIcon, null, { default: () => h(ChatbubbleEllipsesOutline) })
        },
        component: () => import('@/views/index/index.vue')
    },
    {
        path: '/discover',
        name: 'discover',
        meta: {
            title: '发现',
            icon: () => h(NIcon, null, { default: () => h(ChatbubbleEllipsesOutline) })
        },
        component: () => import('@/views/discover/index.vue')
    },
    {
        path: '/my',
        name: 'my',
        meta: {
            title: '我的',
            icon: () => h(NIcon, null, { default: () => h(ChatbubbleEllipsesOutline) })
        },
        component: () => import('@/views/my/index.vue')
    }
]

// 固定路由
const routes: RouteRecordRaw[] = [
    {
        path: '/',
        name: 'layout',
        component: Layout,
        redirect: '/generate',
        children: menuList
    },
    {
        path: '/login',
        name: 'login',
        meta: {
            title: '登录'
        },
        component: () => import('@/views/login/index.vue')
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
    history: createWebHistory(),
    routes
})

router.beforeEach(async (to, _, next) => {
    document.title = `${to.meta.title} - sky3DGen`
    loadingBar.start()
    const token = useUserStore().userInfo.token
    // 当用户未登录时重定向到登录页面
    if (!token) {
        // if (to.path !== '/login') {
        //     next('/login')
        // } else {
        //     next()
        // }
        next()
    }
    // 当用户已登录时重定向到主页
    else {
        if (to.path === '/login') {
            next('/generate')
        } else {
            next()
        }
    }
})

router.afterEach(() => {
    loadingBar.finish()
})

export default router