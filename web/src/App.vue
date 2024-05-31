<template>
  <NConfigProvider class="h-full" :theme="theme" :theme-overrides="themeOverrides">
    <NMessageProvider>
      <NDialogProvider>
        <RouterView />
      </NDialogProvider>
    </NMessageProvider>
  </NConfigProvider>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { NConfigProvider, NMessageProvider, NDialogProvider } from 'naive-ui'
import { useTheme } from '@/hooks/useTheme'

const { theme, themeOverrides } = useTheme()

onMounted(() => {
  document.title = import.meta.env.VITE_APP_TITLE || '统一登录系统'
})
</script>