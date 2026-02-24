<template>
  <div class="flex h-screen overflow-hidden bg-gray-50">
    <Sidebar :show="showSidebar" @close="showSidebar = false" :menuItems="menuItems" />

    <div class="flex-1 flex flex-col overflow-hidden">
      <header class="bg-white shadow-sm">
        <div class="px-4 md:px-6 py-4 flex justify-between items-center">
          <div class="flex items-center">
            <button @click="showSidebar = true" class="mr-3 md:hidden text-gray-600 hover:text-gray-800">
              <i class="fas fa-bars text-xl"></i>
            </button>
            <div>
              <h1 class="text-xl md:text-2xl font-bold text-gray-800">Distribusi Mustahiq</h1>
              <p class="text-sm text-gray-600">Fokus pembagian beras & uang dari fitrah, fidyah, dan zakat mal</p>
            </div>
          </div>
          <button @click="loadData" class="bg-white border border-gray-300 text-gray-700 px-3 py-2 rounded-lg hover:bg-gray-50 text-sm">
            <i class="fas fa-sync-alt mr-2"></i>Refresh
          </button>
        </div>
      </header>

      <main class="flex-1 overflow-y-auto p-4 md:p-6 space-y-4 md:space-y-6">
        <template v-if="isLoading">
          <SkeletonCard v-for="i in 4" :key="i" type="list-item" />
        </template>

        <template v-else>
          <div class="bg-gradient-to-r from-blue-600 to-indigo-600 rounded-2xl p-4 md:p-5 text-white shadow-md">
            <h2 class="text-lg font-bold">Ringkasan Distribusi</h2>
            <p class="text-sm text-blue-100 mt-1">Infaq tidak masuk pembagian mustahiq, otomatis dianggap kas masjid/langgar.</p>
            <div class="grid grid-cols-1 sm:grid-cols-3 gap-3 mt-4">
              <div class="bg-white bg-opacity-15 rounded-lg p-3">
                <p class="text-xs text-blue-100">Mustahiq Aktif</p>
                <p class="text-2xl font-bold mt-1">{{ insight.total_mustahiq_aktif }}</p>
              </div>
              <div class="bg-white bg-opacity-15 rounded-lg p-3">
                <p class="text-xs text-blue-100">Pool Beras</p>
                <p class="text-2xl font-bold mt-1">{{ formatKg(insight.total_pool_beras_kg) }}</p>
              </div>
              <div class="bg-white bg-opacity-15 rounded-lg p-3">
                <p class="text-xs text-blue-100">Pool Uang</p>
                <p class="text-2xl font-bold mt-1">{{ formatCurrency(insight.total_pool_uang) }}</p>
              </div>
            </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div class="bg-white rounded-xl shadow-md p-4 border border-gray-100">
              <p class="text-sm text-gray-500">Rekomendasi per Mustahiq</p>
              <div class="mt-3 space-y-2">
                <div class="flex items-center justify-between bg-amber-50 rounded-lg px-3 py-2 border border-amber-100">
                  <span class="text-sm text-amber-700">Beras</span>
                  <span class="font-bold text-amber-700">{{ formatKg(insight.rekomendasi_beras_per_orang) }}</span>
                </div>
                <div class="flex items-center justify-between bg-green-50 rounded-lg px-3 py-2 border border-green-100">
                  <span class="text-sm text-green-700">Uang</span>
                  <span class="font-bold text-green-700">{{ formatCurrency(insight.rekomendasi_uang_per_orang) }}</span>
                </div>
              </div>
            </div>

            <div :class="`rounded-xl shadow-md p-4 border ${statusCardClass}`">
              <p class="text-sm">Status Sisa Saat Ini</p>
              <div class="mt-3 space-y-2">
                <div class="flex items-center justify-between">
                  <span class="text-sm">Sisa Beras</span>
                  <span class="font-bold">{{ formatKg(insight.sisa_beras_kg) }}</span>
                </div>
                <div class="flex items-center justify-between">
                  <span class="text-sm">Sisa Uang</span>
                  <span class="font-bold">{{ formatCurrency(insight.sisa_uang) }}</span>
                </div>
              </div>
            </div>
          </div>

          <div class="bg-white rounded-xl shadow-md p-4 md:p-5">
            <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-3 mb-4">
              <h2 class="text-lg font-bold text-gray-800">Rencana Distribusi</h2>
              <div class="flex items-center gap-2 text-xs text-gray-500">
                <i class="fas fa-info-circle"></i>
                <span>Satu skema pembagian untuk semua sumber zakat non-infaq</span>
              </div>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-3 gap-2 mb-4 bg-gray-50 p-2 rounded-xl">
              <button @click="mode = 'rekomendasi'" :class="modeButtonClass('rekomendasi')">
                <i class="fas fa-magic mr-2"></i>Rekomendasi
              </button>
              <button @click="mode = 'rata'" :class="modeButtonClass('rata')">
                <i class="fas fa-equals mr-2"></i>Sama Rata
              </button>
              <button @click="mode = 'berbeda'" :class="modeButtonClass('berbeda')">
                <i class="fas fa-sliders-h mr-2"></i>Per Orang
              </button>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-3 gap-3 mb-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">Tanggal Distribusi</label>
                <input v-model="tanggalDistribusi" type="date" class="w-full px-3 py-2 border border-gray-300 rounded-lg" />
              </div>
              <div class="md:col-span-2">
                <label class="block text-sm font-medium text-gray-700 mb-1">Catatan Distribusi</label>
                <input v-model="catatanDistribusi" class="w-full px-3 py-2 border border-gray-300 rounded-lg" placeholder="Contoh: Prioritas fakir lansia" />
              </div>
            </div>

            <div v-if="mode === 'rata'" class="grid grid-cols-1 md:grid-cols-2 gap-3 mb-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">Beras per Mustahiq (kg)</label>
                <input v-model.number="manualRataBeras" type="number" step="0.1" min="0" class="w-full px-3 py-2 border border-gray-300 rounded-lg" />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">Uang per Mustahiq (Rp)</label>
                <input v-model.number="manualRataUang" type="number" min="0" class="w-full px-3 py-2 border border-gray-300 rounded-lg" />
              </div>
            </div>

            <div class="flex flex-wrap gap-2 mb-4">
              <button
                v-if="mode === 'rekomendasi'"
                @click="applyRekomendasi"
                class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 text-sm"
              >
                <i class="fas fa-magic mr-2"></i>Terapkan Rekomendasi
              </button>

              <button
                v-if="mode === 'rata'"
                @click="applyManualRata"
                class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 text-sm"
              >
                <i class="fas fa-equals mr-2"></i>Terapkan Sama Rata
              </button>

              <button
                @click="resetAllocations"
                class="bg-gray-100 text-gray-700 px-4 py-2 rounded-lg hover:bg-gray-200 text-sm"
              >
                <i class="fas fa-undo mr-2"></i>Reset
              </button>
            </div>

            <div v-if="allocations.length === 0" class="text-center text-gray-500 py-6">
              Data mustahiq belum tersedia.
            </div>

            <div v-else class="space-y-2 max-h-80 overflow-y-auto pr-1">
              <div
                v-for="item in allocations"
                :key="item.mustahiq_id"
                class="bg-gray-50 rounded-lg p-3 border border-gray-200"
              >
                <div class="flex items-center justify-between mb-2">
                  <div>
                    <p class="font-semibold text-gray-800">{{ item.nama }}</p>
                    <p class="text-xs text-gray-500 capitalize">{{ item.jenis_penerima || '-' }}</p>
                  </div>
                </div>
                <div class="grid grid-cols-1 sm:grid-cols-2 gap-2">
                  <div>
                    <label class="block text-xs text-gray-600 mb-1">Beras (kg)</label>
                    <input v-model.number="item.beras_kg" type="number" step="0.1" min="0" class="w-full px-3 py-2 border border-gray-300 rounded-lg" />
                  </div>
                  <div>
                    <label class="block text-xs text-gray-600 mb-1">Uang (Rp)</label>
                    <input v-model.number="item.nominal" type="number" min="0" class="w-full px-3 py-2 border border-gray-300 rounded-lg" />
                  </div>
                </div>
              </div>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-3 mt-4">
              <div class="bg-blue-50 rounded-lg p-4 border border-blue-200">
                <p class="text-sm text-blue-700">Total Rencana Dibagikan</p>
                <p class="text-sm mt-1">Beras: <span class="font-bold">{{ formatKg(totalRencanaBeras) }}</span></p>
                <p class="text-sm">Uang: <span class="font-bold">{{ formatCurrency(totalRencanaUang) }}</span></p>
              </div>
              <div :class="`rounded-lg p-4 border ${sisaBoxClass}`">
                <p class="text-sm">Proyeksi Sisa Setelah Rencana</p>
                <p class="text-sm mt-1">Beras: <span class="font-bold">{{ formatKg(proyeksiSisaBeras) }}</span></p>
                <p class="text-sm">Uang: <span class="font-bold">{{ formatCurrency(proyeksiSisaUang) }}</span></p>
              </div>
            </div>

            <button
              @click="saveRencanaDistribusi"
              :disabled="isSaving"
              class="mt-4 w-full md:w-auto bg-green-600 text-white px-5 py-3 rounded-lg hover:bg-green-700 disabled:opacity-50"
            >
              <i class="fas fa-save mr-2"></i>{{ isSaving ? 'Menyimpan...' : 'Simpan ke Realisasi' }}
            </button>
          </div>

          <div class="bg-white rounded-xl shadow-md p-4 md:p-5">
            <h2 class="text-lg font-bold text-gray-800 mb-4">Riwayat Distribusi</h2>

            <div v-if="distribusiList.length === 0" class="text-center text-gray-500 py-4">
              Belum ada data distribusi.
            </div>

            <div v-else class="space-y-3">
              <div v-for="item in distribusiList" :key="item.id" class="border border-gray-200 rounded-lg p-3">
                <div class="flex justify-between items-start gap-3">
                  <div>
                    <p class="font-semibold text-gray-800">{{ item.mustahiq_nama }}</p>
                    <p class="text-xs text-gray-500 capitalize">{{ item.jenis_penerima || '-' }}</p>
                    <p class="text-xs text-gray-500">{{ formatTanggalJam(item.tanggal_distribusi, item.waktu_input, item.created_at) }}</p>
                    <p class="text-sm mt-1">Beras: <span class="font-semibold">{{ formatKg(item.beras_kg || 0) }}</span> | Uang: <span class="font-semibold">{{ formatCurrency(item.nominal || 0) }}</span></p>
                    <p v-if="item.keterangan" class="text-xs text-gray-600 mt-1">{{ item.keterangan }}</p>
                  </div>
                  <div class="flex gap-2">
                    <button @click="openEdit(item)" class="px-3 py-1.5 rounded-lg bg-blue-50 text-blue-600 hover:bg-blue-100 text-sm">
                      <i class="fas fa-edit"></i>
                    </button>
                    <button @click="deleteDistribusi(item.id)" class="px-3 py-1.5 rounded-lg bg-red-50 text-red-600 hover:bg-red-100 text-sm">
                      <i class="fas fa-trash"></i>
                    </button>
                  </div>
                </div>
              </div>
            </div>

            <div class="mt-4 flex flex-col md:flex-row md:items-center md:justify-between gap-3 text-sm text-gray-600">
              <div>
                Menampilkan {{ distribusiList.length }} dari {{ distribusiPagination.total }} data
              </div>
              <div class="flex flex-wrap items-center gap-2">
                <label class="text-xs text-gray-500">Per halaman</label>
                <select v-model="distribusiPagination.pageSize" @change="onDistribusiPageSizeChange" class="px-2 py-1 border border-gray-300 rounded-lg text-sm">
                  <option v-for="size in pageSizeOptions" :key="size" :value="size">
                    {{ size === 'all' ? 'Semua' : size }}
                  </option>
                </select>
                <button
                  class="px-3 py-1 border border-gray-300 rounded-lg disabled:opacity-50"
                  :disabled="distribusiPagination.isAll || distribusiPagination.page <= 1"
                  @click="setDistribusiPage(distribusiPagination.page - 1)"
                >
                  Prev
                </button>
                <span class="text-xs text-gray-500">{{ distribusiPagination.page }} / {{ distribusiPagination.totalPages }}</span>
                <button
                  class="px-3 py-1 border border-gray-300 rounded-lg disabled:opacity-50"
                  :disabled="distribusiPagination.isAll || distribusiPagination.page >= distribusiPagination.totalPages"
                  @click="setDistribusiPage(distribusiPagination.page + 1)"
                >
                  Next
                </button>
              </div>
            </div>
          </div>
        </template>
      </main>
    </div>

    <Modal :show="showEditModal" title="Edit Distribusi" @close="showEditModal = false">
      <form @submit.prevent="saveEdit" class="space-y-3">
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-2">
          <div>
            <label class="block text-sm text-gray-700 mb-1">Beras (kg)</label>
            <input v-model.number="editForm.beras_kg" type="number" step="0.1" min="0" class="w-full px-3 py-2 border border-gray-300 rounded-lg" />
          </div>
          <div>
            <label class="block text-sm text-gray-700 mb-1">Uang (Rp)</label>
            <input v-model.number="editForm.nominal" type="number" min="0" class="w-full px-3 py-2 border border-gray-300 rounded-lg" />
          </div>
        </div>
        <div>
          <label class="block text-sm text-gray-700 mb-1">Tanggal</label>
          <input v-model="editForm.tanggal_distribusi" type="date" class="w-full px-3 py-2 border border-gray-300 rounded-lg" />
        </div>
        <div>
          <label class="block text-sm text-gray-700 mb-1">Catatan</label>
          <input v-model="editForm.keterangan" class="w-full px-3 py-2 border border-gray-300 rounded-lg" />
        </div>
        <div class="flex gap-2 pt-2">
          <button type="submit" class="flex-1 bg-green-600 text-white py-2 rounded-lg hover:bg-green-700">Simpan</button>
          <button type="button" @click="showEditModal = false" class="flex-1 bg-gray-200 text-gray-700 py-2 rounded-lg hover:bg-gray-300">Batal</button>
        </div>
      </form>
    </Modal>

    <Toast :message="toast.message" :type="toast.type" />
    <LoadingOverlay :show="isSaving" message="Menyimpan distribusi..." />
  </div>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue'
