import { defineStore } from 'pinia'
import api from '../api'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: JSON.parse(localStorage.getItem('user')) || null,
    token: localStorage.getItem('token') || null
  }),
  
  getters: {
    isLoggedIn: (state) => !!state.token,
    isSuperadmin: (state) => state.user?.role === 'superadmin',
    isPetugas: (state) => state.user?.role === 'petugas'
  },
  
  actions: {
    async login(credentials) {
      try {
        const { data } = await api.login(credentials)
        if (data.success) {
          this.user = data.data.user
          this.token = data.data.token
          localStorage.setItem('user', JSON.stringify(data.data.user))
          localStorage.setItem('token', data.data.token)
        }
        return data
      } catch (error) {
        return { success: false, message: error.response?.data?.message || 'Terjadi kesalahan. Silakan coba lagi.' }
      }
    },
    
    logout() {
      this.user = null
      this.token = null
      localStorage.removeItem('user')
      localStorage.removeItem('token')
    }
  }
})
