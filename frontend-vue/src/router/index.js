import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const routes = [
  { path: '/', name: 'Home', component: () => import('../views/Home.vue') },
  { path: '/login', name: 'Login', component: () => import('../views/Login.vue') },
  { path: '/public', name: 'Public', component: () => import('../views/Public.vue') },
  
  // Superadmin
  { path: '/superadmin', name: 'SuperadminDashboard', component: () => import('../views/superadmin/Dashboard.vue'), meta: { requiresAuth: true, role: 'superadmin' } },
  { path: '/superadmin/masjid', name: 'SuperadminMasjid', component: () => import('../views/superadmin/Masjid.vue'), meta: { requiresAuth: true, role: 'superadmin' } },
  { path: '/superadmin/users', name: 'SuperadminUsers', component: () => import('../views/superadmin/Users.vue'), meta: { requiresAuth: true, role: 'superadmin' } },
  
  // Petugas
  { path: '/petugas', name: 'PetugasDashboard', component: () => import('../views/petugas/Dashboard.vue'), meta: { requiresAuth: true, role: 'petugas' } },
  { path: '/petugas/pengaturan', name: 'PetugasPengaturan', component: () => import('../views/petugas/Pengaturan.vue'), meta: { requiresAuth: true, role: 'petugas' } },
  { path: '/petugas/mustahiq', name: 'PetugasMustahiq', component: () => import('../views/petugas/Mustahiq.vue'), meta: { requiresAuth: true, role: 'petugas' } },
  { path: '/petugas/transaksi', name: 'PetugasTransaksi', component: () => import('../views/petugas/Transaksi.vue'), meta: { requiresAuth: true, role: 'petugas' } }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  
  if (to.meta.requiresAuth && !authStore.isLoggedIn) {
    next('/login')
  } else if (to.meta.role && authStore.user?.role !== to.meta.role) {
    next('/login')
  } else {
    next()
  }
})

export default router
