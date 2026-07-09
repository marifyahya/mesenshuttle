<template>
  <div class="py-8 bg-gray-50 min-h-screen">
    <UContainer>
      <BookingStepper :current-step="3" />
      
      <div class="max-w-2xl mx-auto">
        <h1 class="text-2xl font-bold mb-6">Passenger Details</h1>
        
        <UCard class="mb-6">
          <template #header>
            <h3 class="font-bold">Contact Information (Guest Checkout)</h3>
          </template>
          <div class="space-y-4">
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
              <UFormField label="Full Name" name="name" required :error="formErrors.name || undefined">
                <UInput v-model="form.name" placeholder="John Doe" size="lg" class="w-full" />
              </UFormField>
              <UFormField label="Phone Number" name="phone" required :error="formErrors.phone || undefined">
                <UInput v-model="form.phone" type="tel" placeholder="08123456789" size="lg" class="w-full" />
              </UFormField>
            </div>
            
            <UFormField label="Email Address" name="email" required help="We will send your e-ticket here" :error="formErrors.email || undefined">
              <UInput v-model="form.email" type="email" placeholder="example@email.com" size="lg" class="w-full" />
            </UFormField>
            
            <UFormField label="Passenger Address" name="address" required help="Only for security" :error="formErrors.address || undefined">
              <UTextarea v-model="form.address" placeholder="Complete address detail" :rows="3" size="lg" class="w-full" />
            </UFormField>
          </div>
        </UCard>
        
        <UCard>
          <div class="flex justify-between items-center mb-4">
            <div>
              <p class="text-sm text-gray-500">Seats</p>
              <p class="font-bold">{{ seatsList }}</p>
            </div>
            <div class="text-right">
              <p class="text-sm text-gray-500">Total Price</p>
              <p class="text-xl font-bold text-primary-600">Rp {{ totalPrice.toLocaleString() }}</p>
            </div>
          </div>
          <UButton block size="lg" color="primary" @click="proceedToPayment">
            Proceed to Payment (Xendit)
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
