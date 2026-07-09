<template>
  <div class="py-8">
    <UContainer>
      <BookingStepper :current-step="2" />
      <UButton variant="ghost" to="/schedules" class="mb-4" icon="i-heroicons-arrow-left">Back to Schedules</UButton>
      
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- Seat Map Area -->
        <div class="lg:col-span-2">
          <UCard>
            <template #header>
              <h3 class="font-bold text-lg">Select Your Seat</h3>
            </template>
            
            <div class="text-center mb-6">
              <h2 class="text-2xl font-bold mb-4">Pilih Seat</h2>
              <div class="flex justify-center gap-6 text-base">
                <div class="flex items-center gap-2">
                  <div class="w-6 h-6 bg-green-700 rounded-full"></div>
                  <span class="font-bold">Tersedia</span>
                </div>
                <div class="flex items-center gap-2">
                  <div class="w-6 h-6 bg-red-600 rounded-full"></div>
                  <span class="font-bold">Dipilih</span>
                </div>
              </div>
            </div>

            <div v-if="scheduleDetails" class="max-w-xs mx-auto">
              
              <div class="bg-gray-100 border border-gray-300 rounded-lg py-4 text-center mb-8">
                <span class="text-xl font-bold text-gray-400">Silakan Pilih Seat</span>
              </div>

              <template v-if="scheduleDetails?.fleet.type === 'Reguler'">
                <!-- Reguler Shuttle (12 Seats) Grid -->
                <div class="grid grid-cols-3 gap-4">
                  <!-- Row 1 -->
                  <button @click="toggleSeat('1')" :class="seatClass('1')" class="h-16 rounded-xl flex items-center justify-center font-bold text-lg transition-colors">1</button>
                  <div></div>
                  <div class="bg-black text-white h-16 rounded-xl flex items-center justify-center font-bold text-lg shadow-sm">Supir</div>
                  
                  <!-- Row 2 (Door on left) -->
                  <div></div>
                  <button @click="toggleSeat('2')" :class="seatClass('2')" class="h-16 rounded-xl flex items-center justify-center font-bold text-lg transition-colors">2</button>
                  <button @click="toggleSeat('3')" :class="seatClass('3')" class="h-16 rounded-xl flex items-center justify-center font-bold text-lg transition-colors">3</button>
                  
                  <!-- Row 3 -->
                  <button @click="toggleSeat('4')" :class="seatClass('4')" class="h-16 rounded-xl flex items-center justify-center font-bold text-lg transition-colors">4</button>
                  <button @click="toggleSeat('5')" :class="seatClass('5')" class="h-16 rounded-xl flex items-center justify-center font-bold text-lg transition-colors">5</button>
                  <button @click="toggleSeat('6')" :class="seatClass('6')" class="h-16 rounded-xl flex items-center justify-center font-bold text-lg transition-colors">6</button>
                  
                  <!-- Row 4 -->
                  <button @click="toggleSeat('7')" :class="seatClass('7')" class="h-16 rounded-xl flex items-center justify-center font-bold text-lg transition-colors">7</button>
                  <button @click="toggleSeat('8')" :class="seatClass('8')" class="h-16 rounded-xl flex items-center justify-center font-bold text-lg transition-colors">8</button>
                  <button @click="toggleSeat('9')" :class="seatClass('9')" class="h-16 rounded-xl flex items-center justify-center font-bold text-lg transition-colors">9</button>
                  
                  <!-- Row 5 -->
                  <button @click="toggleSeat('10')" :class="seatClass('10')" class="h-16 rounded-xl flex items-center justify-center font-bold text-lg transition-colors">10</button>
                  <button @click="toggleSeat('11')" :class="seatClass('11')" class="h-16 rounded-xl flex items-center justify-center font-bold text-lg transition-colors">11</button>
                  <button @click="toggleSeat('12')" :class="seatClass('12')" class="h-16 rounded-xl flex items-center justify-center font-bold text-lg transition-colors">12</button>
                </div>
              </template>

              <template v-else-if="scheduleDetails?.fleet.type === 'Premium'">
                <!-- Premium Shuttle (10 Seats) Grid -->
                <div class="grid grid-cols-4 gap-4">
                  <!-- Row 1 -->
                  <button @click="toggleSeat('1')" :class="seatClass('1')" class="h-16 rounded-xl flex items-center justify-center font-bold text-lg transition-colors">1</button>
                  <div></div>
                  <div></div>
                  <div class="bg-black text-white h-16 rounded-xl flex items-center justify-center font-bold text-lg shadow-sm">Supir</div>
                  
                  <!-- Row 2 -->
                  <div></div>
                  <div></div>
                  <button @click="toggleSeat('2')" :class="seatClass('2')" class="h-16 rounded-xl flex items-center justify-center font-bold text-lg transition-colors">2</button>
                  <button @click="toggleSeat('3')" :class="seatClass('3')" class="h-16 rounded-xl flex items-center justify-center font-bold text-lg transition-colors">3</button>
                  
                  <!-- Row 3 -->
                  <button @click="toggleSeat('4')" :class="seatClass('4')" class="h-16 rounded-xl flex items-center justify-center font-bold text-lg transition-colors">4</button>
                  <div></div>
                  <button @click="toggleSeat('5')" :class="seatClass('5')" class="h-16 rounded-xl flex items-center justify-center font-bold text-lg transition-colors">5</button>
                  <button @click="toggleSeat('6')" :class="seatClass('6')" class="h-16 rounded-xl flex items-center justify-center font-bold text-lg transition-colors">6</button>
                  
                  <!-- Row 4 -->
                  <button @click="toggleSeat('7')" :class="seatClass('7')" class="h-16 rounded-xl flex items-center justify-center font-bold text-lg transition-colors">7</button>
                  <button @click="toggleSeat('8')" :class="seatClass('8')" class="h-16 rounded-xl flex items-center justify-center font-bold text-lg transition-colors">8</button>
                  <button @click="toggleSeat('9')" :class="seatClass('9')" class="h-16 rounded-xl flex items-center justify-center font-bold text-lg transition-colors">9</button>
                  <button @click="toggleSeat('10')" :class="seatClass('10')" class="h-16 rounded-xl flex items-center justify-center font-bold text-lg transition-colors">10</button>
                </div>
              </template>

              <template v-else-if="scheduleDetails?.fleet.type === 'Premium'">
                <!-- Premium Shuttle (8 Seats - Captain Seats) -->
                <!-- Front Row -->
                <div class="flex justify-center gap-8 mb-4">
                  <div class="w-14 h-14"><button @click="toggleSeat('1')" :class="seatClass('1')" class="w-full h-full rounded-lg shadow-sm flex items-center justify-center font-bold text-sm transition-colors border">1</button></div>
                  <div class="w-14 h-14 flex items-center justify-center bg-gray-200 text-gray-500 rounded-lg font-bold text-xs shadow-sm">DRIVER</div>
                </div>
                <!-- Rows 1 to 2 (1-1 layout) -->
                <div v-for="(row, idx) in [[2,3], [4,5]]" :key="idx" class="flex justify-center gap-8 mb-2">
                  <div class="w-14 h-14"><button @click="toggleSeat(String(row[0]))" :class="seatClass(String(row[0]))" class="w-full h-full rounded-lg shadow-sm flex items-center justify-center font-bold text-sm transition-colors border">{{row[0]}}</button></div>
                  <div class="w-14 h-14"><button @click="toggleSeat(String(row[1]))" :class="seatClass(String(row[1]))" class="w-full h-full rounded-lg shadow-sm flex items-center justify-center font-bold text-sm transition-colors border">{{row[1]}}</button></div>
                </div>
                <!-- Back Row (3 seats) -->
                <div class="flex justify-center gap-3 mt-4">
                   <div class="w-14 h-14"><button @click="toggleSeat('6')" :class="seatClass('6')" class="w-full h-full rounded-lg shadow-sm flex items-center justify-center font-bold text-sm transition-colors border">6</button></div>
                   <div class="w-14 h-14"><button @click="toggleSeat('7')" :class="seatClass('7')" class="w-full h-full rounded-lg shadow-sm flex items-center justify-center font-bold text-sm transition-colors border">7</button></div>
                   <div class="w-14 h-14"><button @click="toggleSeat('8')" :class="seatClass('8')" class="w-full h-full rounded-lg shadow-sm flex items-center justify-center font-bold text-sm transition-colors border">8</button></div>
                </div>
              </template>
            </div>
          </UCard>
        </div>

        <!-- Booking Summary -->
        <div>
          <UCard>
            <template #header>
              <h3 class="font-bold">Booking Summary</h3>
            </template>
            <div class="space-y-4" v-if="scheduleDetails">
              <div>
                <p class="text-sm text-gray-500">Route</p>
                <p class="font-medium">{{ scheduleDetails.route.originPool }} &rarr; {{ scheduleDetails.route.destinationPool }}</p>
              </div>
              <div>
                <p class="text-sm text-gray-500">Departure</p>
                <p class="font-medium">{{ new Date(scheduleDetails.departureTime).toLocaleString() }}</p>
              </div>
              <div>
                <p class="text-sm text-gray-500">Bus</p>
                <p class="font-medium">{{ scheduleDetails.fleet.name }} ({{ scheduleDetails.fleet.type }})</p>
              </div>
              <UDivider />
              <div class="flex justify-between items-center">
                <span class="text-gray-700">Selected Seats:</span>
                <span class="font-bold">{{ selectedSeats.join(', ') || 'None' }}</span>
              </div>
              <div class="flex justify-between items-center text-lg">
                <span class="font-bold">Total:</span>
                <span class="font-bold text-primary-600">Rp {{ (selectedSeats.length * scheduleDetails.price).toLocaleString() }}</span>
              </div>
            </div>
            <template #footer>
              <UButton :disabled="selectedSeats.length === 0" block color="primary" @click="proceedToDetails">
                Continue to Details
              </UButton>
            </template>
          </UCard>
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
  if (bookedSeats.value.includes(seatId)) return 'bg-black text-white cursor-not-allowed shadow-sm';
  if (selectedSeats.value.includes(seatId)) return 'bg-red-600 text-white shadow-md transform scale-105';
  return 'bg-green-700 text-white hover:bg-green-800 shadow-sm';
}

const proceedToDetails = () => {
  // Store selected seats in localStorage or query params for simplicity
  router.push({
    path: '/checkout/details',
    query: { scheduleId, seats: selectedSeats.value.join(',') }
  })
}
</script>
