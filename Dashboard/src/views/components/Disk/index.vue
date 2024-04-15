<script setup>
import { computed } from 'vue'
const props = defineProps({
    data: {
        default: {},
    },
    checked: {
        default: false,
    }
})

const indicator = computed(() => ['indicator-item indicator-bottom indicator-center badge h-1.5', props.data.smart_status == 'true' ? 'bg-success' : 'bg-error'])

/*
const data =  {
        index : 1,
        use: 'dpool'
        model_name: 'ST4000VN008-2DR166',
        user_capacity: '4.00 TB',
        smart_status: 'trues',
        dev: 'sdb',
        temperature: '32',
        serial_number: 'ZDHBGR41',
        firmware_version: 'SC60',
        power_on_time: '16308',
        form_factor: '3.5 inches'
    }
*/
</script>

<template>
    <div class="collapse collapse-arrow border border-base-300 bg-base-100 shadow-md mb-2">
        <input type="checkbox" :checked="checked" />
        <!-- <input type="radio" name="disk-accordion" :checked="checked" /> -->
        <div class="collapse-title grid grid-cols-3">
            <div class="flex flex-row items-center col-span-2">
                <div class="indicator">
                    <span :class="indicator" />
                    <Sicon name="hdd" class="h-10 w-10 mr-2 py-0.5" />
                </div>
                <div class="flex flex-col">
                    <span class="font-medium text-base">硬盘 {{ data.index }}</span>
                    <span class="text-xs text-stone-400">{{ data.model_name }}</span>
                </div>
            </div>
            <span class="flex items-center font-medium text-base">{{ data.user_capacity }}</span>
        </div>
        <div class="collapse-content flex-col space-y-2 text-xs text-stone-500">
            <span class="grid grid-cols-2">
                <button class="flex btn-xs btn-link p-0 h-0 min-h-0">SMART:</button>
                <span :class="data.smart_status == 'true' ? 'text-success' : 'text-error'">
                    {{ data.smart_status == 'true' ? '良好' : '故障' }}
                </span>
            </span>
            <span class="grid grid-cols-2">
                <span>用途:</span>
                <span>{{ data.use }}</span>
            </span>
            <span class="grid grid-cols-2">
                <span>盘符:</span>
                <span>{{ data.dev }}</span>
            </span>
            <span class="grid grid-cols-2">
                <span>温度:</span>
                <span>{{ data.temperature }} °C</span>
            </span>
            <span class="grid grid-cols-2">
                <span>序列号:</span>
                <span>{{ data.serial_number }}</span>
            </span>
            <span class="grid grid-cols-2">
                <span>固件版本:</span>
                <span>{{ data.firmware_version }}</span>
            </span>
            <span class="grid grid-cols-2">
                <span>硬盘类型:</span>
                <span>{{ data.form_factor }}</span>
            </span>
            <span class="grid grid-cols-2">
                <span>通电时间:</span>
                <span>{{ data.power_on_time }} 小时</span>
            </span>
        </div>
    </div>

</template>