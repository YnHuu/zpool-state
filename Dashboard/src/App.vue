<script setup>
import { onMounted } from 'vue'
import { RouterView } from 'vue-router'
onMounted(() => {

  // document.body.style.setProperty('--el-fill-color-blank', '#000');
  // document.body.style.setProperty('--el-card-bg-color', '#fff');
})
</script>

<template>
  <RouterView />
</template>
