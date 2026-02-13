<template>
  <div class="flex h-screen overflow-hidden">
    <Sidebar :show="showSidebar" @close="showSidebar = false" :menuItems="menuItems" />
    
    <div class="flex-1 flex flex-col overflow-hidden">
      <header class="bg-white shadow-sm">
        <div class="px-4 md:px-6 py-4 flex justify-between items-center">
          <div class="flex items-center">
            <button @click="showSidebar = true" class="mr-3 md:hidden text-gray-600 hover:text-gray-800">
              <i class="fas fa-bars text-xl"></i>
            </button>
            <div>
              <h1 class="text-xl md:text-2xl font-bold text-gray-800">Dashboard Petugas</h1>
              <p class="text-sm text-gray-600">{{ masjidName }}</p>
            </div>
          </div>
          <div class="flex items-center">
            <div class="text-right mr-4 hidden sm:block">
              <p class="text-sm font-medium text-gray-800">{{ authStore.user?.full_name }}</p>
              <p class="text-xs text-gray-500">Petugas Zakat</p>
            </div>
            <div class="w-10 h-10 bg-green-600 rounded-full flex items-center justify-center">
              <i class="fas fa-user text-white"></i>
            </div>
          </div>
        </div>
      </header>

      <main class="flex-1 overflow-y-auto p-4 md:p-6">
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 md:gap-6 mb-8">
          <div v-for="stat in stats" :key="stat.label" :class="`bg-gradient-to-br ${stat.gradient} rounded-xl shadow-lg p-6 text-white`">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-white text-opacity-90 text-sm">{{ stat.label }}</p>
                <p class="text-2xl md:text-3xl font-bold mt-2">{{ stat.value }}</p>
              </div>
              <div class="w-12 h-12 bg-white bg-opacity-20 rounded-lg flex items-center justify-center">
                <i :class="`${stat.icon} text-2xl`"></i>
              </div>
            </div>
          </div>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-4 md:gap-6">
          <router-link to="/petugas/pengaturan" class="bg-white rounded-xl shadow-md p-6 hover:shadow-lg transition group">
            <div class="flex items-center justify-between mb-4">
              <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center group-hover:bg-blue-200 transition">
                <i class="fas fa-cog text-blue-600 text-xl"></i>
              </div>
              <i class="fas fa-arrow-right text-gray-400 group-hover:text-blue-600 transition"></i>
            </div>
            <h3 class="text-lg font-bold text-gray-800 mb-2">Pengaturan</h3>
            <p class="text-sm text-gray-600">Atur data masjid & kadar zakat</p>
          </router-link>

          <router-link to="/petugas/mustahiq" class="bg-white rounded-xl shadow-md p-6 hover:shadow-lg transition group">
            <div class="flex items-center justify-between mb-4">
              <div class="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center group-hover:bg-green-200 transition">
                <i class="fas fa-hand-holding-heart text-green-600 text-xl"></i>
              </div>
              <i class="fas fa-arrow-right text-gray-400 group-hover:text-green-600 transition"></i>
            </div>
            <h3 class="text-lg font-bold text-gray-800 mb-2">Kelola Mustahiq</h3>
            <p class="text-sm text-gray-600">Data penerima zakat</p>
          </router-link>

          <router-link to="/petugas/transaksi" class="bg-white rounded-xl shadow-md p-6 hover:shadow-lg transition group">
            <div class="flex items-center justify-between mb-4">
              <div class="w-12 h-12 bg-purple-100 rounded-lg flex items-center justify-center group-hover:bg-purple-200 transition">
                <i class="fas fa-receipt text-purple-600 text-xl"></i>
              </div>
              <i class="fas fa-arrow-right text-gray-400 group-hover:text-purple-600 transition"></i>
            </div>
            <h3 class="text-lg font-bold text-gray-800 mb-2">Input Transaksi</h3>
            <p class="text-sm text-gray-600">Catat pembayaran zakat</p>
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

const menuItems = [
  { path: '/petugas', label: 'Dashboard', icon: 'fas fa-home' },
  { path: '/petugas/pengaturan', label: 'Pengaturan', icon: 'fas fa-cog' },
  { path: '/petugas/mustahiq', label: 'Mustahiq', icon: 'fas fa-hand-holding-heart' },
  { path: '/petugas/transaksi', label: 'Transaksi', icon: 'fas fa-receipt' }
]

const masjidName = ref('Loading...')
const showSidebar = ref(false)
const stats = ref([
  { label: 'Total Dizakati', value: 0, icon: 'fas fa-users', gradient: 'from-blue-500 to-blue-600' },
  { label: 'Total Transaksi', value: 0, icon: 'fas fa-receipt', gradient: 'from-purple-500 to-purple-600' },
  { label: 'Total Zakat', value: 'Rp 0', icon: 'fas fa-coins', gradient: 'from-yellow-500 to-yellow-600' },
  { label: 'Muzakki', value: 0, icon: 'fas fa-hand-holding-heart', gradient: 'from-green-500 to-green-600' }
])

const formatCurrency = (amount) => {
  return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR', minimumFractionDigits: 0 }).format(amount)
}

onMounted(async () => {
  try {
    const [masjidData, transaksiData] = await Promise.all([
      api.getMyMasjid(),
      api.getTransaksi()
    ])
    
    if (masjidData.data.success) masjidName.value = masjidData.data.data.nama
    if (transaksiData.data.success) {
      const totalDizakati = transaksiData.data.data.reduce((sum, t) => sum + (t.jumlah_orang || 0), 0)
      stats.value[0].value = totalDizakati
      stats.value[1].value = transaksiData.data.data.length
      const total = transaksiData.data.data.reduce((sum, t) => sum + t.total_dibayar, 0)
      stats.value[2].value = formatCurrency(total)
      stats.value[3].value = new Set(transaksiData.data.data.map(t => t.muzakki_id)).size
    }
  } catch (error) {
    console.error(error)
  }
})
</script>
