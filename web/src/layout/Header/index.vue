<template>
    <NLayoutHeader bordered class="flex justify-between items-center px-4 py-3">
        <div class="flex items-center">
            <span class="ml-2 font-bold">sky3Dgen</span>
        </div>
        <div class="flex items-center">
            <TextMenu v-model:value="activeKey" :options="textMenuOptions" @update:value="selectTextMenu" />
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
import TextMenu from '@/components/TextMenu.vue'
import { menuList } from '@/router'


const { isDark } = useTheme()
const dialog = useDialog()
const appStore = useAppStore()
const userStore = useUserStore()
const router = useRouter()
const activeKey = computed(() => appStore.menuActive)

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

const textMenuOptions = computed(() => menuList.map((item, index) => {
    return {
        icon: item.meta.icon,
        label: item.meta.title,
        key: index + 1,
        path: item.path
    }
}))

const selectTextMenu = (key: number) => {
    appStore.menuActive = key
    router.push(textMenuOptions.value.find((item) => item.key === key)?.path as string)
}

const handleUpdateValue = (value: boolean) => {
    appStore.theme = value ? 'dark' : 'light'
}

const handleSelect = (key: string) => {
    switch (key) {
        case 'my':
            router.push('/my')
            //appStore.setMenuActive(4)
            break;
        case 'logout':
            dialog.warning({
                title: '提示',
                content: '您确定要退出登录？',
                positiveText: '确定',
                negativeText: '取消',
                onPositiveClick: () => {
                    userStore.clearUserInfo()
                    router.push('/login')
                }
            })
            break;
        default:
            break;
    }
}
</script>