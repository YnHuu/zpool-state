<script setup>
import { ref, onMounted } from 'vue'
import Alert from '@/components/Alert';
import Loading from '@/components/Loading'
import { getWechat, setWechat, testWechat } from '@/api'

const data = ref({})
onMounted(async () => {
    await getWechat().then((res) => { data.value = res.data })
})

const save = () => {
    setWechat(data.value).then((res) => {
        Alert({ message: res.data, type: "success" })
    })
}
const test = () => {
    Loading.show()
    testWechat(data.value).then((res) => {
        Alert({ message: res.data })
        Loading.hide()
    })
}
</script>

<template>

    <div class="grid justify-center card border border-base-300 bg-base-100 shadow-md py-4">
        <a href="https://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=sandbox/login" target="_blank"
            rel="noopener noreferrer" class="flex items-center">
            <span class="text-sm italic text-red-400 underline">Go公众号测试平台</span>
            <Sicon name="arrow" class="size-3" />
        </a>
        <div class="label">
            <span class="labelText">APPID:</span>
        </div>
        <input class="inputText" v-model="data.appid" />
        <div class="label">
            <span class="labelText">APPSecret:</span>
        </div>
        <input class="inputText" v-model="data.appsecret" />
        <div class="label">
            <span class="labelText">模板ID:</span>
        </div>
        <input class="inputText" v-model="data.template_id" />
        <div class="label">
            <span class="labelText">用户ID:</span>
        </div>
        <input class="inputText" v-model="data.to_user" />

        <div class="flex justify-end gap-2 my-4">
            <button class="btn btn-neutral" @click="test">测试</button>
            <button class="btn btn-outline btn-success" @click="save">保存</button>
        </div>
        <div class="flex flex-col">
            <span class="spanText">* 将会暴露 http://localhost/wechat 供本地 curl 调用</span>
            <span class="spanText">* curl localhost/wechat -d 'task=测试&status=ok'</span>
        </div>
    </div>
</template>

<style scoped>
.labelText {
    @apply drop-shadow text-stone-500 text-xs italic
}

.inputText {
    @apply input input-sm input-bordered italic text-stone-600 w-72 focus-within:outline-none
}

.spanText {
    @apply text-xs italic drop-shadow text-stone-500
}
</style>