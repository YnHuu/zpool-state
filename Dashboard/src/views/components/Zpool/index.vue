<script setup>
import { computed } from 'vue'
const props = defineProps({
    data: {
        default: {},
    },
    checked: {
        default: false
    }
})

const indicatorClass = computed(() => ['indicator-item indicator-bottom indicator-center badge h-1.5', props.data.state == 'ONLINE' ? 'bg-success' : 'bg-error'])
const progressClass = computed(() => ['progress w-full', props.data.cap >= 90 ? 'progress-warning' : ''])

/*
const data = ({
    pool: 'dpool',
    state: 'online',
    raid: 'raidz1-0',
    size: '10.9T',
    alloc: '181G',
    free: '10.7T',
    cap: 1,
    items: [{ name: 'sdb', state: 'online', read: '0', write: '0', cksum: '0' }]
})
*/
</script>
<template>
    <div class="collapse collapse-arrow border border-base-300 bg-base-100 shadow-md mb-2">
        <input type="checkbox" :checked="checked" />
        <!-- <input type="radio" name="zpool-accordion" :checked="checked" /> -->
        <div class="collapse-title grid grid-cols-2">
            <div class="flex flex-row items-center">
                <div class="indicator">
                    <span :class="indicatorClass" />
                    <Sicon name="zpool" class="size-8 mr-2 py-0.5" />
                </div>
                <div class="flex flex-col">
                    <span class="font-medium text-base">{{ data.pool }}</span>
                    <span class="text-xs text-stone-400 ml-2">{{ data.raid }}</span>
                </div>
            </div>
            <div class="flex flex-col items-center justify-center">
                <span class="text-xs text-stone-400 mb-1">U:{{ data.alloc }} / F:{{ data.free }} / T:{{ data.size }}</span>
                <progress :class="progressClass" :value="data.cap" max="100"></progress>
            </div>
        </div>
        <div class="collapse-content text-xs text-stone-500 overflow-x-auto p-0">
            <table class="table text-center">
                <thead>
                    <tr>
                        <th>NAME</th>
                        <th>STATE</th>
                        <th>READ</th>
                        <th>WRITE</th>
                        <th>CKSUM</th>
                    </tr>
                </thead>
                <tbody>
                    <template v-for=" dev  in  data.items ">
                        <tr>
                            <th>{{ dev.name }}</th>
                            <th>{{ dev.state }}</th>
                            <th>{{ dev.read }}</th>
                            <th>{{ dev.write }}</th>
                            <th>{{ dev.cksum }}</th>
                        </tr>
                    </template>
                </tbody>
            </table>
        </div>
    </div>
</template>