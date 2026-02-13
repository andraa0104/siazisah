<template>
  <div v-if="show" @click="$emit('close')" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
    <div @click.stop class="bg-white rounded-2xl max-w-2xl w-full max-h-[90vh] overflow-y-auto">
      <div class="bg-gradient-to-r from-blue-600 to-indigo-600 text-white p-6 rounded-t-2xl">
        <h2 class="text-2xl font-bold">{{ title }}</h2>
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
