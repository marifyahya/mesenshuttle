<template>
  <div class="py-12">
    <UContainer>
      <div class="text-center mb-12">
        <h1 class="text-4xl font-extrabold tracking-tight text-gray-900 sm:text-5xl md:text-6xl">
          <span class="block">Book your next journey</span>
          <span class="block text-primary-600">with MesenShuttle</span>
        </h1>
        <p class="mt-3 max-w-md mx-auto text-base text-gray-500 sm:text-lg md:mt-5 md:text-xl md:max-w-3xl">
          Fast, reliable, and premium shuttle bookings without the hassle of registration.
        </p>
      </div>

      <UCard class="max-w-4xl mx-auto">
        <template #header>
          <h3 class="text-lg font-medium leading-6 text-gray-900">Search Schedules</h3>
        </template>
        <div class="grid grid-cols-1 gap-y-6 sm:grid-cols-3 sm:gap-x-6">
          <UFormField label="Origin" name="origin" :error="formErrors.origin || undefined">
            <USelect v-model="searchForm.origin" :items="groupedOrigins" placeholder="Select origin pool" size="xl" class="w-full" />
          </UFormField>
          <UFormField label="Destination" name="destination" :error="formErrors.destination || undefined">
            <USelect v-model="searchForm.destination" :items="groupedDestinations" placeholder="Select destination pool" :disabled="!searchForm.origin" size="xl" class="w-full" />
          </UFormField>
          <UFormField label="Date" name="date">
            <UPopover class="w-full">
              <UButton color="neutral" variant="outline" icon="i-lucide-calendar" class="w-full justify-between text-gray-700" size="xl">
                {{ formattedDate }}
              </UButton>
              <template #content>
                <UCalendar v-model="(searchForm.date as any)" :min-value="currentDate" />
              </template>
            </UPopover>
          </UFormField>
        </div>
        <template #footer>
          <div class="flex justify-end">
            <UButton @click="handleSearch" size="lg" color="primary">Search Shuttles</UButton>
          </div>
        </template>
      </UCard>
    </UContainer>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useMockDataStore } from '~/stores/mockData'
import { getLocalTimeZone, today } from '@internationalized/date'

const router = useRouter()
const store = useMockDataStore()

const groupedOrigins = computed(() => {
  const groups = new Map<string, string[]>()
  store.routes.forEach(r => {
    if (!groups.has(r.originCity)) groups.set(r.originCity, [])
    if (!groups.get(r.originCity)!.includes(r.originPool)) {
      groups.get(r.originCity)!.push(r.originPool)
    }
  })

  const result = []
  for (const [city, pools] of groups.entries()) {
    const group = [
      { label: city, type: 'label', class: 'font-bold text-lg text-primary-600' },
      ...pools.map(pool => ({ label: pool, value: pool }))
    ]
    result.push(group)
  }
  return result
})

const selectedOriginCity = computed(() => {
  if (!searchForm.value.origin) return null;
  const route = store.routes.find(r => r.originPool === searchForm.value.origin);
  return route ? route.originCity : null;
})

const selectedDestinationCity = computed(() => {
  if (!searchForm.value.destination) return null;
  const route = store.routes.find(r => r.destinationPool === searchForm.value.destination);
  return route ? route.destinationCity : null;
})

const groupedDestinations = computed(() => {
  const groups = new Map<string, string[]>()
  store.routes.forEach(r => {
    // Do not include any pools from the same city as the selected origin
    if (r.destinationCity !== selectedOriginCity.value) {
      if (!groups.has(r.destinationCity)) groups.set(r.destinationCity, [])
      if (!groups.get(r.destinationCity)!.includes(r.destinationPool)) {
        groups.get(r.destinationCity)!.push(r.destinationPool)
      }
    }
  })

  const result = []
  for (const [city, pools] of groups.entries()) {
    const group = [
      { label: city, type: 'label', class: 'font-bold text-lg text-primary-600' },
      ...pools.map(pool => ({ label: pool, value: pool }))
    ]
    result.push(group)
  }
  return result
})

// Get today's date using internationalized date
const currentDate = today(getLocalTimeZone())

const searchForm = ref({
  origin: '',
  destination: '',
  date: currentDate
})

const formErrors = ref({
  origin: '',
  destination: ''
})

const formattedDate = computed(() => {
  if (!searchForm.value.date) return 'Select date'
  const d = searchForm.value.date
  const day = d.day.toString().padStart(2, '0')
  const month = d.month.toString().padStart(2, '0')
  const year = d.year
  return `${day}-${month}-${year}`
})

// Automatically reset destination if it conflicts with the new origin city, and clear origin error
watch(() => searchForm.value.origin, () => {
  formErrors.value.origin = ''
  if (selectedOriginCity.value && selectedOriginCity.value === selectedDestinationCity.value) {
    searchForm.value.destination = ''
  }
})

// Clear destination error when a destination is selected
watch(() => searchForm.value.destination, () => {
  formErrors.value.destination = ''
})

const handleSearch = () => {
  formErrors.value = { origin: '', destination: '' }
  let hasError = false
  
  if (!searchForm.value.origin) {
    formErrors.value.origin = 'Origin is required'
    hasError = true
  }
  if (!searchForm.value.destination) {
    formErrors.value.destination = 'Destination is required'
    hasError = true
  }
  
  if (hasError) return;
  
  // Pass state via query
  router.push({
    path: '/schedules',
    query: { origin: searchForm.value.origin, destination: searchForm.value.destination, date: searchForm.value.date.toString() }
  })
}
</script>
