import axios from 'axios'
import { createDiscreteApi } from 'naive-ui'
import { useUserStore } from '@/store'

const { dialog, message } = createDiscreteApi(['dialog', 'message'])

const service = axios.create({
	baseURL: '/api/v1',
	timeout: 30000,
	headers: { 'Content-Type': 'application/json;charset=utf-8' }
})

service.interceptors.request.use(
	(config) => {
		config.headers['Authorization'] = `Bearer ${useUserStore().userInfo.token}`
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
		return response.data
	},
	(error) => {
		if (error.response?.status === 401) {
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
			return
		}
		message.error(error.response?.data?.message || error.response?.data || error.message)
		return Promise.reject(error);
	}
)

export default service