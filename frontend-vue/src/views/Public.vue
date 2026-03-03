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
          <i class="fas fa-clock mr-1"></i>Data terakhir diperbarui:
          <span class="font-medium">{{ formatDate(lastUpdate) }}</span>
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
        <p>© 2026 SI-AZISAH - Sistem Informasi Zakat Desa Purwajaya</p>
      </div>
    </footer>

    <div
      v-if="showModal"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-2 sm:p-4"
      @click.self="showModal = false"
    >
      <div
        class="bg-white rounded-lg sm:rounded-xl shadow-2xl max-w-4xl w-full max-h-[95vh] sm:max-h-[90vh] overflow-y-auto"
      >
        <div
          class="sticky top-0 bg-gradient-to-r from-blue-600 to-indigo-600 text-white p-3 sm:p-6 rounded-t-lg sm:rounded-t-xl"
        >
          <div class="flex justify-between items-start gap-2">
            <div class="min-w-0 flex-1">
              <h2 class="text-lg sm:text-2xl font-bold truncate">
                {{ modalData?.masjid.nama }}
              </h2>
              <p class="text-xs sm:text-sm opacity-90 line-clamp-2">
                {{ modalData?.masjid.alamat }}
              </p>
              <p v-if="modalData?.last_update" class="text-xs opacity-75 mt-1">
                <i class="fas fa-clock mr-1"></i>
                <span class="hidden sm:inline">Update terakhir: </span>
                {{ formatDate(modalData.last_update) }}
              </p>
            </div>
            <button
              @click="showModal = false"
              class="flex-shrink-0 text-white hover:bg-white hover:bg-opacity-20 rounded-full w-7 h-7 sm:w-8 sm:h-8 flex items-center justify-center"
            >
              <i class="fas fa-times text-base sm:text-lg"></i>
            </button>
          </div>
        </div>

        <div class="p-3 sm:p-6">
          <div
            class="grid grid-cols-2 lg:grid-cols-5 gap-2 sm:gap-3 md:gap-4 mb-4 sm:mb-6"
          >
            <div class="bg-green-50 p-2 sm:p-3 md:p-4 rounded-lg">
              <p class="text-xs sm:text-sm text-green-600 font-medium">
                Total Muzakki
              </p>
              <p class="text-xl sm:text-2xl font-bold text-green-700">
                {{ modalData?.total_muzakki }}
              </p>
            </div>
            <div class="bg-orange-50 p-2 sm:p-3 md:p-4 rounded-lg">
              <p
                class="text-xs sm:text-sm text-orange-600 font-medium line-clamp-2"
              >
                Total Orang Dizakati
              </p>
              <p class="text-xl sm:text-2xl font-bold text-orange-700">
                {{ modalData?.total_orang_dizakati }}
              </p>
            </div>
            <div class="bg-purple-50 p-2 sm:p-3 md:p-4 rounded-lg">
              <p class="text-xs sm:text-sm text-purple-600 font-medium">
                Total Mustahiq
              </p>
              <p class="text-xl sm:text-2xl font-bold text-purple-700">
                {{ modalData?.total_mustahiq }}
              </p>
            </div>
            <div class="bg-green-50 p-2 sm:p-3 md:p-4 rounded-lg">
              <p class="text-xs sm:text-sm text-green-600 font-medium">
                Fitrah (Uang)
              </p>
              <p
                class="text-sm sm:text-base md:text-lg font-bold text-green-700 line-clamp-2"
              >
                {{ formatCurrency(modalData?.total_zakat_fitrah_uang || 0) }}
              </p>
            </div>
            <div class="bg-amber-50 p-2 sm:p-3 md:p-4 rounded-lg">
              <p class="text-xs sm:text-sm text-amber-600 font-medium">
                Fitrah (Beras)
              </p>
              <p
                class="text-sm sm:text-base md:text-lg font-bold text-amber-700"
              >
                {{ (modalData?.total_zakat_fitrah_beras_kg || 0).toFixed(1)
                }}<span class="text-xs">kg</span>
              </p>
              <p
                v-if="modalData?.total_zakat_fitrah_beras_rupiah > 0"
                class="text-xs text-amber-600 mt-0.5 line-clamp-1"
              >
                {{ formatCurrency(modalData?.total_zakat_fitrah_beras_rupiah) }}
              </p>
            </div>
            <div class="bg-indigo-50 p-2 sm:p-3 md:p-4 rounded-lg">
              <p class="text-xs sm:text-sm text-indigo-600 font-medium">
                Zakat Mal
              </p>
              <p
                class="text-sm sm:text-base md:text-lg font-bold text-indigo-700 line-clamp-2"
              >
                {{ formatCurrency(modalData?.total_zakat_mal || 0) }}
              </p>
            </div>
            <div class="bg-yellow-50 p-2 sm:p-3 md:p-4 rounded-lg">
              <p class="text-xs sm:text-sm text-yellow-600 font-medium">
                Fidyah (Uang)
              </p>
              <p
                class="text-sm sm:text-base md:text-lg font-bold text-yellow-700 line-clamp-2"
              >
                {{
                  formatCurrency(
                    modalData?.total_fidyah_uang ||
                      modalData?.total_fidyah ||
                      0,
                  )
                }}
              </p>
            </div>
            <div class="bg-lime-50 p-2 sm:p-3 md:p-4 rounded-lg">
              <p class="text-xs sm:text-sm text-lime-600 font-medium">
                Fidyah (Beras)
              </p>
              <p
                class="text-sm sm:text-base md:text-lg font-bold text-lime-700"
              >
                {{ (modalData?.total_fidyah_beras_kg || 0).toFixed(1)
                }}<span class="text-xs">kg</span>
              </p>
              <p
                v-if="modalData?.total_fidyah_beras_rupiah > 0"
                class="text-xs text-lime-600 mt-0.5 line-clamp-1"
              >
                {{ formatCurrency(modalData?.total_fidyah_beras_rupiah) }}
              </p>
            </div>
            <div class="bg-pink-50 p-2 sm:p-3 md:p-4 rounded-lg">
              <p class="text-xs sm:text-sm text-pink-600 font-medium">Infaq</p>
              <p
                class="text-sm sm:text-base md:text-lg font-bold text-pink-700 line-clamp-2"
              >
                {{ formatCurrency(modalData?.total_infaq || 0) }}
              </p>
            </div>
          </div>

          <!-- Tab Navigation -->
          <div
            class="flex gap-1 sm:gap-0 border-b border-gray-200 mb-4 sm:mb-6 overflow-x-auto"
          >
            <button
              @click="activeTab = 'zakat'"
              :class="[
                'px-2 sm:px-6 py-2 sm:py-3 text-xs sm:text-base font-medium transition-colors whitespace-nowrap',
                activeTab === 'zakat'
                  ? 'text-blue-600 border-b-2 border-blue-600'
                  : 'text-gray-600 hover:text-gray-800',
              ]"
            >
              <i class="fas fa-receipt mr-1 sm:mr-2 text-xs sm:text-base"></i
              ><span class="hidden sm:inline">Data Zakat</span
              ><span class="sm:hidden">Zakat</span>
            </button>
            <button
              @click="activeTab = 'mustahiq'"
              :class="[
                'px-2 sm:px-6 py-2 sm:py-3 text-xs sm:text-base font-medium transition-colors whitespace-nowrap',
                activeTab === 'mustahiq'
                  ? 'text-blue-600 border-b-2 border-blue-600'
                  : 'text-gray-600 hover:text-gray-800',
              ]"
            >
              <i
                class="fas fa-hand-holding-heart mr-1 sm:mr-2 text-xs sm:text-base"
              ></i
              >Mustahiq
            </button>
            <button
              @click="activeTab = 'pengurus'"
              :class="[
                'px-2 sm:px-6 py-2 sm:py-3 text-xs sm:text-base font-medium transition-colors whitespace-nowrap',
                activeTab === 'pengurus'
                  ? 'text-blue-600 border-b-2 border-blue-600'
                  : 'text-gray-600 hover:text-gray-800',
              ]"
            >
              <i class="fas fa-users mr-1 sm:mr-2 text-xs sm:text-base"></i
              >Pengurus
            </button>
            <button
              @click="activeTab = 'distribusi'"
              :class="[
                'px-2 sm:px-6 py-2 sm:py-3 text-xs sm:text-base font-medium transition-colors whitespace-nowrap',
                activeTab === 'distribusi'
                  ? 'text-blue-600 border-b-2 border-blue-600'
                  : 'text-gray-600 hover:text-gray-800',
              ]"
            >
              <i
                class="fas fa-hand-holding-usd mr-1 sm:mr-2 text-xs sm:text-base"
              ></i
              ><span class="hidden sm:inline">Distribusi</span
              ><span class="sm:hidden">Dist</span>
            </button>
          </div>

          <!-- Tab Content: Data Zakat -->
          <div v-show="activeTab === 'zakat'" class="mb-4 sm:mb-6">
            <h3
              class="text-base sm:text-lg font-bold text-gray-800 mb-2 sm:mb-3"
            >
              <i class="fas fa-receipt mr-2 text-blue-600"></i
              ><span class="hidden sm:inline">Data Transaksi</span>Zakat
            </h3>
            <div
              class="bg-white border border-gray-200 rounded-lg p-2 sm:p-3 mb-3 sm:mb-4 overflow-x-auto"
            >
              <div
                class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-2 sm:gap-3 min-w-max sm:min-w-0"
              >
                <div>
                  <label class="block text-xs text-gray-500 mb-1"
                    >Filter Jenis Zakat</label
                  >
                  <select
                    v-model="transaksiFilters.jenis_zakat"
                    @change="onPublicJenisFilterChange"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm"
                  >
                    <option value="">Semua Jenis</option>
                    <option value="fitrah">Zakat Fitrah</option>
                    <option value="mal">Zakat Mal</option>
                    <option value="fidyah">Fidyah</option>
                    <option value="infaq">Infaq</option>
                  </select>
                </div>
                <div>
                  <label class="block text-xs text-gray-500 mb-1"
                    >Bentuk (Fitrah/Fidyah)</label
                  >
                  <select
                    v-model="transaksiFilters.bentuk_zakat"
                    :disabled="!showPublicBentukFilter"
                    @change="applyPublicTransaksiFilters"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm disabled:bg-gray-100 disabled:text-gray-400"
                  >
                    <option value="">Semua Bentuk</option>
                    <option value="uang">Uang</option>
                    <option value="beras">Beras</option>
                  </select>
                </div>
                <div class="md:col-span-2">
                  <label class="block text-xs text-gray-500 mb-1"
                    >Cari Muzakki</label
                  >
                  <div class="flex gap-2">
                    <input
                      v-model="transaksiFilters.q"
                      @keyup.enter="applyPublicTransaksiFilters"
                      placeholder="Ketik nama muzakki..."
                      class="flex-1 px-3 py-2 border border-gray-300 rounded-lg text-sm"
                    />
                    <button
                      @click="applyPublicTransaksiFilters"
                      class="px-4 py-2 bg-green-600 text-white rounded-lg text-sm hover:bg-green-700"
                    >
                      Cari
                    </button>
                    <button
                      @click="resetPublicTransaksiFilters"
                      class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg text-sm hover:bg-gray-200"
                    >
                      Reset
                    </button>
                  </div>
                </div>
              </div>
            </div>
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
                        Total Uang
                      </th>
                      <th
                        class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase"
                      >
                        Total Beras (kg)
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
                        <span :class="getJenisBadgeClass(t.jenis_zakat)">{{
                          t.jenis_zakat
                        }}</span>
                        <span
                          v-if="t.bentuk_zakat === 'beras'"
                          class="block text-xs text-amber-600 mt-1"
                        >
                          <i class="fas fa-seedling mr-1"></i
                          >{{ t.keterangan || "Beras" }}
                        </span>
                      </td>
                      <td class="px-4 py-2 text-sm text-gray-700">
                        {{ formatJumlah(t) }}
                      </td>
                      <td class="px-4 py-2 text-sm font-medium text-gray-900">
                        {{ formatTotalUang(t) }}
                      </td>
                      <td class="px-4 py-2 text-sm text-gray-900">
                        {{ formatTotalBerasKg(t) }}
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
                      <span
                        v-if="t.bentuk_zakat === 'beras'"
                        class="text-xs text-amber-600 mt-1 block"
                      >
                        <i class="fas fa-seedling mr-1"></i
                        >{{ t.keterangan || "Beras" }}
                      </span>
                    </div>
                    <span :class="getJenisBadgeClass(t.jenis_zakat)">{{
                      t.jenis_zakat
                    }}</span>
                  </div>
                  <div class="space-y-1 text-sm">
                    <div
                      v-if="formatJumlah(t) !== '-'"
                      class="flex justify-between"
                    >
                      <span class="text-gray-600">Jumlah:</span
                      ><span class="font-medium">{{ formatJumlah(t) }}</span>
                    </div>
                    <div class="flex justify-between">
                      <span class="text-gray-600">Total Uang:</span
                      ><span class="font-medium">{{ formatTotalUang(t) }}</span>
                    </div>
                    <div class="flex justify-between">
                      <span class="text-gray-600">Total Beras:</span
                      ><span class="font-medium">{{
                        formatTotalBerasKg(t)
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
              <div
                class="mt-4 flex flex-col md:flex-row md:items-center md:justify-between gap-3 text-sm text-gray-600"
              >
                <div>
                  Menampilkan {{ modalData.transaksi.length }} dari
                  {{ transaksiPagination.total }} data
                </div>
                <div class="flex flex-wrap items-center gap-2">
                  <label class="text-xs text-gray-500">Per halaman</label>
                  <select
                    v-model="transaksiPagination.pageSize"
                    @change="onTransaksiPageSizeChange"
                    class="px-2 py-1 border border-gray-300 rounded-lg text-sm"
                  >
                    <option
                      v-for="size in pageSizeOptions"
                      :key="size"
                      :value="size"
                    >
                      {{ size === "all" ? "Semua" : size }}
                    </option>
                  </select>
                  <button
                    class="px-3 py-1 border border-gray-300 rounded-lg disabled:opacity-50"
                    :disabled="
                      transaksiPagination.isAll || transaksiPagination.page <= 1
                    "
                    @click="setTransaksiPage(transaksiPagination.page - 1)"
                  >
                    Prev
                  </button>
                  <span class="text-xs text-gray-500"
                    >{{ transaksiPagination.page }} /
                    {{ transaksiPagination.totalPages }}</span
                  >
                  <button
                    class="px-3 py-1 border border-gray-300 rounded-lg disabled:opacity-50"
                    :disabled="
                      transaksiPagination.isAll ||
                      transaksiPagination.page >= transaksiPagination.totalPages
                    "
                    @click="setTransaksiPage(transaksiPagination.page + 1)"
                  >
                    Next
                  </button>
                </div>
              </div>
            </div>
            <p v-else class="text-gray-500 text-center py-4">
              Belum ada transaksi
            </p>
          </div>

          <!-- Tab Content: Data Mustahiq -->
          <div v-show="activeTab === 'mustahiq'" class="mb-4 sm:mb-6">
            <h3
              class="text-base sm:text-lg font-bold text-gray-800 mb-2 sm:mb-3"
            >
              <i class="fas fa-hand-holding-heart mr-2 text-purple-600"></i>Data
              Mustahiq
            </h3>
            <div
              class="bg-white border border-gray-200 rounded-lg p-2 sm:p-3 mb-3 sm:mb-4 overflow-x-auto"
            >
              <div
                class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-2 sm:gap-3 min-w-max sm:min-w-0"
              >
                <div>
                  <label class="block text-xs text-gray-500 mb-1"
                    >Filter Jenis Penerima</label
                  >
                  <select
                    v-model="mustahiqFilters.jenis_penerima"
                    @change="applyPublicMustahiqFilters"
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
                  <div class="flex gap-2">
                    <input
                      v-model="mustahiqFilters.q"
                      @keyup.enter="applyPublicMustahiqFilters"
                      placeholder="Ketik nama mustahiq..."
                      class="flex-1 px-3 py-2 border border-gray-300 rounded-lg text-sm"
                    />
                    <button
                      @click="applyPublicMustahiqFilters"
                      class="px-4 py-2 bg-green-600 text-white rounded-lg text-sm hover:bg-green-700"
                    >
                      Cari
                    </button>
                    <button
                      @click="resetPublicMustahiqFilters"
                      class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg text-sm hover:bg-gray-200"
                    >
                      Reset
                    </button>
                  </div>
                </div>
              </div>
            </div>
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
              <div
                class="mt-4 flex flex-col md:flex-row md:items-center md:justify-between gap-3 text-sm text-gray-600"
              >
                <div>
                  Menampilkan {{ modalData.mustahiq.length }} dari
                  {{ mustahiqPagination.total }} data
                </div>
                <div class="flex flex-wrap items-center gap-2">
                  <label class="text-xs text-gray-500">Per halaman</label>
                  <select
                    v-model="mustahiqPagination.pageSize"
                    @change="onMustahiqPageSizeChange"
                    class="px-2 py-1 border border-gray-300 rounded-lg text-sm"
                  >
                    <option
                      v-for="size in pageSizeOptions"
                      :key="size"
                      :value="size"
                    >
                      {{ size === "all" ? "Semua" : size }}
                    </option>
                  </select>
                  <button
                    class="px-3 py-1 border border-gray-300 rounded-lg disabled:opacity-50"
                    :disabled="
                      mustahiqPagination.isAll || mustahiqPagination.page <= 1
                    "
                    @click="setMustahiqPage(mustahiqPagination.page - 1)"
                  >
                    Prev
                  </button>
                  <span class="text-xs text-gray-500"
                    >{{ mustahiqPagination.page }} /
                    {{ mustahiqPagination.totalPages }}</span
                  >
                  <button
                    class="px-3 py-1 border border-gray-300 rounded-lg disabled:opacity-50"
                    :disabled="
                      mustahiqPagination.isAll ||
                      mustahiqPagination.page >= mustahiqPagination.totalPages
                    "
                    @click="setMustahiqPage(mustahiqPagination.page + 1)"
                  >
                    Next
                  </button>
                </div>
              </div>
            </div>
            <p v-else class="text-gray-500 text-center py-4">
              Belum ada data mustahiq
            </p>
          </div>

          <!-- Tab Content: Pengurus -->
          <div v-show="activeTab === 'pengurus'" class="mb-4 sm:mb-6">
            <h3
              class="text-base sm:text-lg font-bold text-gray-800 mb-2 sm:mb-3"
            >
              <i class="fas fa-users mr-2 text-blue-600"></i>Struktur Pengurus
            </h3>
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-3 sm:gap-4">
              <!-- Pengurus Masjid -->
              <div class="bg-blue-50 rounded-lg p-3 sm:p-4">
                <h4
                  class="font-semibold text-sm sm:text-base text-blue-800 mb-2 sm:mb-3"
                >
                  Pengurus Masjid/Langgar
                </h4>
                <div v-if="pengurusMasjidData.length > 0" class="space-y-2">
                  <div
                    v-for="p in pengurusMasjidData"
                    :key="p.id"
                    class="bg-white rounded p-2 sm:p-3"
                  >
                    <p class="text-xs text-gray-600 mb-1">{{ p.jabatan }}</p>
                    <p class="font-medium text-sm sm:text-base text-gray-900">
                      {{ p.nama }}
                    </p>
                    <p
                      v-if="p.telepon"
                      class="text-xs sm:text-sm text-gray-600 mt-1"
                    >
                      <i class="fas fa-phone mr-1"></i>{{ p.telepon }}
                    </p>
                  </div>
                </div>
                <p
                  v-else
                  class="text-gray-500 text-xs sm:text-sm text-center py-3"
                >
                  Belum ada data pengurus
                </p>
              </div>

              <!-- Pengurus UPZ -->
              <div class="bg-green-50 rounded-lg p-3 sm:p-4">
                <h4
                  class="font-semibold text-sm sm:text-base text-green-800 mb-2 sm:mb-3"
                >
                  Pengurus UPZ
                </h4>
                <div v-if="pengurusZakatData.length > 0" class="space-y-2">
                  <div
                    v-for="p in pengurusZakatData"
                    :key="p.id"
                    class="bg-white rounded p-2 sm:p-3"
                  >
                    <p class="text-xs text-gray-600 mb-1">{{ p.jabatan }}</p>
                    <p class="font-medium text-sm sm:text-base text-gray-900">
                      {{ p.nama }}
                    </p>
                    <p
                      v-if="p.telepon"
                      class="text-xs sm:text-sm text-gray-600 mt-1"
                    >
                      <i class="fas fa-phone mr-1"></i>{{ p.telepon }}
                    </p>
                  </div>
                </div>
                <p
                  v-else
                  class="text-gray-500 text-xs sm:text-sm text-center py-3"
                >
                  Belum ada data pengurus UPZ
                </p>
              </div>
            </div>
          </div>
        </div>

        <!-- Tab Content: Distribusi -->
        <div v-show="activeTab === 'distribusi'" class="mb-4 sm:mb-6">
          <h3
            class="text-base sm:text-lg font-bold text-gray-800 mb-2 sm:mb-3 pl-0 sm:pl-2 md:pl-4"
          >
            <i class="fas fa-hand-holding-usd mr-2 text-blue-600"></i>Data
            Distribusi Zakat
          </h3>
          <div
            v-if="modalData?.distribusi?.length > 0"
            class="px-0 sm:px-2 md:px-4"
          >
            <div class="hidden md:block overflow-x-auto">
              <table class="min-w-full divide-y divide-gray-200">
                <thead class="bg-gray-50">
                  <tr>
                    <th
                      class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase"
                    >
                      Mustahiq
                    </th>
                    <th
                      class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase"
                    >
                      Beras
                    </th>
                    <th
                      class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase"
                    >
                      Uang
                    </th>
                    <th
                      class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase"
                    >
                      Keterangan
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
                    v-for="d in modalData?.distribusi"
                    :key="d.id"
                    class="hover:bg-gray-50"
                  >
                    <td class="px-4 py-2 text-sm text-gray-900">
                      {{ d.mustahiq_nama }}
                    </td>
                    <td class="px-4 py-2 text-sm text-gray-700">
                      {{ (d.beras_kg || 0).toFixed(1) }} kg
                    </td>
                    <td class="px-4 py-2 text-sm font-medium text-gray-900">
                      {{ formatCurrency(d.nominal || 0) }}
                    </td>
                    <td class="px-4 py-2 text-sm text-gray-700">
                      {{ d.keterangan || "-" }}
                    </td>
                    <td class="px-4 py-2 text-sm text-gray-600">
                      {{ formatDate(d.tanggal_distribusi) }}
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
            <div class="md:hidden space-y-3">
              <div
                v-for="d in modalData?.distribusi"
                :key="d.id"
                class="bg-gray-50 rounded-lg p-4"
              >
                <div class="flex justify-between items-start mb-2">
                  <div>
                    <p class="font-medium text-gray-900">
                      {{ d.mustahiq_nama }}
                    </p>
                    <p class="text-xs text-gray-500">
                      {{ formatDate(d.tanggal_distribusi) }}
                    </p>
                  </div>
                </div>
                <div class="space-y-1 text-sm">
                  <div class="flex justify-between">
                    <span class="text-gray-600">Beras:</span>
                    <span class="font-medium"
                      >{{ (d.beras_kg || 0).toFixed(1) }} kg</span
                    >
                  </div>
                  <div class="flex justify-between">
                    <span class="text-gray-600">Uang:</span>
                    <span class="font-medium">{{
                      formatCurrency(d.nominal || 0)
                    }}</span>
                  </div>
                  <div class="flex justify-between">
                    <span class="text-gray-600">Keterangan:</span>
                    <span class="font-medium">{{ d.keterangan || "-" }}</span>
                  </div>
                </div>
              </div>
            </div>
            <div
              class="mt-4 flex flex-col md:flex-row md:items-center md:justify-between gap-3 text-sm text-gray-600"
            >
              <div>
                Menampilkan {{ modalData.distribusi.length }} dari
                {{ distribusiPagination.total }} data
              </div>
              <div class="flex flex-wrap items-center gap-2">
                <label class="text-xs text-gray-500">Per halaman</label>
                <select
                  v-model="distribusiPagination.pageSize"
                  @change="onDistribusiPageSizeChange"
                  class="px-2 py-1 border border-gray-300 rounded-lg text-sm"
                >
                  <option
                    v-for="size in pageSizeOptions"
                    :key="size"
                    :value="size"
                  >
                    {{ size === "all" ? "Semua" : size }}
                  </option>
                </select>
                <button
                  class="px-3 py-1 border border-gray-300 rounded-lg disabled:opacity-50"
                  :disabled="
                    distribusiPagination.isAll || distribusiPagination.page <= 1
                  "
                  @click="setDistribusiPage(distribusiPagination.page - 1)"
                >
                  Prev
                </button>
                <span class="text-xs text-gray-500"
                  >{{ distribusiPagination.page }} /
                  {{ distribusiPagination.totalPages }}</span
                >
                <button
                  class="px-3 py-1 border border-gray-300 rounded-lg disabled:opacity-50"
                  :disabled="
                    distribusiPagination.isAll ||
                    distribusiPagination.page >= distribusiPagination.totalPages
                  "
                  @click="setDistribusiPage(distribusiPagination.page + 1)"
                >
                  Next
                </button>
              </div>
            </div>
          </div>
          <div v-else class="text-center text-gray-500 py-6">
            Belum ada data distribusi.
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
    label: "Total Orang Dizakati",
    value: 0,
    icon: "fas fa-person-circle-check",
    bgColor: "bg-orange-100",
    textColor: "text-orange-600",
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
    label: "Fidyah (Uang)",
    value: "Rp 0",
    icon: "fas fa-moon",
    bgColor: "bg-yellow-100",
    textColor: "text-yellow-600",
  },
  {
    label: "Fidyah (Beras)",
    value: "0 kg",
    valueRp: null,
    icon: "fas fa-leaf",
    bgColor: "bg-lime-100",
    textColor: "text-lime-600",
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
const activeTab = ref("zakat");
const pengurusMasjidData = ref([]);
const pengurusZakatData = ref([]);
const isLoadingStats = ref(true);
const isLoadingMasjid = ref(true);
const isLoadingModal = ref(false);
const pageSizeOptions = [5, 10, 25, 50, "all"];
const transaksiPagination = ref({
  page: 1,
  pageSize: 5,
  total: 0,
  totalPages: 1,
  isAll: false,
});
const mustahiqPagination = ref({
  page: 1,
  pageSize: 5,
  total: 0,
  totalPages: 1,
  isAll: false,
});
const distribusiPagination = ref({
  page: 1,
  pageSize: 5,
  total: 0,
  totalPages: 1,
  isAll: false,
});
const transaksiFilters = ref({ jenis_zakat: "", bentuk_zakat: "", q: "" });
const mustahiqFilters = ref({ jenis_penerima: "", q: "" });
const localLists = ref({ transaksi: [], mustahiq: [], distribusi: [] });
const useLocalLists = ref({
  transaksi: false,
  mustahiq: false,
  distribusi: false,
});

const showPublicBentukFilter = ref(false);

const syncPagination = (target, payload, itemCount) => {
  if (payload && payload.pagination) {
    const meta = payload.pagination;
    target.value = {
      page: meta.page || 1,
      pageSize: meta.is_all ? "all" : meta.page_size || target.value.pageSize,
      total: meta.total ?? itemCount,
      totalPages: meta.total_pages || 1,
      isAll: !!meta.is_all,
    };
    return;
  }
  target.value.total = itemCount;
  target.value.totalPages = 1;
  target.value.isAll = target.value.pageSize === "all";
};

const applyLocalPagination = (items, target) => {
  const total = items.length;
  if (target.value.pageSize === "all") {
    target.value.total = total;
    target.value.totalPages = 1;
    target.value.isAll = true;
    return items;
  }

  const pageSize = Number(target.value.pageSize) || 5;
  const totalPages = Math.max(1, Math.ceil(total / pageSize));
  const page = Math.min(target.value.page, totalPages);
  const start = (page - 1) * pageSize;
  const sliced = items.slice(start, start + pageSize);

  target.value.total = total;
  target.value.totalPages = totalPages;
  target.value.page = page;
  target.value.isAll = false;

  return sliced;
};

const fetchTransaksiPage = async (id) => {
  try {
    const normalizedJenis = String(transaksiFilters.value.jenis_zakat || "")
      .trim()
      .toLowerCase();
    const normalizedBentuk = showPublicBentukFilter.value
      ? String(transaksiFilters.value.bentuk_zakat || "")
          .trim()
          .toLowerCase()
      : "";
    const normalizedQ = String(transaksiFilters.value.q || "").trim();

    const params =
      transaksiPagination.value.pageSize === "all"
        ? { page_size: "all" }
        : {
            page: transaksiPagination.value.page,
            page_size: transaksiPagination.value.pageSize,
          };
    if (normalizedJenis) params.jenis_zakat = normalizedJenis;
    if (normalizedBentuk) params.bentuk_zakat = normalizedBentuk;
    if (normalizedQ) params.q = normalizedQ;

    const response = await api.getPublicMasjidTransaksi(id, params);
    if (response.data.success) {
      const payload = response.data.data || {};
      if (Array.isArray(payload.items)) {
        modalData.value.transaksi = payload.items;
        syncPagination(transaksiPagination, payload, payload.items.length);
        useLocalLists.value.transaksi = false;
      } else if (Array.isArray(payload)) {
        modalData.value.transaksi = applyLocalPagination(
          payload,
          transaksiPagination,
        );
      } else {
        modalData.value.transaksi = [];
        syncPagination(transaksiPagination, payload, 0);
      }
    }
    return true;
  } catch (error) {
    console.error("Error loading transaksi:", error);
    if (useLocalLists.value.transaksi) {
      modalData.value.transaksi = applyLocalPagination(
        filterLocalTransaksi(localLists.value.transaksi),
        transaksiPagination,
      );
      return true;
    }
    return false;
  }
};

const fetchMustahiqPage = async (id) => {
  try {
    const normalizedJenis = String(mustahiqFilters.value.jenis_penerima || "")
      .trim()
      .toLowerCase();
    const normalizedQ = String(mustahiqFilters.value.q || "").trim();

    const params =
      mustahiqPagination.value.pageSize === "all"
        ? { page_size: "all" }
        : {
            page: mustahiqPagination.value.page,
            page_size: mustahiqPagination.value.pageSize,
          };
    if (normalizedJenis) params.jenis_penerima = normalizedJenis;
    if (normalizedQ) params.q = normalizedQ;

    const response = await api.getPublicMasjidMustahiq(id, params);
    if (response.data.success) {
      const payload = response.data.data || {};
      if (Array.isArray(payload.items)) {
        modalData.value.mustahiq = payload.items;
        syncPagination(mustahiqPagination, payload, payload.items.length);
        useLocalLists.value.mustahiq = false;
      } else if (Array.isArray(payload)) {
        modalData.value.mustahiq = applyLocalPagination(
          payload,
          mustahiqPagination,
        );
      } else {
        modalData.value.mustahiq = [];
        syncPagination(mustahiqPagination, payload, 0);
      }
    }
    return true;
  } catch (error) {
    console.error("Error loading mustahiq:", error);
    if (useLocalLists.value.mustahiq) {
      modalData.value.mustahiq = applyLocalPagination(
        filterLocalMustahiq(localLists.value.mustahiq),
        mustahiqPagination,
      );
      return true;
    }
    return false;
  }
};

const fetchDistribusiPage = async (id) => {
  try {
    const params =
      distribusiPagination.value.pageSize === "all"
        ? { page_size: "all" }
        : {
            page: distribusiPagination.value.page,
            page_size: distribusiPagination.value.pageSize,
          };
    const response = await api.getPublicMasjidDistribusi(id, params);
    if (response.data.success) {
      const payload = response.data.data || {};
      if (Array.isArray(payload.items)) {
        modalData.value.distribusi = payload.items;
        syncPagination(distribusiPagination, payload, payload.items.length);
        useLocalLists.value.distribusi = false;
      } else if (Array.isArray(payload)) {
        modalData.value.distribusi = applyLocalPagination(
          payload,
          distribusiPagination,
        );
      } else {
        modalData.value.distribusi = [];
        syncPagination(distribusiPagination, payload, 0);
      }
    }
    return true;
  } catch (error) {
    console.error("Error loading distribusi:", error);
    if (useLocalLists.value.distribusi) {
      modalData.value.distribusi = applyLocalPagination(
        localLists.value.distribusi,
        distribusiPagination,
      );
      return true;
    }
    return false;
  }
};

const hydrateListsFromStats = (data) => {
  const transaksi = Array.isArray(data.transaksi) ? data.transaksi : [];
  const mustahiq = Array.isArray(data.mustahiq) ? data.mustahiq : [];
  const distribusi = Array.isArray(data.distribusi) ? data.distribusi : [];

  localLists.value = { transaksi, mustahiq, distribusi };
  useLocalLists.value = { transaksi: true, mustahiq: true, distribusi: true };

  modalData.value.transaksi = applyLocalPagination(
    filterLocalTransaksi(transaksi),
    transaksiPagination,
  );
  modalData.value.mustahiq = applyLocalPagination(
    filterLocalMustahiq(mustahiq),
    mustahiqPagination,
  );
  modalData.value.distribusi = applyLocalPagination(
    distribusi,
    distribusiPagination,
  );
};

const filterLocalTransaksi = (items) => {
  const jenis = String(transaksiFilters.value.jenis_zakat || "")
    .trim()
    .toLowerCase();
  const bentuk = showPublicBentukFilter.value
    ? String(transaksiFilters.value.bentuk_zakat || "")
        .trim()
        .toLowerCase()
    : "";
  const keyword = String(transaksiFilters.value.q || "")
    .trim()
    .toLowerCase();

  return (Array.isArray(items) ? items : []).filter((item) => {
    const itemJenis = String(item.jenis_zakat || "")
      .trim()
      .toLowerCase();
    const itemBentuk = String(item.bentuk_zakat || "")
      .trim()
      .toLowerCase();
    const itemNama = String(item.muzakki_nama || "")
      .trim()
      .toLowerCase();

    if (jenis && itemJenis !== jenis) return false;
    if (bentuk && itemBentuk !== bentuk) return false;
    if (keyword && !itemNama.includes(keyword)) return false;
    return true;
  });
};

const filterLocalMustahiq = (items) => {
  const jenis = String(mustahiqFilters.value.jenis_penerima || "")
    .trim()
    .toLowerCase();
  const keyword = String(mustahiqFilters.value.q || "")
    .trim()
    .toLowerCase();

  return (Array.isArray(items) ? items : []).filter((item) => {
    const itemJenis = String(item.jenis_penerima || item.kategori || "")
      .trim()
      .toLowerCase();
    const itemNama = String(item.nama || "")
      .trim()
      .toLowerCase();

    if (jenis && itemJenis !== jenis) return false;
    if (keyword && !itemNama.includes(keyword)) return false;
    return true;
  });
};

const setTransaksiPage = async (page) => {
  transaksiPagination.value.page = page;
  if (useLocalLists.value.transaksi) {
    modalData.value.transaksi = applyLocalPagination(
      filterLocalTransaksi(localLists.value.transaksi),
      transaksiPagination,
    );
    return;
  }
  if (modalData.value?.masjid?.id)
    await fetchTransaksiPage(modalData.value.masjid.id);
};

const onTransaksiPageSizeChange = async () => {
  transaksiPagination.value.page = 1;
  if (useLocalLists.value.transaksi) {
    modalData.value.transaksi = applyLocalPagination(
      filterLocalTransaksi(localLists.value.transaksi),
      transaksiPagination,
    );
    return;
  }
  if (modalData.value?.masjid?.id)
    await fetchTransaksiPage(modalData.value.masjid.id);
};

const setMustahiqPage = async (page) => {
  mustahiqPagination.value.page = page;
  if (useLocalLists.value.mustahiq) {
    modalData.value.mustahiq = applyLocalPagination(
      filterLocalMustahiq(localLists.value.mustahiq),
      mustahiqPagination,
    );
    return;
  }
  if (modalData.value?.masjid?.id)
    await fetchMustahiqPage(modalData.value.masjid.id);
};

const onMustahiqPageSizeChange = async () => {
  mustahiqPagination.value.page = 1;
  if (useLocalLists.value.mustahiq) {
    modalData.value.mustahiq = applyLocalPagination(
      filterLocalMustahiq(localLists.value.mustahiq),
      mustahiqPagination,
    );
    return;
  }
  if (modalData.value?.masjid?.id)
    await fetchMustahiqPage(modalData.value.masjid.id);
};

const onPublicJenisFilterChange = () => {
  const jenis = String(transaksiFilters.value.jenis_zakat || "")
    .trim()
    .toLowerCase();
  showPublicBentukFilter.value = jenis === "fitrah" || jenis === "fidyah";
  if (!showPublicBentukFilter.value) {
    transaksiFilters.value.bentuk_zakat = "";
  }
  applyPublicTransaksiFilters();
};

const applyPublicTransaksiFilters = async () => {
  transaksiPagination.value.page = 1;
  if (useLocalLists.value.transaksi) {
    modalData.value.transaksi = applyLocalPagination(
      filterLocalTransaksi(localLists.value.transaksi),
      transaksiPagination,
    );
    return;
  }
  if (modalData.value?.masjid?.id)
    await fetchTransaksiPage(modalData.value.masjid.id);
};

const resetPublicTransaksiFilters = async () => {
  transaksiFilters.value = { jenis_zakat: "", bentuk_zakat: "", q: "" };
  showPublicBentukFilter.value = false;
  await applyPublicTransaksiFilters();
};

const applyPublicMustahiqFilters = async () => {
  mustahiqPagination.value.page = 1;
  if (useLocalLists.value.mustahiq) {
    modalData.value.mustahiq = applyLocalPagination(
      filterLocalMustahiq(localLists.value.mustahiq),
      mustahiqPagination,
    );
    return;
  }
  if (modalData.value?.masjid?.id)
    await fetchMustahiqPage(modalData.value.masjid.id);
};

const resetPublicMustahiqFilters = async () => {
  mustahiqFilters.value = { jenis_penerima: "", q: "" };
  await applyPublicMustahiqFilters();
};

const setDistribusiPage = async (page) => {
  distribusiPagination.value.page = page;
  if (useLocalLists.value.distribusi) {
    modalData.value.distribusi = applyLocalPagination(
      localLists.value.distribusi,
      distribusiPagination,
    );
    return;
  }
  if (modalData.value?.masjid?.id)
    await fetchDistribusiPage(modalData.value.masjid.id);
};

const onDistribusiPageSizeChange = async () => {
  distribusiPagination.value.page = 1;
  if (useLocalLists.value.distribusi) {
    modalData.value.distribusi = applyLocalPagination(
      localLists.value.distribusi,
      distribusiPagination,
    );
    return;
  }
  if (modalData.value?.masjid?.id)
    await fetchDistribusiPage(modalData.value.masjid.id);
};

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
    case "fitrah":
      return "px-2 py-1 bg-green-100 text-green-800 rounded-full text-xs";
    case "mal":
      return "px-2 py-1 bg-indigo-100 text-indigo-800 rounded-full text-xs";
    case "fidyah":
      return "px-2 py-1 bg-yellow-100 text-yellow-800 rounded-full text-xs";
    case "infaq":
      return "px-2 py-1 bg-pink-100 text-pink-800 rounded-full text-xs";
    default:
      return "px-2 py-1 bg-gray-100 text-gray-800 rounded-full text-xs";
  }
};

