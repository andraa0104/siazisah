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

          <!-- Pengurus Masjid -->
          <div class="bg-white rounded-xl shadow-md p-6 mb-6">
            <h2 class="text-xl font-bold text-gray-800 mb-4">Pengurus Masjid/Langgar</h2>
            <form @submit.prevent="savePengurusMasjid" class="space-y-4">
              <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <!-- Ketua -->
                <div class="md:col-span-2 border-b pb-3">
                  <h3 class="font-semibold text-gray-700 mb-3">Ketua Pengurus</h3>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">Nama Ketua *</label>
                  <input v-model="pengurusMasjidForm.ketua.nama" required
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500" />
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">Telepon</label>
                  <input v-model="pengurusMasjidForm.ketua.telepon"
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500" />
                </div>
                <div class="md:col-span-2">
                  <label class="block text-sm font-medium text-gray-700 mb-2">Alamat</label>
                  <input v-model="pengurusMasjidForm.ketua.alamat"
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500" />
                </div>

                <!-- Sekretaris -->
                <div class="md:col-span-2 border-b pb-3 pt-2">
                  <h3 class="font-semibold text-gray-700 mb-3">Sekretaris Pengurus</h3>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">Nama Sekretaris *</label>
                  <input v-model="pengurusMasjidForm.sekretaris.nama" required
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500" />
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">Telepon</label>
                  <input v-model="pengurusMasjidForm.sekretaris.telepon"
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500" />
                </div>
                <div class="md:col-span-2">
                  <label class="block text-sm font-medium text-gray-700 mb-2">Alamat</label>
                  <input v-model="pengurusMasjidForm.sekretaris.alamat"
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500" />
                </div>
              </div>
              <button type="submit"
                class="bg-green-600 text-white px-6 py-2 rounded-lg hover:bg-green-700 transition">
                <i class="fas fa-save mr-2"></i>Simpan Pengurus Masjid
              </button>
            </form>
          </div>

          <!-- Pengurus UPZ -->
          <div class="bg-white rounded-xl shadow-md p-6 mb-6">
            <h2 class="text-xl font-bold text-gray-800 mb-4">Pengurus UPZ (Unit Pengumpul Zakat)</h2>
            <form @submit.prevent="saveKetuaUPZ" class="space-y-4 mb-6">
              <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div class="md:col-span-2 border-b pb-3">
                  <h3 class="font-semibold text-gray-700 mb-3">Ketua UPZ</h3>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">Nama Ketua UPZ *</label>
                  <input v-model="pengurusUPZForm.ketua.nama" required
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500" />
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">Telepon</label>
                  <input v-model="pengurusUPZForm.ketua.telepon"
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500" />
                </div>
                <div class="md:col-span-2">
                  <label class="block text-sm font-medium text-gray-700 mb-2">Alamat</label>
                  <input v-model="pengurusUPZForm.ketua.alamat"
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500" />
                </div>
              </div>
              <button type="submit" class="bg-green-600 text-white px-6 py-2 rounded-lg hover:bg-green-700 transition">
                <i class="fas fa-save mr-2"></i>Simpan Ketua UPZ
              </button>
            </form>

            <!-- Anggota UPZ -->
            <div class="border-t pt-6">
              <div class="flex justify-between items-center mb-4">
                <h3 class="font-semibold text-gray-700">Anggota UPZ</h3>
                <button @click="showAnggotaModal = true"
                  class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition text-sm">
                  <i class="fas fa-plus mr-2"></i>Tambah Anggota
                </button>
              </div>
              <div v-if="anggotaUPZList.length === 0" class="text-center text-gray-500 py-4">
                Belum ada anggota UPZ
              </div>
              <div v-else class="space-y-2">
                <div v-for="anggota in anggotaUPZList" :key="anggota.id"
                  class="flex justify-between items-center p-3 bg-gray-50 rounded-lg">
                  <div>
                    <p class="font-medium">{{ anggota.nama }}</p>
                    <p class="text-sm text-gray-600">{{ anggota.telepon || '-' }}</p>
                  </div>
                  <button @click="deleteAnggotaUPZ(anggota.id)"
                    class="text-red-600 hover:text-red-800">
                    <i class="fas fa-trash"></i>
                  </button>
                </div>
              </div>
            </div>
          </div>

          <div class="bg-white rounded-xl shadow-md p-6">
            <h2 class="text-xl font-bold text-gray-800 mb-4">
              Kadar Zakat Fitrah & Fidyah (Uang)
            </h2>
            <form @submit.prevent="saveZakat" class="space-y-4">
              <div>
                <h3 class="text-md font-semibold text-gray-700 mb-3">Zakat Fitrah</h3>
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
              </div>
              <div>
                <h3 class="text-md font-semibold text-gray-700 mb-3">Fidyah</h3>
                <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-2"
                      >Fidyah per Hari (Rp) *</label
                    >
                    <input
                      v-model.number="zakatForm.fidyahPerHari"
                      type="number"
                      required
                      class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500"
                    />
                    <p class="text-xs text-gray-500 mt-1">
                      Kadar fidyah ditentukan oleh Kementerian Agama
                    </p>
                  </div>
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

    <!-- Modal Tambah Anggota UPZ -->
    <div v-if="showAnggotaModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
      <div class="bg-white rounded-xl shadow-xl max-w-md w-full p-6">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-lg font-bold text-gray-800">Tambah Anggota UPZ</h3>
          <button @click="closeAnggotaModal" class="text-gray-500 hover:text-gray-700">
            <i class="fas fa-times text-xl"></i>
          </button>
        </div>
        <form @submit.prevent="saveAnggotaUPZ" class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Nama Anggota *</label>
            <input v-model="anggotaForm.nama" required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Telepon</label>
            <input v-model="anggotaForm.telepon"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Alamat</label>
            <input v-model="anggotaForm.alamat"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500" />
          </div>
          <div class="flex gap-3">
            <button type="submit"
              class="flex-1 bg-green-600 text-white py-2 rounded-lg hover:bg-green-700 transition">
              <i class="fas fa-save mr-2"></i>Simpan
            </button>
            <button type="button" @click="closeAnggotaModal"
              class="flex-1 bg-gray-200 text-gray-700 py-2 rounded-lg hover:bg-gray-300 transition">
              Batal
            </button>
          </div>
        </form>
      </div>
    </div>

    <LoadingOverlay :show="isSaving" message="Menyimpan data..." />
    <Toast :message="toast.message" :type="toast.type" />
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import Sidebar from "../../components/Sidebar.vue";
import Toast from "../../components/Toast.vue";
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
];

