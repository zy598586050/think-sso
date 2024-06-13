import axios from 'axios'
import { createDiscreteApi } from 'naive-ui'
import { useUserStore } from '@/store'
import { getCookie } from '@/utils'

const { dialog, message } = createDiscreteApi(['dialog', 'message'])

const service = axios.create({
	baseURL: import.meta.env.VITE_API_URL,
	timeout: 30000,
	headers: { 'Content-Type': 'application/json;charset=utf-8' }
})

service.interceptors.request.use(
	(config) => {
		if (getCookie('think-sso-token')) {
			config.headers['Authorization'] = `Bearer ${getCookie('think-sso-token')}`
		}
		return config
	},
	(error) => {
		// 对请求错误做些什么
		return Promise.reject(error)
	}
)

service.interceptors.response.use(
	(response) => {
		// 对响应数据做点什么
		const res = response?.data
		const code = response?.data?.code
		if (code === 61) {
			dialog.warning({
				title: '提示',
				content: '登录状态已过期，请重新登录',
				positiveText: '确定',
				maskClosable: false,
				closable: false,
				closeOnEsc: false,
				onPositiveClick: () => {
					useUserStore().clearUserInfo()
					location.href = '/'
				}
			})
		} else if (code !== 0) {
			message.error(res.message)
			return Promise.reject(new Error(res.message))
		} else {
			return res
		}
	},
	(error) => {
		if (error.message.indexOf('timeout') != -1) {
			message.error('网络超时');
		} else if (error.message == 'Network Error') {
			message.error('网络连接错误');
		} else {
			if (error.response?.data) message.error(error.response.statusText);
			else message.error('接口路径找不到');
		}
		return Promise.reject(error);
	}
)

export default service