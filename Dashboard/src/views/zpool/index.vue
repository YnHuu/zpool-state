<script setup>
import { ref, onMounted } from 'vue'
import { getZFS } from '@/api'
import Zpool from '@/views/components/Zpool'

const data = ref([])
onMounted(async () => {
    await getZFS().then((res) => { data.value = res.data })
})


</script>
<template>
    <div>
        <template v-for="zfs, i in data">
            <Zpool :data="zfs" :checked="i == 0 ? true : false" />
        </template>
    </div>
</template>