const parseHariFromKeterangan = (keterangan) => {
  if (!keterangan) return 0;
  const match = keterangan.match(/(\d+)\s*hari/i);
  return match ? Number(match[1]) : 0;
};

const parseKgFromKeterangan = (keterangan) => {
  if (!keterangan) return 0;
  const match = keterangan.match(/(?:Beras|Fidyah Beras):\s*([\d.]+)\s*kg/i);
  return match ? Number(match[1]) : 0;
};

const formatJumlah = (t) => {
  if (!t) return "-";
  if (t.jenis_zakat === "fidyah") {
    const hari =
      Number(t.jumlah_hari_fidyah || 0) ||
      parseHariFromKeterangan(t.keterangan);
    return hari > 0 ? `${hari} hari` : "-";
  }
  return t.jumlah_orang > 0 ? `${t.jumlah_orang} orang` : "-";
};

const getTotalBerasKg = (t) => {
  if (!t) return 0;
  const numericKg = Number(t.kg_beras_dibayar || 0);
  if (numericKg > 0) return numericKg;
  if (String(t.bentuk_zakat || "").toLowerCase() !== "beras") return 0;
  const kg = parseKgFromKeterangan(t.keterangan);
  return kg > 0 ? kg : 0;
};

const formatTotalUang = (t) => {
  if (!t) return "-";
  if (String(t.bentuk_zakat || "").toLowerCase() === "beras") return "-";
  return formatCurrency(t.total_dibayar);
};

