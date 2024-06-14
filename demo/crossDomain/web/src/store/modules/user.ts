import { defineStore } from 'pinia'
import { deleteAllCookies } from '@/utils'

interface Application {
    id: number;
    name: string;
    url: string;
}

interface User {
    id?: number;
    name: string;
    avatar: string;
    phone: string;
    email: string;
    apps: Application[];
}

interface STATE {
    userInfo: User;
}

export const useUserStore = defineStore('userStore', {
    state: (): STATE => {
        return {
            userInfo: {
                id: 0,
                name: '默认昵称',
                avatar: 'http://ai-game-hk.oss-cn-hongkong.aliyuncs.com/common/hao.jiangg/Datasets/Object3D/Avatar/1.webp',
                phone: '',
                email: '',
                apps: []
            }
        }
    },
    actions: {
        setUserInfo(userInfo: User) {
            this.userInfo = {
                ...this.userInfo,
                ...userInfo
            }
        },
        clearUserInfo() {
            this.userInfo = {
                id: 0,
                name: '默认昵称',
                avatar: 'http://ai-game-hk.oss-cn-hongkong.aliyuncs.com/common/hao.jiangg/Datasets/Object3D/Avatar/1.webp',
                phone: '',
                email: '',
                apps: []
            }
            deleteAllCookies()
        }
    },
    persist: {
        key: 'userStore',
        storage: localStorage,
        paths: ['userInfo']
    }
})

export default useUserStore