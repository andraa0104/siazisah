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
            <h1 class="text-xl md:text-2xl font-bold text-gray-800">Kelola Mustahiq</h1>
          </div>
          <button @click="showModal = true" class="bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700 transition">
            <i class="fas fa-plus mr-2"></i>Tambah Mustahiq
          </button>
        </div>
      </header>

      <main class="flex-1 overflow-y-auto p-4 md:p-6">
        <!-- Desktop Table -->
        <div class="hidden md:block bg-white rounded-xl shadow-md overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Nama</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Jenis</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Alamat</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Lokasi</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">RT</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Aksi</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <!-- Loading Skeleton -->
              <template v-if="isLoading">
                <tr v-for="i in 5" :key="i">
                  <td colspan="6" class="px-6 py-4">
                    <SkeletonCard type="table-row" />
                  </td>
                </tr>
              </template>
              
              <!-- Actual Data -->
              <tr v-else v-for="m in mustahiqList" :key="m.id" class="hover:bg-gray-50">
                <td class="px-6 py-4"><div class="font-medium text-gray-900">{{ m.nama }}</div></td>
                <td class="px-6 py-4"><span class="px-2 py-1 text-xs font-semibold rounded-full bg-blue-100 text-blue-800">{{ m.jenis_penerima }}</span></td>
                <td class="px-6 py-4 text-sm text-gray-900">{{ m.alamat }}</td>
                <td class="px-6 py-4 text-sm text-gray-900">{{ m.lokasi === 'dalam_desa' ? 'Dalam Desa' : 'Luar Desa' }}</td>
                <td class="px-6 py-4 text-sm text-gray-900">RT {{ m.rt }}</td>
                <td class="px-6 py-4 text-sm">
                  <button @click="editMustahiq(m)" class="text-blue-600 hover:text-blue-900 mr-3">
                    <i class="fas fa-edit"></i>
                  </button>
                  <button @click="deleteMustahiq(m.id)" class="text-red-600 hover:text-red-900">
                    <i class="fas fa-trash"></i>
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
        <!-- Mobile List -->
        <div class="md:hidden space-y-3">
          <!-- Loading Skeleton -->
          <template v-if="isLoading">
            <SkeletonCard v-for="i in 5" :key="i" type="list-item" />
          </template>
          
          <!-- Actual Data -->
          <div v-else v-for="m in mustahiqList" :key="m.id" class="bg-white rounded-xl shadow-md p-4">
            <div class="flex justify-between items-start mb-2">
              <p class="font-medium text-gray-900">{{ m.nama }}</p>
              <span class="px-2 py-1 text-xs font-semibold rounded-full bg-blue-100 text-blue-800">{{ m.jenis_penerima }}</span>
            </div>
            <div class="space-y-1 text-sm mb-3">
              <p class="text-gray-600">{{ m.alamat }}</p>
              <p class="text-gray-600">{{ m.lokasi === 'dalam_desa' ? 'Dalam Desa' : 'Luar Desa' }} - RT {{ m.rt }}</p>
            </div>
            <div class="flex gap-2">
              <button @click="editMustahiq(m)" class="flex-1 bg-blue-50 text-blue-600 py-2 rounded-lg hover:bg-blue-100 text-sm">
                <i class="fas fa-edit mr-1"></i>Edit
              </button>
              <button @click="deleteMustahiq(m.id)" class="flex-1 bg-red-50 text-red-600 py-2 rounded-lg hover:bg-red-100 text-sm">
                <i class="fas fa-trash mr-1"></i>Hapus
              </button>
            </div>
          </div>
        </div>
      </main>
    </div>

    <Modal :show="showModal" :title="editingId ? 'Edit Mustahiq' : 'Tambah Mustahiq'" @close="closeModal">
      <form @submit.prevent="saveMustahiq" class="space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div class="md:col-span-2">
            <label class="block text-sm font-medium text-gray-700 mb-2">Nama Lengkap *</label>
            <input v-model="form.nama" required class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500">
          </div>
          <div class="md:col-span-2">
            <label class="block text-sm font-medium text-gray-700 mb-2">Jenis Penerima *</label>
            <select v-model="form.jenis_penerima" required class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500">
              <option value="">Pilih Jenis</option>
              <option value="fakir">Fakir</option>
              <option value="miskin">Miskin</option>
              <option value="amil">Amil</option>
              <option value="mualaf">Mualaf</option>
              <option value="riqab">Riqab</option>
              <option value="gharim">Gharim</option>
              <option value="fisabilillah">Fisabilillah</option>
              <option value="ibnu sabil">Ibnu Sabil</option>
            </select>
          </div>
          <div class="md:col-span-2">
            <label class="block text-sm font-medium text-gray-700 mb-2">Alamat *</label>
            <textarea v-model="form.alamat" required rows="3" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500"></textarea>
          </div>
          <div class="md:col-span-2">
            <label class="block text-sm font-medium text-gray-700 mb-2">Lokasi *</label>
            <select v-model="form.lokasi" required @change="form.rt = ''" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500">
              <option value="">Pilih Lokasi</option>
              <option value="dalam_desa">Dalam Desa Purwajaya</option>
              <option value="luar_desa">Luar Desa Purwajaya</option>
            </select>
          </div>
          <div v-if="form.lokasi === 'dalam_desa'" class="md:col-span-2">
            <label class="block text-sm font-medium text-gray-700 mb-2">RT *</label>
            <select v-model="form.rt" required class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500">
              <option value="">Pilih RT</option>
              <option v-for="i in 21" :key="i" :value="String(i).padStart(2, '0')">RT {{ String(i).padStart(2, '0') }}</option>
            </select>
          </div>
          <div v-if="form.lokasi === 'luar_desa'" class="md:col-span-2">
            <label class="block text-sm font-medium text-gray-700 mb-2">RT *</label>
            <input v-model="form.rt" type="number" required min="1" placeholder="Masukkan nomor RT" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Telepon</label>
            <input v-model="form.telepon" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500">
          </div>
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
    <LoadingOverlay :show="isSaving" message="Menyimpan data..." />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import Swal from 'sweetalert2'
