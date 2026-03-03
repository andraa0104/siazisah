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
              Kelola Mustahiq
            </h1>
          </div>
          <button
            @click="showModal = true"
            class="bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700 transition"
          >
            <i class="fas fa-plus mr-2"></i>Tambah Mustahiq
          </button>
        </div>
      </header>

      <main class="flex-1 overflow-y-auto p-4 md:p-6">
        <!-- Category Stats Cards -->
        <div class="grid grid-cols-2 md:grid-cols-4 gap-3 mb-6">
          <div
            v-for="category in mustahiqCategories"
            :key="category.key"
            :class="`${category.bgColor} border ${category.borderColor} rounded-xl p-4 shadow-sm cursor-pointer hover:shadow-md transition-shadow`"
            @click="
              filters.jenis_penerima = category.key;
              applyFilters();
            "
          >
            <div class="flex items-center justify-between">
              <div>
                <p class="text-xs text-gray-600 mb-1">{{ category.label }}</p>
                <p
                  :class="`${category.textColor} text-lg md:text-xl font-bold`"
                >
                  {{ mustahiqStats[category.key] || 0 }}
                </p>
              </div>
              <i
                :class="`${category.icon} ${category.iconColor} text-2xl opacity-50`"
              ></i>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-xl shadow-md p-4 mb-4">
          <div class="space-y-3 md:space-y-0 md:grid md:grid-cols-3 md:gap-3">
            <div>
              <label class="block text-xs text-gray-500 mb-1"
                >Filter Jenis Penerima</label
              >
              <select
                v-model="filters.jenis_penerima"
                @change="applyFilters"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm"
              >
                <option value="">Semua Jenis</option>
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
              <label class="block text-xs text-gray-500 mb-1"
                >Cari Mustahiq</label
              >
              <div class="flex flex-col sm:flex-row gap-2">
                <input
                  v-model="filters.q"
                  @keyup.enter="applyFilters"
                  placeholder="Ketik nama mustahiq..."
                  class="flex-1 px-3 py-2 border border-gray-300 rounded-lg text-sm"
                />
                <div class="flex gap-2">
                  <button
                    @click="applyFilters"
                    class="flex-1 sm:flex-none px-4 py-2 bg-green-600 text-white rounded-lg text-sm hover:bg-green-700 font-medium"
                  >
                    Cari
                  </button>
                  <button
                    @click="resetFilters"
                    class="flex-1 sm:flex-none px-4 py-2 bg-gray-100 text-gray-700 rounded-lg text-sm hover:bg-gray-200 font-medium"
                  >
                    Reset
                  </button>
                </div>
              </div>
            </div>
          </div>

          <div
            v-if="activeFilterChips.length > 0"
            class="mt-3 flex flex-wrap items-center gap-2"
          >
            <span class="text-xs text-gray-500">Filter aktif:</span>
            <span
              v-for="chip in activeFilterChips"
              :key="chip"
              class="px-2 py-1 text-xs rounded-full bg-green-50 text-green-700 border border-green-200"
            >
              {{ chip }}
            </span>
          </div>
        </div>

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
                  Jenis
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase"
                >
                  Alamat
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase"
                >
                  Lokasi
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase"
                >
                  RT
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase"
                >
                  Aksi
                </th>
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
              <tr
                v-else
                v-for="m in mustahiqList"
                :key="m.id"
                class="hover:bg-gray-50"
              >
                <td class="px-6 py-4">
                  <div class="font-medium text-gray-900">{{ m.nama }}</div>
                </td>
                <td class="px-6 py-4">
                  <span
                    class="px-2 py-1 text-xs font-semibold rounded-full bg-blue-100 text-blue-800"
                    >{{ m.jenis_penerima }}</span
                  >
                </td>
                <td class="px-6 py-4 text-sm text-gray-900">{{ m.alamat }}</td>
                <td class="px-6 py-4 text-sm text-gray-900">
                  {{ m.lokasi === "dalam_desa" ? "Dalam Desa" : "Luar Desa" }}
                </td>
                <td class="px-6 py-4 text-sm text-gray-900">RT {{ m.rt }}</td>
                <td class="px-6 py-4 text-sm">
                  <button
                    @click="editMustahiq(m)"
                    class="text-blue-600 hover:text-blue-900 mr-3"
                  >
                    <i class="fas fa-edit"></i>
                  </button>
                  <button
                    @click="deleteMustahiq(m.id)"
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
          <!-- Loading Skeleton -->
          <template v-if="isLoading">
            <SkeletonCard v-for="i in 5" :key="i" type="list-item" />
          </template>

          <!-- Actual Data -->
          <div
            v-else
            v-for="m in mustahiqList"
            :key="m.id"
            class="bg-white rounded-xl shadow-md p-4"
          >
            <div class="flex justify-between items-start mb-2">
              <p class="font-medium text-gray-900">{{ m.nama }}</p>
              <span
                class="px-2 py-1 text-xs font-semibold rounded-full bg-blue-100 text-blue-800"
                >{{ m.jenis_penerima }}</span
              >
            </div>
            <div class="space-y-1 text-sm mb-3">
              <p class="text-gray-600">{{ m.alamat }}</p>
              <p class="text-gray-600">
                {{ m.lokasi === "dalam_desa" ? "Dalam Desa" : "Luar Desa" }} -
                RT {{ m.rt }}
              </p>
            </div>
            <div class="flex gap-2">
              <button
                @click="editMustahiq(m)"
                class="flex-1 bg-blue-50 text-blue-600 py-2 rounded-lg hover:bg-blue-100 text-sm"
              >
                <i class="fas fa-edit mr-1"></i>Edit
              </button>
              <button
                @click="deleteMustahiq(m.id)"
                class="flex-1 bg-red-50 text-red-600 py-2 rounded-lg hover:bg-red-100 text-sm"
              >
                <i class="fas fa-trash mr-1"></i>Hapus
              </button>
            </div>
          </div>
        </div>

        <div
          class="mt-4 flex flex-col md:flex-row md:items-center md:justify-between gap-3 text-sm text-gray-600"
        >
          <div>
            Menampilkan {{ mustahiqList.length }} dari
            {{ pagination.total }} data
          </div>
          <div class="flex flex-wrap items-center gap-2">
            <label class="text-xs text-gray-500">Per halaman</label>
            <select
              v-model="pagination.pageSize"
              @change="onPageSizeChange"
              class="px-2 py-1 border border-gray-300 rounded-lg text-sm"
            >
              <option v-for="size in pageSizeOptions" :key="size" :value="size">
                {{ size === "all" ? "Semua" : size }}
              </option>
            </select>
            <button
              class="px-3 py-1 border border-gray-300 rounded-lg disabled:opacity-50"
              :disabled="pagination.isAll || pagination.page <= 1"
              @click="setPage(pagination.page - 1)"
            >
              Prev
            </button>
            <span class="text-xs text-gray-500"
              >{{ pagination.page }} / {{ pagination.totalPages }}</span
            >
            <button
              class="px-3 py-1 border border-gray-300 rounded-lg disabled:opacity-50"
              :disabled="
                pagination.isAll || pagination.page >= pagination.totalPages
              "
              @click="setPage(pagination.page + 1)"
            >
              Next
            </button>
          </div>
        </div>
      </main>
    </div>

    <Modal
      :show="showModal"
      :title="editingId ? 'Edit Mustahiq' : 'Tambah Mustahiq'"
      @close="closeModal"
    >
      <form @submit.prevent="saveMustahiq" class="space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div class="md:col-span-2">
            <label class="block text-sm font-medium text-gray-700 mb-2"
              >Nama Lengkap *</label
            >
            <input
              v-model="form.nama"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500"
            />
          </div>
          <div class="md:col-span-2">
            <label class="block text-sm font-medium text-gray-700 mb-2"
              >Jenis Penerima *</label
            >
            <select
              v-model="form.jenis_penerima"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500"
            >
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
            <label class="block text-sm font-medium text-gray-700 mb-2"
              >Alamat *</label
            >
            <textarea
              v-model="form.alamat"
              required
              rows="3"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500"
            ></textarea>
          </div>
          <div class="md:col-span-2">
            <label class="block text-sm font-medium text-gray-700 mb-2"
              >Lokasi *</label
            >
            <select
              v-model="form.lokasi"
              required
              @change="form.rt = ''"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500"
            >
              <option value="">Pilih Lokasi</option>
              <option value="dalam_desa">Dalam Desa Purwajaya</option>
              <option value="luar_desa">Luar Desa Purwajaya</option>
            </select>
          </div>
          <div v-if="form.lokasi === 'dalam_desa'" class="md:col-span-2">
            <label class="block text-sm font-medium text-gray-700 mb-2"
              >RT *</label
            >
            <select
              v-model="form.rt"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500"
            >
              <option value="">Pilih RT</option>
              <option
                v-for="i in 21"
                :key="i"
                :value="String(i).padStart(2, '0')"
              >
                RT {{ String(i).padStart(2, "0") }}
              </option>
            </select>
          </div>
          <div v-if="form.lokasi === 'luar_desa'" class="md:col-span-2">
            <label class="block text-sm font-medium text-gray-700 mb-2"
              >RT *</label
            >
            <input
              v-model="form.rt"
              type="number"
              required
              min="1"
              placeholder="Masukkan nomor RT"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2"
              >Telepon</label
            >
            <input
              v-model="form.telepon"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500"
            />
          </div>
        </div>
        <div class="flex gap-3 pt-4">
          <button
            type="submit"
            class="flex-1 bg-green-600 text-white py-3 rounded-lg hover:bg-green-700 transition font-medium"
          >
            <i class="fas fa-save mr-2"></i>Simpan
          </button>
          <button
            type="button"
            @click="closeModal"
            class="flex-1 bg-gray-200 text-gray-700 py-3 rounded-lg hover:bg-gray-300 transition font-medium"
          >
            Batal
          </button>
        </div>
      </form>
    </Modal>

    <Toast :message="toast.message" :type="toast.type" />
    <LoadingOverlay :show="isSaving" message="Menyimpan data..." />
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import Swal from "sweetalert2";
import Sidebar from "../../components/Sidebar.vue";
import Modal from "../../components/Modal.vue";
import Toast from "../../components/Toast.vue";
import SkeletonCard from "../../components/SkeletonCard.vue";
import LoadingOverlay from "../../components/LoadingOverlay.vue";
import api from "../../api";

