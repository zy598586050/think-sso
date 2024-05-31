<template>
    <NLayoutSider bordered :native-scrollbar="false" width="200" show-trigger="arrow-circle" collapse-mode="transform" :collapsed-width="0">
        <NMenu :options="siderMenuOptions" v-model:value="activeKey" @update:value="handleUpdateValue" />
    </NLayoutSider>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { NLayoutSider, NMenu } from 'naive-ui'
import { useRouter } from 'vue-router'
import { useAppStore } from '@/store'
import { menuList } from '@/router'

const appStore = useAppStore()
const activeKey = computed(() => appStore.menuActive)
const router = useRouter()

const siderMenuOptions = computed(() => menuList.map((item, index) => {
    return {
        icon: item.meta.icon,
        label: item.meta.title,
        key: index + 1,
        path: item.path
    }
}))

const handleUpdateValue = (key: number) => {
    appStore.menuActive = key
    router.push(siderMenuOptions.value.find((item) => item.key === key)?.path as string)
}
</script>