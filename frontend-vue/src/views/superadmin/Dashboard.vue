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
          <div class="flex items-center gap-3">
            <button
              @click="openPrintModal"
              class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition text-sm font-medium"
            >
              <i class="fas fa-print mr-2"></i>
              Cetak Laporan Data Zakat
            </button>
            <button
              @click="openPrintMustahiqModal"
              class="bg-indigo-600 text-white px-4 py-2 rounded-lg hover:bg-indigo-700 transition text-sm font-medium"
            >
              <i class="fas fa-print mr-2"></i>
              Cetak Data Mustahiq Global
            </button>
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
          <template v-if="isLoading">
            <SkeletonCard v-for="i in 4" :key="i" type="stat" />
          </template>
          <div v-else v-for="stat in stats" :key="stat.label" class="bg-white rounded-xl shadow-md p-6">
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

    <Modal
      :show="showPrintModal"
      title="Cetak Laporan Data Zakat"
      @close="showPrintModal = false"
    >
      <div class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Tanggal Laporan</label>
          <input
            v-model="printSignDate"
            type="date"
            class="w-full border border-gray-300 rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>
        <div class="flex justify-end gap-2">
          <button
            type="button"
            class="px-4 py-2 rounded-md border border-gray-300 hover:bg-gray-50"
            @click="showPrintModal = false"
          >
            Batal
          </button>
          <button
            type="button"
            class="px-4 py-2 rounded-md bg-blue-600 text-white hover:bg-blue-700"
            @click="printZakatData"
          >
            Cetak
          </button>
        </div>
      </div>
    </Modal>

    <Modal
      :show="showPrintMustahiqModal"
      title="Cetak Data Mustahiq Global"
      @close="showPrintMustahiqModal = false"
    >
      <div class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Tanggal Laporan</label>
          <input
            v-model="printMustahiqSignDate"
            type="date"
            class="w-full border border-gray-300 rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-indigo-500"
          />
        </div>
        <div class="flex justify-end gap-2">
          <button
            type="button"
            class="px-4 py-2 rounded-md border border-gray-300 hover:bg-gray-50"
            @click="showPrintMustahiqModal = false"
          >
            Batal
          </button>
          <button
            type="button"
            class="px-4 py-2 rounded-md bg-indigo-600 text-white hover:bg-indigo-700"
            @click="printMustahiqGlobalData"
          >
            Cetak
          </button>
        </div>
      </div>
    </Modal>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useAuthStore } from '../../stores/auth'
import Sidebar from '../../components/Sidebar.vue'
import SkeletonCard from '../../components/SkeletonCard.vue'
import Modal from '../../components/Modal.vue'
import api from '../../api'

const authStore = useAuthStore()
const showSidebar = ref(false)
const isLoading = ref(true)
const showPrintModal = ref(false)
const printSignDate = ref('')
const showPrintMustahiqModal = ref(false)
const printMustahiqSignDate = ref('')

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

const todayDate = () => {
  const date = new Date()
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')
  return `${year}-${month}-${day}`
}

const openPrintModal = () => {
  printSignDate.value = todayDate()
  showPrintModal.value = true
}

const openPrintMustahiqModal = () => {
  printMustahiqSignDate.value = todayDate()
  showPrintMustahiqModal.value = true
}

const printZakatData = async () => {
  try {
    const selectedDate = printSignDate.value || todayDate()
    const response = await api.getPrintSuperadminZakatData(selectedDate)
    const printWindow = window.open('', '_blank')
    if (!printWindow) {
      alert('Popup diblokir browser. Izinkan popup untuk mencetak laporan.')
      return
    }
    printWindow.document.open()
    printWindow.document.write(response.data)
    printWindow.document.close()
    showPrintModal.value = false
  } catch (error) {
    console.error(error)
    alert('Gagal menyiapkan cetak laporan data zakat')
  }
}

const printMustahiqGlobalData = async () => {
  try {
    const selectedDate = printMustahiqSignDate.value || todayDate()
    const response = await api.getPrintSuperadminMustahiqGlobalData(selectedDate)
    const printWindow = window.open('', '_blank')
    if (!printWindow) {
      alert('Popup diblokir browser. Izinkan popup untuk mencetak laporan.')
      return
    }
    printWindow.document.open()
    printWindow.document.write(response.data)
    printWindow.document.close()
    showPrintMustahiqModal.value = false
  } catch (error) {
    console.error(error)
    alert('Gagal menyiapkan cetak data mustahiq global')
  }
}

onMounted(async () => {
  isLoading.value = true
  try {
    const [publicData, usersData] = await Promise.all([
      api.getPublicDashboard(),
      api.getUsers({ role: 'petugas', page: 1, page_size: 5 })
    ])
    
    if (publicData.data.success) {
      stats.value[0].value = publicData.data.data.total_masjid
      stats.value[2].value = publicData.data.data.total_muzakki
      stats.value[3].value = publicData.data.data.total_mustahiq
    }
    
    if (usersData.data.success) {
      const payload = usersData.data.data || {}
      if (payload.pagination && typeof payload.pagination.total === 'number') {
        stats.value[1].value = payload.pagination.total
      } else {
        const items = Array.isArray(payload.items) ? payload.items : (Array.isArray(payload) ? payload : [])
        stats.value[1].value = items.filter(u => u.role === 'petugas').length
      }
    }
  } catch (error) {
    console.error(error)
  } finally {
    isLoading.value = false
  }
})
</script>
