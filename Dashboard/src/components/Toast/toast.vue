<script setup>
import { onMounted, ref, computed } from 'vue'
const props = defineProps({
    type: {
        default: 'info'
    },
    message: {
        default: '',
        required: true
    },
    duration: {
        default: 3000
    },
    onDestroy: Function
})
const visiable = ref(false)
const toast = computed(() => ['alert shadow-md grid-cols-[auto] py-2 text-xs md:text-base', 'alert-' + props.type])
const close = () => { visiable.value = false }

onMounted(() => { visiable.value = true })

setTimeout(close, props.duration)
</script>
<template>
    <transition @after-leave="onDestroy">
        <div v-show="visiable" class="toast toast-center">
            <div :class="toast">
                <span class="drop-shadow">{{ message }}</span>
            </div>
        </div>
    </transition>
</template>