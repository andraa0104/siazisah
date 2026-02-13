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
            <h1 class="text-xl md:text-2xl font-bold text-gray-800">Transaksi Zakat</h1>
          </div>
          <button @click="showModal = true" class="bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700 transition">
            <i class="fas fa-plus mr-2"></i>Input Transaksi
          </button>
        </div>
      </header>

      <main class="flex-1 overflow-y-auto p-4 md:p-6">
        <!-- Desktop Table -->
        <div class="hidden md:block bg-white rounded-xl shadow-md overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Tanggal</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Muzakki</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Jenis</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Detail</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Total</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Aksi</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="t in transaksiList" :key="t.id" class="hover:bg-gray-50">
                <td class="px-6 py-4 text-sm text-gray-900">{{ formatDate(t.tanggal_bayar) }}</td>
                <td class="px-6 py-4">
                  <div class="font-medium text-gray-900">{{ t.muzakki_nama }}</div>
                  <div class="text-sm text-gray-500">{{ t.muzakki_alamat }}</div>
                </td>
                <td class="px-6 py-4">
                  <span :class="`px-2 py-1 text-xs font-semibold rounded-full ${getZakatBadge(t.jenis_zakat)}`">
                    {{ t.jenis_zakat.toUpperCase() }}
                  </span>
                </td>
                <td class="px-6 py-4 text-sm text-gray-900">{{ getDetail(t) }}</td>
                <td class="px-6 py-4">
                  <div class="font-bold text-gray-900">{{ formatCurrency(t.total_dibayar) }}</div>
                  <div v-if="t.infaq_tambahan > 0" class="text-xs text-blue-600">+Infaq {{ formatCurrency(t.infaq_tambahan) }}</div>
                </td>
                <td class="px-6 py-4 text-sm">
                  <button @click="printReceipt(t.id)" class="text-green-600 hover:text-green-900 mr-3">
                    <i class="fas fa-print"></i>
                  </button>
                  <button @click="deleteTransaksi(t.id)" class="text-red-600 hover:text-red-900">
                    <i class="fas fa-trash"></i>
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
        <!-- Mobile List -->
        <div class="md:hidden space-y-3">
          <div v-for="t in transaksiList" :key="t.id" class="bg-white rounded-xl shadow-md p-4">
            <div class="flex justify-between items-start mb-2">
              <div>
                <p class="font-medium text-gray-900">{{ t.muzakki_nama }}</p>
                <p class="text-xs text-gray-500">{{ formatDate(t.tanggal_bayar) }}</p>
              </div>
              <span :class="`px-2 py-1 text-xs font-semibold rounded-full ${getZakatBadge(t.jenis_zakat)}`">
                {{ t.jenis_zakat.toUpperCase() }}
              </span>
            </div>
            <div class="space-y-1 text-sm mb-3">
              <div class="text-gray-600">{{ getDetail(t) }}</div>
              <div class="font-bold text-gray-900">{{ formatCurrency(t.total_dibayar) }}</div>
              <div v-if="t.infaq_tambahan > 0" class="text-xs text-blue-600">+Infaq {{ formatCurrency(t.infaq_tambahan) }}</div>
            </div>
            <div class="flex gap-2">
              <button @click="printReceipt(t.id)" class="flex-1 bg-green-50 text-green-600 py-2 rounded-lg hover:bg-green-100 text-sm">
                <i class="fas fa-print mr-1"></i>Cetak
              </button>
              <button @click="deleteTransaksi(t.id)" class="flex-1 bg-red-50 text-red-600 py-2 rounded-lg hover:bg-red-100 text-sm">
                <i class="fas fa-trash mr-1"></i>Hapus
              </button>
            </div>
          </div>
        </div>
      </main>
    </div>

    <Modal :show="showModal" title="Input Transaksi Zakat" @close="closeModal">
      <form @submit.prevent="saveTransaksi" class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">Nama Muzakki *</label>
          <input v-model="form.nama_muzakki" required class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500" placeholder="Masukkan nama muzakki">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">Jenis Zakat *</label>
          <select v-model="form.jenis_zakat" required @change="resetForm" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500">
            <option value="">Pilih Jenis Zakat</option>
            <option value="fitrah">Zakat Fitrah</option>
            <option value="mal">Zakat Mal</option>
            <option value="infaq">Infaq</option>
          </select>
        </div>
        <div v-if="form.jenis_zakat === 'fitrah'">
          <label class="block text-sm font-medium text-gray-700 mb-2">Bentuk Zakat *</label>
          <select v-model="form.bentuk_zakat" required @change="calculateTotal" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500">
            <option value="">Pilih Bentuk</option>
            <option value="uang">Uang</option>
            <option value="beras">Beras</option>
          </select>
        </div>
        <div v-if="form.jenis_zakat === 'mal'">
          <label class="block text-sm font-medium text-gray-700 mb-2">Jenis Harta *</label>
          <select v-model="form.jenis_harta" required @change="setPersentase" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500">
            <option value="">Pilih Jenis Harta</option>
            <option value="emas">Emas/Perak (2.5%)</option>
            <option value="uang">Uang/Tabungan (2.5%)</option>
            <option value="perdagangan">Perdagangan (2.5%)</option>
            <option value="pertanian_irigasi">Pertanian dengan Irigasi (5%)</option>
            <option value="pertanian_hujan">Pertanian Tadah Hujan (10%)</option>
            <option value="peternakan">Peternakan (2.5%)</option>
          </select>
        </div>
        <div v-if="form.jenis_harta">
          <label class="block text-sm font-medium text-gray-700 mb-2">Nominal Harta *</label>
          <input v-model.number="form.nominal_harta" type="number" required min="0" @input="calculateMal" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500" placeholder="Masukkan total harta">
          <p class="text-xs text-gray-500 mt-1">Persentase zakat: {{ form.persentase_zakat }}%</p>
        </div>
        <div v-if="form.bentuk_zakat === 'uang'">
          <label class="block text-sm font-medium text-gray-700 mb-2">Kelas Zakat *</label>
          <select v-model="form.kelas_zakat" required @change="calculateTotal" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500">
            <option value="">Pilih Kelas</option>
            <option value="1">Kelas 1 - Rp {{ formatNumber(kadarZakat.kelas1) }}</option>
            <option value="2">Kelas 2 - Rp {{ formatNumber(kadarZakat.kelas2) }}</option>
            <option value="3">Kelas 3 - Rp {{ formatNumber(kadarZakat.kelas3) }}</option>
          </select>
        </div>
        <div v-if="form.jenis_zakat === 'fitrah'">
          <label class="block text-sm font-medium text-gray-700 mb-2">Jumlah Orang *</label>
          <input v-model.number="form.jumlah_orang" type="number" required min="1" @input="calculateTotal" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500">
        </div>
        <div v-if="totalWajib > 0" class="bg-gray-100 p-4 rounded-lg">
          <label class="block text-sm font-medium text-gray-700 mb-2">{{ form.jenis_zakat === 'mal' ? 'Total Zakat Mal' : 'Total Wajib Zakat' }}</label>
          <div class="text-2xl font-bold text-green-600">{{ formatCurrency(totalWajib) }}</div>
          <p v-if="form.jenis_zakat === 'mal'" class="text-xs text-gray-600 mt-1">{{ formatCurrency(form.nominal_harta) }} × {{ form.persentase_zakat }}%</p>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">Total Dibayar *</label>
          <input v-model.number="form.total_dibayar" type="number" required min="0" @input="calculateInfaq" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500">
        </div>
        <div v-if="infaqTambahan > 0" class="bg-blue-50 border border-blue-200 rounded-lg p-4">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-blue-700 font-medium">Infaq Tambahan</p>
              <p class="text-xs text-blue-600">Kelebihan dari total wajib</p>
            </div>
            <p class="text-2xl font-bold text-blue-700">{{ formatCurrency(infaqTambahan) }}</p>
          </div>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">Tanggal Bayar *</label>
          <input v-model="form.tanggal_bayar" type="date" required class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500">
        </div>
        <div class="flex gap-3 pt-4">
          <button type="submit" class="flex-1 bg-green-600 text-white py-3 rounded-lg hover:bg-green-700 transition font-medium">
            <i class="fas fa-save mr-2"></i>Simpan
          </button>
          <button type="button" @click="closeModal" class="flex-1 bg-gray-200 text-gray-700 py-3 rounded-lg hover:bg-gray-300 transition font-medium">Batal</button>
        </div>
      </form>
    </Modal>

    <Toast :message="toast.message" :type="toast.type" />
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import Sidebar from '../../components/Sidebar.vue'
import Modal from '../../components/Modal.vue'
import Toast from '../../components/Toast.vue'
import api from '../../api'

