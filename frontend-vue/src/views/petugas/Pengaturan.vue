<template>
  <div class="flex h-screen overflow-hidden">
    <Sidebar
      :show="showSidebar"
      @close="showSidebar = false"
      :menuItems="menuItems"
    />

    <div class="flex-1 flex flex-col overflow-hidden">
      <header class="bg-white shadow-sm">
        <div class="px-4 md:px-6 py-4">
          <div class="flex items-center">
            <button
              @click="showSidebar = true"
              class="mr-3 md:hidden text-gray-600 hover:text-gray-800"
            >
              <i class="fas fa-bars text-xl"></i>
            </button>
            <h1 class="text-xl md:text-2xl font-bold text-gray-800">
              Pengaturan Masjid & Zakat
            </h1>
          </div>
        </div>
      </header>

      <main class="flex-1 overflow-y-auto p-4 md:p-6">
        <div class="max-w-4xl">
          <div class="bg-white rounded-xl shadow-md p-6 mb-6">
            <h2 class="text-xl font-bold text-gray-800 mb-4">Data Masjid</h2>
            <form @submit.prevent="saveMasjid" class="space-y-4">
              <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div class="md:col-span-2">
                  <label class="block text-sm font-medium text-gray-700 mb-2"
                    >Nama Masjid *</label
                  >
                  <input
                    v-model="masjidForm.nama"
                    required
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500"
                  />
                </div>
                <div class="md:col-span-2">
                  <label class="block text-sm font-medium text-gray-700 mb-2"
                    >Alamat *</label
                  >
                  <textarea
                    v-model="masjidForm.alamat"
                    required
                    rows="3"
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500"
                  ></textarea>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2"
                    >Telepon</label
                  >
                  <input
                    v-model="masjidForm.telepon"
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500"
                  />
                </div>
              </div>
              <button
                type="submit"
                class="bg-green-600 text-white px-6 py-2 rounded-lg hover:bg-green-700 transition"
              >
                <i class="fas fa-save mr-2"></i>Simpan Data Masjid
              </button>
            </form>
          </div>

          <div class="bg-white rounded-xl shadow-md p-6">
            <h2 class="text-xl font-bold text-gray-800 mb-4">
              Kadar Zakat Fitrah (Uang)
            </h2>
            <form @submit.prevent="saveZakat" class="space-y-4">
              <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2"
                    >Kelas 1 (Rp) *</label
                  >
                  <input
                    v-model.number="zakatForm.kelas1"
                    type="number"
                    required
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2"
                    >Kelas 2 (Rp) *</label
                  >
                  <input
                    v-model.number="zakatForm.kelas2"
                    type="number"
                    required
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2"
                    >Kelas 3 (Rp) *</label
                  >
                  <input
                    v-model.number="zakatForm.kelas3"
                    type="number"
                    required
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500"
                  />
                </div>
              </div>
              <button
                type="submit"
                class="bg-green-600 text-white px-6 py-2 rounded-lg hover:bg-green-700 transition"
              >
                <i class="fas fa-save mr-2"></i>Simpan Kadar Zakat
              </button>
            </form>
          </div>
        </div>
      </main>
    </div>

    <Toast :message="toast.message" :type="toast.type" />
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import Sidebar from "../../components/Sidebar.vue";
import Toast from "../../components/Toast.vue";
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
];

const showSidebar = ref(false);

const masjidForm = ref({ nama: "", alamat: "", telepon: "" });
const zakatForm = ref({ kelas1: 35000, kelas2: 45000, kelas3: 55000 });
const toast = ref({ message: "", type: "success" });

const saveMasjid = async () => {
  try {
    const { data } = await api.updateMyMasjid(masjidForm.value);
    if (data.success) {
      toast.value = {
        message: "Data masjid berhasil disimpan",
        type: "success",
      };
    }
  } catch (error) {
    toast.value = { message: "Gagal menyimpan data", type: "error" };
  }
};

const saveZakat = () => {
  localStorage.setItem("kadar_zakat", JSON.stringify(zakatForm.value));
  toast.value = { message: "Kadar zakat berhasil disimpan", type: "success" };
};

onMounted(async () => {
  try {
    const { data } = await api.getMyMasjid();
    console.log("Masjid data:", data);
    if (data.success && data.data) {
      masjidForm.value = {
        nama: data.data.nama || "",
        alamat: data.data.alamat || "",
        telepon: data.data.telepon || "",
      };
    } else {
      console.log("No masjid data or user not assigned to masjid");
    }
  } catch (error) {
    console.error("Error loading masjid:", error);
    if (error.response) {
      console.error("Response data:", error.response.data);
      toast.value = {
        message: error.response.data.message || "Gagal memuat data masjid",
        type: "error",
      };
    }
  }

  const saved = localStorage.getItem("kadar_zakat");
  if (saved) zakatForm.value = JSON.parse(saved);
});
</script>