import Swal from 'sweetalert2'
import Sidebar from '../../components/Sidebar.vue'
import Modal from '../../components/Modal.vue'
import Toast from '../../components/Toast.vue'
import LoadingOverlay from '../../components/LoadingOverlay.vue'
import SkeletonCard from '../../components/SkeletonCard.vue'
import api from '../../api'

const menuItems = [
  { path: '/petugas', label: 'Dashboard', icon: 'fas fa-home' },
  { path: '/petugas/pengaturan', label: 'Pengaturan', icon: 'fas fa-cog' },
  { path: '/petugas/mustahiq', label: 'Mustahiq', icon: 'fas fa-hand-holding-heart' },
  { path: '/petugas/transaksi', label: 'Transaksi', icon: 'fas fa-receipt' },
  { path: '/petugas/distribusi', label: 'Distribusi', icon: 'fas fa-box-open' }
]

const showSidebar = ref(false)
const isLoading = ref(true)
const isSaving = ref(false)
const toast = ref({ message: '', type: 'success' })
const pageSizeOptions = [5, 10, 25, 50, 'all']
const distribusiPagination = ref({ page: 1, pageSize: 5, total: 0, totalPages: 1, isAll: false })

const insight = ref({
  total_mustahiq_aktif: 0,
  total_pool_beras_kg: 0,
  total_pool_uang: 0,
  sisa_beras_kg: 0,
  sisa_uang: 0,
  rekomendasi_beras_per_orang: 0,
  rekomendasi_uang_per_orang: 0
})

