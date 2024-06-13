import { computed } from 'vue'
import { darkTheme, lightTheme, GlobalThemeOverrides } from 'naive-ui'
import { useAppStore } from '@/store'

export const useTheme = () => {
    const appStore = useAppStore()

    const isDark = computed(() => appStore.theme === 'dark')

    const theme = computed(() => isDark.value ? darkTheme : lightTheme)

    const themeOverrides = computed<GlobalThemeOverrides>(() => {
        if (isDark.value) {
            return {
                common: {
                    primaryColor: '#63E2B7FF',
                    primaryColorHover: '#7FE7C4FF',
                    primaryColorPressed: '#5ACEA7FF',
                    primaryColorSuppl: 'rgb(42, 148, 125)'
                },
                Switch: {
                    railColorActive: '#2A947DFF',
                    loadingColor: '#2A947DFF'
                }
            }
        } else {
            return {
                common: {
                    primaryColor: '#18A058FF',
                    primaryColorHover: '#36AD6AFF',
                    primaryColorPressed: '#0C7A43FF',
                    primaryColorSuppl: '#36AD6AFF'
                },
                Switch: {
                    railColorActive: '#18A058FF',
                    loadingColor: '#18A058FF'
                }
            }
        }
    })

    return { isDark, theme, themeOverrides }
}