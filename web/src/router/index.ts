import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router'
import Layout from '@/layout/index.vue'
import { h } from 'vue'
import { createDiscreteApi, NIcon } from 'naive-ui'
import { UserCertification } from '@vicons/carbon'
import { ApiApp } from '@vicons/tabler'
import { PersonOutline } from '@vicons/ionicons5'
import { getCookie, getQueryParam } from '@/utils'
import { GetCode } from '@/api/user'

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
    history: createWebHashHistory(),
    routes
})

router.beforeEach((to, _, next) => {
    document.title = `${to.meta.title} - ${import.meta.env.VITE_APP_TITLE}`
    loadingBar.start()
    const token = getCookie('think-sso-token')
    const redirect_url = getQueryParam('redirect_url')
    if (token) {
        if (redirect_url) {
            GetCode().then((result) => {
                window.location.href = redirect_url + `${redirect_url.includes('?') ? '&' : '?'}code=${result.data.code}`
            })
        } else {
            if (to.path === '/login') {
                next('/user')
            } else {
                next()
            }
        }
    } else {
        if (to.path === '/login') {
            // 不需要校验的路由直接放行渲染
            next()
        } else {
            // 需要校验的都去登录页
            next('/login')
        }
    }
    // B 有token，正常访问 无token,看有没有code,有code执行登录，无code去授权页授权，然后回来登录得到token
})

router.afterEach(() => {
    loadingBar.finish()
})

export default router