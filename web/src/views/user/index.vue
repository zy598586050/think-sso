<template>
    用户管理
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { UserList } from '@/api/user'

onMounted(() => {
    UserList({})
})
</script>

<style scoped lang="scss">

</style>