const showSidebar = ref(false);
const isSaving = ref(false);

const masjidForm = ref({ nama: "", alamat: "", telepon: "" });
const zakatForm = ref({ kelas1: 35000, kelas2: 45000, kelas3: 55000, fidyahPerHari: 30000 });
const pengurusMasjidForm = ref({
  ketua: { nama: "", telepon: "", alamat: "" },
  sekretaris: { nama: "", telepon: "", alamat: "" }
});
const pengurusUPZForm = ref({
  ketua: { nama: "", telepon: "", alamat: "" }
});
const anggotaUPZList = ref([]);
const showAnggotaModal = ref(false);
const anggotaForm = ref({ nama: "", telepon: "", alamat: "" });
const toast = ref({ message: "", type: "success" });

const saveMasjid = async () => {
  isSaving.value = true;
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
  } finally {
    isSaving.value = false;
  }
};

const saveZakat = () => {
  localStorage.setItem("kadar_zakat", JSON.stringify(zakatForm.value));
  toast.value = { message: "Kadar zakat berhasil disimpan", type: "success" };
};

const savePengurusMasjid = async () => {
  isSaving.value = true;
  try {
    // Save Ketua
    await api.savePengurusMasjid({
      nama: pengurusMasjidForm.value.ketua.nama,
      jabatan: "Ketua",
      telepon: pengurusMasjidForm.value.ketua.telepon,
      alamat: pengurusMasjidForm.value.ketua.alamat
    });
    // Save Sekretaris
    await api.savePengurusMasjid({
      nama: pengurusMasjidForm.value.sekretaris.nama,
      jabatan: "Sekretaris",
      telepon: pengurusMasjidForm.value.sekretaris.telepon,
      alamat: pengurusMasjidForm.value.sekretaris.alamat
    });
    toast.value = { message: "Pengurus masjid berhasil disimpan", type: "success" };
  } catch (error) {
    toast.value = { message: "Gagal menyimpan pengurus masjid", type: "error" };
  } finally {
    isSaving.value = false;
  }
};

