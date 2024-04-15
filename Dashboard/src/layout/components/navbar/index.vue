<script setup>
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import screenfull from 'screenfull'
import MessageBox from '@/components/MessageBox'
import { useUserStore } from '@/stores'
import { rebootApi } from '@/api'

const router = useRouter();

const reboot = () => {
    MessageBox({
        type: 'warning',
        title: "警告",
        message: "是否重启系统?",
        okText: "重启",
    }).then(() => {
        rebootApi()
    })
}
const logout = () => {
    useUserStore().delToken()
    router.push('/login')
}

const userAgent = computed(() => (/(iPhone|YnHuu)/i.test(navigator.userAgent)))
</script>
<template>
    <div class="flex justify-end p-2">
        <ul class="shadow-md rounded-md menu menu-sm menu-horizontal bg-base-100">
            <li class="md:hidden">
                <label for="drawer" aria-label="open sidebar">
                    <Sicon name="menu" class="size-4" />
                </label>
            </li>
            <li>
                <a @click="reboot">
                    <Sicon name="reboot" class="size-4" />
                </a>
            </li>
            <li :class="userAgent ? 'hidden' : ''">
                <a class="tooltip tooltip-bottom" data-tip="全屏" @click="screenfull.toggle()">
                    <Sicon name="screenfull" class="size-4" />
                </a>
            </li>
            <li>
                <a @click="logout">
                    <Sicon name="exit" class="size-4" />
                </a>
            </li>
        </ul>
    </div>
</template>