import Sidebar from '../../components/Sidebar.vue'
import Modal from '../../components/Modal.vue'
import Toast from '../../components/Toast.vue'
import SkeletonCard from '../../components/SkeletonCard.vue'
import LoadingOverlay from '../../components/LoadingOverlay.vue'
import api from '../../api'

const menuItems = [
  { path: '/petugas', label: 'Dashboard', icon: 'fas fa-home' },
  { path: '/petugas/pengaturan', label: 'Pengaturan', icon: 'fas fa-cog' },
  { path: '/petugas/mustahiq', label: 'Mustahiq', icon: 'fas fa-hand-holding-heart' },
  { path: '/petugas/transaksi', label: 'Transaksi', icon: 'fas fa-receipt' }
]

const showSidebar = ref(false)

const mustahiqList = ref([])
const showModal = ref(false)
const editingId = ref(null)
const form = ref({ nama: '', jenis_penerima: '', alamat: '', lokasi: '', rt: '', telepon: '' })
const toast = ref({ message: '', type: 'success' })
const isLoading = ref(true)
const isSaving = ref(false)

const loadMustahiq = async () => {
  isLoading.value = true
  try {
    const { data } = await api.getMustahiq()
    console.log('Mustahiq Response:', data)
    console.log('Mustahiq List Length:', data.data?.length || 0)
    if (data.success) mustahiqList.value = data.data || []
  } catch (error) {
    console.error(error)
  } finally {
    isLoading.value = false
  }
}

const editMustahiq = (m) => {
  editingId.value = m.id
  form.value = { ...m }
  showModal.value = true
}

const deleteMustahiq = async (id) => {
  if (!confirm('Yakin ingin menghapus?')) return
  try {
    const { data } = await api.deleteMustahiq(id)
    if (data.success) {
      toast.value = { message: 'Mustahiq berhasil dihapus', type: 'success' }
      loadMustahiq()
    }
  } catch (error) {
    toast.value = { message: 'Gagal menghapus mustahiq', type: 'error' }
  }
}

const saveMustahiq = async () => {
  if (!editingId.value) {
    try {
      const { data } = await api.checkMustahiqDuplicate({ nama: form.value.nama, lokasi: form.value.lokasi, rt: form.value.rt })
      if (data.success && data.data.exists) {
        const result = await Swal.fire({
          title: 'Data Sudah Ada',
          text: `Mustahiq dengan nama, lokasi, dan RT yang sama sudah terdaftar di ${data.data.masjid_nama}. Apakah ingin tetap melanjutkan?`,
          icon: 'warning',
          showCancelButton: true,
          confirmButtonText: 'Ya, Lanjutkan',
          cancelButtonText: 'Batal'
        })
        if (!result.isConfirmed) return
      }
    } catch (error) {
      console.error(error)
    }
  }
  
  try {
    const { data } = editingId.value 
      ? await api.updateMustahiq(editingId.value, form.value)
      : await api.createMustahiq(form.value)
    
    if (data.success) {
      toast.value = { message: editingId.value ? 'Mustahiq berhasil diupdate' : 'Mustahiq berhasil ditambahkan', type: 'success' }
      closeModal()
      loadMustahiq()
    }
  } catch (error) {
    toast.value = { message: 'Gagal menyimpan data', type: 'error' }
  } finally {
    isSaving.value = false
  }
}

const closeModal = () => {
  showModal.value = false
  editingId.value = null
  form.value = { nama: '', jenis_penerima: '', alamat: '', lokasi: '', rt: '', telepon: '' }
}

onMounted(loadMustahiq)
</script>
