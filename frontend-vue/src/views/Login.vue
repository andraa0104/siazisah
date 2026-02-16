<template>
  <div class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center p-4">
    <div class="max-w-md w-full">
      <div class="text-center mb-8">
        <div class="inline-flex items-center justify-center w-16 h-16 bg-gradient-to-br from-blue-600 to-indigo-600 rounded-2xl mb-4 shadow-lg">
          <i class="fas fa-mosque text-white text-2xl"></i>
        </div>
        <h1 class="text-3xl font-bold text-gray-800 mb-2">SI-AZISAH</h1>
        <p class="text-gray-600">Sistem Informasi Zakat</p>
      </div>

      <div class="bg-white rounded-2xl shadow-xl p-8">
        <h2 class="text-2xl font-bold text-gray-800 mb-6">Masuk ke Akun</h2>
        
        <form @submit.prevent="handleLogin" class="space-y-5">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Username</label>
            <div class="relative">
              <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                <i class="fas fa-user text-gray-400"></i>
              </div>
              <input v-model="form.username" type="text" required
                class="block w-full pl-10 pr-3 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition"
                placeholder="Masukkan username">
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Password</label>
            <div class="relative">
              <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                <i class="fas fa-lock text-gray-400"></i>
              </div>
              <input v-model="form.password" type="password" required
                class="block w-full pl-10 pr-3 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition"
                placeholder="Masukkan password">
            </div>
          </div>

          <div v-if="error" class="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded-lg text-sm">
            {{ error }}
          </div>

          <button type="submit" :disabled="loading"
            class="w-full bg-gradient-to-r from-blue-600 to-indigo-600 text-white py-3 rounded-lg font-semibold hover:from-blue-700 hover:to-indigo-700 transition duration-200 shadow-lg hover:shadow-xl disabled:opacity-50">
            <i class="fas fa-sign-in-alt mr-2"></i>{{ loading ? 'Loading...' : 'Masuk' }}
          </button>
        </form>

        <div class="mt-6 text-center space-y-2">
          <div>
            <router-link to="/public" class="text-blue-600 hover:text-blue-700 text-sm font-medium">
              <i class="fas fa-globe mr-1"></i>Lihat Dashboard Publik
            </router-link>
          </div>
          <div>
            <router-link to="/" class="text-gray-600 hover:text-gray-700 text-sm font-medium">
              <i class="fas fa-arrow-left mr-1"></i>Kembali ke Beranda
            </router-link>
          </div>
        </div>
      </div>

      <p class="text-center text-gray-600 text-sm mt-6">© 2024 SI-AZISAH. All rights reserved.</p>
    </div>

    <LoadingOverlay :show="loading" message="Memproses login..." />
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import LoadingOverlay from '../components/LoadingOverlay.vue'

const router = useRouter()
const authStore = useAuthStore()

const form = ref({ username: '', password: '' })
const loading = ref(false)
const error = ref('')

const handleLogin = async () => {
  loading.value = true
  error.value = ''
  try {
    const result = await authStore.login(form.value)
    if (result.success) {
      if (authStore.isSuperadmin) router.push('/superadmin')
      else if (authStore.isPetugas) router.push('/petugas')
    } else {
      error.value = result.message
    }
  } catch (err) {
    error.value = 'Terjadi kesalahan. Silakan coba lagi.'
  } finally {
    loading.value = false
  }
}
</script>
