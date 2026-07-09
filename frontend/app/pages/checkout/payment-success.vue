<template>
  <div class="py-12 bg-gray-50 min-h-screen">
    <UContainer class="max-w-3xl mx-auto">
      <BookingStepper :current-step="4" />

      <div class="mt-8 flex justify-center px-4">
        <UCard class="w-full max-w-lg shadow-xl border-t-4 border-t-green-500 rounded-2xl overflow-hidden relative" :ui="{ body: { padding: '!p-0' } }">
          
          <!-- Success Header -->
          <div class="bg-white p-8 text-center relative z-10">
            <div class="mx-auto flex items-center justify-center h-20 w-20 rounded-full bg-green-50 mb-6 border-4 border-green-100 shadow-sm">
              <UIcon name="i-heroicons-check" class="h-10 w-10 text-green-600 animate-[pulse_2s_ease-in-out_infinite]" />
            </div>
            <h2 class="text-3xl font-extrabold text-gray-900 mb-3 tracking-tight">Payment Successful!</h2>
            <p class="text-gray-500 text-sm">
              Your booking has been confirmed. We've sent the E-Ticket to <br/>
              <strong class="text-gray-900 mt-1 inline-block">{{ email }}</strong>
            </p>
          </div>
          
          <!-- Ticket Details Divider -->
          <div class="relative h-12 bg-gray-50 flex items-center justify-between overflow-hidden">
            <div class="absolute -left-6 w-12 h-12 bg-gray-100 rounded-full shadow-inner"></div>
            <div class="w-full border-t-2 border-dashed border-gray-300 mx-8"></div>
            <div class="absolute -right-6 w-12 h-12 bg-gray-100 rounded-full shadow-inner"></div>
          </div>
          
          <!-- Ticket Details -->
          <div class="bg-gray-50 p-8 pt-2">
            <div class="bg-white rounded-xl p-5 border border-gray-200 shadow-sm mb-8">
              <div class="grid grid-cols-2 gap-4">
                <div>
                  <p class="text-xs font-semibold text-gray-400 uppercase tracking-wider mb-1">Booking ID</p>
                  <p class="font-mono font-bold text-lg text-primary-700">MB-{{ bookingId }}</p>
                </div>
                <div class="text-right">
                  <p class="text-xs font-semibold text-gray-400 uppercase tracking-wider mb-1">Seats</p>
                  <p class="font-bold text-lg text-gray-900">{{ seats }}</p>
                </div>
              </div>
            </div>
            
            <div class="flex flex-col sm:flex-row gap-3">
              <UButton block color="primary" size="lg" class="flex-1 font-bold shadow-md justify-center" icon="i-lucide-download">
                Download E-Ticket
              </UButton>
              <UButton to="/" block variant="outline" size="lg" class="flex-1 font-semibold justify-center">
                Return to Home
              </UButton>
            </div>
          </div>
          
        </UCard>
      </div>
    </UContainer>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const email = computed(() => route.query.email as string || 'passenger@example.com')
const seats = computed(() => route.query.seats as string || '-')

const bookingId = ref('000000')

onMounted(() => {
  bookingId.value = Math.floor(100000 + Math.random() * 900000).toString()
})
</script>
