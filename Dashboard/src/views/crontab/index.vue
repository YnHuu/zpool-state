<script setup>
import { ref, onMounted } from 'vue'
import { getCron, setCron } from '@/api'
import Alert from '@/components/Alert';

const data = ref('')
const selected = ref('daily')
const remark = {
    '15min': '* 每15分钟运行一次',
    'hourly': '* 每一小时运行一次',
    'daily': '* 每天2点运行一次',
    'weekly': '* 每周六3点运行一次',
    'monthly': '* 每月1号5点运行一次',
}
const change = async () => {
    await getCron(selected.value).then((res) => { data.value = res.data })
}

const save = () => {
    setCron(selected.value, data.value).then((res) => {
        Alert({ message: res.data, type: "success" })
    })
}

onMounted(async () => {
    await getCron(selected.value).then((res) => { data.value = res.data })
})
</script>
<template>
    <div class="card border border-base-300 bg-base-100 shadow-md p-2 gap-3">
        <div role="tablist" class="tabs tabs-bordered justify-center drop-shadow">
            <input type="radio" name="cron_tabs" role="tab" class="tab" aria-label="每15分" value="15min"
                v-model="selected" @change="change" />
            <input type="radio" name="cron_tabs" role="tab" class="tab" aria-label="每时" value="hourly"
                v-model="selected" @change="change" />
            <input type="radio" name="cron_tabs" role="tab" class="tab" aria-label="每天" value="daily" v-model="selected"
                @change="change" checked />
            <input type="radio" name="cron_tabs" role="tab" class="tab" aria-label="每周" value="weekly"
                v-model="selected" @change="change" />
            <input type="radio" name="cron_tabs" role="tab" class="tab" aria-label="每月" value="monthly"
                v-model="selected" @change="change" />
        </div>
        <div>
            <textarea class="textarea textarea-bordered w-full min-h-96 leading-normal italic text-stone-600"
                v-model="data" />
            <div class="label">
                <span class="label-text-alt drop-shadow text-stone-500">{{ remark[selected] }}</span>
                <button class="btn btn-outline btn-success" @click="save">保存</button>
            </div>
        </div>
    </div>
</template>