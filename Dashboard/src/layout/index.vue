<script setup>
import { Sidebar, Navbar, Main } from './components'
</script>

<template>
    <div class="drawer md:drawer-open">
        <input id="drawer" type="checkbox" class="drawer-toggle" />
        <div class="drawer-side z-50 md:z-0">
            <label for="drawer" aria-label="close sidebar" class="drawer-overlay"></label>
            <Sidebar />
        </div>
        <div class="drawer-content flex flex-auto flex-col z-0">
            <Navbar />
            <Main />
        </div>
    </div>
</template>