const menuItems = [
  { path: '/petugas', label: 'Dashboard', icon: 'fas fa-home' },
  { path: '/petugas/pengaturan', label: 'Pengaturan', icon: 'fas fa-cog' },
  { path: '/petugas/mustahiq', label: 'Mustahiq', icon: 'fas fa-hand-holding-heart' },
  { path: '/petugas/transaksi', label: 'Transaksi', icon: 'fas fa-receipt' }
]

const showSidebar = ref(false)

const transaksiList = ref([])
const showModal = ref(false)
const form = ref({ 
  nama_muzakki: '', 
  jenis_zakat: '', 
  bentuk_zakat: '', 
  jenis_harta: '',
  nominal_harta: 0,
  persentase_zakat: 0,
  kelas_zakat: '', 
  jumlah_orang: 1, 
  total_dibayar: 0, 
  tanggal_bayar: new Date().toISOString().split('T')[0] 
})
const toast = ref({ message: '', type: 'success' })
const kadarZakat = ref({ kelas1: 35000, kelas2: 45000, kelas3: 55000 })

const totalWajib = computed(() => {
  if (form.value.jenis_zakat === 'fitrah' && form.value.bentuk_zakat === 'uang' && form.value.kelas_zakat && form.value.jumlah_orang) {
    const nominal = kadarZakat.value[`kelas${form.value.kelas_zakat}`]
    return nominal * form.value.jumlah_orang
  }
  if (form.value.jenis_zakat === 'mal' && form.value.nominal_harta && form.value.persentase_zakat) {
    return form.value.nominal_harta * (form.value.persentase_zakat / 100)
  }
  return 0
})