const mustahiqList = ref([])
const distribusiList = ref([])
const allocations = ref([])

const mode = ref('rekomendasi')
const tanggalDistribusi = ref(new Date().toISOString().split('T')[0])
const catatanDistribusi = ref('')
const manualRataBeras = ref(0)
const manualRataUang = ref(0)

const showEditModal = ref(false)
const editForm = ref({
  id: null,
  mustahiq_id: null,
  nominal: 0,
  beras_kg: 0,
  tanggal_distribusi: new Date().toISOString().split('T')[0],
  keterangan: ''
})

const formatCurrency = (amount) => new Intl.NumberFormat('id-ID', {
  style: 'currency',
  currency: 'IDR',
  minimumFractionDigits: 0
}).format(Number(amount || 0))

const formatKg = (kg) => `${Number(kg || 0).toLocaleString('id-ID', { maximumFractionDigits: 2 })} kg`

const formatTanggalJam = (tanggal, jam, createdAt) => {
  if (!tanggal && !jam) return '-'

  const rawTanggal = String(tanggal || '').trim()
  let tanggalText = '-'
  let jamText = String(jam || '').trim().slice(0, 8)

  const rawCreatedAt = String(createdAt || '').trim()
  const createdAtMatch = rawCreatedAt.match(/(\d{2}:\d{2}:\d{2})/)
  const createdAtJam = createdAtMatch ? createdAtMatch[1] : ''

  if (!jamText || jamText === '00:00:00') {
    jamText = createdAtJam
  }

  const parsed = rawTanggal ? new Date(rawTanggal.replace(' ', 'T')) : null
  if (parsed && !Number.isNaN(parsed.getTime())) {
    tanggalText = new Intl.DateTimeFormat('id-ID', {
      day: '2-digit',
      month: 'short',
      year: 'numeric'
    }).format(parsed)

    if (!jamText) {
      jamText = new Intl.DateTimeFormat('id-ID', {
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit',
        hour12: false
      }).format(parsed).replace(/\./g, ':')
    }
  } else if (rawTanggal) {
    tanggalText = rawTanggal.split('T')[0].split(' ')[0]
  }

  if (!jamText && rawTanggal) {
    const match = rawTanggal.match(/(\d{2}:\d{2}:\d{2})/)
    if (match) jamText = match[1]
  }

  if (jamText === '00:00:00') {
    jamText = ''
  }

  return jamText ? `${tanggalText} • ${jamText}` : tanggalText
}

