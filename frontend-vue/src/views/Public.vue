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
            <i class="fas fa-sign-in-alt mr-2"></i>Back to Home
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
        <!-- Loading Skeleton -->
        <template v-if="isLoadingStats">
          <SkeletonCard v-for="i in 7" :key="i" type="stat" />
        </template>
        
        <!-- Actual Stats -->
        <div
          v-else
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
              <p v-if="stat.valueRp" class="text-xs text-gray-500 mt-1">
                {{ stat.valueRp }}
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

      <div v-if="lastUpdate" class="text-center mb-6">
        <p class="text-sm text-gray-500">
          <i class="fas fa-clock mr-1"></i>Data terakhir diperbarui: <span class="font-medium">{{ formatDate(lastUpdate) }}</span>
        </p>
      </div>

      <div class="bg-white rounded-xl shadow-lg p-6 mb-8">
        <h2 class="text-xl md:text-2xl font-bold text-gray-800 mb-6">
          <i class="fas fa-mosque mr-2 text-blue-600"></i>Daftar Masjid/Langgar
        </h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          <!-- Loading Skeleton -->
          <template v-if="isLoadingMasjid">
            <SkeletonCard v-for="i in 3" :key="i" type="card" />
          </template>
          
          <!-- Actual Masjid List -->
          <div
            v-else
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
              <p v-if="modalData?.last_update" class="text-xs opacity-75 mt-1">
                <i class="fas fa-clock mr-1"></i>Update terakhir: {{ formatDate(modalData.last_update) }}
              </p>
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
            <div class="bg-green-50 p-4 rounded-lg">
              <p class="text-sm text-green-600 font-medium">Fitrah (Uang)</p>
              <p class="text-lg font-bold text-green-700">
                {{ formatCurrency(modalData?.total_zakat_fitrah_uang || 0) }}
              </p>
            </div>
            <div class="bg-amber-50 p-4 rounded-lg">
              <p class="text-sm text-amber-600 font-medium">Fitrah (Beras)</p>
              <p class="text-lg font-bold text-amber-700">
                {{ (modalData?.total_zakat_fitrah_beras_kg || 0).toFixed(1) }} kg
              </p>
              <p v-if="modalData?.total_zakat_fitrah_beras_rupiah > 0" class="text-xs text-amber-600 mt-1">
                {{ formatCurrency(modalData?.total_zakat_fitrah_beras_rupiah) }}
              </p>
            </div>
            <div class="bg-indigo-50 p-4 rounded-lg">
              <p class="text-sm text-indigo-600 font-medium">Zakat Mal</p>
              <p class="text-lg font-bold text-indigo-700">
                {{ formatCurrency(modalData?.total_zakat_mal || 0) }}
              </p>
            </div>
            <div class="bg-yellow-50 p-4 rounded-lg">
              <p class="text-sm text-yellow-600 font-medium">Fidyah</p>
              <p class="text-lg font-bold text-yellow-700">
                {{ formatCurrency(modalData?.total_fidyah || 0) }}
              </p>
            </div>
            <div class="bg-pink-50 p-4 rounded-lg">
              <p class="text-sm text-pink-600 font-medium">Infaq</p>
              <p class="text-lg font-bold text-pink-700">
                {{ formatCurrency(modalData?.total_infaq || 0) }}
              </p>
            </div>
          </div>

          <!-- Tab Navigation -->
          <div class="flex border-b border-gray-200 mb-6">
            <button
              @click="activeTab = 'zakat'"
              :class="[
                'px-6 py-3 font-medium transition-colors',
                activeTab === 'zakat'
                  ? 'text-blue-600 border-b-2 border-blue-600'
                  : 'text-gray-600 hover:text-gray-800'
              ]"
            >
              <i class="fas fa-receipt mr-2"></i>Data Zakat
            </button>
            <button
              @click="activeTab = 'mustahiq'"
              :class="[
                'px-6 py-3 font-medium transition-colors',
                activeTab === 'mustahiq'
                  ? 'text-blue-600 border-b-2 border-blue-600'
                  : 'text-gray-600 hover:text-gray-800'
              ]"
            >
              <i class="fas fa-hand-holding-heart mr-2"></i>Mustahiq
            </button>
            <button
              @click="activeTab = 'pengurus'"
              :class="[
                'px-6 py-3 font-medium transition-colors',
                activeTab === 'pengurus'
                  ? 'text-blue-600 border-b-2 border-blue-600'
                  : 'text-gray-600 hover:text-gray-800'
              ]"
            >
              <i class="fas fa-users mr-2"></i>Pengurus
            </button>
          </div>

          <!-- Tab Content: Data Zakat -->
          <div v-show="activeTab === 'zakat'" class="mb-6">
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
                        Jumlah
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
                          :class="getJenisBadgeClass(t.jenis_zakat)"
                          >{{ t.jenis_zakat }}</span
                        >
                        <span v-if="t.bentuk_zakat === 'beras'" class="block text-xs text-amber-600 mt-1">
                          <i class="fas fa-seedling mr-1"></i>{{ t.keterangan || 'Beras' }}
                        </span>
                      </td>
                      <td class="px-4 py-2 text-sm text-gray-700">
                        {{ formatJumlah(t) }}
                      </td>
                      <td class="px-4 py-2 text-sm font-medium text-gray-900">
                        {{ formatTotal(t) }}
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
                    <div>
                      <p class="font-medium text-gray-900">
                        {{ t.muzakki_nama }}
                      </p>
                      <span v-if="t.bentuk_zakat === 'beras'" class="text-xs text-amber-600 mt-1 block">
                        <i class="fas fa-seedling mr-1"></i>{{ t.keterangan || 'Beras' }}
                      </span>
                    </div>
                    <span
                      :class="getJenisBadgeClass(t.jenis_zakat)"
                      >{{ t.jenis_zakat }}</span
                    >
                  </div>
                  <div class="space-y-1 text-sm">
                    <div v-if="formatJumlah(t) !== '-'" class="flex justify-between">
                      <span class="text-gray-600">Jumlah:</span
                      ><span class="font-medium">{{ formatJumlah(t) }}</span>
                    </div>
                    <div class="flex justify-between">
                      <span class="text-gray-600">Total:</span
                      ><span class="font-medium">{{ formatTotal(t) }}</span>
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

          <!-- Tab Content: Data Mustahiq -->
          <div v-show="activeTab === 'mustahiq'">
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

          <!-- Tab Content: Pengurus -->
          <div v-show="activeTab === 'pengurus'" class="mb-6">
            <h3 class="text-lg font-bold text-gray-800 mb-3">
              <i class="fas fa-users mr-2 text-blue-600"></i>Struktur Pengurus
            </h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <!-- Pengurus Masjid -->
              <div class="bg-blue-50 rounded-lg p-4">
                <h4 class="font-semibold text-blue-800 mb-3">Pengurus Masjid/Langgar</h4>
                <div v-if="pengurusMasjidData.length > 0" class="space-y-2">
                  <div v-for="p in pengurusMasjidData" :key="p.id" class="bg-white rounded p-3">
                    <p class="text-xs text-gray-600 mb-1">{{ p.jabatan }}</p>
                    <p class="font-medium text-gray-900">{{ p.nama }}</p>
                    <p v-if="p.telepon" class="text-sm text-gray-600 mt-1">
                      <i class="fas fa-phone mr-1"></i>{{ p.telepon }}
                    </p>
                  </div>
                </div>
                <p v-else class="text-gray-500 text-sm text-center py-3">Belum ada data pengurus</p>
              </div>

              <!-- Pengurus UPZ -->
              <div class="bg-green-50 rounded-lg p-4">
                <h4 class="font-semibold text-green-800 mb-3">Pengurus UPZ</h4>
                <div v-if="pengurusZakatData.length > 0" class="space-y-2">
                  <div v-for="p in pengurusZakatData" :key="p.id" class="bg-white rounded p-3">
                    <p class="text-xs text-gray-600 mb-1">{{ p.jabatan }}</p>
                    <p class="font-medium text-gray-900">{{ p.nama }}</p>
                    <p v-if="p.telepon" class="text-sm text-gray-600 mt-1">
                      <i class="fas fa-phone mr-1"></i>{{ p.telepon }}
                    </p>
                  </div>
                </div>
                <p v-else class="text-gray-500 text-sm text-center py-3">Belum ada data pengurus UPZ</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <LoadingOverlay :show="isLoadingModal" message="Memuat data masjid..." />
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import SkeletonCard from "../components/SkeletonCard.vue";
import LoadingOverlay from "../components/LoadingOverlay.vue";
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
    label: "Zakat Fitrah (Uang)",
    value: "Rp 0",
    icon: "fas fa-money-bill-wave",
    bgColor: "bg-green-100",
    textColor: "text-green-600",
  },
  {
    label: "Zakat Fitrah (Beras)",
    value: "0 kg",
    valueRp: null,
    icon: "fas fa-seedling",
    bgColor: "bg-amber-100",
    textColor: "text-amber-600",
  },
  {
    label: "Zakat Mal",
    value: "Rp 0",
    icon: "fas fa-wallet",
    bgColor: "bg-indigo-100",
    textColor: "text-indigo-600",
  },
  {
    label: "Fidyah",
    value: "Rp 0",
    icon: "fas fa-moon",
    bgColor: "bg-yellow-100",
    textColor: "text-yellow-600",
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
const lastUpdate = ref(null);
const showModal = ref(false);
const modalData = ref(null);
const activeTab = ref('zakat');
const pengurusMasjidData = ref([]);
const pengurusZakatData = ref([]);
const isLoadingStats = ref(true);
const isLoadingMasjid = ref(true);
const isLoadingModal = ref(false);

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

const getJenisBadgeClass = (jenis) => {
  switch (jenis) {
    case 'fitrah':
      return 'px-2 py-1 bg-green-100 text-green-800 rounded-full text-xs'
    case 'mal':
      return 'px-2 py-1 bg-indigo-100 text-indigo-800 rounded-full text-xs'
    case 'fidyah':
      return 'px-2 py-1 bg-yellow-100 text-yellow-800 rounded-full text-xs'
    case 'infaq':
      return 'px-2 py-1 bg-pink-100 text-pink-800 rounded-full text-xs'
    default:
      return 'px-2 py-1 bg-gray-100 text-gray-800 rounded-full text-xs'
  }
}

const parseHariFromKeterangan = (keterangan) => {
  if (!keterangan) return 0
  const match = keterangan.match(/(\d+)\s*hari/i)
  return match ? Number(match[1]) : 0
}

const parseKgFromKeterangan = (keterangan) => {
  if (!keterangan) return 0
  const match = keterangan.match(/(?:Beras|Fidyah Beras):\s*([\d.]+)\s*kg/i)
  return match ? Number(match[1]) : 0
}

const formatJumlah = (t) => {
  if (!t) return '-'
  if (t.jenis_zakat === 'fidyah' && t.bentuk_zakat === 'beras') {
    const hari = Number(t.jumlah_hari_fidyah || 0) || parseHariFromKeterangan(t.keterangan)
    return hari > 0 ? `${hari} hari` : '-'
  }
  return t.jumlah_orang > 0 ? `${t.jumlah_orang} orang` : '-'
}

const formatTotal = (t) => {
  if (!t) return '-'
  if (t.bentuk_zakat === 'beras') {
    const kg = parseKgFromKeterangan(t.keterangan)
    return kg > 0 ? `${kg.toLocaleString('id-ID')} Kg` : '-'
  }
  return formatCurrency(t.total_dibayar)
}

const viewMasjidStats = async (id) => {
  try {
    isLoadingModal.value = true;
    activeTab.value = 'zakat'; // Reset to first tab
    const [statsData, pengurusMasjid, pengurusZakat] = await Promise.all([
      api.getPublicMasjidStats(id),
      api.getPublicPengurusMasjid(id),
      api.getPublicPengurusZakat(id)
    ]);
    
    if (statsData.data.success) {
      modalData.value = statsData.data.data;
      showModal.value = true;
    }
    
    if (pengurusMasjid.data.success) {
      pengurusMasjidData.value = pengurusMasjid.data.data || [];
    }
    
    if (pengurusZakat.data.success) {
      pengurusZakatData.value = pengurusZakat.data.data || [];
    }
  } catch (error) {
    console.error("Error:", error);
  } finally {
    isLoadingModal.value = false;
  }
};

onMounted(async () => {
  isLoadingStats.value = true;
  isLoadingMasjid.value = true;
  try {
    const [dashboardData, masjidData] = await Promise.all([
      api.getPublicDashboard(),
      api.getPublicMasjid(),
    ]);

    if (dashboardData.data.success) {
      const d = dashboardData.data.data;
      console.log("Dashboard data:", d); // Debug log
      stats.value[0].value = d.total_masjid || 0;
      stats.value[1].value = d.total_muzakki || 0;
      stats.value[2].value = d.total_mustahiq || 0;
      
      // Backward compatibility: gunakan field lama jika field baru belum ada
      const fitrahUang = d.total_zakat_fitrah_uang !== undefined ? d.total_zakat_fitrah_uang : d.total_zakat_fitrah || 0;
      const fitrahBerasKg = d.total_zakat_fitrah_beras_kg || 0;
      const fitrahBerasRp = d.total_zakat_fitrah_beras_rupiah || 0;
      
      stats.value[3].value = formatCurrency(fitrahUang);
      stats.value[4].value = `${fitrahBerasKg.toFixed(1)} kg`;
      if (fitrahBerasRp > 0) {
        stats.value[4].valueRp = formatCurrency(fitrahBerasRp);
      }
      stats.value[5].value = formatCurrency(d.total_zakat_mal || 0);
      stats.value[6].value = formatCurrency(d.total_fidyah || 0);
      stats.value[7].value = formatCurrency(d.total_infaq || 0);
      
      // Set last update
      lastUpdate.value = d.last_update || null;
    }

    if (masjidData.data.success) {
      console.log("Masjid data:", masjidData.data.data); // Debug log
      masjidList.value = masjidData.data.data || [];
    }
  } catch (error) {
    console.error("Error loading dashboard:", error);
    console.error("Error details:", error.response || error.message);
  } finally {
    isLoadingStats.value = false;
    isLoadingMasjid.value = false;
  }
});
</script>
