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

        <div class="bg-white rounded-xl shadow-md p-6 mb-8">
          <div class="flex items-start justify-between gap-4 mb-6">
            <div>
              <h2 class="text-lg md:text-xl font-bold text-gray-800">Keterisian Data Per Masjid/Langgar</h2>
              <p class="text-sm text-gray-600 mt-1">Persentase dihitung berdasarkan jumlah masjid/langgar yang sudah mengisi data.</p>
            </div>
            <button
              type="button"
              class="text-sm font-medium text-gray-700 hover:text-gray-900"
              @click="loadDataCompleteness"
              :disabled="completenessLoading"
              title="Refresh"
            >
              <i :class="`fas fa-rotate ${completenessLoading ? 'animate-spin' : ''}`"></i>
            </button>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-3 gap-4 md:gap-6">
            <div class="rounded-2xl border border-gray-100 bg-gradient-to-br from-emerald-50 to-white p-5">
              <div class="flex items-center justify-between">
                <div>
                  <p class="text-sm font-semibold text-emerald-800">Data Muzakki</p>
                  <p class="text-xs text-gray-600 mt-1">Dianggap terisi jika minimal 1 muzakki.</p>
                </div>
                <div class="w-10 h-10 rounded-xl bg-emerald-100 flex items-center justify-center">
                  <i class="fas fa-hand-holding-heart text-emerald-700"></i>
                </div>
              </div>

              <div class="mt-5 flex items-center gap-5">
                <button
                  type="button"
                  class="relative w-24 h-24 rounded-full p-[6px] shadow-sm hover:shadow transition focus:outline-none focus:ring-2 focus:ring-emerald-400"
                  :style="{ background: pieBackground(muzakkiPercent, '#10b981') }"
                  @click="openCompletenessModal('muzakki')"
                  :disabled="completenessLoading"
                >
                  <div class="w-full h-full bg-white rounded-full flex items-center justify-center">
                    <div class="text-center leading-tight">
                      <div class="text-lg font-bold text-gray-900">{{ formatPercent(muzakkiPercent) }}</div>
                      <div class="text-[11px] text-gray-500">terisi</div>
                    </div>
                  </div>
                </button>
                <div class="flex-1">
                  <div class="text-sm text-gray-700">
                    <span class="font-semibold text-gray-900">{{ completeness.filled_muzakki }}</span>
                    dari
                    <span class="font-semibold text-gray-900">{{ completeness.total_masjid }}</span>
                    masjid/langgar
                  </div>
                  <div class="text-xs text-gray-500 mt-1">Klik pie chart untuk lihat detail per masjid.</div>
                </div>
              </div>
            </div>

            <div class="rounded-2xl border border-gray-100 bg-gradient-to-br from-amber-50 to-white p-5">
              <div class="flex items-center justify-between">
                <div>
                  <p class="text-sm font-semibold text-amber-800">Data Mustahiq</p>
                  <p class="text-xs text-gray-600 mt-1">Dianggap terisi jika minimal 1 mustahiq.</p>
                </div>
                <div class="w-10 h-10 rounded-xl bg-amber-100 flex items-center justify-center">
                  <i class="fas fa-users text-amber-700"></i>
                </div>
              </div>

              <div class="mt-5 flex items-center gap-5">
                <button
                  type="button"
                  class="relative w-24 h-24 rounded-full p-[6px] shadow-sm hover:shadow transition focus:outline-none focus:ring-2 focus:ring-amber-400"
                  :style="{ background: pieBackground(mustahiqPercent, '#f59e0b') }"
                  @click="openCompletenessModal('mustahiq')"
                  :disabled="completenessLoading"
                >
                  <div class="w-full h-full bg-white rounded-full flex items-center justify-center">
                    <div class="text-center leading-tight">
                      <div class="text-lg font-bold text-gray-900">{{ formatPercent(mustahiqPercent) }}</div>
                      <div class="text-[11px] text-gray-500">terisi</div>
                    </div>
                  </div>
                </button>
                <div class="flex-1">
                  <div class="text-sm text-gray-700">
                    <span class="font-semibold text-gray-900">{{ completeness.filled_mustahiq }}</span>
                    dari
                    <span class="font-semibold text-gray-900">{{ completeness.total_masjid }}</span>
                    masjid/langgar
                  </div>
                  <div class="text-xs text-gray-500 mt-1">Klik pie chart untuk lihat detail per masjid.</div>
                </div>
              </div>
            </div>

            <div class="rounded-2xl border border-gray-100 bg-gradient-to-br from-sky-50 to-white p-5">
              <div class="flex items-center justify-between">
                <div>
                  <p class="text-sm font-semibold text-sky-800">Data Pengurus</p>
                  <p class="text-xs text-gray-600 mt-1">Terisi jika Ketua Pengurus dan Ketua UPZ sudah diisi.</p>
                </div>
                <div class="w-10 h-10 rounded-xl bg-sky-100 flex items-center justify-center">
                  <i class="fas fa-user-tie text-sky-700"></i>
                </div>
              </div>

              <div class="mt-5 flex items-center gap-5">
                <button
                  type="button"
                  class="relative w-24 h-24 rounded-full p-[6px] shadow-sm hover:shadow transition focus:outline-none focus:ring-2 focus:ring-sky-400"
                  :style="{ background: pieBackground(pengurusPercent, '#0ea5e9') }"
                  @click="openCompletenessModal('pengurus')"
                  :disabled="completenessLoading"
                >
                  <div class="w-full h-full bg-white rounded-full flex items-center justify-center">
                    <div class="text-center leading-tight">
                      <div class="text-lg font-bold text-gray-900">{{ formatPercent(pengurusPercent) }}</div>
                      <div class="text-[11px] text-gray-500">terisi</div>
                    </div>
                  </div>
                </button>
                <div class="flex-1">
                  <div class="text-sm text-gray-700">
                    <span class="font-semibold text-gray-900">{{ completeness.filled_pengurus }}</span>
                    dari
                    <span class="font-semibold text-gray-900">{{ completeness.total_masjid }}</span>
                    masjid/langgar
                  </div>
                  <div class="text-xs text-gray-500 mt-1">Klik pie chart untuk lihat detail per masjid.</div>
                </div>
              </div>
            </div>
          </div>

          <div v-if="completenessLoading" class="mt-4 text-sm text-gray-500">Memuat data keterisian...</div>
          <div v-else-if="completenessError" class="mt-4 text-sm text-red-600">{{ completenessError }}</div>
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

    <Modal
      :show="showCompletenessModal"
      :title="completenessModalTitle"
      @close="closeCompletenessModal"
    >
      <div class="flex items-center justify-between gap-3 mb-4">
        <div class="text-sm text-gray-700">
          Total masjid/langgar:
          <span class="font-semibold text-gray-900">{{ completeness.total_masjid }}</span>
          ,
          Terisi:
          <span class="font-semibold text-gray-900">{{ completenessModalFilled }}</span>
          ({{ formatPercent(completenessModalPercent) }})
        </div>
        <button
          type="button"
          class="text-sm font-medium text-gray-700 hover:text-gray-900"
          @click="loadDataCompleteness"
          :disabled="completenessLoading"
        >
          <i :class="`fas fa-rotate ${completenessLoading ? 'animate-spin' : ''}`"></i>
        </button>
      </div>

      <div class="overflow-x-auto">
        <table class="min-w-full text-sm">
          <thead>
            <tr class="text-left text-gray-600 border-b">
              <th class="py-2 pr-4 font-semibold">Masjid/Langgar</th>
              <th v-if="selectedCompletenessType !== 'pengurus'" class="py-2 pr-4 font-semibold">Jumlah Data</th>
              <th v-else class="py-2 pr-4 font-semibold">Ketua Pengurus</th>
              <th v-if="selectedCompletenessType === 'pengurus'" class="py-2 pr-4 font-semibold">Ketua UPZ</th>
              <th class="py-2 font-semibold">Status</th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="row in completenessModalRows"
              :key="row.id"
              class="border-b last:border-b-0"
            >
              <td class="py-2 pr-4 text-gray-900">{{ row.nama }}</td>
              <td v-if="selectedCompletenessType !== 'pengurus'" class="py-2 pr-4 text-gray-700">{{ formatNumber(row.count) }}</td>
              <td v-else class="py-2 pr-4 text-gray-700">{{ row.ketua_pengurus || '-' }}</td>
              <td v-if="selectedCompletenessType === 'pengurus'" class="py-2 pr-4 text-gray-700">{{ row.ketua_upz || '-' }}</td>
              <td class="py-2">
                <span
                  :class="row.filled ? 'bg-emerald-100 text-emerald-800' : 'bg-gray-100 text-gray-700'"
                  class="inline-flex items-center px-2.5 py-1 rounded-full text-xs font-semibold"
                >
                  {{ row.filled ? 'Terisi' : 'Belum' }}
                </span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </Modal>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
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
const completenessLoading = ref(false)
const completenessError = ref('')
const showCompletenessModal = ref(false)
const selectedCompletenessType = ref('')
const completeness = ref({
  total_masjid: 0,
  filled_muzakki: 0,
  filled_mustahiq: 0,
  filled_pengurus: 0,
  items: []
})

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

