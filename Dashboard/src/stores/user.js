import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', () => {
    const token = ref('')
    const setToken = (newToken) => {
        token.value = newToken;
    }
    const delToken = () => {
        token.value = ''
    }
    return { token, setToken, delToken }
}, { persist: true }
)