const menuItems = [
  { path: "/petugas", label: "Dashboard", icon: "fas fa-home" },
  { path: "/petugas/pengaturan", label: "Pengaturan", icon: "fas fa-cog" },
  {
    path: "/petugas/mustahiq",
    label: "Mustahiq",
    icon: "fas fa-hand-holding-heart",
  },
  { path: "/petugas/transaksi", label: "Transaksi", icon: "fas fa-receipt" },
  { path: "/petugas/distribusi", label: "Distribusi", icon: "fas fa-box-open" },
];

const showSidebar = ref(false);

const mustahiqList = ref([]);
const allMustahiq = ref([]);
const showModal = ref(false);
const editingId = ref(null);
const form = ref({
  nama: "",
  jenis_penerima: "",
  alamat: "",
  lokasi: "",
  rt: "",
  telepon: "",
});
const toast = ref({ message: "", type: "success" });
const isLoading = ref(true);
const isSaving = ref(false);
const pageSizeOptions = [5, 10, 25, 50, "all"];
const pagination = ref({
  page: 1,
  pageSize: 5,
  total: 0,
  totalPages: 1,
  isAll: false,
});
const filters = ref({ jenis_penerima: "", q: "" });

const activeFilterChips = computed(() => {
  const chips = [];
  if (filters.value.jenis_penerima) {
    chips.push(`Jenis: ${filters.value.jenis_penerima}`);
  }
  if (filters.value.q && filters.value.q.trim()) {
    chips.push(`Cari: ${filters.value.q.trim()}`);
  }
  return chips;
});

