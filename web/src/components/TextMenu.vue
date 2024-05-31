<template>
    <div class="flex items-center">
        <div v-for="item in options" :key="item.key" @click="updateValue(item.key)"
            :class="[item.key === value ? 'active-text-menu' : '', isDark ? 'dark-text-menu-item' : 'text-menu-item']">{{
                item.label }}</div>
    </div>
</template>

<script setup lang="ts">
import type { PropType, Component } from 'vue'
import { useTheme } from '@/hooks/useTheme'

const { isDark } = useTheme()

interface Option {
    key: string | number;
    icon: Component;
    label: string;
}

defineProps({
    value: [String, Number] as PropType<string | number>,
    options: Array as () => Option[]
})

const emit = defineEmits(['update:value'])

const updateValue = (key: string | number) => {
    emit('update:value', key)
}
</script>

<style scoped lang="scss">
.dark-text-menu-item {
    position: relative;
    margin-right: 25px;
    cursor: pointer;
}

.dark-text-menu-item::after {
    content: "";
    display: block;
    position: absolute;
    width: 100%;
    height: 2px;
    background-color: $dark-primaryColor;
    bottom: -8px;
    left: 0;
    transform: scaleX(0);
    transition: transform 0.3s ease-in-out;
}

.dark-text-menu-item.active-text-menu::after {
    transform: scaleX(1);
}

.dark-text-menu-item:hover::after {
    transform: scaleX(1);
}

.text-menu-item {
    position: relative;
    margin-right: 25px;
    cursor: pointer;
}

.text-menu-item::after {
    content: "";
    display: block;
    position: absolute;
    width: 100%;
    height: 2px;
    background-color: $primaryColor;
    bottom: -8px;
    left: 0;
    transform: scaleX(0);
    transition: transform 0.3s ease-in-out;
}

.text-menu-item.active-text-menu::after {
    transform: scaleX(1);
}

.text-menu-item:hover::after {
    transform: scaleX(1);
}
</style>