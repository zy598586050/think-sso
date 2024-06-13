<template>
    <NLayoutSider bordered :native-scrollbar="false" show-trigger="arrow-circle" collapse-mode="transform"
        :collapsed-width="0" :width="width">
        <slot></slot>
    </NLayoutSider>
</template>

<script setup lang="ts">
import { NLayoutSider } from 'naive-ui'
defineProps({
    width: {
        type: [Number, String],
        default: 350
    }
})
</script>