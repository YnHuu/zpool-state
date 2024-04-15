<script setup>
import { computed, onMounted, ref } from 'vue'
import Sicon from '../Sicon'

const props = defineProps({
    type: {
        default: 'info',
    },
    top: {
        default: 0
    },
    message: {
        default: '',
        required: true
    },
    duration: {
        default: 5000
    },
    onDestroy: Function,
    onClose: Function
})

const stroke = computed(() => ['size-6', 'stroke-' + props.type])
const visiable = ref(false)

const close = () => { visiable.value = false }

onMounted(() => { visiable.value = true })

setTimeout(close, props.duration)
</script>

<template>
    <transition @before-leave="onClose" @after-leave="onDestroy" name="alert-fade">
        <div v-show="visiable" class="flex justify-center">
            <div role="alert" class="alertClass" :style="{ top: `${top}px` }">
                <Sicon :name="type" :class="stroke" />
                <span class="drop-shadow">{{ message }}</span>
            </div>
        </div>
    </transition>
</template>

<style scoped>
.alertClass {
    @apply alert w-auto flex shadow-md fixed z-50 py-2 transition-all text-xs
}

.alert-fade-enter-from,
.alert-fade-leave-to {
    @apply transform -translate-y-1 opacity-0
}
</style>