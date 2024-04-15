<template>
  <svg aria-hidden="true" class="drop-shadow">
    <use v-if="name" :xlink:href="symbolId" />
  </svg>
</template>

<script setup>
import { computed } from 'vue'
const props = defineProps({
  name: {
    default: '',
    required: true
  }
})

const symbolId = computed(() => `#icon-${props.name}`)
</script>
