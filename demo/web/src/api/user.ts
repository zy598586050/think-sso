import request from '@/utils/request'

// 用户信息
export const UserInfo = () => {
    return request({
        url: '/user/info',
        method: 'POST'
    })
}

// 登录
export const Login = (data: Object) => {
    return request({
        url: '/login/code',
        method: 'POST',
        data
    })
}

// 退出登录
export const Logout = () => {
    return request({
        url: '/logout',
        method: 'POST'
    })
}