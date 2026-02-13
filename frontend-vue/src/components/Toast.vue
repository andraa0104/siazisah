<template>
  <Teleport to="body">
    <div v-if="visible" :class="`fixed top-4 right-4 px-6 py-3 rounded-lg shadow-lg text-white z-50 transition-all ${bgColor}`">
      <div class="flex items-center gap-2">
        <i :class="`fas fa-${icon}`"></i>
        <span>{{ message }}</span>
      </div>
    </div>
  </Teleport>
</template>

<script setup>
import { ref, computed, watch } from 'vue'

const props = defineProps({
  message: String,
  type: { type: String, default: 'success' }
})

const visible = ref(false)

const bgColor = computed(() => {
  const colors = {
    success: 'bg-green-500',
    error: 'bg-red-500',
    warning: 'bg-yellow-500',
    info: 'bg-blue-500'
  }
  return colors[props.type] || colors.success
})

const icon = computed(() => {
  const icons = {
    success: 'check-circle',
    error: 'exclamation-circle',
    warning: 'exclamation-triangle',
    info: 'info-circle'
  }
  return icons[props.type] || icons.success
})

watch(() => props.message, (val) => {
  if (val) {
    visible.value = true
    setTimeout(() => visible.value = false, 3000)
  }
})
</script>