// Kategori Mustahiq dengan icon dan warna
const mustahiqCategories = [
  {
    key: "fakir",
    label: "Fakir",
    icon: "fas fa-person-hiking",
    bgColor: "bg-blue-50",
    borderColor: "border-blue-200",
    textColor: "text-blue-700",
    iconColor: "text-blue-400",
  },
  {
    key: "miskin",
    label: "Miskin",
    icon: "fas fa-hand-holding-heart",
    bgColor: "bg-emerald-50",
    borderColor: "border-emerald-200",
    textColor: "text-emerald-700",
    iconColor: "text-emerald-400",
  },
  {
    key: "amil",
    label: "Amil",
    icon: "fas fa-briefcase",
    bgColor: "bg-amber-50",
    borderColor: "border-amber-200",
    textColor: "text-amber-700",
    iconColor: "text-amber-400",
  },
  {
    key: "mualaf",
    label: "Mualaf",
    icon: "fas fa-heart",
    bgColor: "bg-rose-50",
    borderColor: "border-rose-200",
    textColor: "text-rose-700",
    iconColor: "text-rose-400",
  },
  {
    key: "riqab",
    label: "Riqab",
    icon: "fas fa-unlock",
    bgColor: "bg-red-50",
    borderColor: "border-red-200",
    textColor: "text-red-700",
    iconColor: "text-red-400",
  },
  {
    key: "gharim",
    label: "Gharim",
    icon: "fas fa-coins",
    bgColor: "bg-yellow-50",
    borderColor: "border-yellow-200",
    textColor: "text-yellow-700",
    iconColor: "text-yellow-400",
  },
  {
    key: "fisabilillah",
    label: "Fisabilillah",
    icon: "fas fa-swords",
    bgColor: "bg-purple-50",
    borderColor: "border-purple-200",
    textColor: "text-purple-700",
    iconColor: "text-purple-400",
  },
  {
    key: "ibnu sabil",
    label: "Ibnu Sabil",
    icon: "fas fa-road",
    bgColor: "bg-cyan-50",
    borderColor: "border-cyan-200",
    textColor: "text-cyan-700",
    iconColor: "text-cyan-400",
  },
];

