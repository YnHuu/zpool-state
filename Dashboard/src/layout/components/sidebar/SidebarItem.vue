<script setup>
import { useRoute, RouterLink } from 'vue-router'

const props = defineProps({
    item: {
        required: true
    },
    path: {
        default: ''
    }
})

const route = useRoute()
const open = route.path.includes(props.path)

const hidden = () => {
    if (document.getElementById('drawer').checked) {
        document.getElementById('drawer').click()
    }
}

</script>

<template>
    <template v-if="!item.hidden">
        <li v-if="item.meta.single || !item.children">
            <RouterLink class="flex gap-3 h-12 active:btn-block" activeClass="active" :to=path @click="hidden">
                <Sicon :name=item.meta.icon class="size-5" />
                <span class="text-sm">{{ item.meta.title }}</span>
            </RouterLink>
        </li>
        <li v-else>
            <details :open="open">
                <summary class="gap-3">
                    <Sicon :name=item.meta.icon class="size-5" />
                    <span class="text-sm">{{ item.meta.title }}</span>
                </summary>
                <ul>
                    <SidebarItem v-for="child in item.children" :item="child" :path="item.path + child.path" />
                </ul>
            </details>
        </li>
    </template>
</template>