const infaqTambahan = computed(() => {
  if (totalWajib.value > 0 && form.value.total_dibayar > totalWajib.value) {
    return form.value.total_dibayar - totalWajib.value
  }
  return 0
})

const formatCurrency = (amount) => new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR', minimumFractionDigits: 0 }).format(amount)
const formatNumber = (num) => new Intl.NumberFormat('id-ID').format(num)
const formatDate = (date) => new Date(date).toLocaleDateString('id-ID')

const getZakatBadge = (jenis) => {
  const badges = { fitrah: 'bg-green-100 text-green-800', mal: 'bg-blue-100 text-blue-800', infaq: 'bg-purple-100 text-purple-800' }
  return badges[jenis] || 'bg-gray-100 text-gray-800'
}

const getDetail = (t) => {
  if (t.jenis_zakat === 'fitrah') return `${t.bentuk_zakat} - Kelas ${t.kelas_zakat} - ${t.jumlah_orang} orang`
  if (t.jenis_zakat === 'mal') return `${t.jenis_harta} - ${formatCurrency(t.nominal_harta)} (${t.persentase_zakat}%)`
  return '-'
}

const resetForm = () => {
  form.value.bentuk_zakat = ''
  form.value.jenis_harta = ''
  form.value.nominal_harta = 0
  form.value.persentase_zakat = 0
  form.value.kelas_zakat = ''
  form.value.jumlah_orang = 1
  form.value.total_dibayar = 0
}

