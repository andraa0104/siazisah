import axios from 'axios'

const resolveApiBaseUrl = () => {
  const rawUrl = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8082/api'
  const withoutTrailingSlash = rawUrl.replace(/\/$/, '')
  return withoutTrailingSlash.endsWith('/api')
    ? withoutTrailingSlash
    : `${withoutTrailingSlash}/api`
}

const API_BASE_URL = resolveApiBaseUrl()

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {'Content-Type': 'application/json'}
})

api.interceptors.request.use(config => {
  const token = localStorage.getItem('token')
  if (token) config.headers.Authorization = `Bearer ${token}`
  return config
})

export default {
  // Auth
  login: (data) => api.post('/auth/login', data),
  
  // Public
  getPublicDashboard: () => api.get('/public/dashboard'),
  getPublicMasjid: () => api.get('/public/masjid'),
  getPublicMasjidStats: (id, params) => api.get(`/public/masjid/${id}/stats`, { params }),
  getPublicMasjidTransaksi: (id, params) => api.get(`/public/masjid/${id}/transaksi`, { params }),
  getPublicMasjidMustahiq: (id, params) => api.get(`/public/masjid/${id}/mustahiq`, { params }),
  getPublicMasjidDistribusi: (id, params) => api.get(`/public/masjid/${id}/distribusi`, { params }),
  getPublicPengurusMasjid: (id) => api.get(`/public/masjid/${id}/pengurus-masjid`),
  getPublicPengurusZakat: (id) => api.get(`/public/masjid/${id}/pengurus-zakat`),
  
  // Superadmin
  getSuperadminDataCompleteness: () => api.get('/superadmin/dashboard/data-completeness'),
  getMasjid: (params) => api.get('/superadmin/masjid', { params }),
  createMasjid: (data) => api.post('/superadmin/masjid', data),
  updateMasjid: (id, data) => api.put(`/superadmin/masjid/${id}`, data),
  deleteMasjid: (id) => api.delete(`/superadmin/masjid/${id}`),
  getPrintSuperadminZakatData: (signDate) => api.get('/superadmin/reports/print-zakat', {
    params: { sign_date: signDate, _ts: Date.now() },
    responseType: 'text'
  }),
  getPrintSuperadminMustahiqGlobalData: (signDate) => api.get('/superadmin/reports/print-mustahiq-global', {
    params: { sign_date: signDate, _ts: Date.now() },
    responseType: 'text'
  }),
  
  getUsers: (params) => api.get('/superadmin/users', { params }),
  createUser: (data) => api.post('/superadmin/users', data),
  updateUser: (id, data) => api.put(`/superadmin/users/${id}`, data),
  deleteUser: (id) => api.delete(`/superadmin/users/${id}`),
  
  // Petugas
  getMyMasjid: () => api.get('/petugas/masjid'),
  updateMyMasjid: (data) => api.put('/petugas/masjid', data),
  getPengaturanZakat: () => api.get('/petugas/pengaturan-zakat'),
  updatePengaturanZakat: (data) => api.put('/petugas/pengaturan-zakat', data),
  
  getPengurusMasjid: () => api.get('/petugas/pengurus-masjid'),
  savePengurusMasjid: (data) => api.post('/petugas/pengurus-masjid', data),
  
  getPengurusZakat: () => api.get('/petugas/pengurus-zakat'),
  savePengurusZakat: (data) => api.post('/petugas/pengurus-zakat', data),
  deletePengurusZakat: (id) => api.delete(`/petugas/pengurus-zakat/${id}`),
  
  getMuzakki: () => api.get('/petugas/muzakki'),
  createMuzakki: (data) => api.post('/petugas/muzakki', data),
  
  checkMustahiqDuplicate: (data) => api.post('/petugas/mustahiq/check', data),
  getMustahiq: (params) => api.get('/petugas/mustahiq', { params }),
  createMustahiq: (data) => api.post('/petugas/mustahiq', data),
  updateMustahiq: (id, data) => api.put(`/petugas/mustahiq/${id}`, data),
  deleteMustahiq: (id) => api.delete(`/petugas/mustahiq/${id}`),
  
  getTransaksi: (params) => api.get('/petugas/transaksi', { params }),
  createTransaksi: (data) => api.post('/petugas/transaksi', data),
  deleteTransaksi: (id) => api.delete(`/petugas/transaksi/${id}`),
  printTransaksi: (id) => `${API_BASE_URL}/petugas/transaksi/${id}/print`,
  getPrintTransaksiData: (signDate) => api.get('/petugas/transaksi/print-data', {
    params: { sign_date: signDate, _ts: Date.now() },
    responseType: 'text'
  }),

  getDistribusiInsight: () => api.get('/petugas/distribusi/insight'),
  getDistribusi: (params) => api.get('/petugas/distribusi', { params }),
  getPrintDistribusiData: (signDate) => api.get('/petugas/distribusi/print-data', {
    params: { sign_date: signDate, _ts: Date.now() },
    responseType: 'text'
  }),
  createDistribusi: (data) => api.post('/petugas/distribusi', data),
  updateDistribusi: (id, data) => api.put(`/petugas/distribusi/${id}`, data),
  deleteDistribusi: (id) => api.delete(`/petugas/distribusi/${id}`)
}
