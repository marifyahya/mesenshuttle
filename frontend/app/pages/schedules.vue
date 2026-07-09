<template>
  <div class="py-8 bg-gray-50 min-h-screen">
    <UContainer class="max-w-4xl mx-auto">
      <BookingStepper :current-step="1" />
      
      <div class="mb-6">
        <div class="flex justify-between items-center mb-4">
          <h1 class="text-2xl font-bold text-gray-900">Available Schedules</h1>
          <UButton variant="soft" to="/">Change Search</UButton>
        </div>
        
        <div v-if="searchQuery.origin && searchQuery.destination" class="bg-primary-50 border border-primary-100 rounded-lg p-4 flex flex-col sm:flex-row sm:items-center gap-2 sm:gap-4 shadow-sm">
          <div class="flex items-center gap-3 text-lg font-bold text-primary-900">
            <span>{{ searchQuery.origin }}</span>
            <UIcon name="i-lucide-arrow-right" class="w-5 h-5 text-primary-500" />
            <span>{{ searchQuery.destination }}</span>
          </div>
          <div class="hidden sm:block h-6 w-px bg-primary-200 mx-2"></div>
          <div class="flex items-center gap-2 text-primary-700 font-medium">
            <UIcon name="i-lucide-calendar" class="w-4 h-4" />
            <span>{{ formattedSearchDate }}</span>
          </div>
        </div>
      </div>

      <div class="space-y-2">
        <div v-if="filteredSchedules.length === 0" class="text-center py-12 bg-white rounded-lg shadow-sm border border-gray-100">
          <p class="text-gray-500">No schedules found matching your search.</p>
        </div>
        <template v-else>
          <UCard v-for="item in filteredSchedules" :key="item?.id" class="hover:border-primary-300 transition-colors" :ui="{ body: 'p-4 sm:p-5' }">
            <div v-if="item" class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
              <!-- Left: Time and Fleet Type -->
              <div class="flex items-center gap-4">
                <span class="text-2xl font-black text-gray-900">{{ new Date(item.departureTime).toLocaleTimeString([], {hour: '2-digit', minute:'2-digit', hour12: false}) }}</span>
                <div class="flex items-center gap-2">
                  <span 
                    class="inline-flex items-center gap-1.5 px-2.5 py-1 text-xs font-semibold rounded-md border" 
                    :class="item.fleet?.type === 'Premium' ? 'bg-amber-50 text-amber-700 border-amber-200' : 'bg-slate-50 text-slate-700 border-slate-200'"
                  >
                    <UIcon :name="item.fleet?.type === 'Premium' ? 'i-lucide-crown' : 'i-lucide-armchair'" class="w-4 h-4" />
                    {{ item.fleet?.type }}
                  </span>
                  <span class="inline-flex items-center gap-1 text-[11px] sm:text-xs font-medium text-orange-600 bg-orange-50 px-2 py-1 rounded-md border border-orange-100">
                    <UIcon name="i-lucide-armchair" class="w-3.5 h-3.5" />
                    Only {{ item.availableSeats }} left
                  </span>
                </div>
              </div>
              
              <!-- Right: Price, Button -->
              <div class="flex items-center justify-between sm:justify-end gap-4 sm:gap-6 w-full sm:w-auto border-t sm:border-t-0 pt-3 sm:pt-0 border-gray-100">
                <span class="text-lg sm:text-xl font-extrabold text-gray-900 tracking-tight">Rp {{ item.price.toLocaleString() }}</span>
                <UButton :to="`/seat-selection/${item.id}`" color="primary" size="md" class="font-bold px-4 sm:px-6">Select Seat</UButton>
              </div>
            </div>
          </UCard>
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
