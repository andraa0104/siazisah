<template>
  <div class="bg-gray-50 min-h-screen">
    <nav class="bg-white shadow-md sticky top-0 z-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center h-16">
          <div class="flex items-center">
            <div
              class="w-10 h-10 bg-gradient-to-br from-blue-600 to-indigo-600 rounded-lg flex items-center justify-center"
            >
              <i class="fas fa-mosque text-white"></i>
            </div>
            <span class="ml-3 text-xl font-bold text-gray-800">SI-AZISAH</span>
          </div>
          <router-link
            to="/"
            class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition"
          >
            <i class="fas fa-sign-in-alt mr-2"></i>Login
          </router-link>
        </div>
      </div>
    </nav>

    <div class="bg-gradient-to-r from-blue-600 to-indigo-600 text-white py-12">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl md:text-4xl font-bold mb-4">
          Dashboard Publik Zakat
        </h1>
        <p class="text-lg md:text-xl opacity-90">
          Transparansi Pengelolaan Zakat di Desa Purwajaya
        </p>
      </div>
    </div>

    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 -mt-8">
      <div
        class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 md:gap-6 mb-8"
      >
        <div
          v-for="stat in stats"
          :key="stat.label"
          class="bg-white rounded-xl shadow-lg p-6 transform hover:scale-105 transition"
        >
          <div class="flex items-center justify-between">
            <div>
              <p class="text-gray-500 text-sm font-medium">{{ stat.label }}</p>
              <p class="text-2xl md:text-3xl font-bold text-gray-800 mt-2">
                {{ stat.value }}
              </p>
            </div>
            <div
              :class="`w-14 h-14 ${stat.bgColor} rounded-full flex items-center justify-center`"
            >
              <i :class="`${stat.icon} ${stat.textColor} text-2xl`"></i>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-xl shadow-lg p-6 mb-8">
        <h2 class="text-xl md:text-2xl font-bold text-gray-800 mb-6">
          <i class="fas fa-mosque mr-2 text-blue-600"></i>Daftar Masjid/Langgar
        </h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          <div
            v-for="masjid in masjidList"
            :key="masjid.id"
            @click="viewMasjidStats(masjid.id)"
            class="border border-gray-200 rounded-lg p-4 hover:shadow-lg transition cursor-pointer"
          >
            <h3 class="font-bold text-lg text-gray-800 mb-2">
              {{ masjid.nama }}
            </h3>
            <p class="text-sm text-gray-600 mb-1">
              <i class="fas fa-map-marker-alt mr-2 text-blue-600"></i
              >{{ masjid.alamat }}
            </p>
            <p v-if="masjid.telepon" class="text-sm text-gray-600">
              <i class="fas fa-phone mr-2 text-blue-600"></i
              >{{ masjid.telepon }}
            </p>
            <button
              class="mt-3 w-full bg-blue-50 text-blue-600 py-2 rounded-lg hover:bg-blue-100 transition text-sm font-medium"
            >
              Lihat Detail <i class="fas fa-arrow-right ml-1"></i>
            </button>
          </div>
        </div>
      </div>
    </div>

    <footer class="bg-gray-800 text-white py-8 mt-12">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <p>© 2024 SI-AZISAH - Sistem Informasi Zakat Desa Purwajaya</p>
      </div>
    </footer>

    <div
      v-if="showModal"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4"
      @click.self="showModal = false"
    >
      <div
        class="bg-white rounded-xl shadow-2xl max-w-4xl w-full max-h-[90vh] overflow-y-auto"
      >
        <div
          class="sticky top-0 bg-gradient-to-r from-blue-600 to-indigo-600 text-white p-6 rounded-t-xl"
        >
          <div class="flex justify-between items-center">
            <div>
              <h2 class="text-2xl font-bold">{{ modalData?.masjid.nama }}</h2>
              <p class="text-sm opacity-90">{{ modalData?.masjid.alamat }}</p>
            </div>
            <button
              @click="showModal = false"
              class="text-white hover:bg-white hover:bg-opacity-20 rounded-full w-8 h-8"
            >
              <i class="fas fa-times"></i>
            </button>
          </div>
        </div>

        <div class="p-6">
          <div class="grid grid-cols-2 md:grid-cols-3 gap-4 mb-6">
            <div class="bg-green-50 p-4 rounded-lg">
              <p class="text-sm text-green-600 font-medium">Total Muzakki</p>
              <p class="text-2xl font-bold text-green-700">
                {{ modalData?.total_muzakki }}
              </p>
            </div>
            <div class="bg-purple-50 p-4 rounded-lg">
              <p class="text-sm text-purple-600 font-medium">Total Mustahiq</p>
              <p class="text-2xl font-bold text-purple-700">
                {{ modalData?.total_mustahiq }}
              </p>
            </div>
            <div class="bg-yellow-50 p-4 rounded-lg">
              <p class="text-sm text-yellow-600 font-medium">Zakat Fitrah</p>
              <p class="text-lg font-bold text-yellow-700">
                {{ formatCurrency(modalData?.total_zakat_fitrah || 0) }}
              </p>
            </div>
            <div class="bg-indigo-50 p-4 rounded-lg">
              <p class="text-sm text-indigo-600 font-medium">Zakat Mal</p>
              <p class="text-lg font-bold text-indigo-700">
                {{ formatCurrency(modalData?.total_zakat_mal || 0) }}
              </p>
            </div>
            <div class="bg-pink-50 p-4 rounded-lg">
              <p class="text-sm text-pink-600 font-medium">Infaq</p>
              <p class="text-lg font-bold text-pink-700">
                {{ formatCurrency(modalData?.total_infaq || 0) }}
              </p>
            </div>
          </div>

          <div class="mb-6">
            <h3 class="text-lg font-bold text-gray-800 mb-3">
              <i class="fas fa-receipt mr-2 text-blue-600"></i>Data Transaksi
              Zakat
            </h3>
            <div v-if="modalData?.transaksi?.length > 0">
              <!-- Desktop Table -->
              <div class="hidden md:block overflow-x-auto">
                <table class="min-w-full divide-y divide-gray-200">
                  <thead class="bg-gray-50">
                    <tr>
                      <th
                        class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase"
                      >
                        Muzakki
                      </th>
                      <th
                        class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase"
                      >
                        Jenis
                      </th>
                      <th
                        class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase"
                      >
                        Total
                      </th>
                      <th
                        class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase"
                      >
                        Infaq
                      </th>
                      <th
                        class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase"
                      >
                        Tanggal
                      </th>
                    </tr>
                  </thead>
                  <tbody class="bg-white divide-y divide-gray-200">
                    <tr
                      v-for="t in modalData?.transaksi"
                      :key="t.id"
                      class="hover:bg-gray-50"
                    >
                      <td class="px-4 py-2 text-sm text-gray-900">
                        {{ t.muzakki_nama }}
                      </td>
                      <td class="px-4 py-2 text-sm">
                        <span
                          class="px-2 py-1 bg-blue-100 text-blue-800 rounded-full text-xs"
                          >{{ t.jenis_zakat }}</span
                        >
                      </td>
                      <td class="px-4 py-2 text-sm font-medium text-gray-900">
                        {{ formatCurrency(t.total_dibayar) }}
                      </td>
                      <td class="px-4 py-2 text-sm text-green-600 font-medium">
                        {{
                          t.infaq_tambahan > 0
                            ? formatCurrency(t.infaq_tambahan)
                            : "-"
                        }}
                      </td>
                      <td class="px-4 py-2 text-sm text-gray-600">
                        {{ formatDate(t.tanggal_bayar) }}
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
              <!-- Mobile List -->
              <div class="md:hidden space-y-3">
                <div
                  v-for="t in modalData?.transaksi"
                  :key="t.id"
                  class="bg-gray-50 rounded-lg p-4"
                >
                  <div class="flex justify-between items-start mb-2">
                    <p class="font-medium text-gray-900">
                      {{ t.muzakki_nama }}
                    </p>
                    <span
                      class="px-2 py-1 bg-blue-100 text-blue-800 rounded-full text-xs"
                      >{{ t.jenis_zakat }}</span
                    >
                  </div>
                  <div class="space-y-1 text-sm">
                    <div class="flex justify-between">
                      <span class="text-gray-600">Total:</span
                      ><span class="font-medium">{{
                        formatCurrency(t.total_dibayar)
                      }}</span>
                    </div>
                    <div
                      v-if="t.infaq_tambahan > 0"
                      class="flex justify-between"
                    >
                      <span class="text-gray-600">Infaq:</span
                      ><span class="text-green-600 font-medium">{{
                        formatCurrency(t.infaq_tambahan)
                      }}</span>
                    </div>
                    <div class="flex justify-between">
                      <span class="text-gray-600">Tanggal:</span
                      ><span>{{ formatDate(t.tanggal_bayar) }}</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <p v-else class="text-gray-500 text-center py-4">
              Belum ada transaksi
            </p>
          </div>

          <div>
            <h3 class="text-lg font-bold text-gray-800 mb-3">
              <i class="fas fa-hand-holding-heart mr-2 text-purple-600"></i>Data
              Mustahiq
            </h3>
            <div v-if="modalData?.mustahiq?.length > 0">
              <!-- Desktop Table -->
              <div class="hidden md:block overflow-x-auto">
                <table class="min-w-full divide-y divide-gray-200">
                  <thead class="bg-gray-50">
                    <tr>
                      <th
                        class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase"
                      >
                        Nama
                      </th>
                      <th
                        class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase"
                      >
                        Alamat
                      </th>
                      <th
                        class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase"
                      >
                        RT
                      </th>
                      <th
                        class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase"
                      >
                        Kategori
                      </th>
                      <th
                        class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase"
                      >
                        Keterangan
                      </th>
                    </tr>
                  </thead>
                  <tbody class="bg-white divide-y divide-gray-200">
                    <tr
                      v-for="m in modalData?.mustahiq"
                      :key="m.id"
                      class="hover:bg-gray-50"
                    >
                      <td class="px-4 py-2 text-sm font-medium text-gray-900">
                        {{ m.nama }}
                      </td>
                      <td class="px-4 py-2 text-sm text-gray-600">
                        {{ m.alamat }}
                      </td>
                      <td class="px-4 py-2 text-sm text-gray-600">
                        {{ m.rt }}
                      </td>
                      <td class="px-4 py-2 text-sm">
                        <span
                          class="px-2 py-1 bg-purple-100 text-purple-700 rounded-full text-xs"
                          >{{ m.kategori }}</span
                        >
                      </td>
                      <td class="px-4 py-2 text-sm text-gray-600">
                        {{ m.keterangan || "-" }}
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
              <!-- Mobile List -->
              <div class="md:hidden space-y-3">
                <div
                  v-for="m in modalData?.mustahiq"
                  :key="m.id"
                  class="bg-gray-50 rounded-lg p-4"
                >
                  <p class="font-medium text-gray-900 mb-1">{{ m.nama }}</p>
                  <p class="text-sm text-gray-600 mb-1">{{ m.alamat }}</p>
                  <p class="text-sm text-gray-600 mb-2">RT {{ m.rt }}</p>
                  <div class="flex items-center gap-2">
                    <span
                      class="inline-block px-2 py-1 bg-purple-100 text-purple-700 rounded-full text-xs"
                      >{{ m.kategori }}</span
                    >
                    <span v-if="m.keterangan" class="text-xs text-gray-500">{{
                      m.keterangan
                    }}</span>
                  </div>
                </div>
              </div>
            </div>
            <p v-else class="text-gray-500 text-center py-4">
              Belum ada data mustahiq
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import api from "../api";

