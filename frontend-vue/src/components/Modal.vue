<template>
  <div v-if="show" @click="$emit('close')" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
    <div @click.stop class="bg-white rounded-2xl max-w-2xl w-full max-h-[90vh] overflow-y-auto">
      <div class="bg-gradient-to-r from-blue-600 to-indigo-600 text-white p-6 rounded-t-2xl flex items-start justify-between gap-4">
        <h2 class="text-2xl font-bold leading-tight">{{ title }}</h2>
        <button
          type="button"
          class="shrink-0 w-10 h-10 rounded-xl inline-flex items-center justify-center text-white/90 hover:text-white hover:bg-white/15 transition"
          aria-label="Tutup"
          title="Tutup"
          @click="$emit('close')"
        >
          <i class="fas fa-times"></i>
        </button>
      </div>
      <div class="p-6">
        <slot></slot>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  show: Boolean,
  title: String
})
defineEmits(['close'])
</script>