const modeButtonClass = (modeName) => {
  return mode.value === modeName
    ? 'px-4 py-2 rounded-lg bg-blue-600 text-white text-sm'
    : 'px-4 py-2 rounded-lg bg-white text-gray-700 hover:bg-gray-100 text-sm border border-gray-200'
}

const setDistribusiPage = (page) => {
  distribusiPagination.value.page = page
  loadData()
}

const onDistribusiPageSizeChange = () => {
  distribusiPagination.value.page = 1
  loadData()
}

const syncDistribusiPagination = (payload, itemCount) => {
  if (payload && payload.pagination) {
    const meta = payload.pagination
    distribusiPagination.value = {
      page: meta.page || 1,
      pageSize: meta.is_all ? 'all' : (meta.page_size || distribusiPagination.value.pageSize),
      total: meta.total ?? itemCount,
      totalPages: meta.total_pages || 1,
      isAll: !!meta.is_all
    }
    return
  }
  distribusiPagination.value.total = itemCount
  distribusiPagination.value.totalPages = 1
  distribusiPagination.value.isAll = distribusiPagination.value.pageSize === 'all'
}

const applyDistribusiLocalPagination = (items) => {
  const total = items.length
  if (distribusiPagination.value.pageSize === 'all') {
    distribusiPagination.value.total = total
    distribusiPagination.value.totalPages = 1
    distribusiPagination.value.isAll = true
    return items
  }

  const pageSize = Number(distribusiPagination.value.pageSize) || 5
  const totalPages = Math.max(1, Math.ceil(total / pageSize))
  const page = Math.min(distribusiPagination.value.page, totalPages)
  const start = (page - 1) * pageSize
  const sliced = items.slice(start, start + pageSize)

  distribusiPagination.value.total = total
  distribusiPagination.value.totalPages = totalPages
  distribusiPagination.value.page = page
  distribusiPagination.value.isAll = false

  return sliced
}

