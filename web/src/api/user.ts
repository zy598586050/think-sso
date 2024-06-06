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
export const Logout = (data: Object) => {
    return request({
        url: '/logout',
        method: 'POST',
        data,
    })
}