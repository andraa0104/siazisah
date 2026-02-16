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
            <h1 class="text-xl md:text-2xl font-bold text-gray-800">Kelola Petugas</h1>
          </div>
          <button @click="showModal = true" class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition">
            <i class="fas fa-plus mr-2"></i>Tambah Petugas
          </button>
        </div>
      </header>

      <main class="flex-1 overflow-y-auto p-4 md:p-6">
        <!-- Desktop Table -->
        <div class="hidden md:block bg-white rounded-xl shadow-md overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Username</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Nama Lengkap</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Masjid</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Aksi</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <template v-if="isLoading">
                <SkeletonCard v-for="i in 5" :key="i" type="table-row" />
              </template>
              <tr v-else v-for="u in petugasList" :key="u.id" class="hover:bg-gray-50">
                <td class="px-6 py-4"><div class="font-medium text-gray-900">{{ u.username }}</div></td>
                <td class="px-6 py-4 text-sm text-gray-900">{{ u.full_name }}</td>
                <td class="px-6 py-4 text-sm text-gray-900">{{ getMasjidName(u.masjid_id) }}</td>
                <td class="px-6 py-4 text-sm">
                  <button @click="editUser(u)" class="text-blue-600 hover:text-blue-900 mr-3">
                    <i class="fas fa-edit"></i>
                  </button>
                  <button @click="deleteUser(u.id)" class="text-red-600 hover:text-red-900">
                    <i class="fas fa-trash"></i>
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
        <!-- Mobile List -->
        <div class="md:hidden space-y-3">
          <template v-if="isLoading">
            <SkeletonCard v-for="i in 4" :key="i" type="list-item" />
          </template>
          <div v-else v-for="u in petugasList" :key="u.id" class="bg-white rounded-xl shadow-md p-4">
            <p class="font-medium text-gray-900 mb-1">{{ u.full_name }}</p>
            <div class="space-y-1 text-sm mb-3">
              <p class="text-gray-600">@{{ u.username }}</p>
              <p class="text-gray-600">{{ getMasjidName(u.masjid_id) }}</p>
            </div>
            <div class="flex gap-2">
              <button @click="editUser(u)" class="flex-1 bg-blue-50 text-blue-600 py-2 rounded-lg hover:bg-blue-100 text-sm">
                <i class="fas fa-edit mr-1"></i>Edit
              </button>
              <button @click="deleteUser(u.id)" class="flex-1 bg-red-50 text-red-600 py-2 rounded-lg hover:bg-red-100 text-sm">
                <i class="fas fa-trash mr-1"></i>Hapus
              </button>
            </div>
          </div>
        </div>
      </main>
    </div>

    <Modal :show="showModal" :title="editingId ? 'Edit Petugas' : 'Tambah Petugas'" @close="closeModal">
      <form @submit.prevent="saveUser" class="space-y-4">
        <div class="grid grid-cols-2 gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Username *</label>
            <input v-model="form.username" required class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>
          <!-- Email field removed -->
          <div class="col-span-2">
            <label class="block text-sm font-medium text-gray-700 mb-2">Nama Lengkap *</label>
            <input v-model="form.full_name" required class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Password {{ editingId ? '' : '*' }}</label>
            <input v-model="form.password" type="password" :required="!editingId" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500" :placeholder="editingId ? 'Kosongkan jika tidak ingin mengubah' : ''">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Masjid *</label>
            <select v-model="form.masjid_id" required class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
              <option value="">Pilih Masjid ({{ masjidList.length }} tersedia)</option>
              <option v-for="m in masjidList" :key="m.id" :value="m.id">{{ m.nama }}</option>
            </select>
          </div>
        </div>
        <div class="flex gap-3 pt-4">
          <button type="submit" :disabled="isSaving" class="flex-1 bg-blue-600 text-white py-3 rounded-lg hover:bg-blue-700 transition font-medium disabled:opacity-50">
            <i class="fas fa-save mr-2"></i>{{ isSaving ? 'Menyimpan...' : 'Simpan' }}
          </button>
          <button type="button" @click="closeModal" :disabled="isSaving" class="flex-1 bg-gray-200 text-gray-700 py-3 rounded-lg hover:bg-gray-300 transition font-medium disabled:opacity-50">Batal</button>
        </div>
      </form>
    </Modal>

    <LoadingOverlay :show="isSaving" message="Menyimpan data..." />
    <Toast :message="toast.message" :type="toast.type" />
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import Sidebar from '../../components/Sidebar.vue'
import Modal from '../../components/Modal.vue'
import Toast from '../../components/Toast.vue'
import SkeletonCard from '../../components/SkeletonCard.vue'
import LoadingOverlay from '../../components/LoadingOverlay.vue'
import api from '../../api'

const menuItems = [
  { path: '/superadmin', label: 'Dashboard', icon: 'fas fa-home' },
  { path: '/superadmin/masjid', label: 'Masjid/Langgar', icon: 'fas fa-mosque' },
  { path: '/superadmin/users', label: 'Petugas', icon: 'fas fa-users' }
]

const showSidebar = ref(false)
const usersList = ref([])
const masjidList = ref([])
const showModal = ref(false)
const editingId = ref(null)
const isLoading = ref(true)
const isSaving = ref(false)
const form = ref({ username: '', full_name: '', password: '', masjid_id: '' })
const toast = ref({ message: '', type: 'success' })

const petugasList = computed(() => usersList.value.filter(u => u.role === 'petugas'))

const getMasjidName = (id) => {
  const m = masjidList.value.find(m => m.id === id)
  return m ? m.nama : '-'
}

const loadData = async () => {
  isLoading.value = true
  try {
    const [usersData, masjidData] = await Promise.all([
      api.getUsers(),
      api.getMasjid()
    ])
    console.log('Users Response:', usersData.data)
    console.log('Masjid Response:', masjidData.data)
    if (usersData.data.success) usersList.value = usersData.data.data || []
    if (masjidData.data.success) masjidList.value = masjidData.data.data || []
    console.log('Masjid List:', masjidList.value)
  } catch (error) {
    console.error('Error loading data:', error)
  } finally {
    isLoading.value = false
  }
}

const editUser = (u) => {
  editingId.value = u.id
  form.value = { username: u.username, full_name: u.full_name, password: '', masjid_id: u.masjid_id }
  showModal.value = true
}

const deleteUser = async (id) => {
  if (!confirm('Yakin ingin menghapus?')) return
  try {
    const { data } = await api.deleteUser(id)
    if (data.success) {
      toast.value = { message: 'Petugas berhasil dihapus', type: 'success' }
      loadData()
    }
  } catch (error) {
    toast.value = { message: 'Gagal menghapus petugas', type: 'error' }
  }
}

const saveUser = async () => {
  isSaving.value = true
  try {
    const payload = { ...form.value, role: 'petugas', is_active: true }
    if (editingId.value && !payload.password) delete payload.password
    
    const { data } = editingId.value
      ? await api.updateUser(editingId.value, payload)
      : await api.createUser(payload)
    
    if (data.success) {
      toast.value = { message: editingId.value ? 'Petugas berhasil diupdate' : 'Petugas berhasil ditambahkan', type: 'success' }
      closeModal()
      loadData()
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
  form.value = { username: '', full_name: '', password: '', masjid_id: '' }
}

onMounted(loadData)
</script>