const totalRencanaBeras = computed(() => allocations.value.reduce((sum, item) => sum + Number(item.beras_kg || 0), 0))
const totalRencanaUang = computed(() => allocations.value.reduce((sum, item) => sum + Number(item.nominal || 0), 0))

const proyeksiSisaBeras = computed(() => Number(insight.value.sisa_beras_kg || 0) - totalRencanaBeras.value)
const proyeksiSisaUang = computed(() => Number(insight.value.sisa_uang || 0) - totalRencanaUang.value)

const statusCardClass = computed(() => {
  if (Number(insight.value.sisa_beras_kg) < 0 || Number(insight.value.sisa_uang) < 0) return 'bg-red-50 border-red-200 text-red-700'
  if (Number(insight.value.sisa_beras_kg) === 0 && Number(insight.value.sisa_uang) === 0) return 'bg-gray-50 border-gray-300 text-gray-700'
  return 'bg-green-50 border-green-200 text-green-700'
})

const sisaBoxClass = computed(() => {
  if (proyeksiSisaBeras.value < 0 || proyeksiSisaUang.value < 0) return 'bg-red-50 border-red-200 text-red-700'
  if (proyeksiSisaBeras.value === 0 && proyeksiSisaUang.value === 0) return 'bg-gray-50 border-gray-300 text-gray-700'
  return 'bg-green-50 border-green-200 text-green-700'
})

