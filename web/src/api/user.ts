import request from '@/utils/request'

// 邮箱登录
export const LoginEmail = (data: Object) => {
    return request({
        url: '/login/email',
        method: 'POST',
        data,
    })
}

// 获取Code
export const GetCode = () => {
    return request({
        url: '/code',
        method: 'POST'
    })
}

// 退出登录
export const Logout = () => {
    return request({
        url: '/logout',
        method: 'POST'
    })
}

// 用户信息
export const UserInfo = () => {
    return request({
        url: '/user/info',
        method: 'POST'
    })
}

// 用户列表
export const UserList = (params: Object) => {
    return request({
        url: '/user/list',
        method: 'GET',
        params
    })
}