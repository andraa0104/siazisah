<template>
  <div v-if="show" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-white rounded-xl p-8 shadow-2xl flex flex-col items-center">
      <div class="relative w-16 h-16 mb-4">
        <div class="absolute inset-0 border-4 border-blue-200 rounded-full"></div>
        <div class="absolute inset-0 border-4 border-blue-600 rounded-full border-t-transparent animate-spin"></div>
      </div>
      <p class="text-gray-700 font-medium">{{ message }}</p>
    </div>
  </div>
</template>

<script setup>
defineProps({
  show: {
    type: Boolean,
    default: false
  },
  message: {
    type: String,
    default: 'Memproses data...'
  }
})
</script>