const mapMustahiqToAllocations = () => {
  allocations.value = mustahiqList.value.map(m => ({
    mustahiq_id: m.id,
    nama: m.nama,
    jenis_penerima: m.jenis_penerima,
    beras_kg: 0,
    nominal: 0
  }))
}

const applyRekomendasi = () => {
  allocations.value = allocations.value.map(item => ({
    ...item,
    beras_kg: Number(insight.value.rekomendasi_beras_per_orang || 0),
    nominal: Number(insight.value.rekomendasi_uang_per_orang || 0)
  }))
}

const applyManualRata = () => {
  allocations.value = allocations.value.map(item => ({
    ...item,
    beras_kg: Number(manualRataBeras.value || 0),
    nominal: Number(manualRataUang.value || 0)
  }))
}

const resetAllocations = () => {
  allocations.value = allocations.value.map(item => ({ ...item, beras_kg: 0, nominal: 0 }))
}

const loadData = async () => {
  isLoading.value = true
  try {
    const distribusiParams = distribusiPagination.value.pageSize === 'all'
      ? { page_size: 'all' }
      : { page: distribusiPagination.value.page, page_size: distribusiPagination.value.pageSize }
    const [insightRes, mustahiqRes, distribusiRes] = await Promise.all([
      api.getDistribusiInsight(),
      api.getMustahiq({ page_size: 'all' }),
      api.getDistribusi(distribusiParams)
    ])

    if (insightRes.data.success) insight.value = insightRes.data.data
    if (mustahiqRes.data.success) {
      const payload = mustahiqRes.data.data || {}
      const items = Array.isArray(payload.items) ? payload.items : dataOrEmpty(payload)
      mustahiqList.value = items
      mapMustahiqToAllocations()
    }
    if (distribusiRes.data.success) {
      const payload = distribusiRes.data.data || {}
      if (Array.isArray(payload.items)) {
        distribusiList.value = payload.items
        syncDistribusiPagination(payload, payload.items.length)
      } else if (Array.isArray(payload)) {
        distribusiList.value = applyDistribusiLocalPagination(payload)
      } else {
        distribusiList.value = []
        syncDistribusiPagination(payload, 0)
      }
    }
  } catch (error) {
    toast.value = { message: 'Gagal memuat data distribusi', type: 'error' }
  } finally {
    isLoading.value = false
  }
}

