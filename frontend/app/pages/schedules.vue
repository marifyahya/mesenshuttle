<template>
  <div class="relative py-8 min-h-screen overflow-hidden">
    <!-- Animated Background Gradients -->
    <div class="fixed inset-0 bg-gray-50/80 z-0"></div>
    <div class="fixed -top-40 -right-40 w-96 h-96 bg-primary-300/20 rounded-full blur-3xl mix-blend-multiply animate-blob z-0"></div>
    <div class="fixed top-40 -left-20 w-80 h-80 bg-indigo-300/20 rounded-full blur-3xl mix-blend-multiply animate-blob animation-delay-2000 z-0"></div>

    <UContainer class="relative z-10 max-w-4xl mx-auto">
      <BookingStepper :current-step="1" class="mb-8" />
      
      <div class="mb-8 animate-fade-in-up">
        <div class="flex justify-between items-center mb-6">
          <h1 class="text-3xl font-extrabold text-gray-900 tracking-tight">Available Schedules</h1>
          <UButton variant="soft" color="primary" to="/" icon="i-lucide-search" class="font-medium rounded-xl">Change Search</UButton>
        </div>
        
        <div v-if="searchQuery.origin && searchQuery.destination" class="bg-white/60 backdrop-blur-md border border-white/80 rounded-2xl p-5 flex flex-col sm:flex-row sm:items-center gap-3 sm:gap-6 shadow-xl ring-1 ring-gray-100">
          <div class="flex items-center gap-4 text-xl font-bold text-gray-900">
            <span class="bg-clip-text text-transparent bg-linear-to-r from-primary-600 to-indigo-600">{{ searchQuery.origin }}</span>
            <div class="p-2 bg-primary-50 rounded-full">
              <UIcon name="i-lucide-arrow-right" class="w-5 h-5 text-primary-600" />
            </div>
            <span class="bg-clip-text text-transparent bg-linear-to-r from-primary-600 to-indigo-600">{{ searchQuery.destination }}</span>
          </div>
          <div class="hidden sm:block h-8 w-px bg-gray-200 mx-2"></div>
          <div class="flex items-center gap-2.5 text-gray-600 font-medium bg-white/80 px-4 py-2 rounded-xl shadow-sm">
            <UIcon name="i-lucide-calendar" class="w-4 h-4 text-primary-500" />
            <span>{{ formattedSearchDate }}</span>
          </div>
        </div>
      </div>

      <div class="space-y-4">
        <div v-if="filteredSchedules.length === 0" class="text-center py-16 bg-white/80 backdrop-blur-sm rounded-3xl shadow-xl ring-1 ring-gray-100 animate-fade-in-up animation-delay-300">
          <div class="w-20 h-20 bg-gray-50 rounded-full flex items-center justify-center mx-auto mb-4">
            <UIcon name="i-lucide-calendar-x" class="w-10 h-10 text-gray-400" />
          </div>
          <h3 class="text-xl font-bold text-gray-900 mb-2">No schedules found</h3>
          <p class="text-gray-500">We couldn't find any shuttles matching your search criteria.</p>
          <UButton to="/" color="primary" class="mt-6 rounded-xl font-bold" size="lg">Try Another Date</UButton>
        </div>
        <template v-else>
          <div 
            v-for="(item, index) in filteredSchedules" 
            :key="item?.id" 
            class="relative group animate-fade-in-up"
            :style="`animation-delay: ${300 + index * 100}ms`"
          >
            <!-- Premium Ticket Card Design -->
            <div class="absolute inset-0 bg-linear-to-r from-primary-600 to-indigo-600 rounded-2xl transform translate-y-1 translate-x-1 opacity-0 group-hover:opacity-20 transition-all duration-300"></div>
            
            <div class="relative bg-white/90 backdrop-blur-xl border border-gray-100 rounded-2xl shadow-lg hover:shadow-2xl transition-all duration-300 group-hover:-translate-y-1 overflow-hidden">
              <!-- Left accent border -->
              <div class="absolute left-0 top-0 bottom-0 w-1.5" :class="item?.fleet?.type === 'Premium' ? 'bg-amber-400' : 'bg-primary-500'"></div>
              
              <div class="p-6 flex flex-col sm:flex-row sm:items-center justify-between gap-6 pl-8">
                <!-- Left: Time and Fleet Type -->
                <div class="flex items-center gap-6">
                  <div class="text-center">
                    <span class="block text-3xl font-black text-gray-900 tracking-tight">{{ new Date(item!.departureTime).toLocaleTimeString([], {hour: '2-digit', minute:'2-digit', hour12: false}) }}</span>
                    <span class="text-xs font-semibold text-gray-400 uppercase tracking-wider">Departure</span>
                  </div>
                  
                  <div class="h-12 w-px bg-gray-200"></div>
                  
                  <div class="flex flex-col gap-2">
                    <div class="flex items-center gap-2">
                      <span 
                        class="inline-flex items-center gap-1.5 px-3 py-1.5 text-xs font-bold rounded-lg shadow-sm" 
                        :class="item?.fleet?.type === 'Premium' ? 'bg-amber-100 text-amber-800 ring-1 ring-amber-200' : 'bg-slate-100 text-slate-700 ring-1 ring-slate-200'"
                      >
                        <UIcon :name="item?.fleet?.type === 'Premium' ? 'i-lucide-crown' : 'i-lucide-armchair'" class="w-4 h-4" />
                        {{ item?.fleet?.type }}
                      </span>
                      <span class="inline-flex items-center gap-1.5 text-xs font-bold text-orange-700 bg-orange-100 px-3 py-1.5 rounded-lg ring-1 ring-orange-200 shadow-sm">
                        <UIcon name="i-lucide-users" class="w-4 h-4" />
                        {{ item?.availableSeats }} seats left
                      </span>
                    </div>
                    <span class="text-sm font-medium text-gray-500">{{ item?.fleet?.name || 'MesenShuttle Standard' }}</span>
                  </div>
                </div>
                
                <!-- Right: Price, Button -->
                <div class="flex items-center justify-between sm:justify-end gap-6 w-full sm:w-auto border-t sm:border-t-0 pt-4 sm:pt-0 border-gray-100 border-dashed">
                  <div class="text-right">
                    <span class="block text-xs font-semibold text-gray-400 uppercase tracking-wider mb-0.5">Price</span>
                    <span class="text-2xl font-extrabold text-primary-600 tracking-tight">Rp {{ item?.price.toLocaleString() }}</span>
                  </div>
                  <UButton :to="`/seat-selection/${item?.id}`" color="primary" size="xl" class="font-bold px-8 rounded-xl shadow-lg hover:shadow-xl transition-all" trailing-icon="i-lucide-arrow-right">
                    Select
                  </UButton>
                </div>
              </div>
            </div>
          </div>
        </template>
      </div>
    </UContainer>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import { useMockDataStore } from '~/stores/mockData'

const route = useRoute()
const store = useMockDataStore()

const searchQuery = computed(() => ({
  origin: route.query.origin as string,
  destination: route.query.destination as string,
  date: route.query.date as string
}))

const formattedSearchDate = computed(() => {
  if (!searchQuery.value.date) return '';
  const dateObj = new Date(searchQuery.value.date);
  return dateObj.toLocaleDateString('id-ID', { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' });
})

const filteredSchedules = computed(() => {
  return store.schedules.map(schedule => {
    return store.getScheduleDetails(schedule.id)
  }).filter(s => {
    if (!s) return false;
    let match = true;
    
    // Original Search Query Filters (now searching by specific Pool instead of City)
    if (searchQuery.value.origin && s.route?.originPool !== searchQuery.value.origin) match = false;
    if (searchQuery.value.destination && s.route?.destinationPool !== searchQuery.value.destination) match = false;
    
    return match;
  }).sort((a, b) => new Date(a!.departureTime).getTime() - new Date(b!.departureTime).getTime())
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
.animation-delay-300 {
  animation-delay: 0.3s;
}
</style>
