<script setup>
import { particles } from './particles.js'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores'
import { loginApi } from '@/api'
import { ref } from 'vue';
const router = useRouter();

const data = ref({ password: '' })
const onSubmit = async () => {
    try {
        await loginApi(data.value).then((res) => { useUserStore().setToken(res.data.token); router.push('/') })
    } catch (e) {
        console.log(e);
    }
};

</script>
<template>
    <div class="h-full flex flex-col justify-end items-center">
        <div class="w-80 shadow-xl mb-96 md:mb-40 z-10">
            <label class="h-10 glass input input-md flex items-center ">
                <input autofocus type="password" class="grow" placeholder="请输入管理密码" @submit.native.prevent
                    @keyup.enter.native="onSubmit" v-model="data.password" />
                <Sicon name="enter" class="size-4 opacity-70" @click="onSubmit" />
            </label>
        </div>
        <vue-particles id="tsparticles" :options="particles" />
    </div>

</template>