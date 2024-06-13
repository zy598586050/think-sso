<template>
  <div class="h-full flex flex-col justify-center">
    <NResult status="404" title="404 资源不存在" description="生活总归带点荒谬">
      <template #footer>
        <NButton @click="goHome">回到首页</NButton>
      </template>
    </NResult>
  </div>
</template>

<script setup lang="ts">
import { NResult, NButton } from 'naive-ui'
import { useRouter } from 'vue-router'

const router = useRouter()

const goHome = () => {
  router.push('/home')
}
</script>