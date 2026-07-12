<template>
  <div class="relative py-8 min-h-screen overflow-hidden">
    <!-- Animated Background Gradients -->
    <div class="fixed inset-0 bg-gray-50/80 z-0"></div>
    <div class="fixed top-20 right-20 w-[500px] h-[500px] bg-primary-200/20 rounded-full blur-3xl mix-blend-multiply animate-blob z-0"></div>
    <div class="fixed -bottom-40 -left-20 w-[400px] h-[400px] bg-indigo-200/20 rounded-full blur-3xl mix-blend-multiply animate-blob animation-delay-2000 z-0"></div>

    <UContainer class="relative z-10 max-w-6xl mx-auto">
      <BookingStepper :current-step="2" class="mb-8" />
      <UButton variant="ghost" to="/schedules" class="mb-6 font-semibold hover:-translate-x-1 transition-transform" icon="i-heroicons-arrow-left">Back to Schedules</UButton>
      
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- Seat Map Area -->
        <div class="lg:col-span-2">
          <UCard class="bg-white/80 backdrop-blur-xl ring-1 ring-gray-100 rounded-3xl shadow-xl animate-fade-in-up">
            <template #header>
              <h3 class="font-bold text-xl text-gray-900 flex items-center gap-2">
                <UIcon name="i-lucide-sofa" class="w-6 h-6 text-primary-600" />
                Select Your Seat
              </h3>
            </template>
            
            <div class="text-center mb-8">
              <div class="inline-flex justify-center gap-6 text-sm bg-gray-50 p-3 rounded-2xl ring-1 ring-gray-200 shadow-inner">
                <div class="flex items-center gap-2">
                  <div class="w-5 h-5 bg-white ring-2 ring-primary-500 rounded-md"></div>
                  <span class="font-semibold text-gray-700">Available</span>
                </div>
                <div class="flex items-center gap-2">
                  <div class="w-5 h-5 bg-primary-600 ring-2 ring-primary-600 rounded-md"></div>
                  <span class="font-semibold text-gray-700">Selected</span>
                </div>
                <div class="flex items-center gap-2">
                  <div class="w-5 h-5 bg-gray-200 ring-2 ring-gray-300 rounded-md relative overflow-hidden">
                     <div class="absolute inset-0 bg-[repeating-linear-gradient(45deg,transparent,transparent_2px,#d1d5db_2px,#d1d5db_4px)]"></div>
                  </div>
                  <span class="font-semibold text-gray-700">Booked</span>
                </div>
              </div>
            </div>

            <div v-if="scheduleDetails" class="max-w-md mx-auto bg-gray-50/50 p-6 rounded-[3rem] ring-4 ring-gray-100 relative">
              
              <!-- Front of Shuttle (Driver) -->
              <div class="absolute -top-4 left-1/2 -translate-x-1/2 bg-gray-200 px-6 py-2 rounded-b-xl border-x-4 border-b-4 border-gray-300 flex items-center gap-2 shadow-inner text-gray-500 font-bold text-xs">
                <UIcon name="i-lucide-steering-wheel" class="w-4 h-4" />
                FRONT
              </div>

              <div class="mt-8 mb-6"></div> <!-- Spacer for front -->

              <!-- Layout Standard -->
              <template v-if="scheduleDetails?.fleet?.type === 'Reguler'">
                <div class="grid grid-cols-3 gap-6">
                  <!-- Row 1 -->
                  <button @click="toggleSeat('1')" :class="seatClass('1')" class="h-14 sm:h-16 rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">1</button>
                  <div></div>
                  <div class="h-14 sm:h-16 rounded-2xl flex flex-col items-center justify-center font-bold text-xs bg-gray-100 text-gray-400 border-2 border-gray-200 border-dashed"><UIcon name="i-lucide-steering-wheel" class="w-6 h-6 mb-1"/>Supir</div>
                  
                  <!-- Row 2 -->
                  <div></div>
                  <button @click="toggleSeat('2')" :class="seatClass('2')" class="h-14 sm:h-16 rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">2</button>
                  <button @click="toggleSeat('3')" :class="seatClass('3')" class="h-14 sm:h-16 rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">3</button>
                  
                  <!-- Row 3 -->
                  <button @click="toggleSeat('4')" :class="seatClass('4')" class="h-14 sm:h-16 rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">4</button>
                  <button @click="toggleSeat('5')" :class="seatClass('5')" class="h-14 sm:h-16 rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">5</button>
                  <button @click="toggleSeat('6')" :class="seatClass('6')" class="h-14 sm:h-16 rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">6</button>
                  
                  <!-- Row 4 -->
                  <button @click="toggleSeat('7')" :class="seatClass('7')" class="h-14 sm:h-16 rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">7</button>
                  <button @click="toggleSeat('8')" :class="seatClass('8')" class="h-14 sm:h-16 rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">8</button>
                  <button @click="toggleSeat('9')" :class="seatClass('9')" class="h-14 sm:h-16 rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">9</button>
                  
                  <!-- Row 5 -->
                  <button @click="toggleSeat('10')" :class="seatClass('10')" class="h-14 sm:h-16 rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">10</button>
                  <button @click="toggleSeat('11')" :class="seatClass('11')" class="h-14 sm:h-16 rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">11</button>
                  <button @click="toggleSeat('12')" :class="seatClass('12')" class="h-14 sm:h-16 rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">12</button>
                </div>
              </template>

              <!-- Layout Premium 10 Seats -->
              <template v-else-if="scheduleDetails?.fleet?.type === 'Premium' && scheduleDetails?.fleet?.capacity === 10">
                <div class="grid grid-cols-4 gap-4">
                  <!-- Row 1 -->
                  <button @click="toggleSeat('1')" :class="seatClass('1')" class="h-14 sm:h-16 rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">1</button>
                  <div></div>
                  <div></div>
                  <div class="h-14 sm:h-16 rounded-2xl flex flex-col items-center justify-center font-bold text-xs bg-gray-100 text-gray-400 border-2 border-gray-200 border-dashed"><UIcon name="i-lucide-steering-wheel" class="w-6 h-6 mb-1"/>Supir</div>
                  
                  <!-- Row 2 -->
                  <div></div>
                  <div></div>
                  <button @click="toggleSeat('2')" :class="seatClass('2')" class="h-14 sm:h-16 rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">2</button>
                  <button @click="toggleSeat('3')" :class="seatClass('3')" class="h-14 sm:h-16 rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">3</button>
                  
                  <!-- Row 3 -->
                  <button @click="toggleSeat('4')" :class="seatClass('4')" class="h-14 sm:h-16 rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">4</button>
                  <div></div>
                  <button @click="toggleSeat('5')" :class="seatClass('5')" class="h-14 sm:h-16 rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">5</button>
                  <button @click="toggleSeat('6')" :class="seatClass('6')" class="h-14 sm:h-16 rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">6</button>
                  
                  <!-- Row 4 -->
                  <button @click="toggleSeat('7')" :class="seatClass('7')" class="h-14 sm:h-16 rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">7</button>
                  <button @click="toggleSeat('8')" :class="seatClass('8')" class="h-14 sm:h-16 rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">8</button>
                  <button @click="toggleSeat('9')" :class="seatClass('9')" class="h-14 sm:h-16 rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">9</button>
                  <button @click="toggleSeat('10')" :class="seatClass('10')" class="h-14 sm:h-16 rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">10</button>
                </div>
              </template>

              <!-- Layout Premium 8 Seats (Captain Seats) -->
              <template v-else-if="scheduleDetails?.fleet?.type === 'Premium'">
                <div class="flex justify-center gap-10 mb-6">
                  <div class="w-16 h-16"><button @click="toggleSeat('1')" :class="seatClass('1')" class="w-full h-full rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">1</button></div>
                  <div class="w-16 h-16 driver-seat"><UIcon name="i-lucide-steering-wheel" class="w-6 h-6 mb-1"/>DRIVER</div>
                </div>
                <div v-for="(row, idx) in [[2,3], [4,5]]" :key="idx" class="flex justify-center gap-12 mb-6">
                  <div class="w-16 h-16"><button @click="toggleSeat(String(row[0]))" :class="seatClass(String(row[0]))" class="w-full h-full rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">{{row[0]}}</button></div>
                  <div class="w-16 h-16"><button @click="toggleSeat(String(row[1]))" :class="seatClass(String(row[1]))" class="w-full h-full rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">{{row[1]}}</button></div>
                </div>
                <div class="flex justify-center gap-4 mt-8">
                   <div class="w-16 h-16"><button @click="toggleSeat('6')" :class="seatClass('6')" class="w-full h-full rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">6</button></div>
                   <div class="w-16 h-16"><button @click="toggleSeat('7')" :class="seatClass('7')" class="w-full h-full rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">7</button></div>
                   <div class="w-16 h-16"><button @click="toggleSeat('8')" :class="seatClass('8')" class="w-full h-full rounded-2xl flex items-center justify-center font-bold text-lg transition-all duration-200">8</button></div>
                </div>
              </template>
            </div>
          </UCard>
        </div>

        <!-- Booking Summary -->
        <div class="lg:col-span-1">
          <div class="sticky top-24 animate-fade-in-up animation-delay-300">
            <UCard class="bg-white/80 backdrop-blur-xl ring-1 ring-gray-100 rounded-3xl shadow-2xl">
              <template #header>
                <h3 class="font-bold text-xl text-gray-900 flex items-center gap-2">
                  <UIcon name="i-lucide-receipt" class="w-5 h-5 text-primary-600" />
                  Booking Summary
                </h3>
              </template>
              <div class="space-y-5" v-if="scheduleDetails">
                <div class="bg-gray-50 p-4 rounded-2xl ring-1 ring-gray-100">
                  <p class="text-xs font-semibold text-gray-400 uppercase tracking-wider mb-1">Route</p>
                  <p class="font-bold text-gray-900 flex items-center gap-2">
                    {{ scheduleDetails?.route?.originPool }}
                    <UIcon name="i-lucide-arrow-right" class="w-4 h-4 text-primary-500" />
                    {{ scheduleDetails?.route?.destinationPool }}
                  </p>
                </div>
                <div class="grid grid-cols-2 gap-4">
                  <div class="bg-gray-50 p-4 rounded-2xl ring-1 ring-gray-100">
                    <p class="text-xs font-semibold text-gray-400 uppercase tracking-wider mb-1">Departure</p>
                    <p class="font-bold text-gray-900">{{ new Date(scheduleDetails.departureTime).toLocaleTimeString([], {hour: '2-digit', minute:'2-digit'}) }}</p>
                  </div>
                  <div class="bg-gray-50 p-4 rounded-2xl ring-1 ring-gray-100">
                    <p class="text-xs font-semibold text-gray-400 uppercase tracking-wider mb-1">Shuttle</p>
                    <p class="font-bold text-gray-900">{{ scheduleDetails?.fleet?.type }}</p>
                  </div>
                </div>
                
                <UDivider class="my-4" />
                
                <div class="flex justify-between items-center bg-primary-50 p-4 rounded-2xl ring-1 ring-primary-100">
                  <span class="font-semibold text-primary-900">Selected Seats:</span>
                  <div class="flex gap-1 flex-wrap justify-end">
                    <span v-if="selectedSeats.length === 0" class="text-primary-400 font-medium text-sm">None</span>
                    <span v-for="s in selectedSeats" :key="s" class="bg-primary-600 text-white w-7 h-7 rounded-lg flex items-center justify-center font-bold text-xs shadow-sm">
                      {{ s }}
                    </span>
                  </div>
                </div>
                <div class="flex justify-between items-center text-xl mt-4">
                  <span class="font-bold text-gray-500">Total:</span>
                  <span class="font-black text-2xl text-primary-600">Rp {{ (selectedSeats.length * scheduleDetails.price).toLocaleString() }}</span>
                </div>
              </div>
              <template #footer>
                <UButton :disabled="selectedSeats.length === 0" block color="primary" size="xl" class="font-bold rounded-xl shadow-lg hover:shadow-xl transition-all" @click="proceedToDetails" trailing-icon="i-lucide-arrow-right">
                  Continue to Details
                </UButton>
              </template>
            </UCard>
          </div>
        </div>
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

