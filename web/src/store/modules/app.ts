import { defineStore } from 'pinia'

type Theme = 'light' | 'dark'
interface STATE {
    theme: Theme;
    menuActive: number;
}

export const useAppStore = defineStore('appStore', {
    state: (): STATE => {
        return {
            theme: 'dark',
            menuActive: 1
        }
    },
    actions: {},
    persist: {
        key: 'appStore',
        storage: localStorage,
        paths: ['theme']
    }
})

export default useAppStore