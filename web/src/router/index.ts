import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Layout from '@/layout/index.vue'
import { h } from 'vue'
import { createDiscreteApi, NIcon } from 'naive-ui'
import { UserCertification } from '@vicons/carbon'
import { ApiApp } from '@vicons/tabler'
import { PersonOutline } from '@vicons/ionicons5'
import { useUserStore } from '@/store'

const { loadingBar } = createDiscreteApi(['loadingBar'])

// 菜单配置
export const menuList = [
    {
        path: '/user',
        name: 'user',
        meta: {
            title: '用户管理',
            icon: () => h(NIcon, null, { default: () => h(UserCertification) })
        },
        component: () => import('@/views/user/index.vue')
    },
    {
        path: '/application',
        name: 'application',
        meta: {
            title: '应用管理',
            icon: () => h(NIcon, null, { default: () => h(ApiApp) })
        },
        component: () => import('@/views/application/index.vue')
    },
    {
        path: '/my',
        name: 'my',
        meta: {
            title: '个人中心',
            icon: () => h(NIcon, null, { default: () => h(PersonOutline) })
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
        redirect: '/user',
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

router.beforeEach((to, _, next) => {
    document.title = `${to.meta.title} - ${import.meta.env.VITE_APP_TITLE}`
    loadingBar.start()
    const token = useUserStore().userInfo.token
    // 当用户未登录时重定向到登录页面
    if (!token) {
        if (to.path !== '/login') {
            next('/login')
        } else {
            next()
        }
    }
    // 当用户已登录时重定向到主页
    else {
        if (to.path === '/login') {
            next('/user')
        } else {
            next()
        }
    }
})

router.afterEach(() => {
    loadingBar.finish()
})

export default router