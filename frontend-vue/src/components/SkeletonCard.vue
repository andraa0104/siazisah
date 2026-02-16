<template>
  <div class="animate-pulse">
    <div v-if="type === 'stat'" class="bg-gray-100 rounded-xl p-6">
      <div class="h-4 bg-gray-300 rounded w-3/4 mb-3"></div>
      <div class="h-8 bg-gray-300 rounded w-1/2"></div>
    </div>

    <div v-else-if="type === 'card'" class="bg-white rounded-xl shadow-md p-6">
      <div class="h-6 bg-gray-300 rounded w-1/2 mb-4"></div>
      <div class="space-y-3">
        <div class="h-4 bg-gray-200 rounded"></div>
        <div class="h-4 bg-gray-200 rounded w-5/6"></div>
        <div class="h-4 bg-gray-200 rounded w-4/6"></div>
      </div>
    </div>

    <div v-else-if="type === 'table-row'" class="bg-white border-b border-gray-200">
      <div class="px-6 py-4 flex gap-4">
        <div class="h-4 bg-gray-200 rounded flex-1"></div>
        <div class="h-4 bg-gray-200 rounded flex-1"></div>
        <div class="h-4 bg-gray-200 rounded flex-1"></div>
        <div class="h-8 bg-gray-200 rounded w-20"></div>
      </div>
    </div>

    <div v-else-if="type === 'list-item'" class="bg-gray-50 rounded-lg p-4 mb-3">
      <div class="h-5 bg-gray-300 rounded w-3/4 mb-2"></div>
      <div class="h-3 bg-gray-200 rounded w-1/2 mb-2"></div>
      <div class="h-3 bg-gray-200 rounded w-2/3"></div>
    </div>

    <div v-else class="space-y-3">
      <div class="h-4 bg-gray-200 rounded"></div>
      <div class="h-4 bg-gray-200 rounded w-5/6"></div>
      <div class="h-4 bg-gray-200 rounded w-4/6"></div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  type: {
    type: String,
    default: 'default'
  }
})
</script>