// Computed property untuk statistik mustahiq per kategori
const mustahiqStats = computed(() => {
  const stats = {};
  mustahiqCategories.forEach((cat) => {
    stats[cat.key] = 0;
  });

  // Hitung dari allMustahiq (semua data tanpa filter)
  allMustahiq.value.forEach((item) => {
    if (stats.hasOwnProperty(item.jenis_penerima)) {
      stats[item.jenis_penerima]++;
    }
  });
  return stats;
});

const applyFilters = () => {
  pagination.value.page = 1;
  loadMustahiq();
};

const resetFilters = () => {
  filters.value = { jenis_penerima: "", q: "" };
  pagination.value.page = 1;
  loadMustahiq();
};

const setPage = (page) => {
  pagination.value.page = page;
  loadMustahiq();
};

const onPageSizeChange = () => {
  pagination.value.page = 1;
  loadMustahiq();
};

const syncPagination = (payload, itemCount) => {
  if (payload && payload.pagination) {
    const meta = payload.pagination;
    pagination.value = {
      page: meta.page || 1,
      pageSize: meta.is_all
        ? "all"
        : meta.page_size || pagination.value.pageSize,
      total: meta.total ?? itemCount,
      totalPages: meta.total_pages || 1,
      isAll: !!meta.is_all,
    };
    return;
  }
  pagination.value.total = itemCount;
  pagination.value.totalPages = 1;
  pagination.value.isAll = pagination.value.pageSize === "all";
};

const applyLocalPagination = (items) => {
  const total = items.length;
  if (pagination.value.pageSize === "all") {
    pagination.value.total = total;
    pagination.value.totalPages = 1;
    pagination.value.isAll = true;
    return items;
  }

  const pageSize = Number(pagination.value.pageSize) || 5;
  const totalPages = Math.max(1, Math.ceil(total / pageSize));
  const page = Math.min(pagination.value.page, totalPages);
  const start = (page - 1) * pageSize;
  const sliced = items.slice(start, start + pageSize);

  pagination.value.total = total;
  pagination.value.totalPages = totalPages;
  pagination.value.page = page;
  pagination.value.isAll = false;

  return sliced;
};

