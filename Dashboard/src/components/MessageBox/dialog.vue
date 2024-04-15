<script setup>
import { computed } from 'vue'
import Sicon from '../Sicon'
const props = defineProps({
  type: {
    default: 'info'
  },
  title: {
    default: '提示'
  },
  message: {
    default: ''
  },
  okText: {
    default: '确认'
  },
  cancelText: {
    default: '取消'
  },
  onDestroy: Function,
  onOK: Function,
})
const btn = computed(() => ['btn btn-sm', 'btn-' + props.type])
const stroke = computed(() => ['size-6', 'stroke-' + props.type])
</script>

<template>
  <input type="checkbox" id="messagebox" checked="checked" class="modal-toggle" />
  <div class="modal" role="dialog">
    <div class="modal-box">
      <div class="flex items-center gap-1">
        <Sicon :name="type" :class="stroke" />
        <h3 class="font-bold text-lg">{{ title }}</h3>
      </div>
      <p class="py-4 drop-shadow">{{ message }}</p>
      <div class="modal-action mt-0">
        <label for="messagebox" class="btn btn-sm" @click="onDestroy">{{ cancelText }}</label>
        <label for="messagebox" :class="btn" @click="onOK">{{ okText }}</label>
      </div>
    </div>
  </div>
</template>
