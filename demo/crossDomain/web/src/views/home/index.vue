<template>首页</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { UserInfo } from '@/api/user'

onMounted(() => {
    UserInfo()
})
</script>