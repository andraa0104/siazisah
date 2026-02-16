<template>
  <div class="flex h-screen overflow-hidden">
    <Sidebar
      :show="showSidebar"
      @close="showSidebar = false"
      :menuItems="menuItems"
    />

    <div class="flex-1 flex flex-col overflow-hidden">
      <header class="bg-white shadow-sm">
        <div class="px-4 md:px-6 py-4 flex justify-between items-center">
          <div class="flex items-center">
            <button
              @click="showSidebar = true"
              class="mr-3 md:hidden text-gray-600 hover:text-gray-800"
            >
              <i class="fas fa-bars text-xl"></i>
            </button>
            <h1 class="text-xl md:text-2xl font-bold text-gray-800">
              Kelola Masjid/Langgar
            </h1>
          </div>
          <button
            @click="showModal = true"
            class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition"
          >
            <i class="fas fa-plus mr-2"></i>Tambah Masjid
          </button>
        </div>
      </header>

      <main class="flex-1 overflow-y-auto p-4 md:p-6">
        <!-- Desktop Table -->
        <div
          class="hidden md:block bg-white rounded-xl shadow-md overflow-x-auto"
        >
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase"
                >
                  Nama
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase"
                >
                  Alamat
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase"
                >
                  Telepon
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase"
                >
                  Status
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase"
                >
                  Aksi
                </th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <template v-if="isLoading">
                <SkeletonCard v-for="i in 5" :key="i" type="table-row" />
              </template>
              <tr v-else v-for="m in masjidList" :key="m.id" class="hover:bg-gray-50">
                <td class="px-6 py-4">
                  <div class="font-medium text-gray-900">{{ m.nama }}</div>
                </td>
                <td class="px-6 py-4">
                  <div class="text-sm text-gray-900">{{ m.alamat }}</div>
                </td>
                <td class="px-6 py-4 text-sm text-gray-900">
                  {{ m.telepon || "-" }}
                </td>
                <td class="px-6 py-4">
                  <span
                    :class="`px-2 py-1 text-xs font-semibold rounded-full ${m.is_active ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'}`"
                  >
                    {{ m.is_active ? "Aktif" : "Nonaktif" }}
                  </span>
                </td>
                <td class="px-6 py-4 text-sm">
                  <button
                    @click="editMasjid(m)"
                    class="text-blue-600 hover:text-blue-900 mr-3"
                  >
                    <i class="fas fa-edit"></i>
                  </button>
                  <button
                    @click="deleteMasjid(m.id)"
                    class="text-red-600 hover:text-red-900"
                  >
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
          <div v-else
            v-for="m in masjidList"
            :key="m.id"
            class="bg-white rounded-xl shadow-md p-4"
          >
            <div class="flex justify-between items-start mb-2">
              <p class="font-medium text-gray-900">{{ m.nama }}</p>
              <span
                :class="`px-2 py-1 text-xs font-semibold rounded-full ${m.is_active ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'}`"
              >
                {{ m.is_active ? "Aktif" : "Nonaktif" }}
              </span>
            </div>
            <div class="space-y-1 text-sm mb-3">
              <p class="text-gray-600">{{ m.alamat }}</p>
              <p class="text-gray-600">{{ m.telepon || "-" }}</p>
            </div>
            <div class="flex gap-2">
              <button
                @click="editMasjid(m)"
                class="flex-1 bg-blue-50 text-blue-600 py-2 rounded-lg hover:bg-blue-100 text-sm"
              >
                <i class="fas fa-edit mr-1"></i>Edit
              </button>
              <button
                @click="deleteMasjid(m.id)"
                class="flex-1 bg-red-50 text-red-600 py-2 rounded-lg hover:bg-red-100 text-sm"
              >
                <i class="fas fa-trash mr-1"></i>Hapus
              </button>
            </div>
          </div>
        </div>
      </main>
    </div>

    <Modal
      :show="showModal"
      :title="editingId ? 'Edit Masjid' : 'Tambah Masjid'"
      @close="closeModal"
    >
      <form @submit.prevent="saveMasjid" class="space-y-4">
        <div class="grid grid-cols-2 gap-4">
          <div class="col-span-2">
            <label class="block text-sm font-medium text-gray-700 mb-2"
              >Nama Masjid/Langgar *</label
            >
            <input
              v-model="form.nama"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            />
          </div>
          <div class="col-span-2">
            <label class="block text-sm font-medium text-gray-700 mb-2"
              >Alamat *</label
            >
            <textarea
              v-model="form.alamat"
              required
              rows="3"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            ></textarea>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2"
              >Telepon</label
            >
            <input
              v-model="form.telepon"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2"
              >Status</label
            >
            <select
              v-model="form.is_active"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            >
              <option :value="true">Aktif</option>
              <option :value="false">Nonaktif</option>
            </select>
          </div>
        </div>
        <div class="flex gap-3 pt-4">
          <button
            type="submit"
            :disabled="isSaving"
            class="flex-1 bg-blue-600 text-white py-3 rounded-lg hover:bg-blue-700 transition font-medium disabled:opacity-50"
          >
            <i class="fas fa-save mr-2"></i>{{ isSaving ? 'Menyimpan...' : 'Simpan' }}
          </button>
          <button
            type="button"
            @click="closeModal"
            :disabled="isSaving"
            class="flex-1 bg-gray-200 text-gray-700 py-3 rounded-lg hover:bg-gray-300 transition font-medium disabled:opacity-50"
          >
            Batal
          </button>
        </div>
      </form>
    </Modal>

    <LoadingOverlay :show="isSaving" message="Menyimpan data..." />
    <Toast :message="toast.message" :type="toast.type" />
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import Sidebar from "../../components/Sidebar.vue";
import Modal from "../../components/Modal.vue";
import Toast from "../../components/Toast.vue";
import SkeletonCard from "../../components/SkeletonCard.vue";
import LoadingOverlay from "../../components/LoadingOverlay.vue";
import api from "../../api";