const saveKetuaUPZ = async () => {
  isSaving.value = true;
  try {
    await api.savePengurusZakat({
      nama: pengurusUPZForm.value.ketua.nama,
      jabatan: "Ketua UPZ",
      telepon: pengurusUPZForm.value.ketua.telepon,
      alamat: pengurusUPZForm.value.ketua.alamat
    });
    toast.value = { message: "Ketua UPZ berhasil disimpan", type: "success" };
  } catch (error) {
    toast.value = { message: "Gagal menyimpan ketua UPZ", type: "error" };
  } finally {
    isSaving.value = false;
  }
};

const saveAnggotaUPZ = async () => {
  isSaving.value = true;
  try {
    await api.savePengurusZakat({
      nama: anggotaForm.value.nama,
      jabatan: "Anggota UPZ",
      telepon: anggotaForm.value.telepon,
      alamat: anggotaForm.value.alamat
    });
    toast.value = { message: "Anggota UPZ berhasil ditambahkan", type: "success" };
    closeAnggotaModal();
    loadPengurusZakat();
  } catch (error) {
    toast.value = { message: "Gagal menambahkan anggota UPZ", type: "error" };
  } finally {
    isSaving.value = false;
  }
};

const deleteAnggotaUPZ = async (id) => {
  if (!confirm("Hapus anggota UPZ ini?")) return;
  try {
    await api.deletePengurusZakat(id);
    toast.value = { message: "Anggota UPZ berhasil dihapus", type: "success" };
    loadPengurusZakat();
  } catch (error) {
    toast.value = { message: "Gagal menghapus anggota UPZ", type: "error" };
  }
};

const closeAnggotaModal = () => {
  showAnggotaModal.value = false;
  anggotaForm.value = { nama: "", telepon: "", alamat: "" };
};

const loadPengurusMasjid = async () => {
  try {
    const { data } = await api.getPengurusMasjid();
    if (data.success && data.data) {
      data.data.forEach(p => {
        if (p.jabatan === "Ketua") {
          pengurusMasjidForm.value.ketua = {
            nama: p.nama,
            telepon: p.telepon || "",
            alamat: p.alamat || ""
          };
        } else if (p.jabatan === "Sekretaris") {
          pengurusMasjidForm.value.sekretaris = {
            nama: p.nama,
            telepon: p.telepon || "",
            alamat: p.alamat || ""
          };
        }
      });
    }
  } catch (error) {
    console.error("Error loading pengurus masjid:", error);
  }
};

const loadPengurusZakat = async () => {
  try {
    const { data } = await api.getPengurusZakat();
    if (data.success && data.data) {
      const ketua = data.data.find(p => p.jabatan === "Ketua UPZ");
      if (ketua) {
        pengurusUPZForm.value.ketua = {
          nama: ketua.nama,
          telepon: ketua.telepon || "",
          alamat: ketua.alamat || ""
        };
      }
      anggotaUPZList.value = data.data.filter(p => p.jabatan === "Anggota UPZ");
    }
  } catch (error) {
    console.error("Error loading pengurus zakat:", error);
  }
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

  // Load pengurus data
  loadPengurusMasjid();
  loadPengurusZakat();
});
</script>
