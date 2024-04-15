<script setup>
import { ref, onMounted } from 'vue'
import Disk from '@/views/components/Disk'
import Loading from '@/components/Loading'
import { getDisk } from '@/api'

defineOptions({
    name: 'disk',
})

const data = ref([])
onMounted(async () => {
    Loading.show()
    await getDisk().then((res) => { data.value = res.data; Loading.hide() })
})
</script>
<template>
    <div>
        <template v-for="disk, i in data">
            <Disk :data="disk" :checked="i == 0 ? true : false" />
        </template>
    </div>
</template>