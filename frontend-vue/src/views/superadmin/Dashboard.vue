<template>
  <div class="flex h-screen overflow-hidden">
    <Sidebar :menuItems="menuItems" :show="showSidebar" @close="showSidebar = false" />
    
    <div class="flex-1 flex flex-col overflow-hidden">
      <header class="bg-white shadow-sm">
        <div class="px-4 md:px-6 py-4 flex justify-between items-center">
          <div class="flex items-center">
            <button @click="showSidebar = true" class="md:hidden mr-4 text-gray-600">
              <i class="fas fa-bars text-xl"></i>
            </button>
            <h1 class="text-xl md:text-2xl font-bold text-gray-800">Dashboard Superadmin</h1>
          </div>
          <div class="flex items-center">
            <div class="text-right mr-4 hidden sm:block">
              <p class="text-sm font-medium text-gray-800">{{ authStore.user?.full_name }}</p>
              <p class="text-xs text-gray-500">Administrator</p>
            </div>
            <div class="w-10 h-10 bg-blue-600 rounded-full flex items-center justify-center">
              <i class="fas fa-user text-white"></i>
            </div>
          </div>
        </div>
      </header>

      <main class="flex-1 overflow-y-auto p-4 md:p-6">
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 md:gap-6 mb-8">
          <div v-for="stat in stats" :key="stat.label" class="bg-white rounded-xl shadow-md p-6">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-gray-500 text-sm">{{ stat.label }}</p>
                <p class="text-2xl md:text-3xl font-bold text-gray-800 mt-2">{{ stat.value }}</p>
              </div>
              <div :class="`w-12 h-12 ${stat.bgColor} rounded-lg flex items-center justify-center`">
                <i :class="`${stat.icon} ${stat.textColor} text-xl`"></i>
              </div>
            </div>
          </div>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-4 md:gap-6">
          <router-link to="/superadmin/masjid" class="bg-white rounded-xl shadow-md p-6 hover:shadow-lg transition group">
            <div class="flex items-center justify-between mb-4">
              <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center group-hover:bg-blue-200 transition">
                <i class="fas fa-mosque text-blue-600 text-xl"></i>
              </div>
              <i class="fas fa-arrow-right text-gray-400 group-hover:text-blue-600 transition"></i>
            </div>
            <h3 class="text-lg font-bold text-gray-800 mb-2">Kelola Masjid/Langgar</h3>
            <p class="text-sm text-gray-600">Tambah, edit, atau hapus data masjid dan langgar</p>
          </router-link>

          <router-link to="/superadmin/users" class="bg-white rounded-xl shadow-md p-6 hover:shadow-lg transition group">
            <div class="flex items-center justify-between mb-4">
              <div class="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center group-hover:bg-green-200 transition">
                <i class="fas fa-users text-green-600 text-xl"></i>
              </div>
              <i class="fas fa-arrow-right text-gray-400 group-hover:text-green-600 transition"></i>
            </div>
            <h3 class="text-lg font-bold text-gray-800 mb-2">Kelola Petugas</h3>
            <p class="text-sm text-gray-600">Tambah, edit, atau hapus akun petugas zakat</p>
          </router-link>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useAuthStore } from '../../stores/auth'
import Sidebar from '../../components/Sidebar.vue'
import api from '../../api'

const authStore = useAuthStore()
const showSidebar = ref(false)

const menuItems = [
  { path: '/superadmin', label: 'Dashboard', icon: 'fas fa-home' },
  { path: '/superadmin/masjid', label: 'Masjid/Langgar', icon: 'fas fa-mosque' },
  { path: '/superadmin/users', label: 'Petugas', icon: 'fas fa-users' }
]

const stats = ref([
  { label: 'Total Masjid', value: 0, icon: 'fas fa-mosque', bgColor: 'bg-blue-100', textColor: 'text-blue-600' },
  { label: 'Total Petugas', value: 0, icon: 'fas fa-users', bgColor: 'bg-green-100', textColor: 'text-green-600' },
  { label: 'Total Muzakki', value: 0, icon: 'fas fa-hand-holding-heart', bgColor: 'bg-purple-100', textColor: 'text-purple-600' },
  { label: 'Total Mustahiq', value: 0, icon: 'fas fa-users', bgColor: 'bg-yellow-100', textColor: 'text-yellow-600' }
])

onMounted(async () => {
  try {
    const [publicData, usersData] = await Promise.all([
      api.getPublicDashboard(),
      api.getUsers()
    ])
    
    if (publicData.data.success) {
      stats.value[0].value = publicData.data.data.total_masjid
      stats.value[2].value = publicData.data.data.total_muzakki
      stats.value[3].value = publicData.data.data.total_mustahiq
    }
    
    if (usersData.data.success) {
      stats.value[1].value = usersData.data.data.filter(u => u.role === 'petugas').length
    }
  } catch (error) {
    console.error(error)
  }
})
</script>