const scheduleId = Number(route.params.id)
const scheduleDetails = computed(() => store.getScheduleDetails(scheduleId))

// Mock booked seats
const bookedSeats = ref(['2', '5', '8']) 
const selectedSeats = ref<string[]>([])

const toggleSeat = (seatId: string) => {
  if (bookedSeats.value.includes(seatId)) return;
  
  const index = selectedSeats.value.indexOf(seatId);
  if (index > -1) {
    selectedSeats.value.splice(index, 1);
  } else {
    selectedSeats.value.push(seatId);
  }
}

const seatClass = (seatId: string) => {
  if (bookedSeats.value.includes(seatId)) {
    return 'bg-gray-200 text-gray-400 cursor-not-allowed shadow-inner ring-1 ring-gray-300 relative overflow-hidden after:content-[""] after:absolute after:inset-0 after:bg-[repeating-linear-gradient(45deg,transparent,transparent_2px,#d1d5db_2px,#d1d5db_4px)]';
  }
  if (selectedSeats.value.includes(seatId)) {
    return 'bg-primary-600 text-white shadow-lg ring-2 ring-primary-500 ring-offset-2 transform scale-105';
  }
  return 'bg-white text-gray-700 hover:bg-gray-50 hover:ring-primary-300 ring-1 ring-gray-200 shadow-sm hover:shadow-md';
}

const proceedToDetails = () => {
  // Store selected seats in localStorage or query params for simplicity
  router.push({
    path: '/checkout/details',
    query: { scheduleId, seats: selectedSeats.value.join(',') }
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
.animation-delay-300 {
  animation-delay: 0.3s;
}
</style>