const menuItems = [
  { path: "/superadmin", label: "Dashboard", icon: "fas fa-home" },
  {
    path: "/superadmin/masjid",
    label: "Masjid/Langgar",
    icon: "fas fa-mosque",
  },
  { path: "/superadmin/users", label: "Petugas", icon: "fas fa-users" },
];

const showSidebar = ref(false);
const masjidList = ref([]);
const showModal = ref(false);
const editingId = ref(null);
const isLoading = ref(true);
const isSaving = ref(false);
const form = ref({
  nama: "",
  alamat: "",
  telepon: "",
  is_active: true,
});
const toast = ref({ message: "", type: "success" });

const loadMasjid = async () => {
  isLoading.value = true;
  try {
    const { data } = await api.getMasjid();
    if (data.success) masjidList.value = data.data || [];
  } catch (error) {
    console.error(error);
  } finally {
    isLoading.value = false;
  }
};

const editMasjid = (m) => {
  editingId.value = m.id;
  form.value = { ...m };
  showModal.value = true;
};

const deleteMasjid = async (id) => {
  if (!confirm("Yakin ingin menghapus?")) return;
  try {
    const { data } = await api.deleteMasjid(id);
    if (data.success) {
      toast.value = { message: "Masjid berhasil dihapus", type: "success" };
      loadMasjid();
    }
  } catch (error) {
    toast.value = { message: "Gagal menghapus masjid", type: "error" };
  }
};

const saveMasjid = async () => {
  isSaving.value = true;
  try {
    const { data } = editingId.value
      ? await api.updateMasjid(editingId.value, form.value)
      : await api.createMasjid(form.value);

    if (data.success) {
      toast.value = {
        message: editingId.value
          ? "Masjid berhasil diupdate"
          : "Masjid berhasil ditambahkan",
        type: "success",
      };
      closeModal();
      loadMasjid();
    }
  } catch (error) {
    toast.value = { message: "Gagal menyimpan data", type: "error" };
  } finally {
    isSaving.value = false;
  }
};

const closeModal = () => {
  showModal.value = false;
  editingId.value = null;
  form.value = {
    nama: "",
    alamat: "",
    telepon: "",
    is_active: true,
  };
};

onMounted(loadMasjid);
</script>
