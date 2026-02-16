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
              <template v-if="isLoading">
                <SkeletonCard v-for="i in 5" :key="i" type="table-row" />
              </template>
              <tr v-else v-for="t in transaksiList" :key="t.id" class="hover:bg-gray-50">
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
                  <button @click="viewDetail(t)" class="text-blue-600 hover:text-blue-900 mr-3" title="Lihat Detail">
                    <i class="fas fa-eye"></i>
                  </button>
                  <button @click="printReceipt(t.id)" class="text-green-600 hover:text-green-900 mr-3" title="Cetak">
                    <i class="fas fa-print"></i>
                  </button>
                  <button @click="deleteTransaksi(t.id)" class="text-red-600 hover:text-red-900" title="Hapus">
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
          <div v-else v-for="t in transaksiList" :key="t.id" class="bg-white rounded-xl shadow-md p-4">
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
              <button @click="viewDetail(t)" class="flex-1 bg-blue-50 text-blue-600 py-2 rounded-lg hover:bg-blue-100 text-sm">
                <i class="fas fa-eye mr-1"></i>Detail
              </button>
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
            <option value="fidyah">Fidyah</option>
            <option value="infaq">Infaq</option>
          </select>
        </div>
        <div v-if="form.jenis_zakat === 'fitrah' || form.jenis_zakat === 'fidyah'">
          <label class="block text-sm font-medium text-gray-700 mb-2">Bentuk {{ form.jenis_zakat === 'fidyah' ? 'Fidyah' : 'Zakat' }} *</label>
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
        <div v-if="form.bentuk_zakat === 'uang' && form.jenis_zakat === 'fitrah'">
          <label class="block text-sm font-medium text-gray-700 mb-2">Kelas Zakat *</label>
          <select v-model="form.kelas_zakat" required @change="calculateTotal" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500">
            <option value="">Pilih Kelas</option>
            <option value="1">Kelas 1 - Rp {{ formatNumber(kadarZakat.kelas1) }}</option>
            <option value="2">Kelas 2 - Rp {{ formatNumber(kadarZakat.kelas2) }}</option>
            <option value="3">Kelas 3 - Rp {{ formatNumber(kadarZakat.kelas3) }}</option>
          </select>
        </div>
        <div v-if="form.bentuk_zakat === 'uang' && form.jenis_zakat === 'fidyah'">
          <label class="block text-sm font-medium text-gray-700 mb-2">Jumlah Hari yang Ditinggalkan *</label>
          <input v-model.number="form.jumlah_hari_fidyah" type="number" required min="1" @input="calculateFidyahTotal" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500" placeholder="Masukkan jumlah hari">
          <p class="text-xs text-gray-500 mt-1">Fidyah per hari setara dengan 3/4 gantang beras ({{ formatNumber(kadarZakat.fidyahPerHari) }}/hari)</p>
        </div>
        <div v-if="form.bentuk_zakat === 'beras' && form.jenis_zakat === 'fitrah'">
          <label class="block text-sm font-medium text-gray-700 mb-2">Standar Beras per Jiwa (kg) *</label>
          <input v-model.number="form.standar_beras_per_jiwa" type="number" step="0.1" min="0.1" required @input="calculateBerasTotal" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500" placeholder="Contoh: 2.5 atau 3">
          <p class="text-xs text-gray-500 mt-1">Standar syariat: 1 sha' ≈ 2.5-3 kg beras per jiwa</p>
        </div>
        <div v-if="form.bentuk_zakat === 'beras' && form.jenis_zakat === 'fidyah'">
          <label class="block text-sm font-medium text-gray-700 mb-2">Jumlah Hari Fidyah *</label>
          <input v-model.number="form.jumlah_hari_fidyah" type="number" required min="1" @input="calculateFidyahBerasTotal" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500" placeholder="Masukkan jumlah hari">
          <p class="text-xs text-gray-500 mt-1">Fidyah beras per hari ≈ 0.6 kg (3/4 dari 2.5 kg)</p>
        </div>
        <div v-if="form.jenis_zakat === 'fitrah'">
          <label class="block text-sm font-medium text-gray-700 mb-2">Jumlah Orang *</label>
          <input v-model.number="form.jumlah_orang" type="number" required min="1" @input="calculateTotal" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500">
        </div>
        <div v-if="form.jenis_zakat === 'fidyah' && totalWajib > 0 && form.bentuk_zakat === 'uang'">
          <div class="bg-yellow-50 border border-yellow-200 rounded-lg p-4">
            <p class="text-sm text-yellow-700 font-medium mb-1">Total Wajib Fidyah</p>
            <p class="text-2xl font-bold text-yellow-700">{{ formatCurrency(totalWajib) }}</p>
            <p class="text-xs text-yellow-600 mt-1">{{ form.jumlah_hari_fidyah }} hari × Rp {{ formatNumber(kadarZakat.fidyahPerHari) }}</p>
          </div>
        </div>
        <div v-if="form.jenis_zakat === 'fidyah' && form.bentuk_zakat === 'beras' && fidyahBerasWajib > 0">
          <div class="bg-yellow-50 border border-yellow-200 rounded-lg p-4">
            <p class="text-sm text-yellow-700 font-medium mb-1">Total Wajib Fidyah</p>
            <p class="text-2xl font-bold text-yellow-700">{{ fidyahBerasWajib }} kg</p>
            <p class="text-xs text-yellow-600 mt-1">{{ form.jumlah_hari_fidyah }} hari × 0.6 kg</p>
          </div>
        </div>
        <div v-if="form.bentuk_zakat === 'beras' && form.standar_beras_per_jiwa && form.jumlah_orang">
          <div class="bg-green-50 border border-green-200 rounded-lg p-4">
            <p class="text-sm text-green-700 font-medium mb-1">Total Wajib Beras</p>
            <p class="text-2xl font-bold text-green-700">{{ totalBerasWajib }} kg</p>
            <p class="text-xs text-green-600 mt-1">{{ form.jumlah_orang }} orang × {{ form.standar_beras_per_jiwa }} kg</p>
          </div>
        </div>
        <div v-if="form.bentuk_zakat === 'beras'">
          <label class="block text-sm font-medium text-gray-700 mb-2">Kilogram Beras yang Dibayarkan *</label>
          <input v-model.number="form.kg_beras_dibayar" type="number" step="0.1" min="0.1" required @input="calculateBerasInfaq" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500" placeholder="Masukkan jumlah kg beras">
        </div>
        <div v-if="form.bentuk_zakat === 'beras' && form.kg_beras_dibayar">
          <label class="block text-sm font-medium text-gray-700 mb-2">Harga Beras per Kg (Opsional)</label>
          <input v-model.number="form.harga_beras_per_kg" type="number" min="0" @input="calculateBerasRupiah" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500" placeholder="Untuk konversi ke Rupiah">
          <p class="text-xs text-gray-500 mt-1">Isi jika ingin mencatat nilai dalam rupiah</p>
        </div>
        <div v-if="totalWajib > 0 && form.bentuk_zakat !== 'beras'" class="bg-gray-100 p-4 rounded-lg">
          <label class="block text-sm font-medium text-gray-700 mb-2">{{ form.jenis_zakat === 'mal' ? 'Total Zakat Mal' : 'Total Wajib Zakat' }}</label>
          <div class="text-2xl font-bold text-green-600">{{ formatCurrency(totalWajib) }}</div>
          <p v-if="form.jenis_zakat === 'mal'" class="text-xs text-gray-600 mt-1">{{ formatCurrency(form.nominal_harta) }} × {{ form.persentase_zakat }}%</p>
        </div>
        <div v-if="form.bentuk_zakat !== 'beras'">
          <label class="block text-sm font-medium text-gray-700 mb-2">Total Dibayar *</label>
          <input v-model.number="form.total_dibayar" type="number" required min="0" @input="calculateInfaq" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500">
        </div>
        <div v-if="form.bentuk_zakat === 'beras' && form.harga_beras_per_kg && form.kg_beras_dibayar">
          <div class="bg-blue-50 border border-blue-200 rounded-lg p-4">
            <p class="text-sm text-blue-700 font-medium mb-1">Nilai dalam Rupiah</p>
            <p class="text-2xl font-bold text-blue-700">{{ formatCurrency(totalBerasRupiah) }}</p>
            <p class="text-xs text-blue-600 mt-1">{{ form.kg_beras_dibayar }} kg × Rp {{ formatNumber(form.harga_beras_per_kg) }}</p>
          </div>
        </div>
        <div v-if="infaqTambahan > 0 && form.bentuk_zakat !== 'beras'" class="bg-blue-50 border border-blue-200 rounded-lg p-4">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-blue-700 font-medium">Infaq Tambahan (Uang)</p>
              <p class="text-xs text-blue-600">Kelebihan dari total wajib</p>
            </div>
            <p class="text-2xl font-bold text-blue-700">{{ formatCurrency(infaqTambahan) }}</p>
          </div>
        </div>
        <div v-if="berasLebih > 0" class="bg-blue-50 border border-blue-200 rounded-lg p-4">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-blue-700 font-medium">Sedekah Beras Tambahan</p>
              <p class="text-xs text-blue-600">Kelebihan dari total wajib</p>
            </div>
            <p class="text-2xl font-bold text-blue-700">{{ berasLebih }} kg</p>
          </div>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">Tanggal Bayar *</label>
          <input v-model="form.tanggal_bayar" type="date" required class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500">
        </div>
        <div class="flex gap-3 pt-4">
          <button type="submit" :disabled="isSaving" class="flex-1 bg-green-600 text-white py-3 rounded-lg hover:bg-green-700 transition font-medium disabled:opacity-50">
            <i class="fas fa-save mr-2"></i>{{ isSaving ? 'Menyimpan...' : 'Simpan' }}
          </button>
          <button type="button" @click="closeModal" :disabled="isSaving" class="flex-1 bg-gray-200 text-gray-700 py-3 rounded-lg hover:bg-gray-300 transition font-medium disabled:opacity-50">Batal</button>
        </div>
      </form>
    </Modal>

    <!-- Modal Detail Transaksi -->
    <Modal :show="showDetailModal" title="Detail Transaksi" @close="showDetailModal = false">
      <div v-if="selectedTransaction" class="space-y-4">
        <!-- Info Muzakki -->
        <div class="bg-gray-50 rounded-lg p-4">
          <h3 class="font-semibold text-gray-800 mb-3 flex items-center">
            <i class="fas fa-user mr-2 text-green-600"></i>Data Muzakki
          </h3>
          <div class="space-y-2 text-sm">
            <div class="flex justify-between">
              <span class="text-gray-600">Nama:</span>
              <span class="font-medium text-gray-900">{{ selectedTransaction.muzakki_nama }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-600">Alamat:</span>
              <span class="font-medium text-gray-900">{{ selectedTransaction.muzakki_alamat }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-600">Telepon:</span>
              <span class="font-medium text-gray-900">{{ selectedTransaction.muzakki_telepon || '-' }}</span>
            </div>
          </div>
        </div>

        <!-- Info Transaksi -->
        <div class="bg-green-50 rounded-lg p-4">
          <h3 class="font-semibold text-gray-800 mb-3 flex items-center">
            <i class="fas fa-receipt mr-2 text-green-600"></i>Info Transaksi
          </h3>
          <div class="space-y-2 text-sm">
            <div class="flex justify-between">
              <span class="text-gray-600">Tanggal:</span>
              <span class="font-medium text-gray-900">{{ formatDate(selectedTransaction.tanggal_bayar) }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-600">Jenis Zakat:</span>
              <span :class="`px-2 py-1 text-xs font-semibold rounded-full ${getZakatBadge(selectedTransaction.jenis_zakat)}`">
                {{ selectedTransaction.jenis_zakat.toUpperCase() }}
              </span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-600">Bentuk:</span>
              <span class="font-medium text-gray-900">{{ selectedTransaction.bentuk_zakat || '-' }}</span>
            </div>
          </div>
        </div>

        <!-- Detail Zakat Fitrah -->
        <div v-if="selectedTransaction.jenis_zakat === 'fitrah'" class="bg-blue-50 rounded-lg p-4">
          <h3 class="font-semibold text-gray-800 mb-3 flex items-center">
            <i class="fas fa-info-circle mr-2 text-blue-600"></i>Detail Zakat Fitrah
          </h3>
          
          <!-- Jumlah Orang - Highlighted -->
          <div class="bg-white rounded-lg p-3 mb-3 border-2 border-blue-300">
            <div class="flex justify-between items-center">
              <span class="text-gray-700 font-semibold">Jumlah Orang yang Dizakati:</span>
              <span class="text-2xl font-bold text-blue-700">{{ selectedTransaction.jumlah_orang }} Orang</span>
            </div>
          </div>

          <div class="space-y-2 text-sm">
            <div v-if="selectedTransaction.bentuk_zakat === 'uang'" class="flex justify-between">
              <span class="text-gray-600">Kelas Zakat:</span>
              <span class="font-medium text-gray-900">Kelas {{ selectedTransaction.kelas_zakat }}</span>
            </div>
            <div v-if="selectedTransaction.bentuk_zakat === 'beras'" class="flex justify-between">
              <span class="text-gray-600">Standar per Jiwa:</span>
              <span class="font-medium text-gray-900">{{ selectedTransaction.standar_beras_per_jiwa || 2.5 }} kg</span>
            </div>
            <div v-if="selectedTransaction.bentuk_zakat === 'beras'" class="flex justify-between">
              <span class="text-gray-600">Total Beras Wajib:</span>
              <span class="font-medium text-blue-700">{{ ((selectedTransaction.standar_beras_per_jiwa || 2.5) * selectedTransaction.jumlah_orang).toFixed(1) }} kg</span>
            </div>
            <div v-if="selectedTransaction.bentuk_zakat === 'beras'" class="flex justify-between">
              <span class="text-gray-600">Kg Beras Dibayar:</span>
              <span class="font-medium text-gray-900">{{ selectedTransaction.kg_beras_dibayar || extractKgBerasFromKeterangan(selectedTransaction.keterangan) }} kg</span>
            </div>
            <div v-if="selectedTransaction.bentuk_zakat === 'beras'" class="flex justify-between">
              <span class="text-gray-600">Harga per Kg:</span>
              <span class="font-medium text-gray-900">{{ formatCurrency(selectedTransaction.harga_beras_per_kg || 0) }}</span>
            </div>
          </div>
        </div>

        <!-- Detail Zakat Mal -->
        <div v-if="selectedTransaction.jenis_zakat === 'mal'" class="bg-blue-50 rounded-lg p-4">
          <h3 class="font-semibold text-gray-800 mb-3 flex items-center">
            <i class="fas fa-info-circle mr-2 text-blue-600"></i>Detail Zakat Mal
          </h3>
          
          <div class="space-y-3">
            <!-- Jenis Harta -->
            <div class="bg-white rounded-lg p-3">
              <div class="flex justify-between items-center">
                <span class="text-sm text-gray-600">Jenis Harta:</span>
                <span class="font-semibold text-gray-900 text-lg">{{ selectedTransaction.jenis_harta }}</span>
              </div>
            </div>

            <!-- Breakdown Perhitungan -->
            <div class="bg-white rounded-lg p-3 space-y-2">
              <div class="flex justify-between items-center pb-2">
                <span class="text-sm text-gray-600">Total Nilai Harta:</span>
                <span class="font-medium text-gray-900">{{ formatCurrency(selectedTransaction.nominal_harta) }}</span>
              </div>
              <div class="flex justify-between items-center pb-2">
                <span class="text-sm text-gray-600">Persentase Zakat:</span>
                <span class="font-medium text-gray-900">{{ selectedTransaction.persentase_zakat }}%</span>
              </div>
              <div class="border-t border-gray-200 pt-2">
                <div class="flex justify-between items-center">
                  <span class="text-sm font-semibold text-blue-700">Zakat yang Wajib:</span>
                  <span class="font-bold text-blue-700">{{ formatCurrency(selectedTransaction.nominal_harta * selectedTransaction.persentase_zakat / 100) }}</span>
                </div>
                <p class="text-xs text-gray-500 mt-1 text-right">
                  {{ formatCurrency(selectedTransaction.nominal_harta) }} × {{ selectedTransaction.persentase_zakat }}%
                </p>
              </div>
            </div>

            <!-- Informasi Tambahan -->
            <div class="bg-blue-100 rounded-lg p-3">
              <div class="flex items-start">
                <i class="fas fa-info-circle text-blue-600 mt-1 mr-2"></i>
                <div class="text-xs text-blue-800">
                  <p class="font-medium mb-1">Ketentuan Zakat {{ selectedTransaction.jenis_harta }}:</p>
                  <p v-if="selectedTransaction.jenis_harta.toLowerCase().includes('emas') || selectedTransaction.jenis_harta.toLowerCase().includes('perak')">
                    Nisab emas 85 gram, perak 595 gram. Kadar zakat 2.5%
                  </p>
                  <p v-else-if="selectedTransaction.jenis_harta.toLowerCase().includes('perdagangan') || selectedTransaction.jenis_harta.toLowerCase().includes('simpanan')">
                    Nisab setara 85 gram emas. Haul 1 tahun. Kadar zakat 2.5%
                  </p>
                  <p v-else-if="selectedTransaction.jenis_harta.toLowerCase().includes('pertanian')">
                    Kadar zakat: 5% (dengan irigasi) atau 10% (tadah hujan)
                  </p>
                  <p v-else>
                    Zakat dikeluarkan sesuai ketentuan syariat Islam
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Detail Fidyah -->
        <div v-if="selectedTransaction.jenis_zakat === 'fidyah'" class="bg-yellow-50 rounded-lg p-4">
          <h3 class="font-semibold text-gray-800 mb-3 flex items-center">
            <i class="fas fa-info-circle mr-2 text-yellow-600"></i>Detail Fidyah
          </h3>
          
          <!-- Jumlah Hari - Highlighted -->
          <div class="bg-white rounded-lg p-3 mb-3 border-2 border-yellow-300">
            <div class="flex justify-between items-center">
              <span class="text-gray-700 font-semibold">Jumlah Hari Fidyah:</span>
              <span class="text-2xl font-bold text-yellow-700">{{ selectedTransaction.jumlah_hari_fidyah || 0 }} Hari</span>
            </div>
          </div>

          <div class="space-y-2 text-sm">
            <div class="flex justify-between">
              <span class="text-gray-600">Bentuk Pembayaran:</span>
              <span class="font-medium text-gray-900">{{ selectedTransaction.bentuk_zakat === 'uang' ? 'Uang' : 'Beras' }}</span>
            </div>
            <div v-if="selectedTransaction.bentuk_zakat === 'uang'" class="flex justify-between">
              <span class="text-gray-600">Fidyah per Hari:</span>
              <span class="font-medium text-gray-900">Rp 30,000</span>
            </div>
            <div v-if="selectedTransaction.bentuk_zakat === 'beras'" class="flex justify-between">
              <span class="text-gray-600">Standar per Hari:</span>
              <span class="font-medium text-gray-900">0.6 kg</span>
            </div>
            <div v-if="selectedTransaction.bentuk_zakat === 'beras'" class="flex justify-between">
              <span class="text-gray-600">Total Beras Wajib:</span>
              <span class="font-medium text-yellow-700">{{ ((selectedTransaction.jumlah_hari_fidyah || 0) * 0.6).toFixed(1) }} kg</span>
            </div>
            <div v-if="selectedTransaction.bentuk_zakat === 'beras'" class="flex justify-between">
              <span class="text-gray-600">Kg Beras Dibayar:</span>
              <span class="font-medium text-gray-900">{{ selectedTransaction.kg_beras_dibayar || extractKgBerasFromKeterangan(selectedTransaction.keterangan) }} kg</span>
            </div>
            <div v-if="selectedTransaction.bentuk_zakat === 'beras'" class="flex justify-between">
              <span class="text-gray-600">Harga per Kg:</span>
              <span class="font-medium text-gray-900">{{ formatCurrency(selectedTransaction.harga_beras_per_kg || 0) }}</span>
            </div>
          </div>

          <!-- Info Fidyah -->
          <div class="bg-yellow-100 rounded-lg p-3 mt-3">
            <div class="flex items-start">
              <i class="fas fa-moon text-yellow-600 mt-1 mr-2"></i>
              <div class="text-xs text-yellow-800">
                <p class="font-medium">Tentang Fidyah:</p>
                <p>Fidyah adalah denda atau ganti rugi bagi yang tidak dapat berpuasa karena uzur (sakit permanen, tua, atau kondisi lainnya).</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Total Bayar -->
        <div class="bg-green-100 rounded-lg p-4 border-2 border-green-300">
          <div class="flex justify-between items-center">
            <span class="text-gray-700 font-semibold">Total Dibayar:</span>
            <span class="text-2xl font-bold text-green-700">{{ formatCurrency(selectedTransaction.total_dibayar) }}</span>
          </div>
          <div v-if="selectedTransaction.infaq_tambahan > 0" class="mt-2 pt-2 border-t border-green-200">
            <div class="flex justify-between items-center">
              <span class="text-sm text-blue-700">Infaq Tambahan:</span>
              <span class="text-lg font-semibold text-blue-700">{{ formatCurrency(selectedTransaction.infaq_tambahan) }}</span>
            </div>
          </div>
        </div>

        <!-- Keterangan -->
        <div v-if="selectedTransaction.keterangan" class="bg-gray-50 rounded-lg p-4">
          <h3 class="font-semibold text-gray-800 mb-2 flex items-center">
            <i class="fas fa-sticky-note mr-2 text-gray-600"></i>Keterangan
          </h3>
          <p class="text-sm text-gray-700">{{ selectedTransaction.keterangan }}</p>
        </div>

        <!-- Tombol Aksi -->
        <div class="flex gap-3 pt-4">
          <button @click="printReceipt(selectedTransaction.id)" class="flex-1 bg-green-600 text-white py-3 rounded-lg hover:bg-green-700 transition font-medium">
            <i class="fas fa-print mr-2"></i>Cetak Bukti
          </button>
          <button @click="showDetailModal = false" class="flex-1 bg-gray-200 text-gray-700 py-3 rounded-lg hover:bg-gray-300 transition font-medium">
            Tutup
          </button>
        </div>
      </div>
    </Modal>

    <LoadingOverlay :show="isSaving" message="Menyimpan transaksi..." />
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
  { path: '/petugas', label: 'Dashboard', icon: 'fas fa-home' },
  { path: '/petugas/pengaturan', label: 'Pengaturan', icon: 'fas fa-cog' },
  { path: '/petugas/mustahiq', label: 'Mustahiq', icon: 'fas fa-hand-holding-heart' },
  { path: '/petugas/transaksi', label: 'Transaksi', icon: 'fas fa-receipt' }
]

const showSidebar = ref(false)

const transaksiList = ref([])
const showModal = ref(false)
const showDetailModal = ref(false)
const selectedTransaction = ref(null)
const isLoading = ref(true)
const isSaving = ref(false)
const form = ref({ 
  nama_muzakki: '', 
  jenis_zakat: '', 
  bentuk_zakat: '', 
  jenis_harta: '',
  nominal_harta: 0,
  persentase_zakat: 0,
  kelas_zakat: '', 
  jumlah_orang: 1, 
  jumlah_hari_fidyah: 1,
  total_dibayar: 0, 
  tanggal_bayar: new Date().toISOString().split('T')[0],
  standar_beras_per_jiwa: 2.5,
  kg_beras_dibayar: 0,
  harga_beras_per_kg: 0
})
const toast = ref({ message: '', type: 'success' })
const kadarZakat = ref({ kelas1: 35000, kelas2: 45000, kelas3: 55000, fidyahPerHari: 30000 })

const totalWajib = computed(() => {
  if (form.value.jenis_zakat === 'fitrah' && form.value.bentuk_zakat === 'uang' && form.value.kelas_zakat && form.value.jumlah_orang) {
    const nominal = kadarZakat.value[`kelas${form.value.kelas_zakat}`]
    return nominal * form.value.jumlah_orang
  }
  if (form.value.jenis_zakat === 'mal' && form.value.nominal_harta && form.value.persentase_zakat) {
    return form.value.nominal_harta * (form.value.persentase_zakat / 100)
  }
  if (form.value.jenis_zakat === 'fidyah' && form.value.bentuk_zakat === 'uang' && form.value.jumlah_hari_fidyah) {
    return kadarZakat.value.fidyahPerHari * form.value.jumlah_hari_fidyah
  }
  return 0
})

const fidyahBerasWajib = computed(() => {
  if (form.value.jenis_zakat === 'fidyah' && form.value.bentuk_zakat === 'beras' && form.value.jumlah_hari_fidyah) {
    return form.value.jumlah_hari_fidyah * 0.6 // 0.6 kg per hari (3/4 dari 2.5kg)
  }
  return 0
})

const infaqTambahan = computed(() => {
  if (totalWajib.value > 0 && form.value.total_dibayar > totalWajib.value) {
    return form.value.total_dibayar - totalWajib.value
  }
  return 0
})

const totalBerasWajib = computed(() => {
  if (form.value.bentuk_zakat === 'beras' && form.value.standar_beras_per_jiwa && form.value.jumlah_orang) {
    return form.value.standar_beras_per_jiwa * form.value.jumlah_orang
  }
  return 0
})

const totalBerasRupiah = computed(() => {
  if (form.value.kg_beras_dibayar && form.value.harga_beras_per_kg) {
    return form.value.kg_beras_dibayar * form.value.harga_beras_per_kg
  }
  return 0
})

const berasLebih = computed(() => {
  if (totalBerasWajib.value > 0 && form.value.kg_beras_dibayar > totalBerasWajib.value) {
    return form.value.kg_beras_dibayar - totalBerasWajib.value
  }
  return 0
})

const formatCurrency = (amount) => new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR', minimumFractionDigits: 0 }).format(amount)
const formatNumber = (num) => new Intl.NumberFormat('id-ID').format(num)
const formatDate = (date) => new Date(date).toLocaleDateString('id-ID')

// Helper function to extract kg beras from keterangan (for backward compatibility with old data)
const extractKgBerasFromKeterangan = (keterangan) => {
  if (!keterangan) return 0
  // Format: "Beras: 10 kg" or "Fidyah Beras: 10 kg"
  const match = keterangan.match(/(?:Beras|Fidyah Beras):\s*([\d.]+)\s*kg/i)
  return match ? parseFloat(match[1]) : 0
}

const getZakatBadge = (jenis) => {
  const badges = { 
    fitrah: 'bg-green-100 text-green-800', 
    mal: 'bg-blue-100 text-blue-800', 
    fidyah: 'bg-yellow-100 text-yellow-800',
    infaq: 'bg-purple-100 text-purple-800' 
  }
  return badges[jenis] || 'bg-gray-100 text-gray-800'
}

const getDetail = (t) => {
  if (t.jenis_zakat === 'fitrah') {
    if (t.bentuk_zakat === 'beras') {
      // Untuk beras, tampilkan info dari keterangan
      return t.keterangan || `${t.bentuk_zakat} - ${t.jumlah_orang} orang`
    }
    return `${t.bentuk_zakat} - Kelas ${t.kelas_zakat} - ${t.jumlah_orang} orang`
  }
  if (t.jenis_zakat === 'mal') return `${t.jenis_harta} - ${formatCurrency(t.nominal_harta)} (${t.persentase_zakat}%)`
  if (t.jenis_zakat === 'fidyah') {
    if (t.bentuk_zakat === 'beras') {
      return t.keterangan || `Fidyah Beras`
    }
    return t.keterangan || `Fidyah Uang`
  }
  return '-'
}

const resetForm = () => {
  form.value.bentuk_zakat = ''
  form.value.jenis_harta = ''
  form.value.nominal_harta = 0
  form.value.persentase_zakat = 0
  form.value.kelas_zakat = ''
  form.value.jumlah_orang = 1
  form.value.jumlah_hari_fidyah = 1
  form.value.total_dibayar = 0
  form.value.standar_beras_per_jiwa = 2.5
  form.value.kg_beras_dibayar = 0
  form.value.harga_beras_per_kg = 0
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

const calculateFidyahTotal = () => {
  if (form.value.jumlah_hari_fidyah) {
    form.value.total_dibayar = kadarZakat.value.fidyahPerHari * form.value.jumlah_hari_fidyah
  }
}

const calculateFidyahBerasTotal = () => {
  if (form.value.jumlah_hari_fidyah) {
    // Set kg_beras_dibayar to minimum required
    form.value.kg_beras_dibayar = form.value.jumlah_hari_fidyah * 0.6
  }
}

const calculateTotal = () => {
  if (totalWajib.value > 0 && form.value.bentuk_zakat !== 'beras') {
    form.value.total_dibayar = totalWajib.value
  }
}

const calculateBerasTotal = () => {
  // Menghitung total wajib beras sudah di-handle oleh computed property totalBerasWajib
  // Fungsi ini untuk trigger re-calculation
}

const calculateBerasInfaq = () => {
  // Menghitung sedekah beras tambahan sudah di-handle oleh computed property berasLebih
  // Auto-calculate nilai rupiah jika harga per kg sudah diisi
  if (form.value.kg_beras_dibayar && form.value.harga_beras_per_kg) {
    form.value.total_dibayar = form.value.kg_beras_dibayar * form.value.harga_beras_per_kg
  }
}

const calculateBerasRupiah = () => {
  // Menghitung nilai rupiah beras
  if (form.value.kg_beras_dibayar && form.value.harga_beras_per_kg) {
    form.value.total_dibayar = form.value.kg_beras_dibayar * form.value.harga_beras_per_kg
  }
}

const calculateInfaq = () => {}

const loadTransaksi = async () => {
  isLoading.value = true
  try {
    const { data } = await api.getTransaksi()
    if (data.success) transaksiList.value = data.data || []
  } catch (error) {
    console.error(error)
  } finally {
    isLoading.value = false
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

const viewDetail = (transaction) => {
  selectedTransaction.value = transaction
  showDetailModal.value = true
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
  isSaving.value = true
  try {
    const muzakkiData = await api.createMuzakki({ nama: form.value.nama_muzakki, alamat: '-', telepon: '-' })
    if (!muzakkiData.data.success) {
      toast.value = { message: 'Gagal menyimpan data muzakki', type: 'error' }
      isSaving.value = false
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
      } else if (form.value.bentuk_zakat === 'beras') {
        // Simpan data beras ke field database
        payload.kg_beras_dibayar = form.value.kg_beras_dibayar
        payload.standar_beras_per_jiwa = form.value.standar_beras_per_jiwa
        payload.harga_beras_per_kg = form.value.harga_beras_per_kg
        
        // Untuk beras, simpan data kg beras di keterangan juga (backward compatibility)
        payload.keterangan = `Beras: ${form.value.kg_beras_dibayar} kg (Standar: ${form.value.standar_beras_per_jiwa} kg/jiwa)`
        // Jika ada harga per kg, simpan juga nilai rupiah
        if (form.value.harga_beras_per_kg) {
          payload.total_dibayar = form.value.kg_beras_dibayar * form.value.harga_beras_per_kg
          payload.keterangan += ` - Harga: Rp ${form.value.harga_beras_per_kg}/kg`
        } else {
          payload.total_dibayar = 0
        }
        // Simpan info tambahan di nominal_per_orang untuk tracking (dalam rupiah per kg atau 0)
        payload.nominal_per_orang = form.value.harga_beras_per_kg || 0
      }
    } else if (form.value.jenis_zakat === 'mal') {
      payload.jenis_harta = form.value.jenis_harta
      payload.nominal_harta = form.value.nominal_harta
      payload.persentase_zakat = form.value.persentase_zakat
      payload.jumlah_orang = 0
    } else if (form.value.jenis_zakat === 'fidyah') {
      payload.bentuk_zakat = form.value.bentuk_zakat
      payload.jumlah_hari_fidyah = form.value.jumlah_hari_fidyah
      payload.jumlah_orang = form.value.jumlah_hari_fidyah // Store hari as jumlah_orang for compatibility
      
      if (form.value.bentuk_zakat === 'uang') {
        const totalWajibFidyah = kadarZakat.value.fidyahPerHari * form.value.jumlah_hari_fidyah
        payload.nominal_per_orang = kadarZakat.value.fidyahPerHari
        payload.total_dibayar = form.value.total_dibayar
        payload.keterangan = `Fidyah ${form.value.jumlah_hari_fidyah} hari @ Rp ${kadarZakat.value.fidyahPerHari}`
        
        // Jika bayar lebih, kelebihannya masuk infaq
        if (form.value.total_dibayar > totalWajibFidyah) {
          payload.infaq_tambahan = form.value.total_dibayar - totalWajibFidyah
          payload.keterangan += ` (Kelebihan Rp ${payload.infaq_tambahan} masuk infaq)`
        }
      } else if (form.value.bentuk_zakat === 'beras') {
        const totalWajibBeras = form.value.jumlah_hari_fidyah * 0.6
        
        // Simpan data beras ke field database
        payload.kg_beras_dibayar = form.value.kg_beras_dibayar
        payload.harga_beras_per_kg = form.value.harga_beras_per_kg
        
        payload.keterangan = `Fidyah Beras: ${form.value.kg_beras_dibayar} kg untuk ${form.value.jumlah_hari_fidyah} hari (Wajib: ${totalWajibBeras} kg)`
        
        // Jika ada harga per kg, konversi ke rupiah
        if (form.value.harga_beras_per_kg) {
          payload.total_dibayar = form.value.kg_beras_dibayar * form.value.harga_beras_per_kg
          payload.keterangan += ` - Harga: Rp ${form.value.harga_beras_per_kg}/kg`
        } else {
          payload.total_dibayar = 0
        }
        
        // Jika beras lebih dari wajib, tetap dihitung sebagai beras yang dikasihkan
        if (form.value.kg_beras_dibayar > totalWajibBeras) {
          const kelebihanBeras = form.value.kg_beras_dibayar - totalWajibBeras
          payload.keterangan += ` (Kelebihan ${kelebihanBeras.toFixed(1)} kg tetap diterima)`
        }
        
        payload.nominal_per_orang = form.value.harga_beras_per_kg || 0
      }
    }
    
    const { data } = await api.createTransaksi(payload)
    if (data.success) {
      toast.value = { message: 'Transaksi berhasil disimpan', type: 'success' }
      closeModal()
      loadTransaksi()
    }
  } catch (error) {
    toast.value = { message: 'Gagal menyimpan transaksi', type: 'error' }
  } finally {
    isSaving.value = false
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
    jumlah_hari_fidyah: 1,
    total_dibayar: 0, 
    tanggal_bayar: new Date().toISOString().split('T')[0],
    standar_beras_per_jiwa: 2.5,
    kg_beras_dibayar: 0,
    harga_beras_per_kg: 0
  }
}

onMounted(() => {
  loadTransaksi()
  const saved = localStorage.getItem('kadar_zakat')
  if (saved) kadarZakat.value = JSON.parse(saved)
})
</script>
