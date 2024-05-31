<template>
    <SiderSlot width="200">
        <NMenu :options="menuOptions" v-model:value="activeKey" @update:value="handleUpdateValue" />
    </SiderSlot>
    <NLayoutContent>
        <div class="flex h-full">
            <div class="w-52 px-3 py-2 upload-sider flex flex-col items-center">
                <NUpload ref="upRef" :custom-request="customRequest" accept=".png,.jpg" :max="1" :show-file-list="false"
                    style="width: auto">
                    <div class="n-upload-trigger n-upload-trigger--image-card">
                        <img class="n-upload-dragger" src="" v-if="false" />
                        <div v-else class="n-upload-dragger"><i class="n-base-icon"><svg width="512" height="512"
                                    viewBox="0 0 512 512" fill="none" xmlns="http://www.w3.org/2000/svg">
                                    <path d="M256 112V400M400 256H112" stroke="currentColor" stroke-width="32"
                                        stroke-linecap="round" stroke-linejoin="round"></path>
                                </svg></i>
                        </div>
                    </div>
                </NUpload>
                <NButton type="primary" class="mt-5">
                    <template #icon>
                        <NIcon>
                            <GameControllerOutline />
                        </NIcon>
                    </template>
                    生成
                </NButton>
            </div>
            <div>
                <div class="px-3 box-border">
                    <!-- 自定义 -->
                    <div class="h-[250px] w-[350px]" style="border: 1px solid rgba(63, 63, 68, 1)">
                        <ThreeLoader filePath="http://ai-game-hk.oss-cn-hongkong.aliyuncs.com/common/hao.jiangg/Datasets/Object3D/Raw/objaverse/glbs/000-023/a0c216a7de774aba903f441d7899aa02.glb" @process="onProcess" />
                    </div>
                    <NProgress type="line" :percentage="process" v-if="process < 100" />
                </div>
            </div>
        </div>
    </NLayoutContent>
</template>

<script setup lang="ts">
import { ref, h } from 'vue'
import { NLayoutContent, NIcon, NMenu, NUpload, NButton, NProgress } from 'naive-ui'
import SiderSlot from '@/components/SiderSlot.vue'
import { GameControllerOutline } from '@vicons/ionicons5'
import ThreeLoader from '@/components/ThreeLoader.vue'

const process = ref(0)
const activeKey = ref(1)

const menuOptions = ref([
    {
        label: '图生3D',
        key: 1,
        icon: () => h(NIcon, null, { default: () => h(GameControllerOutline) })
    },
    {
        label: '文本上色',
        key: 2,
        icon: () => h(NIcon, null, { default: () => h(GameControllerOutline) })
    }
])

const onProcess = (event: any) => {
    process.value = Math.floor((event.loaded / event.total) * 100)
}

const handleUpdateValue = (key: number) => {
    activeKey.value = key
}

const customRequest = () => { }
</script>

<style scoped lang="scss">
.upload-sider {
    border-right: 1px solid rgba(255, 255, 255, 0.09);
}

.n-upload-trigger--image-card {
    width: 150px;
    height: 150px
}
</style>