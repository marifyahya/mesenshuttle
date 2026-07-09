<template>
  <div class="w-full mb-8 pt-4">
    <div class="relative flex items-center justify-between max-w-2xl mx-auto">
      <!-- Background Line -->
      <div class="absolute left-0 top-4 sm:top-5 transform -translate-y-1/2 w-full h-1 bg-gray-200 z-0"></div>
      <!-- Progress Line -->
      <div class="absolute left-0 top-4 sm:top-5 transform -translate-y-1/2 h-1 bg-primary-500 z-0 transition-all duration-500" :style="{ width: ((currentStep - 1) / (steps.length - 1)) * 100 + '%' }"></div>
      
      <!-- Steps -->
      <div v-for="(step, index) in steps" :key="step.title" class="relative z-10 flex flex-col items-center">
        <div 
          class="w-8 h-8 sm:w-10 sm:h-10 rounded-full flex items-center justify-center font-bold text-xs sm:text-sm transition-colors duration-300 ring-4 ring-gray-50"
          :class="[
            currentStep > index + 1 ? 'bg-primary-500 text-white' : 
            currentStep === index + 1 ? 'bg-primary-600 text-white shadow-md' : 'bg-gray-200 text-gray-500'
          ]"
        >
          <UIcon v-if="currentStep > index + 1" name="i-lucide-check" class="w-4 h-4 sm:w-5 sm:h-5" />
          <span v-else>{{ index + 1 }}</span>
        </div>
        <span 
          class="mt-2 text-[10px] sm:text-xs font-semibold uppercase tracking-wider"
          :class="currentStep >= index + 1 ? 'text-primary-700' : 'text-gray-400'"
        >
          {{ step.title }}
        </span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps({
  currentStep: {
    type: Number,
    required: true
  }
})

const steps = [
  { title: 'Schedule' },
  { title: 'Seat' },
  { title: 'Details' },
  { title: 'Payment' }
]
</script>
