import request from '@/utils/request'

// 测试业务接口
export const TestHome = () => {
    return request({
        url: '/test',
        method: 'GET'
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