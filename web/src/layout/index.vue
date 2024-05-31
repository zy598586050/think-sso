<template>
    <NLayout class="h-full">
        <Header ref="headerRef" />
        <NLayout has-sider :style="{ height: nLayoutHeight + 'px' }">
            <NLayoutContent has-sider>
                <RouterView />
            </NLayoutContent>
        </NLayout>
    </NLayout>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { NLayout, NLayoutContent } from 'naive-ui'
import Header from './Header/index.vue'

const headerRef = ref()
const nLayoutHeight = ref(0)

onMounted(() => {
    const headerHeight = headerRef.value.$el.offsetHeight
    nLayoutHeight.value = window.innerHeight - headerHeight
})
</script>