<script setup>
import { useRouter } from 'vue-router'
import SidebarItem from './SidebarItem.vue'

const routers = useRouter().options.routes

</script>

<template>
    <div class="h-full p-2 w-52">
        <aside class="h-full overflow-y-auto shadow-lg rounded-lg bg-base-100">
            <ul class="menu">
                <SidebarItem v-for="route in routers" :item="route" :path="route.path" />
            </ul>
        </aside>
    </div>
</template>