const formatTotalBerasKg = (t) => {
  const kg = getTotalBerasKg(t);
  return kg > 0
    ? `${kg.toLocaleString("id-ID", { maximumFractionDigits: 2 })} kg`
    : "-";
};

const viewMasjidStats = async (id) => {
  try {
    isLoadingModal.value = true;
    activeTab.value = "zakat"; // Reset to first tab
    transaksiFilters.value = { jenis_zakat: "", bentuk_zakat: "", q: "" };
    mustahiqFilters.value = { jenis_penerima: "", q: "" };
    showPublicBentukFilter.value = false;
    const [statsData, pengurusMasjid, pengurusZakat] = await Promise.all([
      api.getPublicMasjidStats(id),
      api.getPublicPengurusMasjid(id),
      api.getPublicPengurusZakat(id),
    ]);

    if (statsData.data.success) {
      modalData.value = {
        ...statsData.data.data,
        transaksi: [],
        mustahiq: [],
        distribusi: [],
      };
      showModal.value = true;
    }

    if (pengurusMasjid.data.success) {
      pengurusMasjidData.value = pengurusMasjid.data.data || [];
    }

    if (pengurusZakat.data.success) {
      pengurusZakatData.value = pengurusZakat.data.data || [];
    }
    if (modalData.value?.masjid?.id) {
      transaksiPagination.value.page = 1;
      mustahiqPagination.value.page = 1;
      distribusiPagination.value.page = 1;
      const results = await Promise.all([
        fetchTransaksiPage(modalData.value.masjid.id),
        fetchMustahiqPage(modalData.value.masjid.id),
        fetchDistribusiPage(modalData.value.masjid.id),
      ]);
      const hasFailure = results.some((result) => !result);
      if (hasFailure) {
        const fallback = await api.getPublicMasjidStats(
          modalData.value.masjid.id,
          { include_lists: true },
        );
        if (fallback.data.success) {
          hydrateListsFromStats(fallback.data.data || {});
        }
      }
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
      stats.value[2].value = d.total_orang_dizakati || 0;
      stats.value[3].value = d.total_mustahiq || 0;

      // Backward compatibility: gunakan field lama jika field baru belum ada
      const fitrahUang =
        d.total_zakat_fitrah_uang !== undefined
          ? d.total_zakat_fitrah_uang
          : d.total_zakat_fitrah || 0;
      const fitrahBerasKg = d.total_zakat_fitrah_beras_kg || 0;
      const fitrahBerasRp = d.total_zakat_fitrah_beras_rupiah || 0;

      stats.value[4].value = formatCurrency(fitrahUang);
      stats.value[5].value = `${fitrahBerasKg.toFixed(1)} kg`;
      if (fitrahBerasRp > 0) {
        stats.value[5].valueRp = formatCurrency(fitrahBerasRp);
      }
      stats.value[6].value = formatCurrency(d.total_zakat_mal || 0);

      // Fidyah split into uang and beras
      const fidyahUang = d.total_fidyah_uang || d.total_fidyah || 0;
      const fidyahBerasKg = d.total_fidyah_beras_kg || 0;
      const fidyahBerasRp = d.total_fidyah_beras_rupiah || 0;

      stats.value[7].value = formatCurrency(fidyahUang);
      stats.value[8].value = `${fidyahBerasKg.toFixed(1)} kg`;
      if (fidyahBerasRp > 0) {
        stats.value[8].valueRp = formatCurrency(fidyahBerasRp);
      }

      stats.value[9].value = formatCurrency(d.total_infaq || 0);

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
