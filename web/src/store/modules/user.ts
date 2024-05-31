import { defineStore } from 'pinia'

interface User {
    id?: number;
    name: string;
    avatar: string;
    token?: string;
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
                token: ''
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
                token: ''
            }
        }
    },
    persist: {
        key: 'userStore',
        storage: localStorage,
        paths: ['userInfo']
    }
})

export default useUserStore