const clampPercent = (value) => {
  const n = Number(value || 0)
  if (!Number.isFinite(n)) return 0
  return Math.max(0, Math.min(100, n))
}

const calcPercent = (filled, total) => {
  const f = Number(filled || 0)
  const t = Number(total || 0)
  if (!Number.isFinite(f) || !Number.isFinite(t) || t <= 0) return 0
  return clampPercent((f / t) * 100)
}

const formatPercent = (value) => `${Math.round(clampPercent(value))}%`

const formatNumber = (value) => Number(value || 0).toLocaleString('id-ID')

const pieBackground = (percent, color) => {
  const p = clampPercent(percent)
  return `conic-gradient(${color} 0 ${p}%, #e5e7eb ${p}% 100%)`
}

const muzakkiPercent = computed(() => calcPercent(completeness.value.filled_muzakki, completeness.value.total_masjid))
const mustahiqPercent = computed(() => calcPercent(completeness.value.filled_mustahiq, completeness.value.total_masjid))
const pengurusPercent = computed(() => calcPercent(completeness.value.filled_pengurus, completeness.value.total_masjid))

const completenessModalTitle = computed(() => {
  if (selectedCompletenessType.value === 'muzakki') return 'Detail Keterisian Data Muzakki'
  if (selectedCompletenessType.value === 'mustahiq') return 'Detail Keterisian Data Mustahiq'
  if (selectedCompletenessType.value === 'pengurus') return 'Detail Keterisian Data Pengurus'
  return 'Detail Keterisian Data'
})

