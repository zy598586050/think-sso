<template>
    <NLayoutHeader bordered class="flex justify-between items-center px-4 py-3">
        <div class="flex items-center">
            <span class="ml-2 font-bold">{{ logoTitle }}</span>
        </div>
        <div class="flex items-center">
            <NSwitch v-model:value="isDark" @update:value="handleUpdateValue">
                <template #checked-icon>
                    <NIcon :component="MoonOutline" />
                </template>
                <template #unchecked-icon>
                    <NIcon :component="SunnyOutline" />
                </template>
            </NSwitch>
            <NDropdown trigger="hover" :options="menuOptions" @select="handleSelect">
                <NAvatar class="ml-4 cursor-pointer" :src="userStore.userInfo.avatar" />
            </NDropdown>
        </div>
    </NLayoutHeader>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { NSwitch, NIcon, NAvatar, NDropdown, NLayoutHeader, useDialog } from 'naive-ui'
import { SunnyOutline, MoonOutline } from '@vicons/ionicons5'
import { useAppStore, useUserStore } from '@/store'
import { useRouter } from 'vue-router'
import { useTheme } from '@/hooks/useTheme'
import { Logout } from '@/api/user'


const { isDark } = useTheme()
const dialog = useDialog()
const appStore = useAppStore()
const userStore = useUserStore()
const router = useRouter()
const logoTitle = computed(() => import.meta.env.VITE_APP_TITLE)

const menuOptions = ref([
    {
        label: '个人中心',
        key: 'my'
    },
    {
        label: '退出登录',
        key: 'logout'
    }
])

const handleUpdateValue = (value: boolean) => {
    appStore.theme = value ? 'dark' : 'light'
}

const handleSelect = (key: string) => {
    switch (key) {
        case 'my':
            appStore.menuActive = 3
            router.push('/my')
            break;
        case 'logout':
            dialog.warning({
                title: '提示',
                content: '您确定要退出登录？',
                positiveText: '确定',
                negativeText: '取消',
                onPositiveClick: () => {
                    Logout().then(() => {
                        userStore.clearUserInfo()
                        router.push('/login')
                    })
                }
            })
            break;
        default:
            break;
    }
}
</script>