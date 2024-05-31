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
                    primaryColor: '#c5f955',
                    primaryColorHover: '#c5f955',
                    primaryColorPressed: '#66afd3',
                    primaryColorSuppl: 'rgb(56, 137, 197)'
                },
                Switch: {
                    railColorActive: '#c5f955',
                    loadingColor: '#c5f955'
                }
            }
        } else {
            return {
                common: {
                    primaryColor: '#2080f0',
                    primaryColorHover: '#4098fc',
                    primaryColorPressed: '#1060c9',
                    primaryColorSuppl: '#4098fc'
                },
                Switch: {
                    railColorActive: '#2080f0',
                    loadingColor: '#2080f0'
                }
            }
        }
    })

    return { isDark, theme, themeOverrides }
}