const completenessModalFilled = computed(() => {
  if (selectedCompletenessType.value === 'muzakki') return completeness.value.filled_muzakki
  if (selectedCompletenessType.value === 'mustahiq') return completeness.value.filled_mustahiq
  if (selectedCompletenessType.value === 'pengurus') return completeness.value.filled_pengurus
  return 0
})

const completenessModalPercent = computed(() => calcPercent(completenessModalFilled.value, completeness.value.total_masjid))

const completenessModalRows = computed(() => {
  const items = Array.isArray(completeness.value.items) ? completeness.value.items : []
  if (selectedCompletenessType.value === 'muzakki') {
    return items
      .map((m) => ({
        id: m.id,
        nama: m.nama,
        count: Number(m.muzakki_count || 0),
        filled: Number(m.muzakki_count || 0) > 0
      }))
      .sort((a, b) => (b.filled - a.filled) || (b.count - a.count) || a.nama.localeCompare(b.nama))
  }
  if (selectedCompletenessType.value === 'mustahiq') {
    return items
      .map((m) => ({
        id: m.id,
        nama: m.nama,
        count: Number(m.mustahiq_count || 0),
        filled: Number(m.mustahiq_count || 0) > 0
      }))
      .sort((a, b) => (b.filled - a.filled) || (b.count - a.count) || a.nama.localeCompare(b.nama))
  }
  if (selectedCompletenessType.value === 'pengurus') {
    return items
      .map((m) => {
        const ketuaPengurus = String(m.ketua_pengurus || '').trim()
        const ketuaUpz = String(m.ketua_upz || '').trim()
        return {
          id: m.id,
          nama: m.nama,
          ketua_pengurus: ketuaPengurus,
          ketua_upz: ketuaUpz,
          filled: ketuaPengurus !== '' && ketuaUpz !== ''
        }
      })
      .sort((a, b) => (b.filled - a.filled) || a.nama.localeCompare(b.nama))
  }
  return []
})

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

const loadDataCompleteness = async () => {
  completenessLoading.value = true
  completenessError.value = ''
  try {
    const res = await api.getSuperadminDataCompleteness()
    if (res?.data?.success) {
      const d = res.data.data || {}
      completeness.value = {
        total_masjid: Number(d.total_masjid || 0),
        filled_muzakki: Number(d.filled_muzakki || 0),
        filled_mustahiq: Number(d.filled_mustahiq || 0),
        filled_pengurus: Number(d.filled_pengurus || 0),
        items: Array.isArray(d.items) ? d.items : []
      }
    } else {
      completenessError.value = res?.data?.message || 'Gagal memuat data keterisian'
    }
  } catch (error) {
    console.error(error)
    completenessError.value = 'Gagal memuat data keterisian'
  } finally {
    completenessLoading.value = false
  }
}

const openCompletenessModal = (type) => {
  selectedCompletenessType.value = type
  showCompletenessModal.value = true
}

const closeCompletenessModal = () => {
  showCompletenessModal.value = false
  selectedCompletenessType.value = ''
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

  // Load after initial stats so dashboard tetap cepat tampil
  loadDataCompleteness()
})
</script>