const dataOrEmpty = (value) => Array.isArray(value) ? value : []

const saveRencanaDistribusi = async () => {
  const rows = allocations.value.filter(item => Number(item.beras_kg || 0) > 0 || Number(item.nominal || 0) > 0)
  if (rows.length === 0) {
    toast.value = { message: 'Tidak ada data rencana untuk disimpan', type: 'error' }
    return
  }

  isSaving.value = true
  try {
    for (const row of rows) {
      await api.createDistribusi({
        mustahiq_id: row.mustahiq_id,
        nominal: Number(row.nominal || 0),
        beras_kg: Number(row.beras_kg || 0),
        tanggal_distribusi: tanggalDistribusi.value,
        mode: mode.value,
        keterangan: catatanDistribusi.value
      })
    }

    toast.value = { message: 'Rencana distribusi berhasil disimpan', type: 'success' }
    resetAllocations()
    catatanDistribusi.value = ''
    await loadData()
  } catch (error) {
    toast.value = { message: 'Gagal menyimpan distribusi', type: 'error' }
  } finally {
    isSaving.value = false
  }
}

const openEdit = (item) => {
  editForm.value = {
    id: item.id,
    mustahiq_id: item.mustahiq_id,
    nominal: Number(item.nominal || 0),
    beras_kg: Number(item.beras_kg || 0),
    tanggal_distribusi: item.tanggal_distribusi,
    keterangan: item.keterangan || ''
  }
  showEditModal.value = true
}

const saveEdit = async () => {
  isSaving.value = true
  try {
    await api.updateDistribusi(editForm.value.id, {
      mustahiq_id: editForm.value.mustahiq_id,
      nominal: Number(editForm.value.nominal || 0),
      beras_kg: Number(editForm.value.beras_kg || 0),
      tanggal_distribusi: editForm.value.tanggal_distribusi,
      mode: 'manual',
      keterangan: editForm.value.keterangan
    })
    toast.value = { message: 'Distribusi berhasil diupdate', type: 'success' }
    showEditModal.value = false
    await loadData()
  } catch (error) {
    toast.value = { message: 'Gagal update distribusi', type: 'error' }
  } finally {
    isSaving.value = false
  }
}

const deleteDistribusi = async (id) => {
  const result = await Swal.fire({
    title: 'Hapus Distribusi?',
    text: 'Data distribusi yang dihapus tidak bisa dikembalikan.',
    icon: 'warning',
    showCancelButton: true,
    confirmButtonText: 'Ya, Hapus',
    cancelButtonText: 'Batal'
  })
  if (!result.isConfirmed) return

  try {
    await api.deleteDistribusi(id)
    toast.value = { message: 'Distribusi berhasil dihapus', type: 'success' }
    await loadData()
  } catch (error) {
    toast.value = { message: 'Gagal menghapus distribusi', type: 'error' }
  }
}

onMounted(loadData)
</script>
