<template>
  <div class="flex flex-col md:flex-row md:items-center justify-between gap-4">
    <div>
      <h1 class="text-2xl font-bold text-gray-900 tracking-tight">{{ title }}</h1>
      <p class="text-gray-500 mt-1 text-sm">{{ description }}</p>
    </div>
    <div class="flex items-center gap-3">
      <UButton 
        color="primary" 
        size="lg" 
        :icon="buttonIcon"
        class="font-semibold shadow-md hover:shadow-lg transition-all hover:-translate-y-0.5 rounded-xl px-5"
        @click="$emit('action')"
      >
        {{ buttonLabel }}
      </UButton>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  title: string
  description: string
  buttonLabel: string
  buttonIcon: string
}>()

defineEmits(['action'])
</script>
