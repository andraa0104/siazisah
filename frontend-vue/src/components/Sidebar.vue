<template>
  <aside :class="`w-64 bg-gradient-to-b from-blue-900 to-indigo-900 text-white flex-shrink-0 ${show ? 'fixed inset-y-0 left-0 z-50' : 'hidden'} md:block`">
    <div class="p-6">
      <div class="flex items-center justify-between mb-8">
        <div class="flex items-center">
          <div class="w-10 h-10 bg-white rounded-lg flex items-center justify-center">
            <i class="fas fa-mosque text-blue-900"></i>
          </div>
          <span class="ml-3 text-xl font-bold">SI-AZISAH</span>
        </div>
        <button @click="$emit('close')" class="md:hidden text-white">
          <i class="fas fa-times text-xl"></i>
        </button>
      </div>
      
      <nav class="space-y-2">
        <router-link v-for="item in menuItems" :key="item.path" :to="item.path" @click="$emit('close')"
          class="flex items-center px-4 py-3 rounded-lg transition hover:bg-white hover:bg-opacity-10"
          active-class="bg-white bg-opacity-20">
          <i :class="item.icon + ' w-5'"></i>
          <span class="ml-3">{{ item.label }}</span>
        </router-link>
      </nav>
    </div>
    
    <div class="absolute bottom-0 w-64 p-6">
      <button @click="handleLogout" class="w-full flex items-center px-4 py-3 bg-red-600 hover:bg-red-700 rounded-lg transition">
        <i class="fas fa-sign-out-alt w-5"></i>
        <span class="ml-3">Logout</span>
      </button>
    </div>
  </aside>
  <div v-if="show" @click="$emit('close')" class="fixed inset-0 bg-black bg-opacity-50 z-40 md:hidden"></div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const props = defineProps({
  menuItems: { type: Array, required: true },
  show: { type: Boolean, default: false }
})

defineEmits(['close'])

const router = useRouter()
const authStore = useAuthStore()

const handleLogout = () => {
  authStore.logout()
  router.push('/login')
}
</script>