const filterLocalMustahiq = (items) => {
  const query = String(filters.value.q || "")
    .trim()
    .toLowerCase();
  return items.filter((item) => {
    if (
      filters.value.jenis_penerima &&
      item.jenis_penerima !== filters.value.jenis_penerima
    ) {
      return false;
    }

    if (query) {
      const nama = String(item.nama || "").toLowerCase();
      if (!nama.includes(query)) {
        return false;
      }
    }

    return true;
  });
};

const loadMustahiq = async () => {
  isLoading.value = true;
  try {
    const params =
      pagination.value.pageSize === "all"
        ? { page_size: "all" }
        : { page: pagination.value.page, page_size: pagination.value.pageSize };

    if (filters.value.jenis_penerima) {
      params.jenis_penerima = filters.value.jenis_penerima;
    }
    if (filters.value.q && filters.value.q.trim()) {
      params.q = filters.value.q.trim();
    }

    const { data } = await api.getMustahiq(params);
    console.log("Mustahiq Response:", data);
    console.log("Mustahiq List Length:", data.data?.length || 0);
    if (data.success) {
      const payload = data.data || {};
      if (Array.isArray(payload.items)) {
        const filteredItems = filterLocalMustahiq(payload.items);
        mustahiqList.value = filteredItems;
        syncPagination(payload, filteredItems.length);
      } else if (Array.isArray(payload)) {
        const filteredItems = filterLocalMustahiq(payload);
        mustahiqList.value = applyLocalPagination(filteredItems);
      } else {
        mustahiqList.value = [];
        syncPagination(payload, 0);
      }
    }
  } catch (error) {
    console.error(error);
  } finally {
    isLoading.value = false;
  }
};

// Load semua mustahiq untuk keperluan statistik (tanpa filter)
const loadAllMustahiqForStats = async () => {
  try {
    const { data } = await api.getMustahiq({ page_size: "all" });
    if (data.success) {
      const payload = data.data || {};
      if (Array.isArray(payload.items)) {
        allMustahiq.value = payload.items;
      } else if (Array.isArray(payload)) {
        allMustahiq.value = payload;
      } else {
        allMustahiq.value = [];
      }
    }
  } catch (error) {
    console.error("Error loading stats:", error);
  }
};

const editMustahiq = (m) => {
  editingId.value = m.id;
  form.value = { ...m };
  showModal.value = true;
};

const deleteMustahiq = async (id) => {
  if (!confirm("Yakin ingin menghapus?")) return;
  try {
    const { data } = await api.deleteMustahiq(id);
    if (data.success) {
      toast.value = { message: "Mustahiq berhasil dihapus", type: "success" };
      loadMustahiq();
      loadAllMustahiqForStats();
    }
  } catch (error) {
    toast.value = { message: "Gagal menghapus mustahiq", type: "error" };
  }
};

const saveMustahiq = async () => {
  if (!editingId.value) {
    try {
      const { data } = await api.checkMustahiqDuplicate({
        nama: form.value.nama,
        lokasi: form.value.lokasi,
        rt: form.value.rt,
      });
      if (data.success && data.data.exists) {
        const result = await Swal.fire({
          title: "Data Sudah Ada",
          text: `Mustahiq dengan nama, lokasi, dan RT yang sama sudah terdaftar di ${data.data.masjid_nama}. Apakah ingin tetap melanjutkan?`,
          icon: "warning",
          showCancelButton: true,
          confirmButtonText: "Ya, Lanjutkan",
          cancelButtonText: "Batal",
        });
        if (!result.isConfirmed) return;
      }
    } catch (error) {
      console.error(error);
    }
  }

  try {
    const { data } = editingId.value
      ? await api.updateMustahiq(editingId.value, form.value)
      : await api.createMustahiq(form.value);

    if (data.success) {
      toast.value = {
        message: editingId.value
          ? "Mustahiq berhasil diupdate"
          : "Mustahiq berhasil ditambahkan",
        type: "success",
      };
      closeModal();
      loadMustahiq();
      loadAllMustahiqForStats();
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
    jenis_penerima: "",
    alamat: "",
    lokasi: "",
    rt: "",
    telepon: "",
  };
};

onMounted(() => {
  loadMustahiq();
  loadAllMustahiqForStats();
});
</script>
