<template>
  <div class="relative py-8 min-h-screen overflow-hidden">
    <!-- Animated Background Gradients -->
    <div class="fixed inset-0 bg-gray-50/80 z-0"></div>
    <div class="fixed top-0 left-0 w-[500px] h-[500px] bg-primary-200/20 rounded-full blur-3xl mix-blend-multiply animate-blob z-0"></div>
    <div class="fixed bottom-0 right-0 w-[400px] h-[400px] bg-indigo-200/20 rounded-full blur-3xl mix-blend-multiply animate-blob animation-delay-2000 z-0"></div>

    <UContainer class="relative z-10 max-w-4xl mx-auto">
      <BookingStepper :current-step="3" class="mb-8" />
      
      <div class="max-w-2xl mx-auto animate-fade-in-up">
        <h1 class="text-3xl font-extrabold text-gray-900 tracking-tight mb-8">Passenger Details</h1>
        
        <UCard class="mb-6 bg-white/80 backdrop-blur-xl ring-1 ring-gray-100 rounded-3xl shadow-xl">
          <template #header>
            <h3 class="font-bold text-xl text-gray-900 flex items-center gap-2">
              <div class="p-2 bg-primary-50 rounded-lg">
                <UIcon name="i-lucide-user" class="w-5 h-5 text-primary-600" />
              </div>
              Contact Information
            </h3>
          </template>
          <div class="space-y-6 pt-2">
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
              <UFormField label="Full Name" name="name" required :error="formErrors.name || undefined" class="font-medium">
                <UInput v-model="form.name" placeholder="John Doe" size="xl" class="w-full shadow-sm" icon="i-lucide-user" />
              </UFormField>
              <UFormField label="Phone Number" name="phone" required :error="formErrors.phone || undefined" class="font-medium">
                <UInput v-model="form.phone" type="tel" placeholder="08123456789" size="xl" class="w-full shadow-sm" icon="i-lucide-phone" />
              </UFormField>
            </div>
            
            <UFormField label="Email Address" name="email" required help="We will send your e-ticket here" :error="formErrors.email || undefined" class="font-medium">
              <UInput v-model="form.email" type="email" placeholder="example@email.com" size="xl" class="w-full shadow-sm" icon="i-lucide-mail" />
            </UFormField>
            
            <UFormField label="Passenger Address" name="address" required help="Only for security" :error="formErrors.address || undefined" class="font-medium">
              <UTextarea v-model="form.address" placeholder="Complete address detail" :rows="3" size="xl" class="w-full shadow-sm" />
            </UFormField>
          </div>
        </UCard>
        
        <UCard class="bg-primary-50/50 backdrop-blur-xl ring-1 ring-primary-100 rounded-3xl shadow-2xl overflow-hidden relative">
          <!-- Decorative glow inside card -->
          <div class="absolute -right-20 -top-20 w-40 h-40 bg-primary-200 rounded-full blur-3xl opacity-50"></div>
          
          <div class="relative z-10 flex flex-col sm:flex-row justify-between items-start sm:items-center gap-6 mb-6">
            <div>
              <p class="text-xs font-semibold text-primary-600/80 uppercase tracking-wider mb-1">Selected Seats</p>
              <div class="flex gap-2">
                <span v-for="s in seatsList.split(',')" :key="s" class="bg-white text-primary-900 w-8 h-8 rounded-lg flex items-center justify-center font-bold text-sm shadow-sm ring-1 ring-primary-100">
                  {{ s }}
                </span>
              </div>
            </div>
            <div class="text-left sm:text-right w-full sm:w-auto">
              <p class="text-xs font-semibold text-primary-600/80 uppercase tracking-wider mb-1">Total Payment</p>
              <p class="text-3xl font-black text-primary-700 tracking-tight">Rp {{ totalPrice.toLocaleString() }}</p>
            </div>
          </div>
          <UButton block size="xl" color="primary" class="font-bold rounded-xl shadow-lg hover:shadow-xl transition-all" @click="proceedToPayment" trailing-icon="i-lucide-credit-card">
            Proceed to Payment
          </UButton>
        </UCard>
      </div>
    </UContainer>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useMockDataStore } from '~/stores/mockData'

const route = useRoute()
const router = useRouter()
const store = useMockDataStore()

const scheduleId = Number(route.query.scheduleId)
const seatsList = (route.query.seats as string) || ''
const selectedSeatsCount = seatsList ? seatsList.split(',').length : 0

const scheduleDetails = computed(() => store.getScheduleDetails(scheduleId))
const totalPrice = computed(() => scheduleDetails.value ? scheduleDetails.value.price * selectedSeatsCount : 0)

const form = ref({
  email: '',
  name: '',
  phone: '',
  address: ''
})

const formErrors = ref({
  email: '',
  name: '',
  phone: '',
  address: ''
})

const proceedToPayment = () => {
  formErrors.value = { email: '', name: '', phone: '', address: '' }
  let hasError = false
  
  if (!form.value.name) { formErrors.value.name = 'Full name is required'; hasError = true }
  if (!form.value.phone) { formErrors.value.phone = 'Phone number is required'; hasError = true }
  if (!form.value.email) { formErrors.value.email = 'Email address is required'; hasError = true }
  else if (!/^\S+@\S+\.\S+$/.test(form.value.email)) { formErrors.value.email = 'Valid email is required'; hasError = true }
  if (!form.value.address) { formErrors.value.address = 'Passenger address is required'; hasError = true }
  
  if (hasError) return;

  // Simulate Xendit redirect
  router.push({
    path: '/checkout/payment-success',
    query: { email: form.value.email, seats: seatsList }
  })
}
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
