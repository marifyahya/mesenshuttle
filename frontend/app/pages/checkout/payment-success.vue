<template>
  <div class="relative py-12 min-h-screen overflow-hidden">
    <!-- Animated Background Gradients -->
    <div class="fixed inset-0 bg-gray-50/80 z-0"></div>
    <div class="fixed -top-20 right-0 w-[600px] h-[600px] bg-green-200/30 rounded-full blur-3xl mix-blend-multiply animate-blob z-0"></div>
    <div class="fixed bottom-0 -left-20 w-[500px] h-[500px] bg-emerald-200/30 rounded-full blur-3xl mix-blend-multiply animate-blob animation-delay-2000 z-0"></div>

    <UContainer class="relative z-10 max-w-3xl mx-auto">
      <BookingStepper :current-step="4" class="mb-8" />

      <div class="mt-12 flex justify-center px-4">
        <div class="w-full max-w-lg animate-fade-in-up">
          <UCard class="shadow-2xl ring-1 ring-white/80 rounded-4xl overflow-hidden relative bg-white/90 backdrop-blur-xl border-t-8 border-t-green-500" :ui="{ body: '!p-0' }">
            
            <!-- Success Header -->
            <div class="bg-transparent p-10 text-center relative z-10">
              <div class="mx-auto flex items-center justify-center h-24 w-24 rounded-full bg-green-50 mb-6 ring-8 ring-green-50/50 shadow-inner">
                <UIcon name="i-heroicons-check" class="h-12 w-12 text-green-500" />
              </div>
              <h2 class="text-3xl font-extrabold text-gray-900 mb-4 tracking-tight">Payment Successful!</h2>
              <p class="text-gray-500 text-base leading-relaxed">
                Your booking has been confirmed. We've sent the E-Ticket to <br/>
                <strong class="text-gray-900 mt-2 inline-block font-semibold bg-gray-100 px-3 py-1 rounded-lg">{{ email }}</strong>
              </p>
            </div>
            
            <!-- Ticket Details Divider -->
            <div class="relative h-16 bg-gray-50/50 flex items-center justify-between overflow-hidden">
              <div class="absolute -left-8 w-16 h-16 bg-gray-50 rounded-full shadow-inner ring-1 ring-gray-200/50"></div>
              <div class="w-full border-t-2 border-dashed border-gray-300 mx-10"></div>
              <div class="absolute -right-8 w-16 h-16 bg-gray-50 rounded-full shadow-inner ring-1 ring-gray-200/50"></div>
            </div>
            
            <!-- Ticket Details -->
            <div class="bg-gray-50/50 p-8 pt-4">
              <div class="bg-white/80 rounded-2xl p-6 ring-1 ring-gray-100 shadow-sm mb-8 backdrop-blur-sm relative overflow-hidden">
                <div class="absolute right-0 top-0 w-32 h-32 bg-primary-100/50 rounded-full blur-2xl -mr-16 -mt-16"></div>
                <div class="grid grid-cols-2 gap-6 relative z-10">
                  <div>
                    <p class="text-xs font-semibold text-gray-400 uppercase tracking-wider mb-1.5">Booking ID</p>
                    <p class="font-mono font-bold text-xl text-primary-700">MB-{{ bookingId }}</p>
                  </div>
                  <div class="text-right">
                    <p class="text-xs font-semibold text-gray-400 uppercase tracking-wider mb-1.5">Seats</p>
                    <p class="font-bold text-xl text-gray-900">{{ seats }}</p>
                  </div>
                </div>
              </div>
              
              <div class="flex flex-col sm:flex-row gap-4">
                <UButton block color="primary" size="xl" class="flex-1 font-bold shadow-lg hover:shadow-xl hover:-translate-y-0.5 transition-all justify-center rounded-xl" icon="i-lucide-download">
                  Download E-Ticket
                </UButton>
                <UButton to="/" block variant="soft" color="neutral" size="xl" class="flex-1 font-semibold justify-center rounded-xl bg-white hover:bg-gray-50 ring-1 ring-gray-200 shadow-sm">
                  Return to Home
                </UButton>
              </div>
            </div>
            
          </UCard>
        </div>
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

<style scoped>
@keyframes blob {
  0% { transform: translate(0px, 0px) scale(1); }
  33% { transform: translate(30px, -50px) scale(1.1); }
  66% { transform: translate(-20px, 20px) scale(0.9); }
  100% { transform: translate(0px, 0px) scale(1); }
}
.animate-blob {
  animation: blob 7s infinite;
}
.animation-delay-2000 {
  animation-delay: 2s;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translate3d(0, 20px, 0);
  }
  to {
    opacity: 1;
    transform: translate3d(0, 0, 0);
  }
}
.animate-fade-in-up {
  animation: fadeInUp 0.5s ease-out forwards;
  opacity: 0;
}
</style>