const stats = ref([
  {
    label: "Total Masjid/Langgar",
    value: 0,
    icon: "fas fa-mosque",
    bgColor: "bg-blue-100",
    textColor: "text-blue-600",
  },
  {
    label: "Total Muzakki",
    value: 0,
    icon: "fas fa-users",
    bgColor: "bg-green-100",
    textColor: "text-green-600",
  },
  {
    label: "Total Mustahiq",
    value: 0,
    icon: "fas fa-hand-holding-heart",
    bgColor: "bg-purple-100",
    textColor: "text-purple-600",
  },
  {
    label: "Zakat Fitrah",
    value: "Rp 0",
    icon: "fas fa-coins",
    bgColor: "bg-yellow-100",
    textColor: "text-yellow-600",
  },
  {
    label: "Zakat Mal",
    value: "Rp 0",
    icon: "fas fa-wallet",
    bgColor: "bg-indigo-100",
    textColor: "text-indigo-600",
  },
  {
    label: "Infaq",
    value: "Rp 0",
    icon: "fas fa-donate",
    bgColor: "bg-pink-100",
    textColor: "text-pink-600",
  },
]);

const masjidList = ref([]);
const showModal = ref(false);
const modalData = ref(null);

const formatCurrency = (amount) => {
  return new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
    minimumFractionDigits: 0,
  }).format(amount);
};

const formatDate = (date) => {
  return new Date(date).toLocaleDateString("id-ID");
};

const viewMasjidStats = async (id) => {
  try {
    const { data } = await api.getPublicMasjidStats(id);
    if (data.success) {
      modalData.value = data.data;
      showModal.value = true;
    }
  } catch (error) {
    console.error("Error:", error);
  }
};

onMounted(async () => {
  try {
    const [dashboardData, masjidData] = await Promise.all([
      api.getPublicDashboard(),
      api.getPublicMasjid(),
    ]);

    if (dashboardData.data.success) {
      const d = dashboardData.data.data;
      stats.value[0].value = d.total_masjid;
      stats.value[1].value = d.total_muzakki;
      stats.value[2].value = d.total_mustahiq;
      stats.value[3].value = formatCurrency(d.total_zakat_fitrah);
      stats.value[4].value = formatCurrency(d.total_zakat_mal);
      stats.value[5].value = formatCurrency(d.total_infaq);
    }

    if (masjidData.data.success) {
      masjidList.value = masjidData.data.data;
    }
  } catch (error) {
    console.error(error);
  }
});
</script>