const setPersentase = () => {
  const persentaseMap = {
    emas: 2.5, uang: 2.5, perdagangan: 2.5, peternakan: 2.5,
    pertanian_irigasi: 5, pertanian_hujan: 10
  }
  form.value.persentase_zakat = persentaseMap[form.value.jenis_harta] || 0
  calculateMal()
}

const calculateMal = () => {
  if (form.value.nominal_harta && form.value.persentase_zakat) {
    const total = form.value.nominal_harta * (form.value.persentase_zakat / 100)
    form.value.total_dibayar = total
  }
}

const calculateTotal = () => {
  if (totalWajib.value > 0) {
    form.value.total_dibayar = totalWajib.value
  }
}

const calculateInfaq = () => {}

const loadTransaksi = async () => {
  try {
    const { data } = await api.getTransaksi()
    if (data.success) transaksiList.value = data.data || []
  } catch (error) {
    console.error(error)
  }
}

const deleteTransaksi = async (id) => {
  if (!confirm('Yakin ingin menghapus?')) return
  try {
    const { data } = await api.deleteTransaksi(id)
    if (data.success) {
      toast.value = { message: 'Transaksi berhasil dihapus', type: 'success' }
      loadTransaksi()
    }
  } catch (error) {
    toast.value = { message: 'Gagal menghapus transaksi', type: 'error' }
  }
}

const printReceipt = async (id) => {
  try {
    const token = localStorage.getItem('token')
    const response = await fetch(`http://localhost:8082/api/petugas/transaksi/${id}/print`, {
      headers: { 'Authorization': `Bearer ${token}` }
    })
    const html = await response.text()
    const printWindow = window.open('', '_blank')
    printWindow.document.write(html)
    printWindow.document.close()
  } catch (error) {
    toast.value = { message: 'Gagal mencetak bukti', type: 'error' }
  }
}

const saveTransaksi = async () => {
  try {
    const muzakkiData = await api.createMuzakki({ nama: form.value.nama_muzakki, alamat: '-', telepon: '-' })
    if (!muzakkiData.data.success) {
      toast.value = { message: 'Gagal menyimpan data muzakki', type: 'error' }
      return
    }
    
    const payload = {
      muzakki_id: muzakkiData.data.data.id,
      jenis_zakat: form.value.jenis_zakat,
      tanggal_bayar: form.value.tanggal_bayar,
      total_dibayar: form.value.total_dibayar,
      keterangan: ''
    }
    
    if (form.value.jenis_zakat === 'fitrah') {
      payload.bentuk_zakat = form.value.bentuk_zakat
      payload.jumlah_orang = form.value.jumlah_orang
      if (form.value.bentuk_zakat === 'uang') {
        payload.kelas_zakat = form.value.kelas_zakat
        payload.nominal_per_orang = kadarZakat.value[`kelas${form.value.kelas_zakat}`]
      }
    } else if (form.value.jenis_zakat === 'mal') {
      payload.jenis_harta = form.value.jenis_harta
      payload.nominal_harta = form.value.nominal_harta
      payload.persentase_zakat = form.value.persentase_zakat
      payload.jumlah_orang = 0
    }
    
    const { data } = await api.createTransaksi(payload)
    if (data.success) {
      toast.value = { message: 'Transaksi berhasil disimpan', type: 'success' }
      closeModal()
      loadTransaksi()
    }
  } catch (error) {
    toast.value = { message: 'Gagal menyimpan transaksi', type: 'error' }
  }
}

const closeModal = () => {
  showModal.value = false
  form.value = { 
    nama_muzakki: '', 
    jenis_zakat: '', 
    bentuk_zakat: '', 
    jenis_harta: '',
    nominal_harta: 0,
    persentase_zakat: 0,
    kelas_zakat: '', 
    jumlah_orang: 1, 
    total_dibayar: 0, 
    tanggal_bayar: new Date().toISOString().split('T')[0] 
  }
}

onMounted(() => {
  loadTransaksi()
  const saved = localStorage.getItem('kadar_zakat')
  if (saved) kadarZakat.value = JSON.parse(saved)
})
</script>
