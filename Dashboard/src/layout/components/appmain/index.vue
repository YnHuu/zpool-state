<template>
    <div class="px-2">
        <RouterView v-slot="{ Component, route }">
            <transition name="main-fade">
                <KeepAlive include="disk">
                    <component :is="Component" :key="route.fullPath" />
                </KeepAlive>
            </transition>
        </RouterView>
    </div>
</template>

<style scoped>
.main-fade-leave-active,
.main-fade-enter-active {
    @apply opacity-0 transform duration-300
}

.main-fade-leave-to {
    @apply translate-y-1/4
}
</style>