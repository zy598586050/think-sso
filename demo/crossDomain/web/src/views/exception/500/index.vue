<template>
  <div class="h-full flex flex-col justify-center">
    <NResult status="500" title="500 服务器错误" description="服务器出错可能说明该雇更多程序员了">
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