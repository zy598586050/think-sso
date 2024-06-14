<template>
    <div class="page h-screen flex flex-col justify-center items-center">
        <div class="flex items-center py-10">
            <span class="ml-2 font-bold text-xl text-slate-700">{{ logoTitle }}</span>
        </div>
        <NTabs type="segment" animated class="w-[350px]">
            <NTabPane name="Email" tab="账号登录">
                <NForm ref="formEmailRef" :model="formEmail" label-placement="left" label-width="auto"
                    :rules="rulesEmail" class="mt-1">
                    <NFormItem path="email">
                        <NInput v-model:value="formEmail.email" placeholder="请输入邮箱" clearable>
                            <template #prefix>
                                <NIcon>
                                    <MailOutline />
                                </NIcon>
                            </template>
                        </NInput>
                    </NFormItem>
                    <NFormItem path="password">
                        <NInput v-model:value="formEmail.password" placeholder="请输入密码" type="password" clearable
                            show-password-on="mousedown">
                            <template #prefix>
                                <NIcon>
                                    <LockOpenOutline />
                                </NIcon>
                            </template>
                        </NInput>
                    </NFormItem>
                    <NFormItem>
                        <NButton type="primary" class="!w-full" @click="handleLogin(1)" :loading="isLoading">登 录
                        </NButton>
                    </NFormItem>
                </NForm>
            </NTabPane>
            <NTabPane name="phone" tab="手机号登录">
                <NForm ref="formPhoneRef" :model="formPhone" label-placement="left" label-width="auto"
                    :rules="rulesPhone" class="mt-1">
                    <NFormItem path="phone">
                        <NInput v-model:value="formPhone.phone" placeholder="请输入手机号" @input="handleInput" clearable>
                            <template #prefix>
                                <NIcon>
                                    <PhonePortraitOutline />
                                </NIcon>
                            </template>
                        </NInput>
                    </NFormItem>
                    <NFormItem path="code">
                        <NInputGroup>
                            <NInput v-model:value="formPhone.code" placeholder="请输入验证码" clearable />
                            <NButton type="primary" @click="sendSMS" :disabled="isClick">
                                {{ sendDisabled ? '重新发送' : '发送验证码' }}
                                <NCountdown v-if="sendDisabled" :render="renderCountdown" :duration="60 * 1000"
                                    :active="sendDisabled" :on-finish="countDownFinish" />
                            </NButton>
                        </NInputGroup>
                    </NFormItem>
                    <NFormItem>
                        <NButton type="primary" class="!w-full" @click="handleLogin(2)" :loading="isLoading">登 录
                        </NButton>
                    </NFormItem>
                </NForm>
            </NTabPane>
        </NTabs>
    </div>
</template>

<script lang="ts" setup>
import { reactive, ref, computed } from 'vue'
import { NForm, NFormItem, NInput, NInputGroup, NButton, NCountdown, FormInst, CountdownProps, NIcon, FormItemRule, NTabs, NTabPane } from 'naive-ui'
import { PhonePortraitOutline, MailOutline, LockOpenOutline } from '@vicons/ionicons5'
import { LoginEmail, GetCode, UserInfo } from '@/api/user'
import { getQueryParam } from '@/utils'
import { useUserStore } from '@/store'
import { useRouter } from 'vue-router'

const router = useRouter()
const formPhoneRef = ref<FormInst | null>(null)
const formEmailRef = ref<FormInst | null>(null)
const formPhone = reactive({
    phone: '',
    code: ''
})
const formEmail = reactive({
    email: '',
    password: ''
})
const sendDisabled = ref<boolean>(false)
const isLoading = ref<boolean>(false)
const logoTitle = computed(() => import.meta.env.VITE_APP_TITLE)
const userStore = useUserStore()

const isClick = computed(() => sendDisabled.value || !/^1[3-9]\d{9}$/.test(formPhone.phone))

const rulesPhone = reactive({
    phone: [
        {
            required: true,
            validator: (rule: FormItemRule, value: string) => {
                if (!value) {
                    return new Error('请输入手机号')
                } else if (!/^1[3-9]\d{9}$/.test(value)) {
                    return new Error('请输入正确的手机号')
                }
                console.log(rule)
                return true
            },
            trigger: ['input', 'blur']
        }
    ],
    code: [
        {
            required: true,
            message: '请输入验证码',
            trigger: ['input', 'blur']
        }
    ]
})

const rulesEmail = reactive({
    email: [
        {
            required: true,
            message: '请输入邮箱',
            trigger: ['input', 'blur']
        }
    ],
    password: [
        {
            required: true,
            message: '请输入密码',
            trigger: ['input', 'blur']
        }
    ]
})

const handleInput = (value: string) => {
    formPhone.phone = value.replace(/\D/g, '')
}

const renderCountdown: CountdownProps['render'] = ({ seconds }) => {
    return `(${String(seconds || 60)})`
}

const countDownFinish = () => {
    sendDisabled.value = false
}

const sendSMS = () => {

}

const handleLogin = (type: number) => {
    if (type === 1) {
        formEmailRef.value?.validate((error) => {
            if (error) return
            isLoading.value = true
            LoginEmail({
                email: formEmail.email,
                password: formEmail.password
            }).then(() => {
                isLoading.value = false
                // 获取用户信息
                UserInfo().then((result) => {
                    userStore.setUserInfo(result?.data)
                })
                // 其他系统跳转过来
                const redirect_url = getQueryParam('redirect_url')
                if (redirect_url) {
                    if (import.meta.env.VITE_SSO_IS_SAME === 'true') {
                        window.location.href = redirect_url
                    } else {
                        GetCode().then(res => {
                            window.location.href = redirect_url + `${redirect_url.includes('?') ? '&' : '?'}code=${res.data.code}`
                        })
                    }
                } else {
                    router.push('/user')
                }
            })
        })
    } else {
        formPhoneRef.value?.validate((error) => {
            if (error) return
            isLoading.value = true
        })
    }
}
</script>

<style lang="scss" scoped>
@media (min-width: 768px) {
    .page {
        background-image: url('@/assets/login.svg');
        background-repeat: no-repeat;
        background-position: 50%;
        background-size: 100%;
    }
}
</style>