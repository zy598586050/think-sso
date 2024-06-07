<template>首页</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { TestHome } from '@/api/user'

onMounted(() => {
    TestHome()
})
</script>